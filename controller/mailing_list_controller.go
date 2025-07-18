package controller

import (
	"net/http"
	"GO_LANG_PROJECT_SETUP/utils"
)

func ListMailingLists(w http.ResponseWriter, r *http.Request) {
	utils.RespondSuccess(w, nil) // TODO: implement
}

func GetMailingList(w http.ResponseWriter, r *http.Request) {
	utils.RespondSuccess(w, nil) // TODO: implement
}

func CreateMailingList(w http.ResponseWriter, r *http.Request) {
	utils.RespondSuccessNoData(w) // TODO: implement
}

func UpdateMailingList(w http.ResponseWriter, r *http.Request) {
	utils.RespondSuccessNoData(w) // TODO: implement
}

func DeleteMailingList(w http.ResponseWriter, r *http.Request) {
	utils.RespondSuccessNoData(w) // TODO: implement
} 