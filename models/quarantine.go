package models

type Quarantine struct {
	PartitionTag int    `json:"partition_tag"`
	MailID       string `json:"mail_id"`
	ChunkInd     uint64 `json:"chunk_ind"`
	MailText     []byte `json:"mail_text"`
} 