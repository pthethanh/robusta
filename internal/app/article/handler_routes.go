package article

import (
	"net/http"

	"github.com/pthethanh/robusta/internal/app/auth"
	"github.com/pthethanh/robusta/internal/pkg/http/router"
)

func (h *Handler) Routes() []router.Route {
	return []router.Route{
		{
			Path:    "/api/v1/articles",
			Method:  http.MethodGet,
			Handler: h.Find,
		},
		{
			Path:        "/api/v1/articles",
			Method:      http.MethodPost,
			Handler:     h.Create,
			Middlewares: []router.Middleware{auth.RequiredAuthMiddleware},
		},
		{
			Path:    "/api/v1/articles/{id:[a-z0-9-\\-]+}",
			Method:  http.MethodGet,
			Handler: h.Get,
		},
		{
			Path:        "/api/v1/articles/{id:[a-z0-9-\\-]+}",
			Method:      http.MethodDelete,
			Handler:     h.Delete,
			Middlewares: []router.Middleware{auth.RequiredAuthMiddleware},
		},
		{
			Path:    "/api/v1/articles/{id:[a-z0-9-\\-]+}",
			Method:  http.MethodPut,
			Handler: h.UpdateView,
			Queries: []string{"action", "update_view"},
		},
		{
			Path:        "/api/v1/articles/{id:[a-z0-9-\\-]+}",
			Method:      http.MethodPut,
			Handler:     h.Update,
			Middlewares: []router.Middleware{auth.RequiredAuthMiddleware},
		},
	}
}
