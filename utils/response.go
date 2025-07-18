package utils

import (
	"encoding/json"
	"net/http"
)

func RespondSuccess(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"_success": true,
		"_data":    data,
	})
}

func RespondSuccessNoData(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"_success": true,
	})
}

func RespondError(w http.ResponseWriter, msg string) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"_success": false,
		"_msg":     msg,
	})
} 