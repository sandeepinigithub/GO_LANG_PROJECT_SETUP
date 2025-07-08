package repository

import (
    "GO_LANG_PROJECT_SETUP/config"
    "GO_LANG_PROJECT_SETUP/models"
)

func GetAllUsers() ([]model.User, error) {
    var users []model.User
    err := config.DB.Find(&users).Error
    return users, err
}

func GetUserByID(id uint) (model.User, error) {
    var user model.User
    err := config.DB.First(&user, id).Error
    return user, err
}

func CreateUser(user model.User) (model.User, error) {
    err := config.DB.Create(&user).Error
    return user, err
}

func UpdateUser(id uint, updated model.User) (model.User, error) {
    var user model.User
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
    return config.DB.Delete(&model.User{}, id).Error
}
