package middleware

import (
	"log"
	"net/http"
	"time"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()

		// Call the next handler in the chain
		next.ServeHTTP(w, r)

		// Logging
		endTime := time.Now()
		latency := endTime.Sub(startTime)
		log.Printf("[%s] %s %s %s", r.Method, r.RequestURI, latency, r.RemoteAddr)
	})
}

