package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	models "GO_LANG_PROJECT_SETUP/models"
	"GO_LANG_PROJECT_SETUP/repository"
	"GO_LANG_PROJECT_SETUP/service"

	"github.com/gorilla/mux"
)

var userService = service.UserService{}

// Helper for standard response
func respondJSON(w http.ResponseWriter, status int, success bool, data interface{}, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": success,
		"data":    data,
		"message": message,
	})
}

// Get all users
func GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := repository.GetAllUsers()
	if err != nil {
		respondJSON(w, http.StatusInternalServerError, false, nil, "Failed to fetch users")
		return
	}
	respondJSON(w, http.StatusOK, true, users, "Users fetched successfully")
}

// Get user by ID
func GetUser(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		respondJSON(w, http.StatusBadRequest, false, nil, "Invalid user ID")
		return
	}
	user, err := repository.GetUserByID(uint(id))
	if err != nil {
		respondJSON(w, http.StatusNotFound, false, nil, "User not found")
		return
	}
	respondJSON(w, http.StatusOK, true, user, "User fetched successfully")
}

// Create new user
func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		respondJSON(w, http.StatusBadRequest, false, nil, "Invalid request payload")
		return
	}
	createdUser, err := repository.CreateUser(user)
	if err != nil {
		respondJSON(w, http.StatusInternalServerError, false, nil, "Failed to create user")
		return
	}
	respondJSON(w, http.StatusCreated, true, createdUser, "User created successfully")
}

// Update user by ID
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		respondJSON(w, http.StatusBadRequest, false, nil, "Invalid user ID")
		return
	}
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		respondJSON(w, http.StatusBadRequest, false, nil, "Invalid request payload")
		return
	}
	updatedUser, err := repository.UpdateUser(uint(id), user)
	if err != nil {
		respondJSON(w, http.StatusInternalServerError, false, nil, "Failed to update user")
		return
	}
	respondJSON(w, http.StatusOK, true, updatedUser, "User updated successfully")
}

// Delete user by ID
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		respondJSON(w, http.StatusBadRequest, false, nil, "Invalid user ID")
		return
	}
	if err := repository.DeleteUser(uint(id)); err != nil {
		respondJSON(w, http.StatusInternalServerError, false, nil, "Failed to delete user")
		return
	}
	respondJSON(w, http.StatusOK, true, nil, "User deleted successfully")
}
