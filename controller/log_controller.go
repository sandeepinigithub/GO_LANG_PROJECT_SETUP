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

var logService = service.LogService{}

func ListLogs(w http.ResponseWriter, r *http.Request) {
	logs, err := logService.ListLogs()
	if err != nil {
		utils.RespondError(w, "Failed to fetch logs")
		return
	}
	utils.RespondSuccess(w, logs)
}

func GetLog(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		utils.RespondError(w, "Invalid ID")
		return
	}
	logEntry, err := logService.GetLogByID(id)
	if err != nil {
		utils.RespondError(w, "Log entry not found")
		return
	}
	utils.RespondSuccess(w, logEntry)
}

func CreateLog(w http.ResponseWriter, r *http.Request) {
	var req dto.LogRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.RespondError(w, "Invalid request payload")
		return
	}
	_, err := logService.CreateLogDTO(req)
	if err != nil {
		utils.RespondError(w, err.Error())
		return
	}
	utils.RespondSuccessNoData(w)
}

func UpdateLog(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		utils.RespondError(w, "Invalid ID")
		return
	}
	var req dto.LogRequest
	if err = json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.RespondError(w, "Invalid request payload")
		return
	}
	_, err = logService.UpdateLogDTO(id, req)
	if err != nil {
		utils.RespondError(w, err.Error())
		return
	}
	utils.RespondSuccessNoData(w)
}

func DeleteLog(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		utils.RespondError(w, "Invalid ID")
		return
	}
	if err := logService.DeleteLog(id); err != nil {
		utils.RespondError(w, "Failed to delete log entry")
		return
	}
	utils.RespondSuccessNoData(w)
} 