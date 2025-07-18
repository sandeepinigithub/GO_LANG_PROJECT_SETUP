package models

type NewsletterSubunsubConfirm struct {
	ID         uint64 `gorm:"primaryKey" json:"id"`
	Mail       string `json:"mail"`
	Mlid       string `json:"mlid"`
	Subscriber string `json:"subscriber"`
	Kind       string `json:"kind"`
	Token      string `json:"token"`
	Expired    uint64 `json:"expired"`
} 