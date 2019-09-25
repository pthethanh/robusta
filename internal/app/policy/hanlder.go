package policy

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/pthethanh/robusta/internal/app/types"
	"github.com/pthethanh/robusta/internal/pkg/http/respond"
)

type (
	service interface {
		AssignPolicy(ctx context.Context, req AssignPolicyRequest) error
		AssignGroupPolicy(ctx context.Context, req AssignGroupPolicyRequest) error
		ListActions(ctx context.Context) ([]string, error)
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

func (h *Handler) AssignPolicy(w http.ResponseWriter, r *http.Request) {
	var req AssignPolicyRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respond.Error(w, err, http.StatusBadRequest)
		return
	}
	if err := h.srv.AssignPolicy(r.Context(), req); err != nil {
		respond.Error(w, err, http.StatusInternalServerError)
		return
	}
	respond.JSON(w, http.StatusOK, types.BaseResponse{})
}

func (h *Handler) AssignGroupPolicy(w http.ResponseWriter, r *http.Request) {
	var req AssignGroupPolicyRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respond.Error(w, err, http.StatusBadRequest)
		return
	}
	if err := h.srv.AssignGroupPolicy(r.Context(), req); err != nil {
		respond.Error(w, err, http.StatusInternalServerError)
		return
	}
	respond.JSON(w, http.StatusOK, types.BaseResponse{})
}

func (h *Handler) ListActions(w http.ResponseWriter, r *http.Request) {
	actions, err := h.srv.ListActions(r.Context())
	if err != nil {
		respond.Error(w, err, http.StatusInternalServerError)
		return
	}
	respond.JSON(w, http.StatusOK, types.BaseResponse{
		Data: actions,
	})
}
