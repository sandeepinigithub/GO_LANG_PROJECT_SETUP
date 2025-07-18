package repository

import (
	"GO_LANG_PROJECT_SETUP/config"
	"GO_LANG_PROJECT_SETUP/models"
)

func GetAllBanned() ([]models.Banned, error) {
	var banned []models.Banned
	err := config.DB.Table("banned").Find(&banned).Error
	return banned, err
}

func GetBannedByIP(ip string) (models.Banned, error) {
	var banned models.Banned
	err := config.DB.Table("banned").Where("ip = ?", ip).First(&banned).Error
	return banned, err
}

func UnbanByID(id uint64) error {
	return config.DB.Table("banned").Where("id = ?", id).Delete(nil).Error
} 