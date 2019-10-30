package article

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/pthethanh/robusta/internal/app/auth"
	"github.com/pthethanh/robusta/internal/app/status"
	"github.com/pthethanh/robusta/internal/app/types"
	"github.com/pthethanh/robusta/internal/pkg/http/respond"
	"github.com/pthethanh/robusta/internal/pkg/log"
	"github.com/pthethanh/robusta/internal/pkg/util/handlerutil"
)

type (
	service interface {
		FindAll(ctx context.Context, req FindRequest) ([]*Article, error)
		IncreaseView(ctx context.Context, id string) error
		Create(ctx context.Context, a *Article) error
		FindByID(ctx context.Context, id string) (*Article, error)
		FindByArticleID(ctx context.Context, id string) (*Article, error)
		Update(ctx context.Context, id string, a *Article) error
		ChangeStatus(ctx context.Context, id string, status Status) error
	}
	// Handler is friend web handler
	Handler struct {
		srv service
	}
)

const (
	previewMode = "preview"
)

// NewHTTPHandler return new rest api friend handler
func NewHTTPHandler(s service) *Handler {
	return &Handler{
		srv: s,
	}
}

// Find find the articles base on the given query.
// Supported queries: offset, limit, tags, created_by_id, sort_by
// Mode: preview, full if empty
// Strategy: cache -> for fetching result from cache.
func (h *Handler) Find(w http.ResponseWriter, r *http.Request) {
	queries := r.URL.Query()
	req := FindRequest{
		Offset:      handlerutil.IntFromQuery("offset", queries, 0),
		Limit:       handlerutil.IntFromQuery("limit", queries, 15),
		Status:      StatusPublic, // query only public articles
		Tags:        queries["tags"],
		CreatedByID: queries.Get("created_by_id"),
		SortBy:      queries["sort_by"],
	}
	mode := queries.Get("mode")
	// fetch the list of articles from db
	list, err := h.srv.FindAll(r.Context(), req)
	if err != nil {
		log.WithContext(r.Context()).Infof("failed to execute FinAll, err: %v", err)
		respond.Error(w, err, http.StatusInternalServerError)
		return
	}
	if mode == previewMode {
		list = convertArticlesToReviewMode(list)
	}
	respond.JSON(w, http.StatusOK, types.BaseResponse{
		Data: list,
	})
}

func (h *Handler) UpdateView(w http.ResponseWriter, r *http.Request) {
	err := h.srv.IncreaseView(r.Context(), mux.Vars(r)["id"])
	if err != nil {
		log.WithContext(r.Context()).Infof("failed to execute View, err: %v", err)
		respond.Error(w, err, http.StatusInternalServerError)
		return
	}
	respond.JSON(w, http.StatusOK, types.BaseResponse{})
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	var a Article
	if err := json.NewDecoder(r.Body).Decode(&a); err != nil {
		log.WithContext(r.Context()).Infof("failed to decode body, err: %v", err)
		respond.Error(w, err, http.StatusBadRequest)
		return
	}
	defer r.Body.Close()
	if err := h.srv.Create(r.Context(), &a); err != nil {
		log.WithContext(r.Context()).Errorf("could not create article, err: %v", err)
		respond.Error(w, err, http.StatusInternalServerError)
		return
	}
	respond.JSON(w, http.StatusOK, types.BaseResponse{
		Data: a,
	})
}

func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	if id == "" {
		log.WithContext(r.Context()).Infof("id is not valid")
		respond.Error(w, fmt.Errorf("id is not valid"), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()
	var find func(ctx context.Context, id string) (*Article, error)
	find = h.srv.FindByID
	if isArticleID(id) {
		find = h.srv.FindByArticleID
	}
	a, err := find(r.Context(), id)
	if err != nil {
		log.WithContext(r.Context()).Errorf("could not found the article, err: %v", err)
		respond.Error(w, err, http.StatusInternalServerError)
		return
	}
	if a.Status == StatusPublic {
		respond.JSON(w, http.StatusOK, types.BaseResponse{
			Data: a,
		})
		return
	}
	// if it's the owner of the article return a result event if it's a draft
	user := auth.FromContext(r.Context())
	if a.Status == StatusDraft && user != nil && user.UserID == a.CreatedByID {
		respond.JSON(w, http.StatusOK, types.BaseResponse{
			Data: a,
		})
		return
	}
	// don't allow to find by ID for deleted article
	respond.JSON(w, http.StatusOK, types.BaseResponse{
		Status: status.Gen().NotFound,
		Data:   a,
	})
	return
}

func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	if id == "" {
		log.WithContext(r.Context()).Infof("id is not valid")
		respond.Error(w, fmt.Errorf("id is not valid"), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()
	// go ahead to update
	var req Article
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respond.Error(w, err, http.StatusBadRequest)
		return
	}
	if err := h.srv.Update(r.Context(), id, &req); err != nil {
		respond.Error(w, err, http.StatusInternalServerError)
		return
	}
	respond.JSON(w, http.StatusOK, types.BaseResponse{
		Data: types.IDResponse{
			ID: id,
		},
	})
}

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	if id == "" {
		log.WithContext(r.Context()).Infof("id is not valid")
		respond.Error(w, fmt.Errorf("id is not valid"), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()
	if err := h.srv.ChangeStatus(r.Context(), id, StatusDeleted); err != nil {
		respond.Error(w, err, http.StatusInternalServerError)
		return
	}
	respond.JSON(w, http.StatusOK, types.BaseResponse{
		Data: types.IDResponse{
			ID: id,
		},
	})
}

// convertArticlesToReviewMode convert the given articles to preview mode
// by stripping the content down to 2 blocks for a lighter package and
// better performance when sending the articles in the network
func convertArticlesToReviewMode(articles []*Article) []*Article {
	rs := make([]*Article, 0)
	for _, a := range articles {
		v := *a
		if len(v.Content.Blocks) > 2 {
			v.Content.Blocks = v.Content.Blocks[0:2] // max 2 blocks...
		}
		rs = append(rs, &v)
	}
	return rs
}
