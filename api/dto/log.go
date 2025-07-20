package dto

type LogRequest struct {
	Timestamp string `json:"timestamp"`
	Admin     string `json:"admin"`
	IP        string `json:"ip"`
	Domain    string `json:"domain"`
	Username  string `json:"username"`
	Event     string `json:"event"`
	Loglevel  string `json:"loglevel"`
	Msg       string `json:"msg"`
}

type LogResponse struct {
	ID        uint64 `json:"id"`
	Timestamp string `json:"timestamp"`
	Admin     string `json:"admin"`
	IP        string `json:"ip"`
	Domain    string `json:"domain"`
	Username  string `json:"username"`
	Event     string `json:"event"`
	Loglevel  string `json:"loglevel"`
	Msg       string `json:"msg"`
} 