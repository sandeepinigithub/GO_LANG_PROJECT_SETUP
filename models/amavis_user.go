package models

type AmavisUser struct {
	ID       uint64 `gorm:"primaryKey" json:"id"`
	Priority int    `json:"priority"`
	PolicyID uint64 `json:"policy_id"`
	Email    string `json:"email"`
	Fullname string `json:"fullname"`
} 