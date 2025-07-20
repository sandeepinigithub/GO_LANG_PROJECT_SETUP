package controller

import (
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
	"GO_LANG_PROJECT_SETUP/service"
	"GO_LANG_PROJECT_SETUP/utils"
	"GO_LANG_PROJECT_SETUP/api/dto"
)

var mailingListService = service.MailingListService{}



func ListMailingLists(w http.ResponseWriter, r *http.Request) {
	lists, err := mailingListService.ListMailingLists()
	if err != nil {
		utils.RespondError(w, "Failed to fetch mailing lists")
		return
	}
	utils.RespondSuccess(w, lists)
}

func GetMailingList(w http.ResponseWriter, r *http.Request) {
	address := mux.Vars(r)["address"]
	list, err := mailingListService.GetMailingListByAddress(address)
	if err != nil {
		utils.RespondError(w, "Mailing list not found")
		return
	}
	utils.RespondSuccess(w, list)
}

func CreateMailingList(w http.ResponseWriter, r *http.Request) {
	address := mux.Vars(r)["address"]
	var req dto.MailingListRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.RespondError(w, "Invalid request payload")
		return
	}
	_, err := mailingListService.CreateMailingListDTO(address, req)
	if err != nil {
		utils.RespondError(w, err.Error())
		return
	}
	utils.RespondSuccessNoData(w)
}

func UpdateMailingList(w http.ResponseWriter, r *http.Request) {
	address := mux.Vars(r)["address"]
	var req dto.MailingListRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.RespondError(w, "Invalid request payload")
		return
	}
	_, err := mailingListService.UpdateMailingListDTO(address, req)
	if err != nil {
		utils.RespondError(w, err.Error())
		return
	}
	utils.RespondSuccessNoData(w)
}

func DeleteMailingList(w http.ResponseWriter, r *http.Request) {
	address := mux.Vars(r)["address"]
	if err := mailingListService.DeleteMailingList(address); err != nil {
		utils.RespondError(w, "Failed to delete mailing list")
		return
	}
	utils.RespondSuccessNoData(w)
} 