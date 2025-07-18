package models

type ThrottleTracking struct {
	ID            uint64 `gorm:"primaryKey" json:"id"`
	Tid           uint64 `json:"tid"`
	Account       string `json:"account"`
	Period        uint64 `json:"period"`
	CurMsgs       uint64 `json:"cur_msgs"`
	CurQuota      int64  `json:"cur_quota"`
	InitTime      uint64 `json:"init_time"`
	LastTime      uint64 `json:"last_time"`
	LastNotifyTime uint64 `json:"last_notify_time"`
} 