package playground

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pthethanh/robusta/internal/app/status"
	"github.com/pthethanh/robusta/internal/app/types"
	"github.com/pthethanh/robusta/internal/pkg/http/respond"
	"github.com/pthethanh/robusta/internal/pkg/log"
	"github.com/pthethanh/robusta/internal/pkg/playground"
)

type (
	service interface {
		Run(ctx context.Context, r *Request) (*Response, error)
		Evaluate(ctx context.Context, r *EvaluateRequest) (*playground.EvaluateResponse, error)
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
	var req Request
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respond.Error(w, fmt.Errorf("invalid request: %w", err), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()
	res, err := h.srv.Run(r.Context(), &req)
	if err != nil {
		respond.Error(w, fmt.Errorf("failed to run: %w", err), http.StatusInternalServerError)
		return
	}
	res.Code = int(status.Success().Code())
	respond.JSON(w, http.StatusOK, res)
}

func (h *Handler) Evaluate(w http.ResponseWriter, r *http.Request) {
	var req EvaluateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respond.Error(w, err, http.StatusBadRequest)
		return
	}
	defer r.Body.Close()
	res, err := h.srv.Evaluate(r.Context(), &req)
	if err != nil {
		log.WithContext(r.Context()).Errorf("evaluate failed, err: %v", err)
		respond.Error(w, err, http.StatusInternalServerError)
		return
	}
	respond.JSON(w, http.StatusOK, types.BaseResponse{
		Data: res,
	})
}
