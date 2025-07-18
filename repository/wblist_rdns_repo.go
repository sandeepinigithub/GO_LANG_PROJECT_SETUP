package repository

import (
	"GO_LANG_PROJECT_SETUP/config"
	"GO_LANG_PROJECT_SETUP/models"
)

func GetAllWblistRDNS() ([]models.WblistRDNS, error) {
	var entries []models.WblistRDNS
	err := config.DB.Table("wblist_rdns").Find(&entries).Error
	return entries, err
} 