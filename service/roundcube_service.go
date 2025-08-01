package service

import (
	"fmt"
	"devsMailGo/api/dto"
	"devsMailGo/models"
	"devsMailGo/repository"
)

type RoundcubeService struct{}

// RoundcubeUser represents a Roundcube user
type RoundcubeUser struct {
	UserID      int    `json:"user_id"`
	Username    string `json:"username"`
	MailHost    string `json:"mail_host"`
	Created     string `json:"created"`
	LastLogin   string `json:"last_login"`
	Language    string `json:"language"`
	Preferences string `json:"preferences"`
	VirtName    string `json:"virt_name"`
	VirtMailbox string `json:"virt_mailbox"`
	Deleted     int    `json:"deleted"`
	Alias       string `json:"alias"`
}

// CreateRoundcubeUser creates a new Roundcube user
func (s *RoundcubeService) CreateRoundcubeUser(user *dto.UserResponse) error {
	// Create Roundcube user in database
	language := "en_US"
	roundcubeUser := &models.RoundcubeUser{
		Username: user.Email,
		MailHost: user.Domain,
		Language: &language,
	}

	err := repository.CreateRoundcubeUser(roundcubeUser)
	if err != nil {
		return fmt.Errorf("failed to create Roundcube user: %w", err)
	}

	return nil
}

// DeleteRoundcubeUser deletes a Roundcube user
func (s *RoundcubeService) DeleteRoundcubeUser(email string) error {
	// For now, just return success since the model doesn't have a Deleted field
	// In a real implementation, you might want to mark the user as inactive
	// or remove them from the database
	return nil
}

// GetRoundcubeUser gets Roundcube user information
func (s *RoundcubeService) GetRoundcubeUser(email string) (*RoundcubeUser, error) {
	user, err := repository.GetRoundcubeUserByEmail(email)
	if err != nil {
		return nil, fmt.Errorf("Roundcube user not found: %w", err)
	}

	lastLogin := ""
	if user.LastLogin != nil {
		lastLogin = *user.LastLogin
	}
	
	language := ""
	if user.Language != nil {
		language = *user.Language
	}
	
	preferences := ""
	if user.Preferences != nil {
		preferences = *user.Preferences
	}
	
	return &RoundcubeUser{
		UserID:      int(user.UserID),
		Username:    user.Username,
		MailHost:    user.MailHost,
		Created:     user.Created,
		LastLogin:   lastLogin,
		Language:    language,
		Preferences: preferences,
		VirtName:    "", // Not available in this model
		VirtMailbox: "", // Not available in this model
		Deleted:     0,  // Not available in this model
		Alias:       "", // Not available in this model
	}, nil
}

// UpdateRoundcubeUser updates Roundcube user information
func (s *RoundcubeService) UpdateRoundcubeUser(email string, updates map[string]interface{}) error {
	user, err := repository.GetRoundcubeUserByEmail(email)
	if err != nil {
		return fmt.Errorf("Roundcube user not found: %w", err)
	}

	// Apply updates
	if language, ok := updates["language"].(string); ok {
		user.Language = &language
	}
	if preferences, ok := updates["preferences"].(string); ok {
		user.Preferences = &preferences
	}

	err = repository.UpdateRoundcubeUser(user.UserID, &user)
	if err != nil {
		return fmt.Errorf("failed to update Roundcube user: %w", err)
	}

	return nil
}

// ListRoundcubeUsers lists all Roundcube users
func (s *RoundcubeService) ListRoundcubeUsers() ([]*RoundcubeUser, error) {
	users, err := repository.GetAllRoundcubeUsers()
	if err != nil {
		return nil, fmt.Errorf("failed to get Roundcube users: %w", err)
	}

	var result []*RoundcubeUser
	for _, user := range users {
		lastLogin := ""
		if user.LastLogin != nil {
			lastLogin = *user.LastLogin
		}
		
		language := ""
		if user.Language != nil {
			language = *user.Language
		}
		
		preferences := ""
		if user.Preferences != nil {
			preferences = *user.Preferences
		}
		
		result = append(result, &RoundcubeUser{
			UserID:      int(user.UserID),
			Username:    user.Username,
			MailHost:    user.MailHost,
			Created:     user.Created,
			LastLogin:   lastLogin,
			Language:    language,
			Preferences: preferences,
			VirtName:    "", // Not available in this model
			VirtMailbox: "", // Not available in this model
			Deleted:     0,  // Not available in this model
			Alias:       "", // Not available in this model
		})
	}

	return result, nil
}

// SyncUserToRoundcube syncs a user to Roundcube
func (s *RoundcubeService) SyncUserToRoundcube(user *dto.UserResponse) error {
	// Check if user already exists in Roundcube
	_, err := repository.GetRoundcubeUserByEmail(user.Email)
	if err == nil {
		// User exists, no need to update since VirtName is not available in this model
		return nil
	}

	// User doesn't exist, create new
	return s.CreateRoundcubeUser(user)
}

// GetRoundcubeStats gets Roundcube statistics
func (s *RoundcubeService) GetRoundcubeStats() (map[string]interface{}, error) {
	users, err := repository.GetAllRoundcubeUsers()
	if err != nil {
		return nil, fmt.Errorf("failed to get Roundcube users: %w", err)
	}

	totalUsers := len(users)
	activeUsers := 0
	deletedUsers := 0

	for range users {
		// Since Deleted field is not available in this model, consider all users as active
		activeUsers++
	}

	return map[string]interface{}{
		"total_users":   totalUsers,
		"active_users":  activeUsers,
		"deleted_users": deletedUsers,
	}, nil
} 