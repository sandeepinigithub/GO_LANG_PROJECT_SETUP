package repository

import (
	"devsMailGo/config"
	"devsMailGo/models"
)

func GetSenderscoreByClientAddress(addr string) (models.SenderscoreCache, error) {
	var entry models.SenderscoreCache
	err := config.DB.Table("senderscore_cache").Where("client_address = ?", addr).First(&entry).Error
	return entry, err
} 