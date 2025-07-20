package controller

import (
	"encoding/json"
	"net/http"
	"GO_LANG_PROJECT_SETUP/service"
	"GO_LANG_PROJECT_SETUP/utils"
)

var authUserService = service.UserService{}

// DTO for login request
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// DTO for login response
type LoginResponse struct {
	Token string `json:"token"`
}

// Login authenticates a user and returns a JWT token
func Login(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.RespondError(w, "Invalid request payload")
		return
	}
	user, err := authUserService.AuthenticateUser(req.Username, req.Password)
	if err != nil {
		utils.RespondError(w, err.Error())
		return
	}
	token, err := utils.GenerateJWT(user.Email)
	if err != nil {
		utils.RespondError(w, "Failed to generate token")
		return
	}
	resp := LoginResponse{Token: token}
	utils.RespondSuccess(w, resp)
} 