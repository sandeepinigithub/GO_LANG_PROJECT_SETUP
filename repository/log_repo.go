package repository

import (
	"GO_LANG_PROJECT_SETUP/config"
	"GO_LANG_PROJECT_SETUP/models"
)

func GetAllLogs() ([]models.Log, error) {
	var logs []models.Log
	err := config.DB.Table("log").Find(&logs).Error
	return logs, err
}

func GetLogByID(id uint64) (models.Log, error) {
	var logEntry models.Log
	err := config.DB.Table("log").Where("id = ?", id).First(&logEntry).Error
	return logEntry, err
}

func CreateLog(entry *models.Log) error {
	return config.DB.Table("log").Create(entry).Error
}

func UpdateLog(id uint64, updated *models.Log) error {
	return config.DB.Table("log").Where("id = ?", id).Updates(updated).Error
}

func DeleteLog(id uint64) error {
	return config.DB.Table("log").Where("id = ?", id).Delete(&models.Log{}).Error
} 