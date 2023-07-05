package middleware

import (
	"context"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/coderj001/phoneguardian/app/auth"
	"github.com/coderj001/phoneguardian/app/handler"
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


func AuthMiddleware(applyAuth bool) func(http.Handler) http.HandlerFunc {
	return func(next http.Handler) http.HandlerFunc {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if applyAuth {
				tokenString := r.Header.Get("Authorization")
				if tokenString == "" {
					handler.RespondError(w, http.StatusUnauthorized, "Missing authorization token")
					return
				}
				bearerToken := strings.Split(tokenString, " ")
				if len(bearerToken) == 2 {
					tokenString = bearerToken[1]
				}

				claims, err := auth.ValidateToken(tokenString)

				if err != nil {
					handler.RespondError(w, http.StatusUnauthorized, "Invalid authorization token")
					return
				}

				ctx := context.WithValue(r.Context(), "userID", claims.UserID)
				r = r.WithContext(ctx)
			}

			next.ServeHTTP(w, r)
		})
	}
}
