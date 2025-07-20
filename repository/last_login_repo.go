package repository

import (
	"devsMailGo/config"
	"devsMailGo/models"
)

func GetLastLoginByUsername(username string) (models.LastLogin, error) {
	var login models.LastLogin
	err := config.DB.Table("last_login").Where("username = ?", username).First(&login).Error
	return login, err
}

func GetAllLastLogins() ([]models.LastLogin, error) {
	var logins []models.LastLogin
	err := config.DB.Table("last_login").Find(&logins).Error
	return logins, err
} 