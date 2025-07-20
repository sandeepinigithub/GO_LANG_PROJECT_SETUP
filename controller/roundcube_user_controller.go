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

var roundcubeUserService = service.RoundcubeUserService{}

func ListRoundcubeUsers(w http.ResponseWriter, r *http.Request) {
	users, err := roundcubeUserService.ListRoundcubeUsers()
	if err != nil {
		utils.RespondError(w, "Failed to fetch roundcube users")
		return
	}
	utils.RespondSuccess(w, users)
}

func GetRoundcubeUser(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		utils.RespondError(w, "Invalid ID")
		return
	}
	user, err := roundcubeUserService.GetRoundcubeUserByID(id)
	if err != nil {
		utils.RespondError(w, "Roundcube user not found")
		return
	}
	utils.RespondSuccess(w, user)
}

func CreateRoundcubeUser(w http.ResponseWriter, r *http.Request) {
	var req dto.RoundcubeUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.RespondError(w, "Invalid request payload")
		return
	}
	_, err := roundcubeUserService.CreateRoundcubeUserDTO(req)
	if err != nil {
		utils.RespondError(w, err.Error())
		return
	}
	utils.RespondSuccessNoData(w)
}

func UpdateRoundcubeUser(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		utils.RespondError(w, "Invalid ID")
		return
	}
	var req dto.RoundcubeUserRequest
	if err = json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.RespondError(w, "Invalid request payload")
		return
	}
	_, err = roundcubeUserService.UpdateRoundcubeUserDTO(id, req)
	if err != nil {
		utils.RespondError(w, err.Error())
		return
	}
	utils.RespondSuccessNoData(w)
}

func DeleteRoundcubeUser(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		utils.RespondError(w, "Invalid ID")
		return
	}
	if err := roundcubeUserService.DeleteRoundcubeUser(id); err != nil {
		utils.RespondError(w, "Failed to delete roundcube user")
		return
	}
	utils.RespondSuccessNoData(w)
} 