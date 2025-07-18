package models

type Msgrcpt struct {
	PartitionTag int     `json:"partition_tag"`
	MailID       string  `json:"mail_id"`
	Rseqnum      int     `json:"rseqnum"`
	Rid          uint64  `json:"rid"`
	IsLocal      string  `json:"is_local"`
	Content      string  `json:"content"`
	Ds           string  `json:"ds"`
	Rs           string  `json:"rs"`
	Bl           string  `json:"bl"`
	Wl           string  `json:"wl"`
	BspamLevel   float64 `json:"bspam_level"`
	SmtpResp     string  `json:"smtp_resp"`
} 