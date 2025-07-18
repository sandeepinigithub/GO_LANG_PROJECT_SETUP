package repository

import (
    "GO_LANG_PROJECT_SETUP/config"
    "GO_LANG_PROJECT_SETUP/models"
)

type RoundcubeUser struct {
	UserID   int
	Username string
	Password string
	MailHost string
}

func GetAllUsers() ([]models.User, error) {
    var users []models.User
    err := config.DB.Find(&users).Error
    return users, err
}

func GetUserByID(id uint) (models.User, error) {
    var user models.User
    err := config.DB.First(&user, id).Error
    return user, err
}

func CreateUser(user models.User) (models.User, error) {
    err := config.DB.Create(&user).Error
    return user, err
}

func UpdateUser(id uint, updated models.User) (models.User, error) {
    var user models.User
    err := config.DB.First(&user, id).Error
    if err != nil {
        return user, err
    }
    user.Name = updated.Name
    user.Email = updated.Email
    config.DB.Save(&user)
    return user, nil
}

func DeleteUser(id uint) error {
    return config.DB.Delete(&models.User{}, id).Error
}

func GetRoundcubeUserByUsername(username string) (RoundcubeUser, error) {
	var user RoundcubeUser
	db := config.DB.Table("users").Select("user_id, username, password, mail_host").Where("username = ?", username).First(&user)
	return user, db.Error
}
