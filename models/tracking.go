package models

type Tracking struct {
	ID   uint64 `gorm:"primaryKey" json:"id"`
	K    string `json:"k"`
	V    string `json:"v"`
	Time string `json:"time"`
} 