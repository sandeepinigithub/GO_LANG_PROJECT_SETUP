package models

type GreylistingWhitelistDomain struct {
	ID     uint64 `gorm:"primaryKey" json:"id"`
	Domain string `json:"domain"`
} 