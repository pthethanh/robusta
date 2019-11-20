package auth

import (
	"encoding/base64"
	"encoding/json"
	"net/http"
	"net/url"
	"time"

	"github.com/markbates/goth"
	"github.com/pthethanh/robusta/internal/app/types"
	"github.com/pthethanh/robusta/internal/pkg/jwt"
	"github.com/pthethanh/robusta/internal/pkg/log"
	"github.com/pthethanh/robusta/internal/pkg/util/timeutil"
)

const (
	TokenCookieName = "_r_token"
	UserCookieName  = "_r_user"
)

func updateUserInfo(dst *types.User, src *types.User) {
	dst.FirstName = src.FirstName
	dst.LastName = src.LastName
	dst.Location = src.Location
	dst.Name = src.Name
	dst.NickName = src.NickName
	dst.Provider = src.Provider
	dst.UserID = src.UserID
	dst.AvatarURL = src.AvatarURL
	dst.Description = src.Description
	dst.Email = src.Email
	dst.UpdateAt = timeutil.Now()
}

func remoteUserToUser(remoteUser *goth.User) *types.User {
	return &types.User{
		AvatarURL:   remoteUser.AvatarURL,
		Description: remoteUser.Description,
		Email:       remoteUser.Email,
		FirstName:   remoteUser.FirstName,
		LastName:    remoteUser.LastName,
		Location:    remoteUser.Location,
		Name:        remoteUser.Name,
		NickName:    remoteUser.NickName,
		Provider:    remoteUser.Provider,
		UserID:      remoteUser.UserID,
	}
}

// getRedirectURL return redirect URL
func getRedirectURL(r *http.Request) string {
	path := r.URL.Query().Get("redirect")
	if path == "" {
		return "/"
	}
	parsedURL, err := url.Parse(path)
	if err != nil {
		log.WithContext(r.Context()).Errorf("invalid redirect url, err: %v", err)
		return "/"
	}
	if parsedURL.IsAbs() {
		log.WithContext(r.Context()).Errorf("absolute redirect url is not allowed, err: %v", err)
		return "/"
	}
	return path
}

// userToClaims convert user to a JWTClaims for signing
func userToClaims(user *types.User, lifeTime time.Duration) jwt.Claims {
	return jwt.Claims{
		AvatarURL: user.AvatarURL,
		Roles:     user.Roles,
		Name:      user.Name,
		LastName:  user.LastName,
		FirstName: user.FirstName,
		UserID:    user.UserID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(lifeTime).Unix(),
			Id:        user.UserID,
			IssuedAt:  time.Now().Unix(),
			Issuer:    jwt.DefaultIssuer,
			Subject:   user.UserID,
		},
	}
}

// claimsToUser extract information from a JWT Claims and return a user
func claimsToUser(claims *jwt.Claims) *types.User {
	return &types.User{
		AvatarURL: claims.AvatarURL,
		Roles:     claims.Roles,
		Name:      claims.Name,
		LastName:  claims.LastName,
		FirstName: claims.FirstName,
		UserID:    claims.UserID,
	}
}

// createTokenCookie create a token cookie which will expires after 24h
func createTokenCookie(token string, r *http.Request) *http.Cookie {
	return &http.Cookie{
		Name:     TokenCookieName,
		Value:    token,
		Expires:  time.Now().Add(24 * time.Hour),
		Domain:   r.Host,
		Path:     "/",
		Secure:   false,
		HttpOnly: false, // allow client to access this cookie
	}
}

// createUserInfoCookie create a user info (base64) cookie which will expires after 24h
func createUserInfoCookie(user *types.User, r *http.Request) (*http.Cookie, error) {
	info, err := userInfoCookieValue(user)
	if err != nil {
		return nil, err
	}
	return &http.Cookie{
		Name:     UserCookieName,
		Value:    info,
		Expires:  time.Now().Add(24 * time.Hour),
		Domain:   r.Host,
		Path:     "/",
		HttpOnly: false, // allow client to access this cookie
	}, nil
}

func setCookies(w http.ResponseWriter, r *http.Request, token string, user *types.User) {
	tokenCookie := createTokenCookie(token, r)
	http.SetCookie(w, tokenCookie)
	// set user info to cookie
	userCookie, err := createUserInfoCookie(user.Strip(), r)
	if err != nil {
		log.WithContext(r.Context()).Errorf("failed to create user cookie, err: %v", err)
		http.Redirect(w, r, getRedirectURL(r), http.StatusInternalServerError)
		return
	}
	http.SetCookie(w, userCookie)
}

func userInfoCookieValue(user *types.User) (string, error) {
	b, err := json.Marshal(user)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(b), nil
}
