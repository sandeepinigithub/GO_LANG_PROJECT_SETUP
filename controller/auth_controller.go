package controller

import (
	"encoding/json"
	"net/http"
	"GO_LANG_PROJECT_SETUP/utils"
	"GO_LANG_PROJECT_SETUP/repository"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Login(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.RespondError(w, "Invalid request payload")
		return
	}
	user, err := repository.GetRoundcubeUserByUsername(req.Username)
	if err != nil {
		utils.RespondError(w, "User not found")
		return
	}
	// For demo: assume password is plain text
	if user.Password != req.Password {
		utils.RespondError(w, "Invalid password")
		return
	}
	utils.RespondSuccessNoData(w)
} 