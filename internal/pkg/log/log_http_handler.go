package log

import (
	"net/http"

	"github.com/pthethanh/robusta/internal/pkg/uuid"
)

// NewHTTPContextHandler adds a context logger based on the given logger to
// each request. After a request passes through this handler,
// WithContext(req.Context()).Error(, "foo") will log to that logger and add useful context
// to each log entry.
func NewHTTPContextHandler(l Logger) func(http.Handler) http.Handler {
	return func(inner http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			logger := l.WithFields(Fields{
				"request_id":  uuid.New(),
				"path":        r.URL.Path,
				"remote_addr": r.RemoteAddr,
				"method":      r.Method,
			})
			r = r.WithContext(NewContext(ctx, logger))
			inner.ServeHTTP(w, r)
		})
	}
}
