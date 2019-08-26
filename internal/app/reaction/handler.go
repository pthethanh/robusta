package reaction

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/pthethanh/robusta/internal/app/types"
	"github.com/pthethanh/robusta/internal/pkg/http/respond"
	"github.com/pthethanh/robusta/internal/pkg/log"
)

type (
	service interface {
		Create(ctx context.Context, reaction *types.Reaction) (*types.ReactionDetail, error)
		Find(ctx context.Context, queries string) (types.Reactions, error)
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

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	var req types.Reaction
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respond.Error(w, err, http.StatusBadRequest)
		return
	}
	defer r.Body.Close()
	reactions, err := h.srv.Create(r.Context(), &req)
	if err != nil {
		log.WithContext(r.Context()).Errorf("failed to create reaction, err: %v", err)
		respond.Error(w, err, http.StatusInternalServerError)
		return
	}
	respond.JSON(w, http.StatusCreated, types.BaseResponse{
		Data: reactions,
	})
}

func (h *Handler) Find(w http.ResponseWriter, r *http.Request) {
	reactions, err := h.srv.Find(r.Context(), r.URL.RawQuery)
	if err != nil {
		log.WithContext(r.Context()).Errorf("failed to create reaction, err: %v", err)
		respond.Error(w, err, http.StatusInternalServerError)
		return
	}
	respond.JSON(w, http.StatusCreated, types.BaseResponse{
		Data: reactions,
	})
}
