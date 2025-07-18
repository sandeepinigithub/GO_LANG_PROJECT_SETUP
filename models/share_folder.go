package models

type ShareFolder struct {
	FromUser string `gorm:"primaryKey" json:"from_user"`
	ToUser   string `gorm:"primaryKey" json:"to_user"`
	Dummy    string `json:"dummy"`
} 