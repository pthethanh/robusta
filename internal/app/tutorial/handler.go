package tutorial

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/pthethanh/robusta/internal/app/types"
	"github.com/pthethanh/robusta/internal/pkg/http/respond"
	"github.com/pthethanh/robusta/internal/pkg/log"
)

type (
	service interface {
		FindAll(ctx context.Context, offset, limit int) ([]*types.Tutorial, error)
		View(ctx context.Context, id string) error
		Create(ctx context.Context, a *types.Tutorial) error
		FindByID(ctx context.Context, id string) (*types.Tutorial, error)
		Update(ctx context.Context, id string, a *types.Tutorial) error
		Delete(ctx context.Context, id string) error
	}
	// Handler is friend web handler
	Handler struct {
		srv service
	}
)

// NewHTTPHandler return new rest api friend handler
func NewHTTPHandler(s service) *Handler {
	return &Handler{
		srv: s,
	}
}

func (h *Handler) List(w http.ResponseWriter, r *http.Request) {
	list, err := h.srv.FindAll(r.Context(), 0, 100)
	if err != nil {
		log.WithContext(r.Context()).Infof("failed to execute FinAll, err: %v", err)
		respond.Error(w, err, http.StatusInternalServerError)
		return
	}
	respond.JSON(w, http.StatusOK, types.BaseResponse{
		Data: list,
	})
}

func (h *Handler) View(w http.ResponseWriter, r *http.Request) {
	err := h.srv.View(r.Context(), mux.Vars(r)["id"])
	if err != nil {
		log.WithContext(r.Context()).Infof("failed to execute View, err: %v", err)
		respond.Error(w, err, http.StatusInternalServerError)
		return
	}
	respond.JSON(w, http.StatusOK, types.BaseResponse{})
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	var a types.Tutorial
	if err := json.NewDecoder(r.Body).Decode(&a); err != nil {
		log.WithContext(r.Context()).Infof("failed to decode body, err: %v", err)
		respond.Error(w, err, http.StatusBadRequest)
		return
	}
	defer r.Body.Close()
	if err := h.srv.Create(r.Context(), &a); err != nil {
		log.WithContext(r.Context()).Errorf("could not create tutorial, err: %v", err)
		respond.Error(w, err, http.StatusInternalServerError)
		return
	}
	respond.JSON(w, http.StatusOK, types.BaseResponse{
		Data: types.IDResponse{
			ID: a.ID,
		},
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
	a, err := h.srv.FindByID(r.Context(), id)
	if err != nil {
		log.WithContext(r.Context()).Errorf("could not found the tutorial, err: %v", err)
		respond.Error(w, err, http.StatusInternalServerError)
		return
	}
	respond.JSON(w, http.StatusOK, types.BaseResponse{
		Data: a,
	})
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
	var req types.Tutorial
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
	defer r.Body.Close()
	if err := h.srv.Delete(r.Context(), id); err != nil {
		respond.Error(w, err, http.StatusInternalServerError)
		return
	}
	respond.JSON(w, http.StatusOK, types.BaseResponse{
		Data: types.IDResponse{
			ID: id,
		},
	})
}
