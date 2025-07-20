package controller

import (
	"encoding/json"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
	"devsMailGo/service"
	"devsMailGo/utils"
	"devsMailGo/api/dto"
)

var throttleService = service.ThrottleService{}

func ListThrottle(w http.ResponseWriter, r *http.Request) {
	entries, err := throttleService.ListThrottle()
	if err != nil {
		utils.RespondError(w, "Failed to fetch throttle entries")
		return
	}
	utils.RespondSuccess(w, entries)
}

func GetThrottle(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		utils.RespondError(w, "Invalid ID")
		return
	}
	entry, err := throttleService.GetThrottleByID(id)
	if err != nil {
		utils.RespondError(w, "Throttle entry not found")
		return
	}
	utils.RespondSuccess(w, entry)
}

func CreateThrottle(w http.ResponseWriter, r *http.Request) {
	var req dto.ThrottleRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.RespondError(w, "Invalid request payload")
		return
	}
	_, err := throttleService.CreateThrottleDTO(req)
	if err != nil {
		utils.RespondError(w, err.Error())
		return
	}
	utils.RespondSuccessNoData(w)
}

func UpdateThrottle(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		utils.RespondError(w, "Invalid ID")
		return
	}
	var req dto.ThrottleRequest
	if err = json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.RespondError(w, "Invalid request payload")
		return
	}
	_, err = throttleService.UpdateThrottleDTO(id, req)
	if err != nil {
		utils.RespondError(w, err.Error())
		return
	}
	utils.RespondSuccessNoData(w)
}

func DeleteThrottle(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		utils.RespondError(w, "Invalid ID")
		return
	}
	if err := throttleService.DeleteThrottle(id); err != nil {
		utils.RespondError(w, "Failed to delete throttle entry")
		return
	}
	utils.RespondSuccessNoData(w)
} 