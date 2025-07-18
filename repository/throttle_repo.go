package repository

import (
	"GO_LANG_PROJECT_SETUP/config"
	"GO_LANG_PROJECT_SETUP/models"
)

func GetAllThrottle() ([]models.Throttle, error) {
	var entries []models.Throttle
	err := config.DB.Table("throttle").Find(&entries).Error
	return entries, err
}

func GetThrottleByID(id uint64) (models.Throttle, error) {
	var entry models.Throttle
	err := config.DB.Table("throttle").Where("id = ?", id).First(&entry).Error
	return entry, err
} 