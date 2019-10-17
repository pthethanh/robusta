package user

import (
	"net/http"

	"github.com/pthethanh/robusta/internal/app/auth"
	"github.com/pthethanh/robusta/internal/pkg/http/router"
)

func (h *Handler) Routes() []router.Route {
	return []router.Route{
		{
			Path:        "/api/v1/users",
			Method:      http.MethodGet,
			Handler:     h.FindAll,
			Middlewares: []router.Middleware{auth.RequiredAuthMiddleware},
		},
		{
			Path:    "/api/v1/users/registration",
			Method:  http.MethodPost,
			Handler: h.Register,
		},
		{
			Path:    "/api/v1/users",
			Method:  http.MethodPut,
			Handler: h.GenerateResetPasswordToken,
			Queries: []string{"action", "request-reset-password"},
		},
		{
			Path:    "/api/v1/users",
			Method:  http.MethodPut,
			Handler: h.ResetPassword,
			Queries: []string{"action", "reset-password"},
		},
	}
}
