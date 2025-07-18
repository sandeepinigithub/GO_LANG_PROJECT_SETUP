package models

type GreylistingTracking struct {
	ID            uint64 `gorm:"primaryKey" json:"id"`
	Sender        string `json:"sender"`
	Recipient     string `json:"recipient"`
	ClientAddress string `json:"client_address"`
	SenderDomain  string `json:"sender_domain"`
	RcptDomain    string `json:"rcpt_domain"`
	InitTime      uint64 `json:"init_time"`
	BlockExpired  uint64 `json:"block_expired"`
	RecordExpired uint64 `json:"record_expired"`
	BlockedCount  uint64 `json:"blocked_count"`
	Passed        bool   `json:"passed"`
} 