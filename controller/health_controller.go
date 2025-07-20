package controller

import (
	"net/http"
	"GO_LANG_PROJECT_SETUP/utils"
)

// HealthCheck provides system health status
func HealthCheck(w http.ResponseWriter, r *http.Request) {
	response := map[string]interface{}{
		"status":    "healthy",
		"service":   "devsMailGo API",
		"version":   "1.0.0",
		"timestamp": "2024-01-01T00:00:00Z",
	}
	utils.RespondSuccess(w, response)
} 