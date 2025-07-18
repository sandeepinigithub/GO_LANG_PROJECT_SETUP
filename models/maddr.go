package models

type Maddr struct {
	PartitionTag int    `json:"partition_tag"`
	ID           uint64 `gorm:"primaryKey" json:"id"`
	Email        string `json:"email"`
	EmailRaw     string `json:"email_raw"`
	Domain       string `json:"domain"`
} 