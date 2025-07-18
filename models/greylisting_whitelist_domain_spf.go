package models

type GreylistingWhitelistDomainSPF struct {
	ID      uint64 `gorm:"primaryKey" json:"id"`
	Account string `json:"account"`
	Sender  string `json:"sender"`
	Comment string `json:"comment"`
} 