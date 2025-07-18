package models

type WblistRDNS struct {
	ID   uint64 `gorm:"primaryKey" json:"id"`
	RDNS string `json:"rdns"`
	WB   string `json:"wb"`
} 