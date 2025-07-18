package models

type Log struct {
	ID        uint64 `gorm:"primaryKey" json:"id"`
	Timestamp string `json:"timestamp"`
	Admin     string `json:"admin"`
	IP        string `json:"ip"`
	Domain    string `json:"domain"`
	Username  string `json:"username"`
	Event     string `json:"event"`
	Loglevel  string `json:"loglevel"`
	Msg       string `json:"msg"`
} 