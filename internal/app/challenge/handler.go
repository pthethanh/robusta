package challenge

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/pthethanh/robusta/internal/app/types"
	"github.com/pthethanh/robusta/internal/pkg/http/respond"
	"github.com/pthethanh/robusta/internal/pkg/log"
	"github.com/pthethanh/robusta/internal/pkg/util/handlerutil"
)

type (
	service interface {
		Create(ctx context.Context, c *types.Challenge) error
		Get(ctx context.Context, id string) (*types.Challenge, error)
		Delete(ctx context.Context, id string) error
		FindAll(ctx context.Context, r FindRequest) ([]*types.Challenge, error)
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
		Tags:        queries["tags"],
		CreatedByID: queries.Get("created_by_id"),
		SortBy:      queries["sort_by"],
		IDs:         queries["ids"],
		FolderID:    queries.Get("folder_id"),
		Title:       queries.Get("title"),
		Selects:     queries["selects"],
	}
	challenges, err := h.srv.FindAll(r.Context(), req)
	if err != nil {
		respond.Error(w, err, http.StatusInternalServerError)
		return
	}
	for _, c := range challenges {
		c.Test = ""
	}
	respond.JSON(w, http.StatusOK, types.BaseResponse{
		Data: challenges,
	})
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	var c types.Challenge
	if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
		respond.Error(w, err, http.StatusBadRequest)
		return
	}
	defer r.Body.Close()
	if err := h.srv.Create(r.Context(), &c); err != nil {
		log.WithContext(r.Context()).Errorf("failed to create challenge, err: %v", err)
		respond.Error(w, err, http.StatusInternalServerError)
		return
	}
	respond.JSON(w, http.StatusOK, types.BaseResponse{
		Data: types.IDResponse{
			ID: c.ID,
		},
	})
}

// Get return full information of a challenge (including its test detail).
// This API requires read permission on the requested challenge.
func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	if id == "" {
		respond.Error(w, errors.New("invalid id"), http.StatusBadRequest)
		return
	}
	c, err := h.srv.Get(r.Context(), id)
	if err != nil {
		respond.Error(w, err, http.StatusInternalServerError)
		return
	}
	respond.JSON(w, http.StatusOK, types.BaseResponse{
		Data: c,
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

func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	if id == "" {
		respond.Error(w, errors.New("invalid id"), http.StatusBadRequest)
		return
	}
	var req UpdateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respond.Error(w, err, http.StatusBadRequest)
		return
	}
	defer r.Body.Close()
	req.ID = id
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
