package controller

import (
	"net/http"
	"GO_LANG_PROJECT_SETUP/utils"
)

func ListDomains(w http.ResponseWriter, r *http.Request) {
	utils.RespondSuccess(w, nil) // TODO: implement
}

func GetDomain(w http.ResponseWriter, r *http.Request) {
	utils.RespondSuccess(w, nil) // TODO: implement
}

func CreateDomain(w http.ResponseWriter, r *http.Request) {
	utils.RespondSuccessNoData(w) // TODO: implement
}

func UpdateDomain(w http.ResponseWriter, r *http.Request) {
	utils.RespondSuccessNoData(w) // TODO: implement
}

func DeleteDomain(w http.ResponseWriter, r *http.Request) {
	utils.RespondSuccessNoData(w) // TODO: implement
} 