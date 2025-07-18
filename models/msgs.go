package models

type Msgs struct {
	PartitionTag int     `json:"partition_tag"`
	MailID       string  `json:"mail_id"`
	SecretID     string  `json:"secret_id"`
	AmID         string  `json:"am_id"`
	TimeNum      uint64  `json:"time_num"`
	TimeIso      string  `json:"time_iso"`
	Sid          uint64  `json:"sid"`
	Policy       string  `json:"policy"`
	ClientAddr   string  `json:"client_addr"`
	Size         uint64  `json:"size"`
	Originating  string  `json:"originating"`
	Content      string  `json:"content"`
	QuarType     string  `json:"quar_type"`
	QuarLoc      string  `json:"quar_loc"`
	DsnSent      string  `json:"dsn_sent"`
	SpamLevel    float64 `json:"spam_level"`
	MessageID    string  `json:"message_id"`
	FromAddr     string  `json:"from_addr"`
	Subject      string  `json:"subject"`
	Host         string  `json:"host"`
} 