package repository

import (
	"devsMailGo/config"
	"devsMailGo/models"
)

func GetAllJails() ([]models.Jail, error) {
	var jails []models.Jail
	err := config.DB.Table("jails").Find(&jails).Error
	return jails, err
} 