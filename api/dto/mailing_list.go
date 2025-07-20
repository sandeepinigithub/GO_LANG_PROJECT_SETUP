package dto

// MailingListRequest represents the request structure for mailing list operations
type MailingListRequest struct {
	Address string `json:"address" binding:"required"`
	Domain  string `json:"domain"`
	Members string `json:"members"`
	Type    string `json:"type"`
	Active  bool   `json:"active"`
}

// MailingListResponse represents the response structure for mailing list operations
type MailingListResponse struct {
	ID      uint64 `json:"id"`
	Address string `json:"address"`
	Domain  string `json:"domain"`
	Members string `json:"members"`
	Type    string `json:"type"`
	Active  bool   `json:"active"`
}

// MailingListListResponse represents the response structure for listing mailing lists
type MailingListListResponse struct {
	MailingLists []MailingListResponse `json:"mailing_lists"`
} 