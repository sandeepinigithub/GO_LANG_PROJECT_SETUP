package controller

import (
	"encoding/json"
	"net/http"
	"strings"
	"devsMailGo/service"
	"devsMailGo/utils"
)

var authUserService = service.UserService{}
var ldapService = service.LDAPService{}

// LoginRequest represents the login request structure
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	AuthType string `json:"auth_type"` // "local" or "ldap"
}

// LoginResponse represents the login response structure
type LoginResponse struct {
	Token     string                 `json:"token"`
	User      map[string]interface{} `json:"user"`
	ExpiresAt string                 `json:"expires_at"`
	AuthType  string                 `json:"auth_type"`
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

	// Default to local authentication if not specified
	if req.AuthType == "" {
		req.AuthType = "local"
	}

	var user *service.LDAPUser
	var err error

	// Authenticate based on auth type
	switch req.AuthType {
	case "ldap":
		// Authenticate against LDAP
		user, err = ldapService.AuthenticateUser(req.Username, req.Password)
		if err != nil {
			utils.RespondWithError(w, http.StatusUnauthorized, "LDAP authentication failed: "+err.Error())
			return
		}
		
		// Sync LDAP user to local database if needed
		_, err = ldapService.SyncUserToDatabase(user)
		if err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, "Failed to sync LDAP user")
			return
		}
		
	default:
		// Local authentication
		localUser, err := authUserService.AuthenticateUser(req.Username, req.Password)
		if err != nil {
			utils.RespondWithError(w, http.StatusUnauthorized, "Invalid credentials")
			return
		}
		
		// Convert to LDAPUser format for consistency
		user = &service.LDAPUser{
			UID:    req.Username,
			Email:  localUser.Email,
			Name:   localUser.Name,
			Groups: []string{"local_users"},
		}
	}

	// Use default domain if user domain is empty
	domain := extractDomain(user.Email)
	if domain == "" {
		domain = "example.com" // Default domain
	}

	// Determine user role based on groups
	role := "user"
	for _, group := range user.Groups {
		if group == "admins" || group == "mail_admins" {
			role = "admin"
			break
		}
	}

	// Generate JWT token with user information
	token, err := utils.GenerateJWTWithClaims(user.Email, role, domain)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed to generate token")
		return
	}

	// Prepare response
	resp := LoginResponse{
		Token: token,
		User: map[string]interface{}{
			"email":  user.Email,
			"name":   user.Name,
			"domain": domain,
			"role":   role,
			"groups": user.Groups,
		},
		ExpiresAt: "24h",
		AuthType:  req.AuthType,
	}

	utils.RespondSuccess(w, resp)
}

// extractDomain extracts domain from email
func extractDomain(email string) string {
	parts := strings.Split(email, "@")
	if len(parts) == 2 {
		return parts[1]
	}
	return ""
} 