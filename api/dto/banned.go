package dto

// BannedRequest represents the request structure for banned operations
type BannedRequest struct {
	IP string `json:"ip" binding:"required"`
}

// BannedResponse represents the response structure for banned operations
type BannedResponse struct {
	ID uint64 `json:"id"`
	IP string `json:"ip"`
}

// BannedListResponse represents the response structure for listing banned IPs
type BannedListResponse struct {
	Banned []BannedResponse `json:"banned"`
} 