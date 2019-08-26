package editor

import (
	"net/http"

	"github.com/pthethanh/robusta/internal/app/auth"
	"github.com/pthethanh/robusta/internal/pkg/http/router"
)

func (h *Handler) Routes() []router.Route {
	return []router.Route{
		{
			Path:        "/api/v1/editor/image_by_file",
			Method:      http.MethodPost,
			Handler:     h.UploadImageByFile,
			Middlewares: []router.Middleware{auth.RequiredAuthMiddleware},
		},
		{
			Path:        "/api/v1/editor/image_by_url",
			Method:      http.MethodPost,
			Handler:     h.UploadImageByURL,
			Middlewares: []router.Middleware{auth.RequiredAuthMiddleware},
		},
		{
			Path:    "/api/v1/editor/fetch-url",
			Method:  http.MethodGet,
			Handler: h.FetchURL,
		},
	}
}
