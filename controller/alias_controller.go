package controller

import (
	"net/http"
	"GO_LANG_PROJECT_SETUP/utils"
)

func ListAliases(w http.ResponseWriter, r *http.Request) {
	utils.RespondSuccess(w, nil) // TODO: implement
}

func GetAlias(w http.ResponseWriter, r *http.Request) {
	utils.RespondSuccess(w, nil) // TODO: implement
}

func CreateAlias(w http.ResponseWriter, r *http.Request) {
	utils.RespondSuccessNoData(w) // TODO: implement
}

func UpdateAlias(w http.ResponseWriter, r *http.Request) {
	utils.RespondSuccessNoData(w) // TODO: implement
}

func DeleteAlias(w http.ResponseWriter, r *http.Request) {
	utils.RespondSuccessNoData(w) // TODO: implement
} 