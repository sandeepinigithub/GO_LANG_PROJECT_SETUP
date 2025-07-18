package models

type Banned struct {
	ID        uint64 `gorm:"primaryKey" json:"id"`
	IP        string `json:"ip"`
	Ports     string `json:"ports"`
	Protocol  string `json:"protocol"`
	Jail      string `json:"jail"`
	Hostname  string `json:"hostname"`
	Country   string `json:"country"`
	RDNS      string `json:"rdns"`
	Failures  int    `json:"failures"`
	Loglines  string `json:"loglines"`
	Timestamp string `json:"timestamp"`
	Remove    bool   `json:"remove"`
} 