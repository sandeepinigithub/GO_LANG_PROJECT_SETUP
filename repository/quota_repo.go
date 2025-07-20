package repository

import (
	"GO_LANG_PROJECT_SETUP/config"
	"GO_LANG_PROJECT_SETUP/models"
)

func GetAllUsedQuota() ([]models.UsedQuota, error) {
	var quotas []models.UsedQuota
	err := config.DB.Find(&quotas).Error
	return quotas, err
}

func GetUsedQuotaByUsername(username string) (models.UsedQuota, error) {
	var quota models.UsedQuota
	err := config.DB.Where("username = ?", username).First(&quota).Error
	return quota, err
}

func CreateUsedQuota(quota *models.UsedQuota) error {
	return config.DB.Create(quota).Error
}

func UpdateUsedQuota(username string, updated *models.UsedQuota) error {
	return config.DB.Model(&models.UsedQuota{}).Where("username = ?", username).Updates(updated).Error
}

func DeleteUsedQuota(username string) error {
	return config.DB.Where("username = ?", username).Delete(&models.UsedQuota{}).Error
} 