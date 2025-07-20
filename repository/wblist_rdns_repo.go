package repository

import (
	"devsMailGo/config"
	"devsMailGo/models"
)

func GetAllWblistRDNS() ([]models.WblistRDNS, error) {
	var entries []models.WblistRDNS
	err := config.DB.Table("wblist_rdns").Find(&entries).Error
	return entries, err
} 