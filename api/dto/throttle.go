package dto

type ThrottleRequest struct {
	Account  string `json:"account"`
	Kind     string `json:"kind"`
	Priority int    `json:"priority"`
	Period   uint64 `json:"period"`
	MsgSize  int64  `json:"msg_size"`
	MaxMsgs  int64  `json:"max_msgs"`
	MaxQuota int64  `json:"max_quota"`
	MaxRcpts int64  `json:"max_rcpts"`
}

type ThrottleResponse struct {
	ID       uint64 `json:"id"`
	Account  string `json:"account"`
	Kind     string `json:"kind"`
	Priority int    `json:"priority"`
	Period   uint64 `json:"period"`
	MsgSize  int64  `json:"msg_size"`
	MaxMsgs  int64  `json:"max_msgs"`
	MaxQuota int64  `json:"max_quota"`
	MaxRcpts int64  `json:"max_rcpts"`
} 