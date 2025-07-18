package models

type RoundcubeUser struct {
	UserID             uint64 `gorm:"primaryKey" json:"user_id"`
	Username           string `json:"username"`
	MailHost           string `json:"mail_host"`
	Created            string `json:"created"`
	LastLogin          *string `json:"last_login"`
	FailedLogin        *string `json:"failed_login"`
	FailedLoginCounter *uint64 `json:"failed_login_counter"`
	Language           *string `json:"language"`
	Preferences        *string `json:"preferences"`
} 