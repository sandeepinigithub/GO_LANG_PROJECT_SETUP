package models

type LastLogin struct {
	Username string `gorm:"primaryKey" json:"username"`
	Domain   string `json:"domain"`
	Imap     *int   `json:"imap"`
	Pop3     *int   `json:"pop3"`
	Lda      *int   `json:"lda"`
} 