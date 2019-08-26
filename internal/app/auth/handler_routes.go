package auth

import (
	"net/http"

	"github.com/pthethanh/robusta/internal/pkg/http/router"
)

func (h *Handler) Routes() []router.Route {
	return []router.Route{
		{
			Path:    "/api/v1/auth",
			Method:  http.MethodPost,
			Handler: h.Auth,
		},
	}
}
