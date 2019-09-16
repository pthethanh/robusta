package challenge

import (
	"net/http"

	"github.com/pthethanh/robusta/internal/pkg/http/router"
)

func (h *Handler) Routes() []router.Route {
	return []router.Route{
		{
			Path:    "/api/v1/challenges",
			Method:  http.MethodGet,
			Handler: h.FindAll,
		},
		{
			Path:    "/api/v1/challenges",
			Method:  http.MethodPost,
			Handler: h.Create,
		},
		{
			Path:    "/api/v1/challenges/{id:[a-z0-9-\\-]+}",
			Method:  http.MethodGet,
			Handler: h.Get,
		},
	}
}
