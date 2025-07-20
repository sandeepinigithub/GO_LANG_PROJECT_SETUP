package repository

import (
	"GO_LANG_PROJECT_SETUP/config"
	"GO_LANG_PROJECT_SETUP/models"
)

func CreateMailingList(list *models.MailingList) error {
	return config.DB.Create(list).Error
}

func GetAllMailingLists() ([]models.MailingList, error) {
	var lists []models.MailingList
	err := config.DB.Find(&lists).Error
	return lists, err
}

func GetMailingListByAddress(address string) (models.MailingList, error) {
	var list models.MailingList
	err := config.DB.Where("address = ?", address).First(&list).Error
	return list, err
}

func UpdateMailingList(address string, updated *models.MailingList) error {
	return config.DB.Model(&models.MailingList{}).Where("address = ?", address).Updates(updated).Error
}

func DeleteMailingList(address string) error {
	return config.DB.Where("address = ?", address).Delete(&models.MailingList{}).Error
} 