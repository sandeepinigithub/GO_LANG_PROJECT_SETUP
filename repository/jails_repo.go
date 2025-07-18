package repository

import (
	"GO_LANG_PROJECT_SETUP/config"
	"GO_LANG_PROJECT_SETUP/models"
)

func GetAllJails() ([]models.Jail, error) {
	var jails []models.Jail
	err := config.DB.Table("jails").Find(&jails).Error
	return jails, err
} 