package models

import "gorm.io/gorm"

type MailingList struct {
	gorm.Model
	Address   string `gorm:"unique;not null" json:"address"`
	Domain    string `json:"domain"`
	Members   string `json:"members"` // Comma-separated list of emails
	Type      string `json:"type"` // subscribable, unsubscribable
	Active    bool   `json:"active"`
} 