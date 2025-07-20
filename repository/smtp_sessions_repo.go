package repository

import (
	"devsMailGo/config"
	"devsMailGo/models"
)

func GetAllSMTPSessions() ([]models.SMTPSession, error) {
	var entries []models.SMTPSession
	err := config.DB.Table("smtp_sessions").Find(&entries).Error
	return entries, err
} 