package repository

import (
	"GO_LANG_PROJECT_SETUP/config"
	"GO_LANG_PROJECT_SETUP/models"
)

func GetRoundcubeUserByID(userID uint64) (models.RoundcubeUser, error) {
	var user models.RoundcubeUser
	err := config.DB.Table("users").Where("user_id = ?", userID).First(&user).Error
	return user, err
}

func GetAllRoundcubeUsers() ([]models.RoundcubeUser, error) {
	var users []models.RoundcubeUser
	err := config.DB.Table("users").Find(&users).Error
	return users, err
} 