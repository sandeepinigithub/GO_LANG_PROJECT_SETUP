package models

type DomainOwnership struct {
	ID         uint64 `gorm:"primaryKey" json:"id"`
	Admin      string `json:"admin"`
	Domain     string `json:"domain"`
	AliasDomain string `json:"alias_domain"`
	VerifyCode string `json:"verify_code"`
	Verified   bool   `json:"verified"`
	Message    string `json:"message"`
	LastVerify string `json:"last_verify"`
	Expire     uint64 `json:"expire"`
} 