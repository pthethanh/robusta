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
		AssignPolicy(ctx context.Context, req AssignPolicyRequest) error
		AssignGroupPolicy(ctx context.Context, req AssignGroupPolicyRequest) error
		ListActions(ctx context.Context) ([]string, error)
		GetRoles(ctx context.Context) ([]string, error)
		GetUsersForRole(ctx context.Context, role string) ([]string, error)
		FindPolicies(ctx context.Context, req FindPolicyRequest) ([]types.Policy, error)
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
	})
	if err != nil {
		respond.Error(w, err, http.StatusInternalServerError)
		return
	}
	respond.JSON(w, http.StatusOK, types.BaseResponse{
		Data: policies,
	})
}
