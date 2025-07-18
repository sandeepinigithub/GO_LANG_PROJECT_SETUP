package models

type SRSExcludeDomain struct {
	ID     uint64 `gorm:"primaryKey" json:"id"`
	Domain string `json:"domain"`
} 