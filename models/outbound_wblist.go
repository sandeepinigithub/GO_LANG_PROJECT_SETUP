package models

type OutboundWblist struct {
	Rid uint64 `json:"rid"`
	Sid uint64 `json:"sid"`
	Wb  string `json:"wb"`
} 