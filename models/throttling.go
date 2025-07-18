package models

import "gorm.io/gorm"

type Throttling struct {
	gorm.Model
	Domain      string `json:"domain"`
	User        string `json:"user"`
	Type        string `json:"type"` // inbound, outbound
	Period      int    `json:"period"` // in seconds
	MsgSize     int    `json:"msg_size"` // in KB
	MsgCount    int    `json:"msg_count"`
	Active      bool   `json:"active"`
} 