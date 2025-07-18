package models

import "gorm.io/gorm"

type DomainAdmin struct {
	gorm.Model
	Email      string `gorm:"unique;not null" json:"email"`
	Password   string `json:"-"`
	Domains    string `json:"domains"` // Comma-separated list of domains
	Privileges string `json:"privileges"` // Comma-separated list of privileges
	Active     bool   `json:"active"`
} 