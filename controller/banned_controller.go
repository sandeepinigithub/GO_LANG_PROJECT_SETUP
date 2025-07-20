package controller

import (
	"net/http"
	"strconv"
	"devsMailGo/service"
	"devsMailGo/utils"
)

var bannedService = service.BannedService{}

type BannedResponse struct {
	IP string `json:"ip"`
}

func GetBanned(w http.ResponseWriter, r *http.Request) {
	banned, err := bannedService.ListBanned()
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
	if err := bannedService.UnbanByID(id); err != nil {
		utils.RespondError(w, "Failed to unban IP")
		return
	}
	utils.RespondSuccessNoData(w)
} 