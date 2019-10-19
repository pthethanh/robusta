package policy

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/pthethanh/robusta/internal/app/types"
	"github.com/pthethanh/robusta/internal/pkg/http/respond"
)

type (
	service interface {
		AddPolicy(ctx context.Context, req types.Policy) error
		AddGroupPolicy(ctx context.Context, req GroupPolicy) error
		ListActions(ctx context.Context) ([]string, error)
		GetRoles(ctx context.Context) ([]string, error)
		GetUsersForRole(ctx context.Context, role string) ([]string, error)
		FindPolicies(ctx context.Context, req FindPolicyRequest) ([]types.Policy, error)
		RemovePolicy(ctx context.Context, req types.Policy) error
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

func (h *Handler) AddPolicy(w http.ResponseWriter, r *http.Request) {
	var req types.Policy
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respond.Error(w, err, http.StatusBadRequest)
		return
	}
	if err := h.srv.AddPolicy(r.Context(), req); err != nil {
		respond.Error(w, err, http.StatusInternalServerError)
		return
	}
	respond.JSON(w, http.StatusOK, types.BaseResponse{})
}

func (h *Handler) AddGroupPolicy(w http.ResponseWriter, r *http.Request) {
	var req GroupPolicy
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respond.Error(w, err, http.StatusBadRequest)
		return
	}
	if err := h.srv.AddGroupPolicy(r.Context(), req); err != nil {
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

func (h *Handler) GetRoles(w http.ResponseWriter, r *http.Request) {
	roles, err := h.srv.GetRoles(r.Context())
	if err != nil {
		respond.Error(w, err, http.StatusInternalServerError)
		return
	}
	respond.JSON(w, http.StatusOK, types.BaseResponse{
		Data: roles,
	})
}

func (h *Handler) GetUsersForRole(w http.ResponseWriter, r *http.Request) {
	role := mux.Vars(r)["role"]
	users, err := h.srv.GetUsersForRole(r.Context(), role)
	if err != nil {
		respond.Error(w, err, http.StatusInternalServerError)
		return
	}
	respond.JSON(w, http.StatusOK, types.BaseResponse{
		Data: users,
	})
}

func (h *Handler) FindPolicies(w http.ResponseWriter, r *http.Request) {
	queries := r.URL.Query()
	policies, err := h.srv.FindPolicies(r.Context(), FindPolicyRequest{
		Subjects: queries["subjects"],
		Actions:  queries["actions"],
		Objects:  queries["objects"],
	})
	if err != nil {
		respond.Error(w, err, http.StatusInternalServerError)
		return
	}
	respond.JSON(w, http.StatusOK, types.BaseResponse{
		Data: policies,
	})
}

func (h *Handler) RemovePolicy(w http.ResponseWriter, r *http.Request) {
	var req types.Policy
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respond.Error(w, err, http.StatusBadRequest)
		return
	}
	if err := h.srv.RemovePolicy(r.Context(), req); err != nil {
		respond.Error(w, err, http.StatusInternalServerError)
		return
	}
	respond.JSON(w, http.StatusOK, types.BaseResponse{})
}
