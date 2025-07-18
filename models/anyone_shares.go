package models

type AnyoneShare struct {
	FromUser string `gorm:"primaryKey" json:"from_user"`
	Dummy    string `json:"dummy"`
} 