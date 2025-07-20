package dto

type WblistRequest struct {
	Sid uint64 `json:"sid"`
	Wb  string `json:"wb"`
}

type WblistResponse struct {
	Rid uint64 `json:"rid"`
	Sid uint64 `json:"sid"`
	Wb  string `json:"wb"`
} 