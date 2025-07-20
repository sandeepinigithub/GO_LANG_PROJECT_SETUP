package dto

type GreylistingRequest struct {
	Account        string `json:"account"`
	Priority       int    `json:"priority"`
	Sender         string `json:"sender"`
	SenderPriority int    `json:"sender_priority"`
	Comment        string `json:"comment"`
	Active         bool   `json:"active"`
}

type GreylistingResponse struct {
	ID             uint64 `json:"id"`
	Account        string `json:"account"`
	Priority       int    `json:"priority"`
	Sender         string `json:"sender"`
	SenderPriority int    `json:"sender_priority"`
	Comment        string `json:"comment"`
	Active         bool   `json:"active"`
} 