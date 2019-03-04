package middleware

import (
	"context"
	"net/http"

	"github.com/pthethanh/robusta/internal/pkg/uuid"
)

func RequestID(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), "request_id", uuid.New())
		r = r.WithContext(ctx)
		h.ServeHTTP(w, r)
	})
}
