package dto

type DomainRequest struct {
	Name            string `json:"name"`
	Description     string `json:"description"`
	Quota           int    `json:"quota"`
	Language        string `json:"language"`
	Transport       string `json:"transport"`
	DefaultQuota    int    `json:"default_quota"`
	MaxUserQuota    int    `json:"max_user_quota"`
	NumberOfUsers   int    `json:"number_of_users"`
	NumberOfAliases int    `json:"number_of_aliases"`
	NumberOfLists   int    `json:"number_of_lists"`
	AccountStatus   string `json:"account_status"`
	SenderBcc       string `json:"sender_bcc"`
	RecipientBcc    string `json:"recipient_bcc"`
}

type DomainResponse struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Quota       int    `json:"quota"`
} 