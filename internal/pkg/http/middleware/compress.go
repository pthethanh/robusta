package middleware

import (
	"net/http"

	"github.com/NYTimes/gziphandler"
)

func Compress(h http.Handler) http.Handler {
	return gziphandler.GzipHandler(h)
}
