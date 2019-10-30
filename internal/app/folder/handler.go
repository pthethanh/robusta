package folder

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pthethanh/robusta/internal/app/status"
	"github.com/pthethanh/robusta/internal/app/types"
	"github.com/pthethanh/robusta/internal/pkg/http/respond"
	"github.com/pthethanh/robusta/internal/pkg/log"
	"github.com/pthethanh/robusta/internal/pkg/util/handlerutil"
)

type (
	service interface {
		Create(ctx context.Context, f *CreateRequest) error
		Get(ctx context.Context, id string) (*Folder, error)
		Delete(ctx context.Context, id string) error
		FindAll(ctx context.Context, r FindRequest) ([]*Folder, error)
		AddChildren(ctx context.Context, req AddChildrenRequest) error
		Update(ctx context.Context, req UpdateRequest) error
	}
	Handler struct {
		srv service
	}
)

func NewHandler(srv service) *Handler {
	return &Handler{
		srv: srv,
	}
}

func (h *Handler) FindAll(w http.ResponseWriter, r *http.Request) {
	queries := r.URL.Query()
	req := FindRequest{
		Offset:      handlerutil.IntFromQuery("offset", queries, 0),
		Limit:       handlerutil.IntFromQuery("limit", queries, 15),
		Type:        Type(queries.Get("type")),
		CreatedByID: queries.Get("created_by_id"),
		SortBy:      queries["sort_by"],
	}
	folders, err := h.srv.FindAll(r.Context(), req)
	if err != nil {
		respond.Error(w, err, http.StatusInternalServerError)
		return
	}
	respond.JSON(w, http.StatusOK, types.BaseResponse{
		Data: folders,
	})
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	var f CreateRequest
	if err := json.NewDecoder(r.Body).Decode(&f); err != nil {
		respond.Error(w, err, http.StatusBadRequest)
		return
	}
	defer r.Body.Close()
	if err := h.srv.Create(r.Context(), &f); err != nil {
		log.WithContext(r.Context()).Errorf("failed to create folder, err: %v", err)
		respond.Error(w, err, http.StatusInternalServerError)
		return
	}
	respond.JSON(w, http.StatusOK, types.BaseResponse{
		Data: types.IDResponse{
			ID: f.ID,
		},
	})
}

func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	if id == "" {
		respond.Error(w, errors.New("invalid id"), http.StatusBadRequest)
		return
	}
	f, err := h.srv.Get(r.Context(), id)
	if err != nil {
		respond.Error(w, err, http.StatusInternalServerError)
		return
	}
	respond.JSON(w, http.StatusOK, types.BaseResponse{
		Data: f,
	})
}

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	if id == "" {
		respond.Error(w, errors.New("invalid id"), http.StatusBadRequest)
		return
	}
	if err := h.srv.Delete(r.Context(), id); err != nil {
		respond.Error(w, err, http.StatusInternalServerError)
		return
	}
	respond.JSON(w, http.StatusOK, types.BaseResponse{
		Data: types.IDResponse{
			ID: id,
		},
	})
}

func (h *Handler) AddChildren(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	if id == "" {
		respond.Error(w, errors.New("invalid id"), http.StatusBadRequest)
		return
	}
	var req AddChildrenRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respond.Error(w, status.Gen().BadRequest, http.StatusBadRequest)
		return
	}
	defer r.Body.Close()
	if err := h.srv.AddChildren(r.Context(), req); err != nil {
		respond.Error(w, err, http.StatusInternalServerError)
		return
	}
	respond.JSON(w, http.StatusOK, types.BaseResponse{
		Data: types.IDResponse{
			ID: id,
		},
	})
}

func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	if id == "" {
		respond.Error(w, errors.New("invalid id"), http.StatusBadRequest)
		return
	}
	var req UpdateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respond.Error(w, status.Gen().BadRequest, http.StatusBadRequest)
		return
	}
	defer r.Body.Close()
	if err := h.srv.Update(r.Context(), req); err != nil {
		respond.Error(w, err, http.StatusInternalServerError)
		return
	}
	respond.JSON(w, http.StatusOK, types.BaseResponse{
		Data: types.IDResponse{
			ID: id,
		},
	})
}
