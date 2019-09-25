package policy

import (
	"net/http"

	"github.com/pthethanh/robusta/internal/app/auth"
	"github.com/pthethanh/robusta/internal/pkg/http/router"
)

func (h *Handler) Routes() []router.Route {
	return []router.Route{
		{
			Path:        "/api/v1/policies",
			Method:      http.MethodPost,
			Handler:     h.AssignPolicy,
			Queries:     []string{"action", "add-policy"},
			Middlewares: []router.Middleware{auth.RequiredAuthMiddleware},
		},
		{
			Path:        "/api/v1/policies",
			Method:      http.MethodPost,
			Handler:     h.AssignGroupPolicy,
			Queries:     []string{"action", "add-group-policy"},
			Middlewares: []router.Middleware{auth.RequiredAuthMiddleware},
		},
		{
			Path:        "/api/v1/policies/actions",
			Method:      http.MethodGet,
			Handler:     h.ListActions,
			Middlewares: []router.Middleware{auth.RequiredAuthMiddleware},
		},
	}
}
