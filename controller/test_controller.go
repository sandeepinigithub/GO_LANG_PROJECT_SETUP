package controller

import (
	"devsMailGo/config"
	"devsMailGo/service"
	"devsMailGo/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"os/exec"
	"strings"
	// "bufio"
)

var testLDAPService = service.LDAPService{}
var testLocalMailService *service.LocalMailService

// getLocalMailService returns the local mail service instance, initializing it if needed
func getLocalMailService() *service.LocalMailService {
	if testLocalMailService == nil {
		// Ensure config is loaded before creating service
		if config.AppConfig == nil {
			if err := config.LoadConfig(); err != nil {
				panic(fmt.Sprintf("Failed to load config: %v", err))
			}
		}
		testLocalMailService = service.NewLocalMailService()
	}
	return testLocalMailService
}

// TestLDAPRequest represents the LDAP test request
type TestLDAPRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// TestMailRequest represents the mail test request
type TestMailRequest struct {
	From    string `json:"from"`
	To      string `json:"to"`
	Subject string `json:"subject"`
	Body    string `json:"body"`
}

// TestResponse represents the test response
type TestResponse struct {
	Success bool                   `json:"success"`
	Message string                 `json:"message"`
	Details map[string]interface{} `json:"details,omitempty"`
}

// TestLDAPConnection tests LDAP connectivity and authentication
func TestLDAPConnection(w http.ResponseWriter, r *http.Request) {
	var req TestLDAPRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	// Validate required fields
	if req.Username == "" || req.Password == "" {
		utils.RespondWithError(w, http.StatusBadRequest, "Username and password are required")
		return
	}

	response := TestResponse{
		Success: false,
		Details: make(map[string]interface{}),
	}

	// Test 1: LDAP Connection
	ldapConfig := config.AppConfig.LDAP
	response.Details["ldap_config"] = map[string]interface{}{
		"host":         ldapConfig.Host,
		"port":         ldapConfig.Port,
		"base_dn":      ldapConfig.BaseDN,
		"bind_dn":      ldapConfig.BindDN,
		"use_ssl":      ldapConfig.UseSSL,
		"use_tls":      ldapConfig.UseTLS,
		"user_filter":  ldapConfig.UserFilter,
		"group_filter": ldapConfig.GroupFilter,
	}

	// Test 2: LDAP Authentication
	user, err := testLDAPService.AuthenticateUser(req.Username, req.Password)
	if err != nil {
		response.Message = "LDAP authentication failed: " + err.Error()
		response.Details["authentication"] = map[string]interface{}{
			"success": false,
			"error":   err.Error(),
		}
	} else {
		response.Success = true
		response.Message = "LDAP authentication successful"
		response.Details["authentication"] = map[string]interface{}{
			"success": true,
			"user": map[string]interface{}{
				"dn":        user.DN,
				"uid":       user.UID,
				"email":     user.Email,
				"name":      user.Name,
				"groups":    user.Groups,
				"is_active": user.IsActive,
			},
		}
	}

	// Test 3: LDAP Search
	if searchUser, err := testLDAPService.SearchUser(req.Username); err != nil {
		response.Details["search"] = map[string]interface{}{
			"success": false,
			"error":   err.Error(),
		}
	} else {
		response.Details["search"] = map[string]interface{}{
			"success": true,
			"user": map[string]interface{}{
				"dn":        searchUser.DN,
				"uid":       searchUser.UID,
				"email":     searchUser.Email,
				"name":      searchUser.Name,
				"is_active": searchUser.IsActive,
			},
		}
	}

	// Test 4: LDAP Groups
	if groups, err := testLDAPService.GetUserGroups(req.Username); err != nil {
		response.Details["groups"] = map[string]interface{}{
			"success": false,
			"error":   err.Error(),
		}
	} else {
		response.Details["groups"] = map[string]interface{}{
			"success": true,
			"groups":  groups,
		}
	}

	if response.Success {
		utils.RespondSuccess(w, response)
	} else {
		utils.RespondWithError(w, http.StatusUnauthorized, response.Message)
	}
}

// TestLDAPConnectionSimple tests basic LDAP connectivity without authentication
func TestLDAPConnectionSimple(w http.ResponseWriter, r *http.Request) {
	response := TestResponse{
		Success: false,
		Details: make(map[string]interface{}),
	}

	// Get LDAP configuration
	ldapConfig := config.AppConfig.LDAP
	response.Details["ldap_config"] = map[string]interface{}{
		"host":         ldapConfig.Host,
		"port":         ldapConfig.Port,
		"base_dn":      ldapConfig.BaseDN,
		"bind_dn":      ldapConfig.BindDN,
		"use_ssl":      ldapConfig.UseSSL,
		"use_tls":      ldapConfig.UseTLS,
		"user_filter":  ldapConfig.UserFilter,
		"group_filter": ldapConfig.GroupFilter,
	}

	// Test basic connection
	conn, err := testLDAPService.ConnectToLDAP()
	if err != nil {
		response.Message = "LDAP connection failed: " + err.Error()
		response.Details["connection"] = map[string]interface{}{
			"success": false,
			"error":   err.Error(),
		}
	} else {
		defer conn.Close()
		response.Success = true
		response.Message = "LDAP connection successful"
		response.Details["connection"] = map[string]interface{}{
			"success": true,
			"message": "Connected to LDAP server",
		}
	}

	if response.Success {
		utils.RespondSuccess(w, response)
	} else {
		utils.RespondWithError(w, http.StatusInternalServerError, response.Message)
	}
}

// TestMailSending tests mail sending functionality
func TestMailSending(w http.ResponseWriter, r *http.Request) {
	var req TestMailRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	// Set defaults if not provided
	if req.From == "" {
		req.From = "sandeep@dev.local"
	}
	if req.To == "" {
		req.To = "sandeep@dev.local"
	}
	if req.Subject == "" {
		req.Subject = "Test Mail from API"
	}
	if req.Body == "" {
		req.Body = "This is a test email sent from the API to verify mail functionality."
	}

	response := TestResponse{
		Success: false,
		Details: make(map[string]interface{}),
	}

	// Test 1: Check platform and available services
	response.Details["platform"] = map[string]interface{}{
		"os":      "windows",
		"message": "Running on Windows platform",
	}

	// Test 2: Check local mail services first
	localMailServices, err := getLocalMailService().CheckMailServices()
	if err != nil {
		response.Details["local_mail_services"] = map[string]interface{}{
			"success": false,
			"error":   err.Error(),
		}
	} else {
		response.Details["local_mail_services"] = localMailServices
	}

	// Test 3: Send email via SMTP
	if err := getLocalMailService().SendEmail(req.From, req.To, req.Subject, req.Body); err == nil {
		response.Success = true
		response.Message = "Test email sent successfully via SMTP"
		response.Details["mail_sending"] = map[string]interface{}{
			"success": true,
			"method":  "smtp",
			"from":    req.From,
			"to":      req.To,
			"subject": req.Subject,
			"body":    req.Body,
		}
	} else {
		response.Message = "Mail sending failed"
		response.Details["mail_sending"] = map[string]interface{}{
			"success": false,
			"method":  "smtp",
			"error":   err.Error(),
			"from":    req.From,
			"to":      req.To,
			"subject": req.Subject,
			"body":    req.Body,
		}
	}

	// Add configuration information
	response.Details["configuration"] = map[string]interface{}{
		"mail_config": map[string]interface{}{
			"smtp_host":    config.AppConfig.Mail.SMTPHost,
			"smtp_port":    config.AppConfig.Mail.SMTPPort,
			"use_tls":      config.AppConfig.Mail.UseTLS,
			"from_address": config.AppConfig.Mail.FromAddress,
			"to_address":   config.AppConfig.Mail.ToAddress,
		},
	}

	if response.Success {
		utils.RespondSuccess(w, response)
	} else {
		// Ensure we have a message even if it's empty
		if response.Message == "" {
			response.Message = "Mail test failed - check details for specific errors"
		}
		utils.RespondWithError(w, http.StatusInternalServerError, response.Message)
	}
}

// TestSystemHealth provides comprehensive system health check
func TestSystemHealth(w http.ResponseWriter, r *http.Request) {
	response := TestResponse{
		Success: true,
		Message: "System health check completed",
		Details: make(map[string]interface{}),
	}

	// Test 1: Database connection
	if config.DB != nil {
		sqlDB, err := config.DB.DB()
		if err != nil {
			response.Details["database"] = map[string]interface{}{
				"success": false,
				"error":   err.Error(),
			}
		} else {
			if err := sqlDB.Ping(); err != nil {
				response.Details["database"] = map[string]interface{}{
					"success": false,
					"error":   err.Error(),
				}
			} else {
				response.Details["database"] = map[string]interface{}{
					"success": true,
					"message": "Database connection healthy",
				}
			}
		}
	} else {
		response.Details["database"] = map[string]interface{}{
			"success": false,
			"error":   "Database not initialized",
		}
	}

	// Test 2: LDAP connection (without authentication)
	ldapConfig := config.AppConfig.LDAP
	response.Details["ldap_config"] = map[string]interface{}{
		"host":    ldapConfig.Host,
		"port":    ldapConfig.Port,
		"base_dn": ldapConfig.BaseDN,
		"bind_dn": ldapConfig.BindDN,
		"use_ssl": ldapConfig.UseSSL,
		"use_tls": ldapConfig.UseTLS,
	}

	// Test 3: Mail services on Linux server
	// Postfix
	cmd := exec.Command("ssh", "root@172.105.49.219", "postfix", "status")
	if err := cmd.Run(); err != nil {
		response.Details["postfix"] = map[string]interface{}{
			"success": false,
			"error":   err.Error(),
			"note":    "Testing Postfix on Linux server 172.105.49.219",
		}
	} else {
		response.Details["postfix"] = map[string]interface{}{
			"success": true,
			"message": "Postfix is running on Linux server",
		}
	}

	// Dovecot
	cmd = exec.Command("ssh", "root@172.105.49.219", "systemctl", "is-active", "dovecot")
	if output, err := cmd.Output(); err != nil {
		response.Details["dovecot"] = map[string]interface{}{
			"success": false,
			"error":   err.Error(),
			"note":    "Testing Dovecot on Linux server 172.105.49.219",
		}
	} else {
		response.Details["dovecot"] = map[string]interface{}{
			"success": true,
			"status":  strings.TrimSpace(string(output)),
		}
	}

	// Test 4: System resources (Windows-compatible)
	// Check disk space using Windows command
	cmd = exec.Command("wmic", "logicaldisk", "get", "size,freespace,caption")
	if output, err := cmd.Output(); err != nil {
		response.Details["disk_space"] = map[string]interface{}{
			"success": false,
			"error":   err.Error(),
			"note":    "Using Windows disk space check",
		}
	} else {
		response.Details["disk_space"] = map[string]interface{}{
			"success": true,
			"output":  strings.TrimSpace(string(output)),
			"note":    "Windows disk space information",
		}
	}

	// Check memory usage using Windows command
	cmd = exec.Command("wmic", "computersystem", "get", "TotalPhysicalMemory")
	if output, err := cmd.Output(); err != nil {
		response.Details["memory"] = map[string]interface{}{
			"success": false,
			"error":   err.Error(),
			"note":    "Using Windows memory check",
		}
	} else {
		response.Details["memory"] = map[string]interface{}{
			"success": true,
			"output":  strings.TrimSpace(string(output)),
			"note":    "Windows memory information",
		}
	}

	utils.RespondSuccess(w, response)
}
