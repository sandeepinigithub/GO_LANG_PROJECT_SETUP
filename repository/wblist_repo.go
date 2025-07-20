package repository

import (
	"devsMailGo/config"
	"devsMailGo/models"
)

func GetAllWblist() ([]models.Wblist, error) {
	var entries []models.Wblist
	err := config.DB.Table("wblist").Find(&entries).Error
	return entries, err
}

func GetWblistByRid(rid uint64) (models.Wblist, error) {
	var entry models.Wblist
	err := config.DB.Table("wblist").Where("rid = ?", rid).First(&entry).Error
	return entry, err
}

func CreateWblist(entry *models.Wblist) error {
	return config.DB.Table("wblist").Create(entry).Error
}

func UpdateWblist(rid uint64, updated *models.Wblist) error {
	return config.DB.Table("wblist").Where("rid = ?", rid).Updates(updated).Error
}

func DeleteWblist(rid uint64) error {
	return config.DB.Table("wblist").Where("rid = ?", rid).Delete(&models.Wblist{}).Error
} 