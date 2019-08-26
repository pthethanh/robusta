package middleware

import (
	"fmt"
	"net/http"
)

func StaticCache(h http.Handler, maxAge int) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Cache-Control", fmt.Sprintf("public, max-age=%d", maxAge)) // 1 day
			h.ServeHTTP(w, r)
		})
}
