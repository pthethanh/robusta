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
		FindSolutionInfo(ctx context.Context, req FindRequest) ([]SolutionInfo, error)
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

func (h *Handler) FindSolutionInfo(w http.ResponseWriter, r *http.Request) {
	queries := r.URL.Query()
	req := FindRequest{
		Offset:      handlerutil.IntFromQuery("offset", queries, 0),
		Limit:       handlerutil.IntFromQuery("limit", queries, 15),
		ChallengeID: queries.Get("challenge_id"),
		CreatedByID: queries.Get("created_by_id"),
		SortBy:      queries["sort_by"],
	}
	solutions, err := h.srv.FindSolutionInfo(r.Context(), req)
	if err != nil {
		respond.Error(w, err, http.StatusInternalServerError)
		return
	}
	respond.JSON(w, http.StatusOK, types.BaseResponse{
		Data: solutions,
	})
}
