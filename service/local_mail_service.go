package service

import (
	"crypto/tls"
	"fmt"
	"net"
	"net/smtp"
	"strconv"
	"time"

	"devsMailGo/config"
)

type LocalMailService struct {
	config *config.MailConfig
}

// NewLocalMailService creates a new local mail service instance
func NewLocalMailService() *LocalMailService {
	return &LocalMailService{
		config: &config.AppConfig.Mail,
	}
}

// SendEmail sends an email using SMTP
func (s *LocalMailService) SendEmail(from, to, subject, body string) error {
	return s.sendViaSMTP(from, to, subject, body)
}

// sendViaSMTP sends email via SMTP
func (s *LocalMailService) sendViaSMTP(from, to, subject, body string) error {
	// Create email message with proper headers
	message := fmt.Sprintf("From: %s\r\nTo: %s\r\nSubject: %s\r\nContent-Type: text/plain; charset=UTF-8\r\n\r\n%s", 
		from, to, subject, body)
	
	// Convert port to int for proper handling
	port, err := strconv.Atoi(s.config.SMTPPort)
	if err != nil {
		return fmt.Errorf("invalid SMTP port: %w", err)
	}
	
	// SMTP authentication
	var auth smtp.Auth
	if s.config.SMTPUsername != "" && s.config.SMTPPassword != "" {
		auth = smtp.PlainAuth("", s.config.SMTPUsername, s.config.SMTPPassword, s.config.SMTPHost)
	}
	
	// Handle different connection types based on port and TLS setting
	if port == 587 && !s.config.UseTLS {
		// STARTTLS on port 587 (like your JavaScript config)
		return s.sendWithSTARTTLS(from, to, message, auth)
	} else if port == 465 && s.config.UseTLS {
		// SSL/TLS on port 465
		return s.sendWithSSL(from, to, message, auth)
	} else {
		// Plain connection (fallback)
		addr := fmt.Sprintf("%s:%s", s.config.SMTPHost, s.config.SMTPPort)
		if err := smtp.SendMail(addr, auth, from, []string{to}, []byte(message)); err != nil {
			return fmt.Errorf("failed to send email via SMTP: %w", err)
		}
	}
	
	return nil
}

// sendWithSTARTTLS sends email using STARTTLS (port 587)
func (s *LocalMailService) sendWithSTARTTLS(from, to, message string, auth smtp.Auth) error {
	host := s.config.SMTPHost
	addr := fmt.Sprintf("%s:587", host)
	
	// Connect to server
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return fmt.Errorf("failed to connect to SMTP server: %w", err)
	}
	defer conn.Close()
	
	// Create SMTP client
	client, err := smtp.NewClient(conn, host)
	if err != nil {
		return fmt.Errorf("failed to create SMTP client: %w", err)
	}
	defer client.Close()
	
	// Start TLS with insecure skip verify
	tlsConfig := &tls.Config{
		ServerName:         host,
		InsecureSkipVerify: true,
	}
	
	if err = client.StartTLS(tlsConfig); err != nil {
		return fmt.Errorf("failed to start TLS: %w", err)
	}
	
	// Authenticate if credentials provided
	if auth != nil {
		if err = client.Auth(auth); err != nil {
			return fmt.Errorf("failed to authenticate: %w", err)
		}
	}
	
	// Send email
	if err = client.Mail(from); err != nil {
		return fmt.Errorf("failed to set sender: %w", err)
	}
	
	if err = client.Rcpt(to); err != nil {
		return fmt.Errorf("failed to set recipient: %w", err)
	}
	
	w, err := client.Data()
	if err != nil {
		return fmt.Errorf("failed to start data transfer: %w", err)
	}
	
	_, err = w.Write([]byte(message))
	if err != nil {
		return fmt.Errorf("failed to write message: %w", err)
	}
	
	err = w.Close()
	if err != nil {
		return fmt.Errorf("failed to close data transfer: %w", err)
	}
	
	return nil
}

// sendWithSSL sends email using SSL/TLS (port 465)
func (s *LocalMailService) sendWithSSL(from, to, message string, auth smtp.Auth) error {
	host := s.config.SMTPHost
	addr := fmt.Sprintf("%s:465", host)
	
	// Connect with TLS
	tlsConfig := &tls.Config{
		ServerName:         host,
		InsecureSkipVerify: true,
	}
	
	conn, err := tls.Dial("tcp", addr, tlsConfig)
	if err != nil {
		return fmt.Errorf("failed to connect to SMTP server with TLS: %w", err)
	}
	defer conn.Close()
	
	// Create SMTP client
	client, err := smtp.NewClient(conn, host)
	if err != nil {
		return fmt.Errorf("failed to create SMTP client: %w", err)
	}
	defer client.Close()
	
	// Authenticate if credentials provided
	if auth != nil {
		if err = client.Auth(auth); err != nil {
			return fmt.Errorf("failed to authenticate: %w", err)
		}
	}
	
	// Send email
	if err = client.Mail(from); err != nil {
		return fmt.Errorf("failed to set sender: %w", err)
	}
	
	if err = client.Rcpt(to); err != nil {
		return fmt.Errorf("failed to set recipient: %w", err)
	}
	
	w, err := client.Data()
	if err != nil {
		return fmt.Errorf("failed to start data transfer: %w", err)
	}
	
	_, err = w.Write([]byte(message))
	if err != nil {
		return fmt.Errorf("failed to write message: %w", err)
	}
	
	err = w.Close()
	if err != nil {
		return fmt.Errorf("failed to close data transfer: %w", err)
	}
	
	return nil
}

// CheckMailServices checks if mail services are available
func (s *LocalMailService) CheckMailServices() (map[string]interface{}, error) {
	results := make(map[string]interface{})
	
	// Check SMTP connectivity
	results["smtp"] = s.checkSMTPConnectivity()
	
	return results, nil
}

// checkSMTPConnectivity checks SMTP connectivity
func (s *LocalMailService) checkSMTPConnectivity() map[string]interface{} {
	// Try to connect to SMTP server
	addr := fmt.Sprintf("%s:%s", s.config.SMTPHost, s.config.SMTPPort)
	
	// Simple TCP connection test
	conn, err := net.DialTimeout("tcp", addr, 5*time.Second)
	if err != nil {
		return map[string]interface{}{
			"success": false,
			"error":   fmt.Sprintf("SMTP connection failed: %v", err),
		}
	}
	defer conn.Close()
	
	return map[string]interface{}{
		"success": true,
		"message": fmt.Sprintf("SMTP connection successful to %s", addr),
	}
}

// GetMailQueue gets the current mail queue status
func (s *LocalMailService) GetMailQueue() ([]string, error) {
	// For SMTP-only setup, return connection status
	status := s.checkSMTPConnectivity()
	if status["success"].(bool) {
		return []string{"SMTP connection is healthy"}, nil
	} else {
		return []string{fmt.Sprintf("SMTP connection failed: %v", status["error"])}, nil
	}
} 