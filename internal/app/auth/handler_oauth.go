// Package auth provide OAuth mechanism such as Google, Facebook, Github,...
package auth

import (
	"context"
	"errors"
	"net/http"

	"github.com/pthethanh/robusta/internal/app/status"
	"github.com/pthethanh/robusta/internal/app/types"
	"github.com/pthethanh/robusta/internal/pkg/config/envconfig"
	"github.com/pthethanh/robusta/internal/pkg/http/respond"
	"github.com/pthethanh/robusta/internal/pkg/jwt"
	"github.com/pthethanh/robusta/internal/pkg/log"
	"github.com/pthethanh/robusta/internal/pkg/uuid"

	"github.com/gorilla/sessions"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/google"
)

type (
	// Provider define what an OAuth provider needs to have
	Provider = goth.Provider
	// Google configuration for Google OAuth
	Google struct {
		ClientID     string   `envconfig:"GOOGLE_CLIENT_ID"`
		ClientSecret string   `envconfig:"GOOGLE_CLIENT_SECRET"`
		RedirectURL  string   `envconfig:"GOOGLE_REDIRECT_URL"`
		Scopes       []string `envconfig:"GOOGLE_SCOPES" default:"email,profile"`
	}

	// Config hold auth configuration
	Config struct {
		Enable  bool `envconfig:"AUTH_ENABLED" default:"true"`
		Session struct {
			Secret   string `envconfig:"SESSION_SECRET" default:"robusta - 1186"`
			Type     string `envconfig:"SESSION_STORE_TYPE" default:"memory"`
			FilePath string `envconfig:"SESSION_FILE_PATH" default:"./sessions"`
		}
	}

	UserService interface {
		FindByUserID(ctx context.Context, userID string) (*types.User, error)
		Create(ctx context.Context, user *types.User) (string, error)
		Update(ctx context.Context, userID string, user *types.User) error
	}

	OAuth2Handler struct {
		userService UserService
		jwtSigner   jwt.Signer
		session     sessions.Store
	}
)

const (
	redirectURIKey = "_r_auth_redirect_uri"
)

// NewGoogleProvider return Google OAuth provider
func NewGoogleProvider() Provider {
	var googleConf Google
	envconfig.Load(&googleConf)
	return google.New(googleConf.ClientID, googleConf.ClientSecret, googleConf.RedirectURL, "profile", "email")
}

// NewOAuth2Handler return new OAuth2Handler
func NewOAuth2Handler(conf Config, srv UserService, jwtSigner jwt.Signer, providers ...Provider) *OAuth2Handler {
	goth.UseProviders(providers...)
	session := createSessionStore(conf)
	gothic.Store = session
	return &OAuth2Handler{
		userService: srv,
		jwtSigner:   jwtSigner,
		session:     session,
	}
}

// Callback handle OAuth callback
func (h *OAuth2Handler) Callback(w http.ResponseWriter, r *http.Request) {
	redirectURI := "/"
	session, err := h.session.Get(r, gothic.GetState(r))
	if err != nil {
		log.WithContext(r.Context()).Errorf("failed to get state session, err: %v", err)
		respond.Error(w, errors.New("invalid session"), http.StatusBadRequest)
		return
	}
	if v, ok := session.Values[redirectURIKey]; ok {
		redirectURI = v.(string)
	}
	// we don't need the state session anymore, destroy it
	session.Options.MaxAge = -1
	if err := session.Save(r, w); err != nil {
		log.WithContext(r.Context()).Errorf("failed to destroy state session, err: %v", err)
	}
	// completes the auth process
	remoteUser, err := gothic.CompleteUserAuth(w, r)
	if err != nil {
		log.WithContext(r.Context()).Errorf("failed to login, err: %v", err)
		respond.Error(w, status.Policy().Unauthorized, http.StatusUnauthorized)
		return
	}
	// update user info to our DB
	user := remoteUserToUser(&remoteUser)
	existingUser, err := h.upsertUser(r.Context(), user)
	if err != nil {
		log.WithContext(r.Context()).Errorf("failed to upsert user, err: %v", err)
		http.Redirect(w, r, redirectURI, http.StatusInternalServerError)
		return
	}
	// check if user is locked
	if existingUser.IsLocked() {
		log.WithContext(r.Context()).Errorf("user is locked for login, user_id: %s, name: %s", existingUser.UserID, existingUser.Name)
		http.Redirect(w, r, redirectURI, http.StatusUnauthorized)
		return
	}
	// gen token
	user.Roles = existingUser.Roles
	token, err := h.jwtSigner.Sign(userToClaims(user, jwt.DefaultLifeTime))
	if err != nil {
		log.WithContext(r.Context()).Errorf("failed to generate JWT token, err: %v", err)
		http.Redirect(w, r, getRedirectURL(r), http.StatusInternalServerError)
		return
	}
	// set necessary cookies for frondend usages...
	setCookies(w, r, token, user)

	// redirect
	http.Redirect(w, r, redirectURI, http.StatusFound)
}

// Logout handler OAuth logout
func (h *OAuth2Handler) Logout(w http.ResponseWriter, r *http.Request) {
	gothic.Logout(w, r)
	respond.JSON(w, http.StatusOK, status.Success())
}

// Login handle OAuth login
func (h *OAuth2Handler) Login(w http.ResponseWriter, r *http.Request) {
	// try to get the user without re-authenticating
	if user, err := gothic.CompleteUserAuth(w, r); err == nil {
		log.WithContext(r.Context()).Debugf("user already login: %v", user)
		return
	}
	// override the default state of gothic
	state := uuid.New()
	q := r.URL.Query()
	q.Add("state", state)
	r.URL.RawQuery = q.Encode()

	// store state values for later usages
	session, err := h.session.New(r, state)
	if err != nil {
		log.WithContext(r.Context()).Errorf("failed create new state session, err: %v", err)
		respond.Error(w, errors.New("invalid session"), http.StatusInternalServerError)
		return
	}
	session.Values[redirectURIKey] = getRedirectURL(r)
	if err := session.Save(r, w); err != nil {
		log.WithContext(r.Context()).Errorf("failed to store redirect uri, err: %v", err)
		respond.Error(w, errors.New("invalid session"), http.StatusInternalServerError)
		return
	}

	// re-login
	log.WithContext(r.Context()).Debugf("user not login yet, let go to login...")
	gothic.BeginAuthHandler(w, r)
}

// upsertUser perform an update if user already exist, otherwise new user will be created
func (h *OAuth2Handler) upsertUser(ctx context.Context, newUser *types.User) (*types.User, error) {
	existingUser, err := h.userService.FindByUserID(ctx, newUser.UserID)
	if err != nil && !errors.Is(err, status.Gen().NotFound) {
		log.WithContext(ctx).Errorf("failed to find user, err: %v", err)
		return nil, err
	}
	log.WithContext(ctx).WithFields(log.Fields{"provider": newUser.Provider, "user_id": newUser.UserID, "email": newUser.Email, "is_new": existingUser == nil}).Debugf("upsert user")
	if existingUser == nil {
		if _, err := h.userService.Create(ctx, newUser); err != nil {
			return nil, err
		}
		return newUser, nil
	}
	// update existing user
	updateUserInfo(existingUser, newUser)
	if err := h.userService.Update(ctx, newUser.UserID, existingUser); err != nil {
		return nil, err
	}
	return existingUser, nil
}
