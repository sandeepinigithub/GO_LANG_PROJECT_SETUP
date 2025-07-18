package controller

import (
	"net/http"
	// "encoding/json"
	"strconv"
	"GO_LANG_PROJECT_SETUP/repository"
	"GO_LANG_PROJECT_SETUP/utils"
)

func GetBanned(w http.ResponseWriter, r *http.Request) {
	banned, err := repository.GetAllBanned()
	if err != nil {
		utils.RespondError(w, "Failed to fetch banned IPs")
		return
	}
	utils.RespondSuccess(w, banned)
}

func Unban(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		utils.RespondError(w, "Invalid ID")
		return
	}
	err = repository.UnbanByID(id)
	if err != nil {
		utils.RespondError(w, "Failed to unban IP")
		return
	}
	utils.RespondSuccessNoData(w)
} 