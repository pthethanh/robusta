package solution

import (
	"context"
	"net/http"

	"github.com/pthethanh/robusta/internal/app/types"
	"github.com/pthethanh/robusta/internal/pkg/http/respond"
	"github.com/pthethanh/robusta/internal/pkg/util/handlerutil"
)

type (
	service interface {
		FindAll(ctx context.Context, req FindRequest) ([]*types.Solution, error)
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
		ChallengeID: queries.Get("challenge_id"),
		CreatedByID: queries.Get("created_by_id"),
		SortBy:      queries["sort_by"],
	}
	maxLimit := 15
	if req.Limit > maxLimit {
		req.Limit = maxLimit
	}
	solutions, err := h.srv.FindAll(r.Context(), req)
	if err != nil {
		respond.Error(w, err, http.StatusInternalServerError)
		return
	}
	respond.JSON(w, http.StatusOK, types.BaseResponse{
		Data: solutions,
	})
}
