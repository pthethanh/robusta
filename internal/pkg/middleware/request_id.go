package middleware

import (
	"net/http"

	"github.com/pthethanh/robusta/internal/pkg/kontext"
)

// RequestID is a http.Handler that attach a UUID into the context of the request
// as a request id for tracing
func RequestID(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := kontext.NewRequestIDContext(r.Context())
		r = r.WithContext(ctx)
		h.ServeHTTP(w, r)
	})
}
