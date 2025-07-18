package models

type MailAddr struct {
	ID       uint64 `gorm:"primaryKey" json:"id"`
	Priority int    `json:"priority"`
	Email    string `json:"email"`
} 