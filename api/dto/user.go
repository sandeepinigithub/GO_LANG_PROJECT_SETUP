package dto

type UserRequest struct {
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Domain   string `json:"domain"`
	Quota    int    `json:"quota"`
	Language string `json:"language"`
}

type UserResponse struct {
	ID     uint   `json:"id"`
	Email  string `json:"email"`
	Name   string `json:"name"`
	Domain string `json:"domain"`
} 