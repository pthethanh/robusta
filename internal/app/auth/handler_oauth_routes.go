package auth

import (
	"net/http"

	"github.com/pthethanh/robusta/internal/pkg/http/router"
)

func (h *OAuth2Handler) Routes() []router.Route {
	return []router.Route{
		{
			Path:    "/auth/{provider}/callback",
			Method:  http.MethodGet,
			Handler: h.Callback,
		},
		{
			Path:        "/logout/{provider}",
			Method:      http.MethodPost,
			Handler:     h.Logout,
			Middlewares: []router.Middleware{RequiredAuthMiddleware},
		},
		{
			Path:    "/auth/{provider}",
			Method:  http.MethodGet,
			Handler: h.Login,
		},
	}
}
