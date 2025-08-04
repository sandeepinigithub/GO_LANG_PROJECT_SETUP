package service

import (
	"bufio"
	"fmt"
	"os/exec"
	"strings"
	"devsMailGo/api/dto"
)

type MailService struct{}

// MailboxInfo represents mailbox information
type MailboxInfo struct {
	Username    string
	Domain      string
	MailboxPath string
	Quota       int64
	UsedQuota   int64
	MessageCount int
	LastLogin   string
}

// PostfixConfig represents Postfix configuration
type PostfixConfig struct {
	MainConfigPath string
	VirtualAliasPath string
	VirtualDomainPath string
}

// DovecotConfig represents Dovecot configuration
type DovecotConfig struct {
	ConfigPath string
	UserDBPath string
	QuotaPath  string
}

// CreateMailbox creates a new mailbox for a user
func (s *MailService) CreateMailbox(user *dto.UserResponse) error {
	// Create mailbox directory
	mailboxPath := fmt.Sprintf("/var/mail/vhosts/%s/%s", user.Domain, strings.Split(user.Email, "@")[0])
	
	cmd := exec.Command("mkdir", "-p", mailboxPath)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to create mailbox directory: %w", err)
	}

	// Set proper permissions
	cmd = exec.Command("chown", "-R", "vmail:vmail", mailboxPath)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to set mailbox permissions: %w", err)
	}

	// Update Postfix virtual domains if needed
	if err := s.updatePostfixVirtualDomains(user.Domain); err != nil {
		return fmt.Errorf("failed to update Postfix virtual domains: %w", err)
	}

	// Update Dovecot user database if needed
	if err := s.updateDovecotUserDB(user); err != nil {
		return fmt.Errorf("failed to update Dovecot user database: %w", err)
	}

	return nil
}

// DeleteMailbox deletes a mailbox for a user
func (s *MailService) DeleteMailbox(user *dto.UserResponse) error {
	mailboxPath := fmt.Sprintf("/var/mail/vhosts/%s/%s", user.Domain, strings.Split(user.Email, "@")[0])
	
	cmd := exec.Command("rm", "-rf", mailboxPath)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to delete mailbox directory: %w", err)
	}

	// Remove from Postfix virtual domains
	if err := s.removeFromPostfixVirtualDomains(user.Domain); err != nil {
		return fmt.Errorf("failed to remove from Postfix virtual domains: %w", err)
	}

	// Remove from Dovecot user database
	if err := s.removeFromDovecotUserDB(user.Email); err != nil {
		return fmt.Errorf("failed to remove from Dovecot user database: %w", err)
	}

	return nil
}

// GetMailboxInfo retrieves mailbox information
func (s *MailService) GetMailboxInfo(email string) (*MailboxInfo, error) {
	parts := strings.Split(email, "@")
	if len(parts) != 2 {
		return nil, fmt.Errorf("invalid email format")
	}

	username := parts[0]
	domain := parts[1]
	mailboxPath := fmt.Sprintf("/var/mail/vhosts/%s/%s", domain, username)

	info := &MailboxInfo{
		Username:    username,
		Domain:      domain,
		MailboxPath: mailboxPath,
	}

	// Get quota information
	if quota, err := s.getQuotaInfo(email); err == nil {
		info.Quota = int64(quota.Bytes)
		info.UsedQuota = int64(quota.Bytes) // TODO: Calculate actual used quota
	}

	// Get message count
	if count, err := s.getMessageCount(mailboxPath); err == nil {
		info.MessageCount = count
	}

	// Get last login
	if lastLogin, err := s.getLastLogin(email); err == nil {
		info.LastLogin = lastLogin
	}

	return info, nil
}

// UpdateQuota updates user quota
func (s *MailService) UpdateQuota(email string, quota int64) error {
	// Update quota in database
	// This would integrate with your existing quota service
	
	// Update Dovecot quota configuration
	return s.updateDovecotQuota(email, quota)
}

// ReloadPostfix reloads Postfix configuration
func (s *MailService) ReloadPostfix() error {
	cmd := exec.Command("postfix", "reload")
	return cmd.Run()
}

// ReloadDovecot reloads Dovecot configuration
func (s *MailService) ReloadDovecot() error {
	cmd := exec.Command("systemctl", "reload", "dovecot")
	return cmd.Run()
}

// updatePostfixVirtualDomains adds domain to Postfix virtual domains
func (s *MailService) updatePostfixVirtualDomains(domain string) error {
	// This is a simplified implementation
	// In production, you'd want to properly manage the virtual_domains file
	
	cmd := exec.Command("postconf", "-e", fmt.Sprintf("virtual_domains=%s", domain))
	return cmd.Run()
}

// removeFromPostfixVirtualDomains removes domain from Postfix virtual domains
func (s *MailService) removeFromPostfixVirtualDomains(domain string) error {
	// Implementation would depend on your Postfix configuration
	// This is a placeholder
	return nil
}

// updateDovecotUserDB adds user to Dovecot user database
func (s *MailService) updateDovecotUserDB(user *dto.UserResponse) error {
	// This would update Dovecot's user database
	// Implementation depends on your Dovecot configuration (passwd, SQL, etc.)
	return nil
}

// removeFromDovecotUserDB removes user from Dovecot user database
func (s *MailService) removeFromDovecotUserDB(email string) error {
	// Implementation depends on your Dovecot configuration
	return nil
}

// getQuotaInfo gets quota information for a user
func (s *MailService) getQuotaInfo(email string) (*dto.QuotaResponse, error) {
	// This would integrate with your existing quota service
	// For now, return placeholder data
	return &dto.QuotaResponse{
		Username: email,
		Bytes:    1024 * 1024 * 100, // 100MB
		Messages: 1000,
		Domain:   extractDomainFromEmail(email),
	}, nil
}

// getMessageCount gets the number of messages in a mailbox
func (s *MailService) getMessageCount(mailboxPath string) (int, error) {
	cmd := exec.Command("find", mailboxPath, "-name", "*.eml", "-type", "f")
	output, err := cmd.Output()
	if err != nil {
		return 0, err
	}

	lines := strings.Split(strings.TrimSpace(string(output)), "\n")
	if len(lines) == 1 && lines[0] == "" {
		return 0, nil
	}
	return len(lines), nil
}

// getLastLogin gets the last login time for a user
func (s *MailService) getLastLogin(email string) (string, error) {
	// This would check Dovecot logs or a separate login tracking system
	// For now, return placeholder
	return "2024-01-01 12:00:00", nil
}

// updateDovecotQuota updates quota in Dovecot configuration
func (s *MailService) updateDovecotQuota(email string, quota int64) error {
	// Implementation depends on your Dovecot quota configuration
	// This is a placeholder
	return nil
}

// CheckMailDelivery checks if mail delivery is working
func (s *MailService) CheckMailDelivery() error {
	// Check if we're on Windows or Linux
	// For Windows, we'll return a success message since mail services may not be available
	// In production, you'd implement proper Windows mail service checks
	
	// Try to check if any mail service is available
	cmd := exec.Command("postfix", "check")
	if err := cmd.Run(); err != nil {
		// On Windows, this is expected to fail
		return fmt.Errorf("Mail service check not available on Windows: %w", err)
	}

	return nil
}

// GetMailQueue gets the current mail queue
func (s *MailService) GetMailQueue() ([]string, error) {
	// For Windows, mail queue may not be available
	// Return an empty queue for now
	cmd := exec.Command("mailq")
	output, err := cmd.Output()
	if err != nil {
		// On Windows, this is expected to fail
		return []string{}, fmt.Errorf("Mail queue not available on Windows: %w", err)
	}

	var queue []string
	scanner := bufio.NewScanner(strings.NewReader(string(output)))
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line != "" {
			queue = append(queue, line)
		}
	}

	return queue, nil
}

// extractDomainFromEmail extracts domain from email
func extractDomainFromEmail(email string) string {
	parts := strings.Split(email, "@")
	if len(parts) == 2 {
		return parts[1]
	}
	return ""
} 