package models

type UpdateLog struct {
	ID   uint64 `gorm:"primaryKey" json:"id"`
	Date string `json:"date"`
} 