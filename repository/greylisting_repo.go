package repository

import (
	"GO_LANG_PROJECT_SETUP/config"
	"GO_LANG_PROJECT_SETUP/models"
)

func GetAllGreylisting() ([]models.Greylisting, error) {
	var entries []models.Greylisting
	err := config.DB.Table("greylisting").Find(&entries).Error
	return entries, err
}

func GetGreylistingByID(id uint64) (models.Greylisting, error) {
	var entry models.Greylisting
	err := config.DB.Table("greylisting").Where("id = ?", id).First(&entry).Error
	return entry, err
} 