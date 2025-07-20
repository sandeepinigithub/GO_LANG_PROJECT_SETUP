package controller

import (
	"encoding/json"
	"net/http"
	"devsMailGo/service"
	"devsMailGo/utils"
)

var authUserService = service.UserService{}

// LoginRequest represents the login request structure
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// LoginResponse represents the login response structure
type LoginResponse struct {
	Token     string                 `json:"token"`
	User      map[string]interface{} `json:"user"`
	ExpiresAt string                 `json:"expires_at"`
}

// Login authenticates a user and returns a JWT token
func Login(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	// Validate required fields
	if req.Username == "" || req.Password == "" {
		utils.RespondWithError(w, http.StatusBadRequest, "Username and password are required")
		return
	}

	// Authenticate user
	user, err := authUserService.AuthenticateUser(req.Username, req.Password)
	if err != nil {
		utils.RespondWithError(w, http.StatusUnauthorized, "Invalid credentials")
		return
	}

	// Use default domain if user domain is empty
	domain := user.Domain
	if domain == "" {
		domain = "example.com" // Default domain
	}

	// Generate JWT token with user information
	token, err := utils.GenerateJWTWithClaims(user.Email, "admin", domain)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed to generate token")
		return
	}

	// Prepare response
	resp := LoginResponse{
		Token: token,
		User: map[string]interface{}{
			"id":     user.ID,
			"email":  user.Email,
			"name":   user.Name,
			"domain": user.Domain,
			"role":   "admin",
		},
		ExpiresAt: "24h",
	}

	utils.RespondSuccess(w, resp)
} 