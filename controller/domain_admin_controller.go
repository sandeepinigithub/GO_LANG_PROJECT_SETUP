package controller

import (
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
	"GO_LANG_PROJECT_SETUP/models"
	"GO_LANG_PROJECT_SETUP/repository"
	"GO_LANG_PROJECT_SETUP/utils"
	"log"
)

func ListDomainAdmins(w http.ResponseWriter, r *http.Request) {
	admins, err := repository.GetAllDomainAdmins()
	if err != nil {
		utils.RespondError(w, "Failed to fetch domain admins")
		return
	}
	utils.RespondSuccess(w, admins)
}

func GetDomainAdmin(w http.ResponseWriter, r *http.Request) {
	email := mux.Vars(r)["email"]
	admin, err := repository.GetDomainAdminByEmail(email)
	if err != nil {
		utils.RespondError(w, "Domain admin not found")
		return
	}
	utils.RespondSuccess(w, admin)
}

func CreateDomainAdmin(w http.ResponseWriter, r *http.Request) {
	var admin models.DomainAdmin
	if err := json.NewDecoder(r.Body).Decode(&admin); err != nil {
		utils.RespondError(w, "Invalid request payload")
		return
	}
	if err := repository.CreateDomainAdmin(&admin); err != nil {
		log.Printf("Failed to create domain admin: %v", err)
		utils.RespondError(w, "Failed to create domain admin")
		return
	}
	utils.RespondSuccess(w, admin)
}

func UpdateDomainAdmin(w http.ResponseWriter, r *http.Request) {
	email := mux.Vars(r)["email"]
	var updated models.DomainAdmin
	if err := json.NewDecoder(r.Body).Decode(&updated); err != nil {
		utils.RespondError(w, "Invalid request payload")
		return
	}
	if err := repository.UpdateDomainAdmin(email, &updated); err != nil {
		log.Printf("Failed to update domain admin: %v", err)
		utils.RespondError(w, "Failed to update domain admin")
		return
	}
	utils.RespondSuccessNoData(w)
}

func DeleteDomainAdmin(w http.ResponseWriter, r *http.Request) {
	email := mux.Vars(r)["email"]
	if err := repository.DeleteDomainAdmin(email); err != nil {
		log.Printf("Failed to delete domain admin: %v", err)
		utils.RespondError(w, "Failed to delete domain admin")
		return
	}
	utils.RespondSuccessNoData(w)
} 