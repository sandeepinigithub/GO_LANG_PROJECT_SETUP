package models

type Jail struct {
	ID      uint64 `gorm:"primaryKey" json:"id"`
	Name    string `json:"name"`
	Enabled bool   `json:"enabled"`
} 