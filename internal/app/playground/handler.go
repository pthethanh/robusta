package playground

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/pthethanh/robusta/internal/app/types"
	"github.com/pthethanh/robusta/internal/pkg/http/respond"

	"github.com/pkg/errors"
)

type (
	service interface {
		Run(ctx context.Context, r *types.PlaygroundRequest) (*types.PlaygroundResponse, error)
	}

	Handler struct {
		srv service
	}
)

func New(s service) *Handler {
	return &Handler{
		srv: s,
	}
}

func (h *Handler) Run(w http.ResponseWriter, r *http.Request) {
	var req types.PlaygroundRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respond.Error(w, errors.Wrap(err, "invalid request"), http.StatusBadRequest)
		return
	}
	res, err := h.srv.Run(r.Context(), &req)
	if err != nil {
		respond.Error(w, errors.Wrap(err, "failed to run"), http.StatusInternalServerError)
		return
	}
	res.Code = types.AppCodeSuccess
	respond.JSON(w, http.StatusOK, res)
}
