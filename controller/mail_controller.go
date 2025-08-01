package controller

import (
	"encoding/json"
	"net/http"
	"devsMailGo/service"
	"devsMailGo/utils"
	"devsMailGo/api/dto"
	"github.com/gorilla/mux"
)

var mailService = service.MailService{}

// CreateMailboxRequest represents the request to create a mailbox
type CreateMailboxRequest struct {
	Email string `json:"email"`
	Quota int64  `json:"quota"`
}

// UpdateQuotaRequest represents the request to update quota
type UpdateQuotaRequest struct {
	Quota int64 `json:"quota"`
}

// CreateMailbox creates a new mailbox for a user
func CreateMailbox(w http.ResponseWriter, r *http.Request) {
	var req CreateMailboxRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	// Validate email
	if req.Email == "" {
		utils.RespondWithError(w, http.StatusBadRequest, "Email is required")
		return
	}

	// Create user response object
	user := &dto.UserResponse{
		Email: req.Email,
	}

	// Create mailbox
	if err := mailService.CreateMailbox(user); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed to create mailbox: "+err.Error())
		return
	}

	// Set quota if provided
	if req.Quota > 0 {
		if err := mailService.UpdateQuota(req.Email, req.Quota); err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, "Failed to set quota: "+err.Error())
			return
		}
	}

	utils.RespondSuccess(w, map[string]interface{}{
		"message": "Mailbox created successfully",
		"email":   req.Email,
	})
}

// DeleteMailbox deletes a mailbox for a user
func DeleteMailbox(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	email := vars["email"]

	if email == "" {
		utils.RespondWithError(w, http.StatusBadRequest, "Email parameter is required")
		return
	}

	user := &dto.UserResponse{
		Email: email,
	}

	if err := mailService.DeleteMailbox(user); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed to delete mailbox: "+err.Error())
		return
	}

	utils.RespondSuccess(w, map[string]interface{}{
		"message": "Mailbox deleted successfully",
		"email":   email,
	})
}

// GetMailboxInfo retrieves mailbox information
func GetMailboxInfo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	email := vars["email"]

	if email == "" {
		utils.RespondWithError(w, http.StatusBadRequest, "Email parameter is required")
		return
	}

	info, err := mailService.GetMailboxInfo(email)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed to get mailbox info: "+err.Error())
		return
	}

	utils.RespondSuccess(w, info)
}

// UpdateMailboxQuota updates user quota
func UpdateMailboxQuota(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	email := vars["email"]

	if email == "" {
		utils.RespondWithError(w, http.StatusBadRequest, "Email parameter is required")
		return
	}

	var req UpdateQuotaRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	if req.Quota <= 0 {
		utils.RespondWithError(w, http.StatusBadRequest, "Quota must be greater than 0")
		return
	}

	if err := mailService.UpdateQuota(email, req.Quota); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed to update quota: "+err.Error())
		return
	}

	utils.RespondSuccess(w, map[string]interface{}{
		"message": "Quota updated successfully",
		"email":   email,
		"quota":   req.Quota,
	})
}

// ReloadMailServices reloads Postfix and Dovecot configurations
func ReloadMailServices(w http.ResponseWriter, r *http.Request) {
	// Reload Postfix
	if err := mailService.ReloadPostfix(); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed to reload Postfix: "+err.Error())
		return
	}

	// Reload Dovecot
	if err := mailService.ReloadDovecot(); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed to reload Dovecot: "+err.Error())
		return
	}

	utils.RespondSuccess(w, map[string]interface{}{
		"message": "Mail services reloaded successfully",
	})
}

// CheckMailDelivery checks if mail delivery is working
func CheckMailDelivery(w http.ResponseWriter, r *http.Request) {
	if err := mailService.CheckMailDelivery(); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Mail delivery check failed: "+err.Error())
		return
	}

	utils.RespondSuccess(w, map[string]interface{}{
		"message": "Mail delivery is working correctly",
	})
}

// GetMailQueue gets the current mail queue
func GetMailQueue(w http.ResponseWriter, r *http.Request) {
	queue, err := mailService.GetMailQueue()
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed to get mail queue: "+err.Error())
		return
	}

	utils.RespondSuccess(w, map[string]interface{}{
		"queue":   queue,
		"count":   len(queue),
	})
} 