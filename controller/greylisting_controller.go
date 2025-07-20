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

var greylistingService = service.GreylistingService{}

func ListGreylisting(w http.ResponseWriter, r *http.Request) {
	entries, err := greylistingService.ListGreylisting()
	if err != nil {
		utils.RespondError(w, "Failed to fetch greylisting entries")
		return
	}
	utils.RespondSuccess(w, entries)
}

func GetGreylisting(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		utils.RespondError(w, "Invalid ID")
		return
	}
	entry, err := greylistingService.GetGreylistingByID(id)
	if err != nil {
		utils.RespondError(w, "Greylisting entry not found")
		return
	}
	utils.RespondSuccess(w, entry)
}

func CreateGreylisting(w http.ResponseWriter, r *http.Request) {
	var req dto.GreylistingRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.RespondError(w, "Invalid request payload")
		return
	}
	_, err := greylistingService.CreateGreylistingDTO(req)
	if err != nil {
		utils.RespondError(w, err.Error())
		return
	}
	utils.RespondSuccessNoData(w)
}

func UpdateGreylisting(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		utils.RespondError(w, "Invalid ID")
		return
	}
	var req dto.GreylistingRequest
	if err = json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.RespondError(w, "Invalid request payload")
		return
	}
	_, err = greylistingService.UpdateGreylistingDTO(id, req)
	if err != nil {
		utils.RespondError(w, err.Error())
		return
	}
	utils.RespondSuccessNoData(w)
}

func DeleteGreylisting(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		utils.RespondError(w, "Invalid ID")
		return
	}
	if err := greylistingService.DeleteGreylisting(id); err != nil {
		utils.RespondError(w, "Failed to delete greylisting entry")
		return
	}
	utils.RespondSuccessNoData(w)
} 