package middleware

import (
	"net/http"
	"time"

	"github.com/pthethanh/robusta/internal/pkg/glog"
)

// Logging is a handler that log request information
func Logging(logger glog.Logger) func(inner http.Handler) http.Handler {
	return func(inner http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			bg := time.Now()
			logger.WithField("method", r.Method).WithField("path", r.URL.Path).
				WithField("stage", "started").
				Infoc(r.Context(), "")
			inner.ServeHTTP(w, r)
			code := http.StatusOK
			if mw, ok := w.(interface {
				Status() int
			}); ok {
				code = mw.Status()
			}
			logger.WithField("method", r.Method).WithField("path", r.URL.Path).
				WithField("time_spent", time.Since(bg)).
				WithField("status", code).
				WithField("stage", "finished").
				Infoc(r.Context(), "")
		})
	}
}
