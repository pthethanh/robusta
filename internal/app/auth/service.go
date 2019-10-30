package auth

import (
	"context"

	"github.com/pthethanh/robusta/internal/app/status"
	"github.com/pthethanh/robusta/internal/app/types"
	"github.com/pthethanh/robusta/internal/pkg/jwt"
	"github.com/pthethanh/robusta/internal/pkg/log"
)

type (
	Authenticator interface {
		Auth(ctx context.Context, username, password string) (*types.User, error)
	}
	Service struct {
		jwtSigner      jwt.Signer
		authenticators map[string]Authenticator
	}
)

func NewService(signer jwt.Signer) *Service {
	return &Service{
		jwtSigner:      signer,
		authenticators: make(map[string]Authenticator),
	}
}

// RegisterAuthenticator register an authenticator.
func (s *Service) RegisterAuthenticator(name string, a Authenticator) {
	s.authenticators[name] = a
}

// Auth try to use the given information to login with the configured authenticators by priority.
// Auth return JWT token, user information and error if any.
func (s *Service) Auth(ctx context.Context, username, password string) (string, *types.User, error) {
	for name, authenticator := range s.authenticators {
		user, err := authenticator.Auth(ctx, username, password)
		if err != nil {
			log.WithContext(ctx).Errorf("failed to login with %s, err: %#v", name, err)
			continue
		}
		token, err := s.jwtSigner.Sign(userToClaims(user, jwt.DefaultLifeTime))
		if err != nil {
			log.WithContext(ctx).Errorf("failed to generate JWT token, err: %v", err)
			return "", nil, err
		}
		return token, user, nil
	}
	return "", nil, status.Auth().InvalidUserPassword
}
