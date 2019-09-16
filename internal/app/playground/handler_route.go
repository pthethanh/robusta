package playground

import (
	"net/http"

	"github.com/pthethanh/robusta/internal/pkg/http/router"
)

func (h *Handler) Routes() []router.Route {
	return []router.Route{
		{
			Path:    "/api/v1/playground/run",
			Method:  http.MethodPost,
			Handler: h.Run,
		},
		{
			Path:    "/api/v1/playground/evaluate",
			Method:  http.MethodPost,
			Handler: h.Evaluate,
		},
	}
}
