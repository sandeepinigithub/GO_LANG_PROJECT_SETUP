package models

type UsedQuota struct {
	Username string `gorm:"primaryKey" json:"username"`
	Bytes    uint64 `json:"bytes"`
	Messages uint64 `json:"messages"`
	Domain   string `json:"domain"`
} 