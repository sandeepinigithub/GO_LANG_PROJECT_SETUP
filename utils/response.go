package utils

import (
	"encoding/json"
	"net/http"
	"github.com/golang-jwt/jwt/v4"
	"time"
	"context"
	"os"
)

var jwtSecret = []byte(getJWTSecret())

func getJWTSecret() string {
	if s := os.Getenv("JWT_SECRET"); s != "" {
		return s
	}
	return "default_secret" // fallback for dev
}

type JWTClaims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func GenerateJWT(username string) (string, error) {
	claims := JWTClaims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

func ValidateJWT(tokenString string) (*JWTClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(*JWTClaims)
	if !ok || !token.Valid {
		return nil, err
	}
	return claims, nil
}

// Context helpers for passing claims
type contextKey string
const userClaimsKey = contextKey("userClaims")

func SetUserClaimsInContext(r *http.Request, claims *JWTClaims) *http.Request {
	ctx := context.WithValue(r.Context(), userClaimsKey, claims)
	return r.WithContext(ctx)
}

func GetUserClaimsFromContext(r *http.Request) *JWTClaims {
	claims, _ := r.Context().Value(userClaimsKey).(*JWTClaims)
	return claims
}

func RespondSuccess(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"_success": true,
		"_data":    data,
	})
}

func RespondSuccessNoData(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"_success": true,
	})
}

func RespondError(w http.ResponseWriter, msg string) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"_success": false,
		"_msg":     msg,
	})
} 