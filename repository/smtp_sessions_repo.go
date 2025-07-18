package repository

import (
	"GO_LANG_PROJECT_SETUP/config"
	"GO_LANG_PROJECT_SETUP/models"
)

func GetAllSMTPSessions() ([]models.SMTPSession, error) {
	var entries []models.SMTPSession
	err := config.DB.Table("smtp_sessions").Find(&entries).Error
	return entries, err
} 