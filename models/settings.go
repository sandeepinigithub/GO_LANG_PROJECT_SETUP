package models

type Setting struct {
	ID      uint64 `gorm:"primaryKey" json:"id"`
	Account string `json:"account"`
	K       string `json:"k"`
	V       string `json:"v"`
} 