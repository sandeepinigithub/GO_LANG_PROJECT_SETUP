package controller

import (
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
	"GO_LANG_PROJECT_SETUP/service"
	"GO_LANG_PROJECT_SETUP/utils"
	"GO_LANG_PROJECT_SETUP/api/dto"
)

var aliasService = service.AliasService{}

func ListAliases(w http.ResponseWriter, r *http.Request) {
	aliases, err := aliasService.ListAliases()
	if err != nil {
		utils.RespondError(w, "Failed to fetch aliases")
		return
	}
	utils.RespondSuccess(w, aliases)
}

func GetAlias(w http.ResponseWriter, r *http.Request) {
	address := mux.Vars(r)["address"]
	alias, err := aliasService.GetAliasByAddress(address)
	if err != nil {
		utils.RespondError(w, "Alias not found")
		return
	}
	utils.RespondSuccess(w, alias)
}

func CreateAlias(w http.ResponseWriter, r *http.Request) {
	address := mux.Vars(r)["address"]
	var req dto.AliasRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.RespondError(w, "Invalid request payload")
		return
	}
	_, err := aliasService.CreateAliasDTO(address, req)
	if err != nil {
		utils.RespondError(w, err.Error())
		return
	}
	utils.RespondSuccessNoData(w)
}

func UpdateAlias(w http.ResponseWriter, r *http.Request) {
	address := mux.Vars(r)["address"]
	var req dto.AliasRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.RespondError(w, "Invalid request payload")
		return
	}
	_, err := aliasService.UpdateAliasDTO(address, req)
	if err != nil {
		utils.RespondError(w, err.Error())
		return
	}
	utils.RespondSuccessNoData(w)
}

func DeleteAlias(w http.ResponseWriter, r *http.Request) {
	address := mux.Vars(r)["address"]
	if err := aliasService.DeleteAlias(address); err != nil {
		utils.RespondError(w, "Failed to delete alias")
		return
	}
	utils.RespondSuccessNoData(w)
} 