package editor

import (
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/pthethanh/robusta/internal/app/status"
	"github.com/pthethanh/robusta/internal/pkg/log"

	"github.com/pthethanh/robusta/internal/pkg/http/respond"
)

type (
	service interface {
		FetchURL(ctx context.Context, url string) (*Link, error)
		UploadImageByURL(ctx context.Context, url string) (string, error)
		UploadImageByFile(ctx context.Context, name string, image io.Reader) (string, error)
	}

	Handler struct {
		srv  service
		conf Config
	}
)

// New return new handler with configuration from environment variables
func New(conf Config, srv service) *Handler {
	return &Handler{
		srv:  srv,
		conf: conf,
	}
}

// UploadImageByFile handle for uploading file supports for EditorJS - image plugin
func (h *Handler) UploadImageByFile(w http.ResponseWriter, r *http.Request) {
	r.Body = http.MaxBytesReader(w, r.Body, h.conf.MaxUploadSize)
	if err := r.ParseMultipartForm(h.conf.MaxUploadSize); err != nil {
		log.WithContext(r.Context()).Errorf("failed to parse request", err)
		respond.JSON(w, http.StatusOK, ImageToolResponse{
			Status:  status.Editor().FileSizeExceedLimit,
			Success: ImageToolStatusFailed,
		})
		return
	}
	defer r.Body.Close()
	file, handler, err := r.FormFile("image")
	if err != nil {
		log.WithContext(r.Context()).Errorf("failed to get form file")
		respond.JSON(w, http.StatusOK, ImageToolResponse{
			Status:  status.Success(),
			Success: ImageToolStatusFailed,
		})
		return
	}
	uploadStatus := ImageToolStatusSuccess
	uploadedPath, err := h.srv.UploadImageByFile(r.Context(), handler.Filename, file)
	if err != nil {
		uploadStatus = ImageToolStatusFailed
		log.WithContext(r.Context()).Errorf("failed to upload image %v, err: %v", handler.Filename, err)
	}
	respond.JSON(w, http.StatusOK, ImageToolResponse{
		Status:  status.Success(),
		Success: uploadStatus,
		File: ImageToolFile{
			URL: uploadedPath,
		},
	})
}

// FetchURL handle fetch URL information support for EditorJS - link plugin
func (h *Handler) FetchURL(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Query().Get("url")
	link, err := h.srv.FetchURL(r.Context(), url)
	if err != nil {
		respond.Error(w, err, http.StatusInternalServerError)
		return
	}
	respond.JSON(w, http.StatusOK, link)
}

// UploadImageByURL handle for uploading file supports for EditorJS - image plugin
func (h *Handler) UploadImageByURL(w http.ResponseWriter, r *http.Request) {
	req := struct {
		URL string `json:"url"`
	}{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.WithContext(r.Context()).Errorf("failed to decode request, err: %v", err)
		respond.JSON(w, http.StatusOK, ImageToolResponse{
			Status:  status.Success(),
			Success: ImageToolStatusFailed,
		})
		return
	}
	if req.URL == "" {
		log.WithContext(r.Context()).Errorf("url is missing")
		respond.JSON(w, http.StatusOK, ImageToolResponse{
			Status:  status.Success(),
			Success: ImageToolStatusFailed,
		})
		return
	}
	uploadedURL, err := h.srv.UploadImageByURL(r.Context(), req.URL)
	uploadStatus := ImageToolStatusSuccess
	if err != nil {
		uploadStatus = ImageToolStatusFailed
		log.WithContext(r.Context()).Errorf("failed to upload image %v, err: %v", r.URL, err)
	}
	respond.JSON(w, http.StatusOK, ImageToolResponse{
		Status:  status.Success(),
		Success: uploadStatus,
		File: ImageToolFile{
			URL: uploadedURL,
		},
	})

}
