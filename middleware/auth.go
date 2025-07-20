package middleware

import (
	"net/http"
	"strings"
	"context"
	"time"
	"GO_LANG_PROJECT_SETUP/utils"
	"github.com/gorilla/mux"
)

// UserClaims represents the JWT claims structure
type UserClaims struct {
	Email    string    `json:"email"`
	Role     string    `json:"role"`
	Domain   string    `json:"domain"`
	ExpiresAt time.Time `json:"expires_at"`
}

// ContextKey is a custom type for context keys
type ContextKey string

const (
	UserClaimsKey ContextKey = "user_claims"
)

// AuthMiddleware validates JWT tokens and sets user claims in context
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Skip authentication for OPTIONS requests (CORS preflight)
		if r.Method == "OPTIONS" {
			next.ServeHTTP(w, r)
			return
		}

		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			utils.RespondWithError(w, http.StatusUnauthorized, "Authorization header required")
			return
		}

		if !strings.HasPrefix(authHeader, "Bearer ") {
			utils.RespondWithError(w, http.StatusUnauthorized, "Invalid authorization format. Use 'Bearer <token>'")
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == "" {
			utils.RespondWithError(w, http.StatusUnauthorized, "Token is required")
			return
		}

		// Validate JWT token
		claims, err := utils.ValidateJWT(tokenString)
		if err != nil {
			utils.RespondWithError(w, http.StatusUnauthorized, "Invalid or expired token")
			return
		}

		// Check if token is expired
		if time.Now().After(claims.ExpiresAt) {
			utils.RespondWithError(w, http.StatusUnauthorized, "Token has expired")
			return
		}

		// Set user claims in context for downstream handlers
		ctx := context.WithValue(r.Context(), UserClaimsKey, claims)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}

// RoleMiddleware checks if the user has the required role
func RoleMiddleware(requiredRoles ...string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			claims, ok := r.Context().Value(UserClaimsKey).(*utils.UserClaims)
			if !ok {
				utils.RespondWithError(w, http.StatusForbidden, "User claims not found")
				return
			}

			// Check if user has any of the required roles
			hasRole := false
			for _, role := range requiredRoles {
				if claims.Role == role {
					hasRole = true
					break
				}
			}

			if !hasRole {
				utils.RespondWithError(w, http.StatusForbidden, "Insufficient permissions")
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}

// DomainMiddleware checks if the user has access to the specified domain
func DomainMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		claims, ok := r.Context().Value(UserClaimsKey).(*utils.UserClaims)
		if !ok {
			utils.RespondWithError(w, http.StatusForbidden, "User claims not found")
			return
		}

		// Extract domain from URL path
		vars := mux.Vars(r)
		requestedDomain := vars["domain"]

		// Super admin can access all domains
		if claims.Role == "super_admin" {
			next.ServeHTTP(w, r)
			return
		}

		// Domain admin can only access their assigned domain
		if claims.Role == "domain_admin" && claims.Domain == requestedDomain {
			next.ServeHTTP(w, r)
			return
		}

		utils.RespondWithError(w, http.StatusForbidden, "Access denied to this domain")
	})
}

// GetUserClaims retrieves user claims from context
func GetUserClaims(r *http.Request) (*utils.UserClaims, bool) {
	claims, ok := r.Context().Value(UserClaimsKey).(*utils.UserClaims)
	return claims, ok
} 