package controller

import (
	"encoding/json"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
	"GO_LANG_PROJECT_SETUP/service"
	"GO_LANG_PROJECT_SETUP/utils"
	"GO_LANG_PROJECT_SETUP/api/dto"
)

var wblistService = service.WblistService{}

func ListWblist(w http.ResponseWriter, r *http.Request) {
	entries, err := wblistService.ListWblist()
	if err != nil {
		utils.RespondError(w, "Failed to fetch wblist entries")
		return
	}
	utils.RespondSuccess(w, entries)
}

func GetWblist(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	rid, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		utils.RespondError(w, "Invalid ID")
		return
	}
	entry, err := wblistService.GetWblistByRid(rid)
	if err != nil {
		utils.RespondError(w, "Wblist entry not found")
		return
	}
	utils.RespondSuccess(w, entry)
}

func CreateWblist(w http.ResponseWriter, r *http.Request) {
	var req dto.WblistRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.RespondError(w, "Invalid request payload")
		return
	}
	_, err := wblistService.CreateWblistDTO(req)
	if err != nil {
		utils.RespondError(w, err.Error())
		return
	}
	utils.RespondSuccessNoData(w)
}

func UpdateWblist(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	rid, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		utils.RespondError(w, "Invalid ID")
		return
	}
	var req dto.WblistRequest
	if err = json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.RespondError(w, "Invalid request payload")
		return
	}
	_, err = wblistService.UpdateWblistDTO(rid, req)
	if err != nil {
		utils.RespondError(w, err.Error())
		return
	}
	utils.RespondSuccessNoData(w)
}

func DeleteWblist(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	rid, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		utils.RespondError(w, "Invalid ID")
		return
	}
	if err := wblistService.DeleteWblist(rid); err != nil {
		utils.RespondError(w, "Failed to delete wblist entry")
		return
	}
	utils.RespondSuccessNoData(w)
} 