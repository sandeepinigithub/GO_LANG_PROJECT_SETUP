package repository

import (
	"devsMailGo/config"
	"devsMailGo/models"
)

func CreateAlias(alias *models.Alias) error {
	return config.DB.Create(alias).Error
}

func GetAllAliases() ([]models.Alias, error) {
	var aliases []models.Alias
	err := config.DB.Find(&aliases).Error
	return aliases, err
}

func GetAliasByAddress(address string) (models.Alias, error) {
	var alias models.Alias
	err := config.DB.Where("address = ?", address).First(&alias).Error
	return alias, err
}

func UpdateAlias(address string, updated *models.Alias) error {
	return config.DB.Model(&models.Alias{}).Where("address = ?", address).Updates(updated).Error
}

func DeleteAlias(address string) error {
	return config.DB.Where("address = ?", address).Delete(&models.Alias{}).Error
} 