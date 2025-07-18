package models

type DeletedMailbox struct {
	ID        uint64 `gorm:"primaryKey" json:"id"`
	Timestamp string `json:"timestamp"`
	Username  string `json:"username"`
	Domain    string `json:"domain"`
	Maildir   string `json:"maildir"`
	Admin     string `json:"admin"`
	DeleteDate string `json:"delete_date"`
} 