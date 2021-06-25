package middleware

import (
	"github.com/gorilla/handlers"
	"net/http"
	"os"
)

// Logger Is a utility middleware that coerces gorilla's logging middleware
// into the right type to work with mux.Use()
func Logger(next http.Handler) http.Handler {
	loggingHandler := handlers.LoggingHandler(os.Stdout, next)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		loggingHandler.ServeHTTP(w, r)
	})
}
