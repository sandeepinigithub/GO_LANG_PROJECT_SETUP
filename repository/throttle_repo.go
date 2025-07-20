package repository

import (
	"devsMailGo/config"
	"devsMailGo/models"
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

func CreateThrottle(entry *models.Throttle) error {
	return config.DB.Table("throttle").Create(entry).Error
}

func UpdateThrottle(id uint64, updated *models.Throttle) error {
	return config.DB.Table("throttle").Where("id = ?", id).Updates(updated).Error
}

func DeleteThrottle(id uint64) error {
	return config.DB.Table("throttle").Where("id = ?", id).Delete(&models.Throttle{}).Error
} 