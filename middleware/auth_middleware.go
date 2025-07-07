package middleware

import (
	"airline-system/utils"
	"context"
	"net/http"
	"strings"
)

type key int

const UserCtxKey key = 0

// JWTAuthMiddleware verifies the JWT token and adds user ID to context
func JWTAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			http.Error(w, "Missing or invalid Authorization header", http.StatusUnauthorized)
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		userID, err := utils.ValidateJWT(tokenString)
		if err != nil {
			http.Error(w, "Invalid token: "+err.Error(), http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), UserCtxKey, userID) // ✅ userID directly
		next.ServeHTTP(w, r.WithContext(ctx))

	})
}

// Helper to retrieve user ID from context
func GetUserIDFromContext(ctx context.Context) string {
	userID, _ := ctx.Value(UserCtxKey).(string)
	return userID
}
