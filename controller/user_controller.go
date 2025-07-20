package controller

import (
	"encoding/json"
	"net/http"
	"strconv"
	"GO_LANG_PROJECT_SETUP/api/dto"
	"GO_LANG_PROJECT_SETUP/service"
	"github.com/gorilla/mux"
)

var userService = service.UserService{}

// DTOs for user API
// (Add more fields as needed for your API)
type UserRequest struct {
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Domain   string `json:"domain"`
	Quota    int    `json:"quota"`
	Language string `json:"language"`
}

type UserResponse struct {
	ID    uint   `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
	// Add more fields as needed
}

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

// List all users
func GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := userService.ListUsers()
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
	user, err := userService.GetUserByID(uint(id))
	if err != nil {
		respondJSON(w, http.StatusNotFound, false, nil, "User not found")
		return
	}
	respondJSON(w, http.StatusOK, true, user, "User fetched successfully")
}

// Create new user
func CreateUser(w http.ResponseWriter, r *http.Request) {
	var req dto.UserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondJSON(w, http.StatusBadRequest, false, nil, "Invalid request payload")
		return
	}
	user, err := userService.RegisterUserDTO(req)
	if err != nil {
		respondJSON(w, http.StatusInternalServerError, false, nil, err.Error())
		return
	}
	respondJSON(w, http.StatusCreated, true, user, "User created successfully")
}

// Update user by ID
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		respondJSON(w, http.StatusBadRequest, false, nil, "Invalid user ID")
		return
	}
	var req dto.UserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondJSON(w, http.StatusBadRequest, false, nil, "Invalid request payload")
		return
	}
	user, err := userService.UpdateUserDTO(uint(id), req)
	if err != nil {
		respondJSON(w, http.StatusInternalServerError, false, nil, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, true, user, "User updated successfully")
}

// Delete user by ID
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		respondJSON(w, http.StatusBadRequest, false, nil, "Invalid user ID")
		return
	}
	if err := userService.DeleteUser(uint(id)); err != nil {
		respondJSON(w, http.StatusInternalServerError, false, nil, "Failed to delete user")
		return
	}
	respondJSON(w, http.StatusOK, true, nil, "User deleted successfully")
}
