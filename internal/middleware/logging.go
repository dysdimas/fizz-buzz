package middleware

import (
	"log"
	"net/http"
	"time"
)

// LoggingMiddleware logs the request and response details.
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		latency := time.Since(start)
		log.Printf("Request: %s, Latency: %s\n", r.URL.Path, latency)
	})
}
