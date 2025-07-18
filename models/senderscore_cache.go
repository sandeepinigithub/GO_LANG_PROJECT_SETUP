package models

type SenderscoreCache struct {
	ClientAddress string `gorm:"primaryKey" json:"client_address"`
	Score         int    `json:"score"`
	Time          uint64 `json:"time"`
} 