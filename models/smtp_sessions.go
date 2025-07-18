package models

type SMTPSession struct {
	ID                uint64 `gorm:"primaryKey" json:"id"`
	Time              string `json:"time"`
	TimeNum           uint64 `json:"time_num"`
	Action            string `json:"action"`
	Reason            string `json:"reason"`
	Instance          string `json:"instance"`
	ClientAddress     string `json:"client_address"`
	ClientName        string `json:"client_name"`
	ReverseClientName string `json:"reverse_client_name"`
	HeloName          string `json:"helo_name"`
	Sender            string `json:"sender"`
	SenderDomain      string `json:"sender_domain"`
	SaslUsername      string `json:"sasl_username"`
	SaslDomain        string `json:"sasl_domain"`
	Recipient         string `json:"recipient"`
	RecipientDomain   string `json:"recipient_domain"`
	EncryptionProtocol string `json:"encryption_protocol"`
	EncryptionCipher  string `json:"encryption_cipher"`
	ServerAddress     string `json:"server_address"`
	ServerPort        string `json:"server_port"`
} 