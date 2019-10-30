package auth

import (
	"context"
	"net/http"

	"github.com/pthethanh/robusta/internal/app/status"
	"github.com/pthethanh/robusta/internal/app/types"
	"github.com/pthethanh/robusta/internal/pkg/http/respond"
	"github.com/pthethanh/robusta/internal/pkg/jwt"
	"github.com/pthethanh/robusta/internal/pkg/log"
)

type (
	contextKey string
)

const (
	authContextKey contextKey = "r_auth_user"
	adminContext   contextKey = "r_auth_admin"
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

// RequiredAuthMiddleware reject request that has not authenticated
func RequiredAuthMiddleware(inner http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if user := FromContext(r.Context()); user == nil {
			respond.JSON(w, http.StatusUnauthorized, status.Policy().Unauthorized)
			return
		}
		inner.ServeHTTP(w, r)
	})
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

// NewAdminContext return a context indicates that it's an admin call.
// This should be only used for internal call between services.
func NewAdminContext(ctx context.Context) context.Context {
	return context.WithValue(ctx, adminContext, true)
}

// IsAdminContext check whether the given context is an admin call.
func IsAdminContext(ctx context.Context) bool {
	if v, ok := ctx.Value(adminContext).(bool); ok {
		return v
	}
	return false
}
