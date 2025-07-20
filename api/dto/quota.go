package dto

type QuotaRequest struct {
	Username string `json:"username"`
	Bytes    uint64 `json:"bytes"`
	Messages uint64 `json:"messages"`
	Domain   string `json:"domain"`
}

type QuotaResponse struct {
	Username string `json:"username"`
	Bytes    uint64 `json:"bytes"`
	Messages uint64 `json:"messages"`
	Domain   string `json:"domain"`
} 