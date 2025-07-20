package repository

import (
	"GO_LANG_PROJECT_SETUP/config"
	"GO_LANG_PROJECT_SETUP/models"
)

func CreateDomainAdmin(admin *models.DomainAdmin) error {
	return config.DB.Create(admin).Error
}

func GetDomainAdminByEmail(email string) (models.DomainAdmin, error) {
	var admin models.DomainAdmin
	err := config.DB.Where("email = ?", email).First(&admin).Error
	return admin, err
}

func GetAllDomainAdmins() ([]models.DomainAdmin, error) {
	var admins []models.DomainAdmin
	err := config.DB.Find(&admins).Error
	return admins, err
}

func UpdateDomainAdmin(email string, updated *models.DomainAdmin) error {
	return config.DB.Model(&models.DomainAdmin{}).Where("email = ?", email).Updates(updated).Error
}

func DeleteDomainAdmin(email string) error {
	return config.DB.Where("email = ?", email).Delete(&models.DomainAdmin{}).Error
} 