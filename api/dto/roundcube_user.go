package dto

type RoundcubeUserRequest struct {
	Username           string  `json:"username"`
	MailHost           string  `json:"mail_host"`
	Created            string  `json:"created"`
	LastLogin          *string `json:"last_login"`
	FailedLogin        *string `json:"failed_login"`
	FailedLoginCounter *uint64 `json:"failed_login_counter"`
	Language           *string `json:"language"`
	Preferences        *string `json:"preferences"`
}

type RoundcubeUserResponse struct {
	UserID             uint64  `json:"user_id"`
	Username           string  `json:"username"`
	MailHost           string  `json:"mail_host"`
	Created            string  `json:"created"`
	LastLogin          *string `json:"last_login"`
	FailedLogin        *string `json:"failed_login"`
	FailedLoginCounter *uint64 `json:"failed_login_counter"`
	Language           *string `json:"language"`
	Preferences        *string `json:"preferences"`
} 