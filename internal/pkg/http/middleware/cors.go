package middleware

import (
	"net/http"
	"os"

	"github.com/gorilla/handlers"
)

func CORS(h http.Handler) http.Handler {
	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Accept", "Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization"})
	originsOk := handlers.AllowedOrigins([]string{os.Getenv("HTTP_ALLOWED_ORIGINS")})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS", "DELETE"})
	return handlers.CORS(headersOk, originsOk, methodsOk)(h)
}
