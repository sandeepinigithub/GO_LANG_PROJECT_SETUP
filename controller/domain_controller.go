package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"GO_LANG_PROJECT_SETUP/service"
	"GO_LANG_PROJECT_SETUP/utils"
	"GO_LANG_PROJECT_SETUP/api/dto"
)

var domainService = service.DomainService{}

func ListDomains(w http.ResponseWriter, r *http.Request) {
	domains, err := domainService.ListDomains()
	if err != nil {
		log.Printf("Failed to fetch domains: %v", err)
		utils.RespondError(w, "Failed to fetch domains")
		return
	}
	utils.RespondSuccess(w, domains)
}

func GetDomain(w http.ResponseWriter, r *http.Request) {
	name := mux.Vars(r)["domain"]
	domain, err := domainService.GetDomainByName(name)
	if err != nil {
		log.Printf("Domain not found: %v", err)
		utils.RespondError(w, "Domain not found")
		return
	}
	utils.RespondSuccess(w, domain)
}

func CreateDomain(w http.ResponseWriter, r *http.Request) {
	name := mux.Vars(r)["domain"]
	var req dto.DomainRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.RespondError(w, "Invalid request payload")
		return
	}
	_, err := domainService.CreateDomainDTO(name, req)
	if err != nil {
		log.Printf("Failed to create domain: %v", err)
		utils.RespondError(w, err.Error())
		return
	}
	utils.RespondSuccessNoData(w)
}

func UpdateDomain(w http.ResponseWriter, r *http.Request) {
	name := mux.Vars(r)["domain"]
	var req dto.DomainRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.RespondError(w, "Invalid request payload")
		return
	}
	_, err := domainService.UpdateDomainDTO(name, req)
	if err != nil {
		log.Printf("Failed to update domain: %v", err)
		utils.RespondError(w, err.Error())
		return
	}
	utils.RespondSuccessNoData(w)
}

func DeleteDomain(w http.ResponseWriter, r *http.Request) {
	name := mux.Vars(r)["domain"]
	if err := domainService.DeleteDomain(name); err != nil {
		log.Printf("Failed to delete domain: %v", err)
		utils.RespondError(w, "Failed to delete domain")
		return
	}
	utils.RespondSuccessNoData(w)
} 