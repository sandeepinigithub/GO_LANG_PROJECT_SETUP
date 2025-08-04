package service

import (
	"fmt"
	"devsMailGo/config"
	"devsMailGo/api/dto"
	"gopkg.in/ldap.v3"
	"strings"
)

type LDAPService struct{}

// LDAPUser represents a user from LDAP
type LDAPUser struct {
	DN       string
	UID      string
	Email    string
	Name     string
	Groups   []string
	IsActive bool
}

// AuthenticateUser authenticates a user against LDAP
func (s *LDAPService) AuthenticateUser(username, password string) (*LDAPUser, error) {
	// Connect to LDAP server
	conn, err := s.ConnectToLDAP()
	if err != nil {
		return nil, fmt.Errorf("failed to connect to LDAP: %w", err)
	}
	defer conn.Close()

	// Search for user
	user, err := s.searchUser(conn, username)
	if err != nil {
		return nil, fmt.Errorf("user not found: %w", err)
	}

	// Attempt to bind with user credentials
	err = conn.Bind(user.DN, password)
	if err != nil {
		return nil, fmt.Errorf("invalid credentials: %w", err)
	}

	// Get user groups
	groups, err := s.getUserGroups(conn, username)
	if err != nil {
		// Log error but don't fail authentication
		fmt.Printf("Warning: failed to get user groups: %v\n", err)
	}

	return &LDAPUser{
		DN:       user.DN,
		UID:      user.UID,
		Email:    user.Email,
		Name:     user.Name,
		Groups:   groups,
		IsActive: true,
	}, nil
}

// SearchUser searches for a user in LDAP
func (s *LDAPService) SearchUser(username string) (*LDAPUser, error) {
	conn, err := s.ConnectToLDAP()
	if err != nil {
		return nil, fmt.Errorf("failed to connect to LDAP: %w", err)
	}
	defer conn.Close()

	return s.searchUser(conn, username)
}

// GetUserGroups retrieves groups for a user
func (s *LDAPService) GetUserGroups(username string) ([]string, error) {
	conn, err := s.ConnectToLDAP()
	if err != nil {
		return nil, fmt.Errorf("failed to connect to LDAP: %w", err)
	}
	defer conn.Close()

	return s.getUserGroups(conn, username)
}

// ConnectToLDAP establishes connection to LDAP server (public method)
func (s *LDAPService) ConnectToLDAP() (*ldap.Conn, error) {
	ldapConfig := config.AppConfig.LDAP
	
	var conn *ldap.Conn
	var err error
	
	// Build LDAP URL
	ldapURL := fmt.Sprintf("%s:%s", ldapConfig.Host, ldapConfig.Port)
	
	if ldapConfig.UseSSL {
		conn, err = ldap.DialTLS("tcp", ldapURL, nil)
	} else {
		conn, err = ldap.Dial("tcp", ldapURL)
	}
	
	if err != nil {
		return nil, err
	}

	// Start TLS if configured
	if ldapConfig.UseTLS {
		err = conn.StartTLS(nil)
		if err != nil {
			conn.Close()
			return nil, err
		}
	}

	// Bind with service account
	err = conn.Bind(ldapConfig.BindDN, ldapConfig.BindPassword)
	if err != nil {
		conn.Close()
		return nil, err
	}

	return conn, nil
}

// searchUser searches for a user in LDAP
func (s *LDAPService) searchUser(conn *ldap.Conn, username string) (*LDAPUser, error) {
	ldapConfig := config.AppConfig.LDAP
	
	// Build search filter
	searchFilter := fmt.Sprintf(ldapConfig.UserFilter, username)
	
	// Search for user
	searchRequest := ldap.NewSearchRequest(
		ldapConfig.BaseDN,
		ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
		searchFilter,
		[]string{"uid", "mail", "cn", "sn", "givenName", "userPassword", "accountStatus"},
		nil,
	)

	result, err := conn.Search(searchRequest)
	if err != nil {
		return nil, err
	}

	if len(result.Entries) == 0 {
		return nil, fmt.Errorf("user not found")
	}

	entry := result.Entries[0]
	
	// Extract user attributes
	user := &LDAPUser{
		DN:  entry.DN,
		UID: entry.GetAttributeValue("uid"),
	}

	// Get email
	if emails := entry.GetAttributeValues("mail"); len(emails) > 0 {
		user.Email = emails[0]
	}

	// Get name
	if cn := entry.GetAttributeValue("cn"); cn != "" {
		user.Name = cn
	} else {
		// Fallback to givenName + sn
		givenName := entry.GetAttributeValue("givenName")
		sn := entry.GetAttributeValue("sn")
		if givenName != "" && sn != "" {
			user.Name = givenName + " " + sn
		}
	}

	// Check if account is active
	accountStatus := entry.GetAttributeValue("accountStatus")
	user.IsActive = accountStatus != "inactive" && accountStatus != "locked"

	return user, nil
}

// getUserGroups retrieves groups for a user
func (s *LDAPService) getUserGroups(conn *ldap.Conn, username string) ([]string, error) {
	ldapConfig := config.AppConfig.LDAP
	
	// Build search filter for groups
	searchFilter := fmt.Sprintf(ldapConfig.GroupFilter, username)
	
	// Search for groups
	searchRequest := ldap.NewSearchRequest(
		ldapConfig.BaseDN,
		ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
		searchFilter,
		[]string{"cn", "memberUid"},
		nil,
	)

	result, err := conn.Search(searchRequest)
	if err != nil {
		return nil, err
	}

	var groups []string
	for _, entry := range result.Entries {
		if cn := entry.GetAttributeValue("cn"); cn != "" {
			groups = append(groups, cn)
		}
	}

	return groups, nil
}

// SyncUserToDatabase syncs LDAP user to local database
func (s *LDAPService) SyncUserToDatabase(ldapUser *LDAPUser) (*dto.UserResponse, error) {
	// This would integrate with your existing user service
	// to create/update local user records based on LDAP data
	// userService := &UserService{} // TODO: Implement user sync
	
	// Check if user exists in database
	// If not, create new user
	// If exists, update user information
	
	// For now, return a basic response
	return &dto.UserResponse{
		Email:  ldapUser.Email,
		Name:   ldapUser.Name,
		Domain: extractDomain(ldapUser.Email),
	}, nil
}

// extractDomain extracts domain from email
func extractDomain(email string) string {
	parts := strings.Split(email, "@")
	if len(parts) == 2 {
		return parts[1]
	}
	return ""
} 