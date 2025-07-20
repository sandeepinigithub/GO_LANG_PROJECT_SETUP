package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"GO_LANG_PROJECT_SETUP/routes"
	"GO_LANG_PROJECT_SETUP/utils"
	"github.com/gorilla/mux"
)

// TestSetup represents test setup
type TestSetup struct {
	Router *mux.Router
	Token  string
}

// SetupTest creates a test setup
func SetupTest() *TestSetup {
	router := routes.SetupRoutes()
	return &TestSetup{
		Router: router,
	}
}

// TestLogin tests the login endpoint
func TestLogin(t *testing.T) {
	setup := SetupTest()
	
	// Test valid login
	loginData := map[string]interface{}{
		"username": "admin@example.com",
		"password": "password",
	}
	
	jsonData, _ := json.Marshal(loginData)
	req := httptest.NewRequest("POST", "/api/login", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	
	w := httptest.NewRecorder()
	setup.Router.ServeHTTP(w, req)
	
	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}
	
	var response map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &response)
	
	if !response["success"].(bool) {
		t.Error("Expected successful login")
	}
	
	// Test invalid login
	invalidData := map[string]interface{}{
		"username": "invalid@example.com",
		"password": "wrongpassword",
	}
	
	jsonData, _ = json.Marshal(invalidData)
	req = httptest.NewRequest("POST", "/api/login", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	
	w = httptest.NewRecorder()
	setup.Router.ServeHTTP(w, req)
	
	if w.Code != http.StatusUnauthorized {
		t.Errorf("Expected status 401, got %d", w.Code)
	}
}

// TestProtectedEndpoint tests a protected endpoint
func TestProtectedEndpoint(t *testing.T) {
	setup := SetupTest()
	
	// Test without authentication
	req := httptest.NewRequest("GET", "/api/users", nil)
	w := httptest.NewRecorder()
	setup.Router.ServeHTTP(w, req)
	
	if w.Code != http.StatusUnauthorized {
		t.Errorf("Expected status 401, got %d", w.Code)
	}
	
	// Test with invalid token
	req = httptest.NewRequest("GET", "/api/users", nil)
	req.Header.Set("Authorization", "Bearer invalid-token")
	w = httptest.NewRecorder()
	setup.Router.ServeHTTP(w, req)
	
	if w.Code != http.StatusUnauthorized {
		t.Errorf("Expected status 401, got %d", w.Code)
	}
}

// TestHealthCheck tests the health check endpoint
func TestHealthCheck(t *testing.T) {
	setup := SetupTest()
	
	req := httptest.NewRequest("GET", "/api/health", nil)
	w := httptest.NewRecorder()
	setup.Router.ServeHTTP(w, req)
	
	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}
	
	var response map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &response)
	
	if !response["success"].(bool) {
		t.Error("Expected successful health check")
	}
} 