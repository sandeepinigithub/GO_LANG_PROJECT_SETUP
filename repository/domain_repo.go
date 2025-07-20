package repository

import (
	"devsMailGo/config"
	"devsMailGo/models"
)

func CreateDomain(domain *models.Domain) error {
	return config.DB.Create(domain).Error
}

func GetAllDomains() ([]models.Domain, error) {
	var domains []models.Domain
	err := config.DB.Find(&domains).Error
	return domains, err
}

func GetDomainByName(name string) (models.Domain, error) {
	var domain models.Domain
	err := config.DB.Where("name = ?", name).First(&domain).Error
	return domain, err
}

func UpdateDomain(name string, updated *models.Domain) error {
	return config.DB.Model(&models.Domain{}).Where("name = ?", name).Updates(updated).Error
}

func DeleteDomain(name string) error {
	return config.DB.Where("name = ?", name).Delete(&models.Domain{}).Error
} 