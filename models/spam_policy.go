package models

import "gorm.io/gorm"

type SpamPolicy struct {
	gorm.Model
	Name        string `gorm:"unique;not null" json:"name"`
	Domain      string `json:"domain"`
	User        string `json:"user"`
	Policy      string `json:"policy"`
	Active      bool   `json:"active"`
} 