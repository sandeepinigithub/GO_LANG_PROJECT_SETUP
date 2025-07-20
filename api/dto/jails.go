package dto

// JailRequest represents the request structure for jail operations
type JailRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
}

// JailResponse represents the response structure for jail operations
type JailResponse struct {
	ID          uint64 `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// JailListResponse represents the response structure for listing jails
type JailListResponse struct {
	Jails []JailResponse `json:"jails"`
} 