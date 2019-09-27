package solution

import (
	"context"
	"net/http"

	"github.com/pthethanh/robusta/internal/app/types"
	"github.com/pthethanh/robusta/internal/pkg/http/respond"
	"github.com/pthethanh/robusta/internal/pkg/log"
	"github.com/pthethanh/robusta/internal/pkg/util/handlerutil"
)

type (
	service interface {
		FindSolutionInfo(ctx context.Context, req FindRequest) ([]SolutionInfo, error)
		GetCompletionReport(ctx context.Context, r CompletionReportRequest) ([]SolutionInfo, error)
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
		Offset:        handlerutil.IntFromQuery("offset", queries, 0),
		Limit:         handlerutil.IntFromQuery("limit", queries, 15),
		ChallengeIDs:  queries["challenge_ids"],
		CreatedByID:   queries.Get("created_by_id"),
		SortBy:        queries["sort_by"],
		Status:        queries.Get("status"),
		IncludeDetail: queries.Get("include_detail") == "true",
	}
	solutions, err := h.srv.FindSolutionInfo(r.Context(), req)
	if err != nil {
		log.WithContext(r.Context()).Errorf("failed to find solution info, err: %v", err)
		respond.Error(w, err, http.StatusInternalServerError)
		return
	}
	respond.JSON(w, http.StatusOK, types.BaseResponse{
		Data: solutions,
	})
}

func (h *Handler) GetCompletionReport(w http.ResponseWriter, r *http.Request) {
	queries := r.URL.Query()
	req := CompletionReportRequest{
		ChallengeIDs:  queries["challenge_ids"],
		CreatedByID:   queries.Get("created_by_id"),
		IncludeDetail: queries.Get("include_detail") == "true",
		Status:        queries.Get("status"),
	}
	solutions, err := h.srv.GetCompletionReport(r.Context(), req)
	if err != nil {
		log.WithContext(r.Context()).Errorf("failed to get completion report, err: %v", err)
		respond.Error(w, err, http.StatusInternalServerError)
		return
	}
	respond.JSON(w, http.StatusOK, types.BaseResponse{
		Data: solutions,
	})
}
