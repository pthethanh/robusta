package middleware

import (
	"net/http"
	"strings"
	"time"

	"github.com/pthethanh/robusta/internal/pkg/log"
)

// HTTPRequestResponseInfo is a handler that log request/response information
func HTTPRequestResponseInfo(ignorePrefixes []string) func(inner http.Handler) http.Handler {
	return func(inner http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			for _, ignore := range ignorePrefixes {
				if strings.HasPrefix(r.URL.Path, ignore) {
					inner.ServeHTTP(w, r)
					return
				}
			}
			bg := time.Now()
			log.WithContext(r.Context()).WithField("stage", "request").Info("request started")

			inner.ServeHTTP(w, r)

			code := http.StatusOK
			if mw, ok := w.(interface {
				Status() int
			}); ok {
				code = mw.Status()
			}
			responseTime := time.Since(bg)
			log.WithContext(r.Context()).WithFields(log.Fields{
				"status":        code,
				"response_time": responseTime,
				"stage":         "response",
			}).Info("request finished")
		})
	}
}
