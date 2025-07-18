package models

type Greylisting struct {
	ID             uint64 `gorm:"primaryKey" json:"id"`
	Account        string `json:"account"`
	Priority       int    `json:"priority"`
	Sender         string `json:"sender"`
	SenderPriority int    `json:"sender_priority"`
	Comment        string `json:"comment"`
	Active         bool   `json:"active"`
} 