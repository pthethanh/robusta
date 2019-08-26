package auth

import (
	"context"
	"net/http"

	"github.com/pthethanh/robusta/internal/pkg/http/respond"

	"github.com/pthethanh/robusta/internal/app/types"
	"github.com/pthethanh/robusta/internal/pkg/jwt"
	"github.com/pthethanh/robusta/internal/pkg/log"
)

type (
	contextKey string
)

const (
	authContextKey contextKey = "r_auth_user"
)

// UserInfoMiddleware decode user info in  Authorization header
// and attach it into HTTP context
func UserInfoMiddleware(verifier jwt.Verifier) func(http.Handler) http.Handler {
	return func(inner http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			key := r.Header.Get("Authorization")
			if key == "" {
				inner.ServeHTTP(w, r)
				return
			}
			claims, err := verifier.Verify(key)
			if err != nil {
				log.WithContext(r.Context()).Errorf("invalid JWT token, err: %v", err)
				inner.ServeHTTP(w, r)
				return
			}
			newCtx := NewContext(r.Context(), claimsToUser(claims))
			r = r.WithContext(newCtx)
			log.WithContext(r.Context()).WithFields(log.Fields{"user_id": claims.UserID, "name": claims.Name}).Debugf("decode JWT successfully")
			inner.ServeHTTP(w, r)
		})
	}
}

// NewContext return a new context with user inside
func NewContext(ctx context.Context, user *types.User) context.Context {
	return context.WithValue(ctx, authContextKey, user)
}

// FromContext extract user information from the given context
func FromContext(ctx context.Context) *types.User {
	if v, ok := ctx.Value(authContextKey).(*types.User); ok {
		return v
	}
	return nil
}

// RequiredAuthMiddleware reject request that has not authenticated
func RequiredAuthMiddleware(inner http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if user := FromContext(r.Context()); user == nil {
			respond.JSON(w, http.StatusUnauthorized, ErrUnauthorized)
			return
		}
		inner.ServeHTTP(w, r)
	})
}
