package models

type Session struct {
	SessionID string `gorm:"primaryKey" json:"session_id"`
	Atime     string `json:"atime"`
	Data      string `json:"data"`
} 