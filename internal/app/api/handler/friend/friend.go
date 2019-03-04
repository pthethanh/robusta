package friendhandler

import (
	"context"
	"net/http"

	"github.com/pthethanh/robusta/internal/app/types"
	"github.com/pthethanh/robusta/internal/pkg/glog"
	"github.com/pthethanh/robusta/internal/pkg/respond"

	"github.com/gorilla/mux"
)

type (
	service interface {
		Get(ctx context.Context, id string) (*types.Friend, error)
	}

	// Handler is friend web handler
	Handler struct {
		srv    service
		logger glog.Logger
	}
)

// New return new rest api friend handler
func New(s service, l glog.Logger) *Handler {
	return &Handler{
		srv:    s,
		logger: l,
	}
}

// Get handle get friend HTTP request
func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
	friend, err := h.srv.Get(r.Context(), mux.Vars(r)["id"])
	if err != nil {
		respond.Error(w, err, http.StatusInternalServerError)
		return
	}
	respond.JSON(w, http.StatusOK, friend)
}
