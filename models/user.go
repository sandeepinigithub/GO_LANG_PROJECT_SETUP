package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email         string `gorm:"unique;not null" json:"email"`
	Name          string `json:"name"`
	Password      string `json:"-"`
	Domain        string `json:"domain"`
	Quota         int    `json:"quota"` // in MB
	Language      string `json:"language"`
	AccountStatus string `json:"account_status"` // active, disabled
	IsAdmin       bool   `json:"is_admin"`
}
