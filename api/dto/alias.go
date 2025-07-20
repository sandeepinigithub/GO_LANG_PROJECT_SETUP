package dto

type AliasRequest struct {
	Address string `json:"address"`
	Domain  string `json:"domain"`
	Goto    string `json:"goto"`
	Active  bool   `json:"active"`
}

type AliasResponse struct {
	Address string `json:"address"`
	Domain  string `json:"domain"`
	Goto    string `json:"goto"`
	Active  bool   `json:"active"`
} 