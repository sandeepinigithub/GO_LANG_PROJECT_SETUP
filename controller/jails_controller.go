package controller

import (
	"net/http"
	"GO_LANG_PROJECT_SETUP/service"
	"GO_LANG_PROJECT_SETUP/utils"
)

var jailsService = service.JailsService{}

type JailResponse struct {
	// Add fields as needed
}

func GetJails(w http.ResponseWriter, r *http.Request) {
	jails, err := jailsService.ListJails()
	if err != nil {
		utils.RespondError(w, "Failed to fetch jails")
		return
	}
	utils.RespondSuccess(w, jails)
} 