package api

import (
	"github.com/pthethanh/robusta/internal/app/auth"
	"github.com/pthethanh/robusta/internal/pkg/config/envconfig"
	"github.com/pthethanh/robusta/internal/pkg/jwt"
)

func newOAuth2Handler(signer jwt.Signer, userService auth.UserService) *auth.OAuth2Handler {
	var conf auth.Config
	envconfig.Load(&conf)
	googleProvider := auth.NewGoogleProvider()
	return auth.NewOAuth2Handler(conf, userService, signer, googleProvider)
}

func newAuthHandler(signer jwt.Signer, authenticators map[string]auth.Authenticator) *auth.Handler {
	srv := auth.NewService(signer)
	for name, authenticator := range authenticators {
		srv.RegisterAuthenticator(name, authenticator)
	}
	return auth.NewHandler(srv)
}

func newJWTSignVerifier() jwt.SignVerifier {
	var conf jwt.Config
	envconfig.Load(&conf)
	return jwt.New(conf)
}
