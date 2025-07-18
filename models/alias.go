package models

import "gorm.io/gorm"

type Alias struct {
	gorm.Model
	Address   string `gorm:"unique;not null" json:"address"`
	Domain    string `json:"domain"`
	Goto      string `json:"goto"` // Comma-separated list of destination emails
	Active    bool   `json:"active"`
} 