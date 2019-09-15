package solution

import (
	"net/http"

	"github.com/pthethanh/robusta/internal/pkg/http/router"
)

func (h *Handler) Routes() []router.Route {
	return []router.Route{
		{
			Path:    "/api/v1/solutions",
			Method:  http.MethodGet,
			Handler: h.FindAll,
		},
	}
}
