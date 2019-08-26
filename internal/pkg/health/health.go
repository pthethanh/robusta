package health

import (
	"context"
	"errors"
	"net/http"
	"sync"

	"github.com/pthethanh/robusta/internal/pkg/http/respond"
)

var (
	isReadyMu sync.RWMutex
	isReady   bool
)

// CheckFunc function signature for health checks.
type CheckFunc func(context.Context) error

// Ready marks the service as ready to receive traffic.
func Ready() {
	isReadyMu.Lock()
	isReady = true
	isReadyMu.Unlock()
}

// Readiness returns an HTTP handler for checking Readiness state.
// Will return 503 untill Ready() is called
func Readiness() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		isReadyMu.RLock()
		defer isReadyMu.RUnlock()
		if !isReady {
			respond.Error(w, errors.New("not ready"), http.StatusServiceUnavailable)
			return
		}
		respond.JSON(w, http.StatusOK, "OK")
	})
}

// Liveness returns an HTTP handler for checking the health of the service.
func Liveness(cf ...CheckFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		for _, c := range cf {
			if err := c(ctx); err != nil {
				respond.Error(w, err, http.StatusServiceUnavailable)
				return
			}
		}
		respond.JSON(w, http.StatusOK, "OK")
	})
}
