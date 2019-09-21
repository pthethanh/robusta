package comment

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pthethanh/robusta/internal/app/types"
	"github.com/pthethanh/robusta/internal/pkg/http/respond"
	"github.com/pthethanh/robusta/internal/pkg/log"
	"github.com/pthethanh/robusta/internal/pkg/util/handlerutil"
)

type (
	service interface {
		Create(ctx context.Context, c *types.Comment) error
		FindAll(ctx context.Context, req FindRequest) ([]*types.Comment, error)
		Update(ctx context.Context, id string, c *types.Comment) error
		Delete(ctx context.Context, id string) error
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

func (h *Handler) Find(w http.ResponseWriter, r *http.Request) {
	queries := r.URL.Query()
	req := FindRequest{
		Offset:      handlerutil.IntFromQuery("offset", queries, 0),
		Limit:       handlerutil.IntFromQuery("limit", queries, 15),
		Target:      queries.Get("target"),
		ReplyToID:   queries.Get("reply_to_id"),
		ThreadID:    queries.Get("thread_id"),
		CreatedByID: queries.Get("created_by_id"),
		SortBy:      queries["sort_by"],
	}
	list, err := h.srv.FindAll(r.Context(), req)
	if err != nil {
		log.WithContext(r.Context()).Infof("failed to execute FinAll, err: %v", err)
		respond.Error(w, err, http.StatusInternalServerError)
		return
	}
	respond.JSON(w, http.StatusOK, types.BaseResponse{
		Data: list,
	})
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	var c types.Comment
	if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
		log.WithContext(r.Context()).Infof("failed to decode body, err: %v", err)
		respond.Error(w, err, http.StatusBadRequest)
		return
	}
	defer r.Body.Close()
	if err := h.srv.Create(r.Context(), &c); err != nil {
		log.WithContext(r.Context()).Errorf("could not create comment, err: %v", err)
		respond.Error(w, err, http.StatusInternalServerError)
		return
	}
	respond.JSON(w, http.StatusOK, types.BaseResponse{
		Data: c,
	})
}

func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	if id == "" {
		log.WithContext(r.Context()).Infof("id is not valid")
		respond.Error(w, fmt.Errorf("id is not valid"), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()
	// go ahead to update
	var req types.Comment
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respond.Error(w, err, http.StatusBadRequest)
		return
	}
	if err := h.srv.Update(r.Context(), id, &req); err != nil {
		respond.Error(w, err, http.StatusInternalServerError)
		return
	}
	respond.JSON(w, http.StatusOK, types.BaseResponse{
		Data: types.IDResponse{
			ID: id,
		},
	})
}

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	if id == "" {
		log.WithContext(r.Context()).Infof("id is not valid")
		respond.Error(w, fmt.Errorf("id is not valid"), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()
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
