package middleware

import (
	"net/http"
	"strings"
	"GO_LANG_PROJECT_SETUP/utils"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			utils.RespondError(w, "Missing or invalid Authorization header")
			return
		}
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		claims, err := utils.ValidateJWT(tokenString)
		if err != nil {
			utils.RespondError(w, "Invalid or expired token")
			return
		}
		// Optionally set claims in context for downstream handlers
		r = utils.SetUserClaimsInContext(r, claims)
		next.ServeHTTP(w, r)
	})
} 