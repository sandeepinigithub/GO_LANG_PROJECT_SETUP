package controller

import (
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
	"devsMailGo/service"
	"devsMailGo/utils"
	"devsMailGo/api/dto"
)

var quotaService = service.QuotaService{}

func ListQuota(w http.ResponseWriter, r *http.Request) {
	quotas, err := quotaService.ListQuota()
	if err != nil {
		utils.RespondError(w, "Failed to fetch quota entries")
		return
	}
	utils.RespondSuccess(w, quotas)
}

func GetQuota(w http.ResponseWriter, r *http.Request) {
	username := mux.Vars(r)["username"]
	quota, err := quotaService.GetQuotaByUsername(username)
	if err != nil {
		utils.RespondError(w, "Quota entry not found")
		return
	}
	utils.RespondSuccess(w, quota)
}

func CreateQuota(w http.ResponseWriter, r *http.Request) {
	username := mux.Vars(r)["username"]
	var req dto.QuotaRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.RespondError(w, "Invalid request payload")
		return
	}
	_, err := quotaService.CreateQuotaDTO(username, req)
	if err != nil {
		utils.RespondError(w, err.Error())
		return
	}
	utils.RespondSuccessNoData(w)
}

func UpdateQuota(w http.ResponseWriter, r *http.Request) {
	username := mux.Vars(r)["username"]
	var req dto.QuotaRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.RespondError(w, "Invalid request payload")
		return
	}
	_, err := quotaService.UpdateQuotaDTO(username, req)
	if err != nil {
		utils.RespondError(w, err.Error())
		return
	}
	utils.RespondSuccessNoData(w)
}

func DeleteQuota(w http.ResponseWriter, r *http.Request) {
	username := mux.Vars(r)["username"]
	if err := quotaService.DeleteQuota(username); err != nil {
		utils.RespondError(w, "Failed to delete quota entry")
		return
	}
	utils.RespondSuccessNoData(w)
} 