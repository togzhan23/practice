package auth

import (
	"context"
	"net/http"
	"strings"
)

type contextKey string

const userContextKey = contextKey("user")

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Unauthorized: missing Authorization header", http.StatusUnauthorized)
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			http.Error(w, "Unauthorized: invalid Authorization header format", http.StatusUnauthorized)
			return
		}
		token := parts[1]

		if !validateToken(token) {
			http.Error(w, "Unauthorized: invalid token", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), userContextKey, token)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func validateToken(token string) bool {
	return token == "valid-token"
}

func FromContext(ctx context.Context) (string, bool) {
	user, ok := ctx.Value(userContextKey).(string)
	return user, ok
}
