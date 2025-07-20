package utils

import (
	"encoding/json"
	"net/http"
	"time"
	"github.com/golang-jwt/jwt/v4"
)

// Response represents a standard API response structure
type Response struct {
	Success   bool        `json:"success"`
	Data      interface{} `json:"data,omitempty"`
	Message   string      `json:"message,omitempty"`
	Error     string      `json:"error,omitempty"`
	Timestamp string      `json:"timestamp"`
}

// UserClaims represents JWT claims structure
type UserClaims struct {
	Email     string    `json:"email"`
	Role      string    `json:"role"`
	Domain    string    `json:"domain"`
	ExpiresAt time.Time `json:"expires_at"`
	jwt.RegisteredClaims
}

// JWTSecret should be stored in environment variables in production
var JWTSecret = []byte("your-secret-key-change-in-production")

// RespondSuccess sends a successful response
func RespondSuccess(w http.ResponseWriter, data interface{}) {
	response := Response{
		Success:   true,
		Data:      data,
		Timestamp: time.Now().UTC().Format(time.RFC3339),
	}
	
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

// RespondSuccessNoData sends a successful response without data
func RespondSuccessNoData(w http.ResponseWriter) {
	response := Response{
		Success:   true,
		Message:   "Operation completed successfully",
		Timestamp: time.Now().UTC().Format(time.RFC3339),
	}
	
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

// RespondError sends an error response
func RespondError(w http.ResponseWriter, message string) {
	RespondWithError(w, http.StatusInternalServerError, message)
}

// RespondWithError sends an error response with custom status code
func RespondWithError(w http.ResponseWriter, statusCode int, message string) {
	response := Response{
		Success:   false,
		Error:     message,
		Timestamp: time.Now().UTC().Format(time.RFC3339),
	}
	
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(response)
}

// GenerateJWT creates a new JWT token for the given email
func GenerateJWT(email string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	
	claims := &UserClaims{
		Email:     email,
		Role:      "admin", // Default role, should be fetched from user data
		Domain:    "example.com", // Default domain, should be fetched from user data
		ExpiresAt: expirationTime,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "devsMailGo",
			Subject:   email,
		},
	}
	
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(JWTSecret)
}

// ValidateJWT validates and parses a JWT token
func ValidateJWT(tokenString string) (*UserClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return JWTSecret, nil
	})
	
	if err != nil {
		return nil, err
	}
	
	if claims, ok := token.Claims.(*UserClaims); ok && token.Valid {
		return claims, nil
	}
	
	return nil, jwt.ErrSignatureInvalid
}

// GenerateJWTWithClaims creates a JWT token with custom claims
func GenerateJWTWithClaims(email, role, domain string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	
	claims := &UserClaims{
		Email:     email,
		Role:      role,
		Domain:    domain,
		ExpiresAt: expirationTime,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "devsMailGo",
			Subject:   email,
		},
	}
	
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(JWTSecret)
} 