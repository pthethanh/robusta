package articlehandler

import (
	"context"
	"net/http"

	"github.com/pthethanh/robusta/internal/app/types"
	"github.com/pthethanh/robusta/internal/pkg/glog"
	"github.com/pthethanh/robusta/internal/pkg/respond"
)

type (
	service interface {
		FindAll(ctx context.Context, offset, limit int) ([]*types.Article, error)
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

func (h *Handler) List(w http.ResponseWriter, r *http.Request) {
	list, err := h.srv.FindAll(r.Context(), 0, 100)
	if err != nil {
		respond.Error(w, err, http.StatusInternalServerError)
		return
	}
	respond.JSON(w, http.StatusOK, map[string]interface{}{
		"code":    20000,
		"message": "success",
		"items":   list,
	})
}
