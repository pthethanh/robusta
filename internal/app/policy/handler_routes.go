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
			Method:      http.MethodGet,
			Handler:     h.FindPolicies,
			Middlewares: []router.Middleware{auth.RequiredAuthMiddleware},
		},
		{
			Path:        "/api/v1/policies",
			Method:      http.MethodDelete,
			Handler:     h.RemovePolicy,
			Middlewares: []router.Middleware{auth.RequiredAuthMiddleware},
		},
		{
			Path:        "/api/v1/policies",
			Method:      http.MethodPost,
			Handler:     h.AddPolicy,
			Queries:     []string{"action", "add-policy"},
			Middlewares: []router.Middleware{auth.RequiredAuthMiddleware},
		},
		{
			Path:        "/api/v1/policies",
			Method:      http.MethodPost,
			Handler:     h.AddGroupPolicy,
			Queries:     []string{"action", "add-group-policy"},
			Middlewares: []router.Middleware{auth.RequiredAuthMiddleware},
		},
		{
			Path:        "/api/v1/policies/actions",
			Method:      http.MethodGet,
			Handler:     h.ListActions,
			Middlewares: []router.Middleware{auth.RequiredAuthMiddleware},
		},
		{
			Path:        "/api/v1/policies/roles",
			Method:      http.MethodGet,
			Handler:     h.GetRoles,
			Middlewares: []router.Middleware{auth.RequiredAuthMiddleware},
		},
		{
			Path:        "/api/v1/policies/roles/{role:[a-zA-Z0-9_]+}/users",
			Method:      http.MethodGet,
			Handler:     h.GetUsersForRole,
			Middlewares: []router.Middleware{auth.RequiredAuthMiddleware},
		},
	}
}
