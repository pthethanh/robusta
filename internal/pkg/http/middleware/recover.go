package middleware

import (
	"fmt"
	"net/http"
	"runtime"

	"github.com/pthethanh/robusta/internal/pkg/http/respond"
	"github.com/pthethanh/robusta/internal/pkg/log"
)

// Recover is a middleware that recover a handler from a panic
func Recover(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if rec := recover(); rec != nil {
				err, ok := rec.(error)
				if !ok {
					err = fmt.Errorf("%v", rec)
				}
				stack := make([]byte, 4<<10) // 4KB
				length := runtime.Stack(stack, false)

				log.WithContext(r.Context()).Errorf("panic recover, err: %v, stack: %s", err, stack[:length])
				respond.Error(w,
					fmt.Errorf(http.StatusText(http.StatusInternalServerError)),
					http.StatusInternalServerError)
			}
		}()

		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}
