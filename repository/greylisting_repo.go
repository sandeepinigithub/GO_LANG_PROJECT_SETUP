package repository

import (
	"devsMailGo/config"
	"devsMailGo/models"
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

func CreateGreylisting(entry *models.Greylisting) error {
	return config.DB.Table("greylisting").Create(entry).Error
}

func UpdateGreylisting(id uint64, updated *models.Greylisting) error {
	return config.DB.Table("greylisting").Where("id = ?", id).Updates(updated).Error
}

func DeleteGreylisting(id uint64) error {
	return config.DB.Table("greylisting").Where("id = ?", id).Delete(&models.Greylisting{}).Error
} 