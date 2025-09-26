package api

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Kranold/hyrox/internal/auth"
)

func (cfg *APIConfig) AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//get authorization header
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Missing Authorization header", http.StatusUnauthorized)
			return
		}
		// get token string

		tokenString := authHeader[len("Bearer "):]

		// validate jwt and handle errors
		userID, err := auth.ValidateJWT(tokenString, cfg.JWTSecret)
		if err != nil {
			fmt.Println(err)
			http.Error(w, "Invalid or expired token", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), "userID", userID)
		next.ServeHTTP(w, r.WithContext(ctx))

	})
}
