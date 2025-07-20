package repository

import (
	"GO_LANG_PROJECT_SETUP/config"
	"GO_LANG_PROJECT_SETUP/models"
)

func GetAllRoundcubeUsers() ([]models.RoundcubeUser, error) {
	var users []models.RoundcubeUser
	err := config.DB.Table("users").Find(&users).Error
	return users, err
}

func GetRoundcubeUserByID(userID uint64) (models.RoundcubeUser, error) {
	var user models.RoundcubeUser
	err := config.DB.Table("users").Where("user_id = ?", userID).First(&user).Error
	return user, err
}

func CreateRoundcubeUser(user *models.RoundcubeUser) error {
	return config.DB.Table("users").Create(user).Error
}

func UpdateRoundcubeUser(userID uint64, updated *models.RoundcubeUser) error {
	return config.DB.Table("users").Where("user_id = ?", userID).Updates(updated).Error
}

func DeleteRoundcubeUser(userID uint64) error {
	return config.DB.Table("users").Where("user_id = ?", userID).Delete(&models.RoundcubeUser{}).Error
} 