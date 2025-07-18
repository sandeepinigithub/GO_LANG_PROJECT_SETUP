package controller

import (
	"net/http"
	"GO_LANG_PROJECT_SETUP/repository"
	"GO_LANG_PROJECT_SETUP/utils"
)

func GetJails(w http.ResponseWriter, r *http.Request) {
	jails, err := repository.GetAllJails()
	if err != nil {
		utils.RespondError(w, "Failed to fetch jails")
		return
	}
	utils.RespondSuccess(w, jails)
} 