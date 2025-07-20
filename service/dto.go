package service

type UserRequest struct {
	Email    string
	Name     string
	Password string
	Domain   string
	Quota    int
	Language string
}

type UserResponse struct {
	ID    uint
	Email string
	Name  string
}

type AliasRequest struct {
	Address string
	Domain  string
	Goto    string
	Active  bool
}

type AliasResponse struct {
	Address string
	Domain  string
	Goto    string
	Active  bool
}

type DomainRequest struct {
	Name            string
	Description     string
	Quota           int
	Language        string
	Transport       string
	DefaultQuota    int
	MaxUserQuota    int
	NumberOfUsers   int
	NumberOfAliases int
	NumberOfLists   int
	AccountStatus   string
	SenderBcc       string
	RecipientBcc    string
}

type DomainResponse struct {
	Name        string
	Description string
	Quota       int
}

type GreylistingRequest struct {
	Account        string
	Priority       int
	Sender         string
	SenderPriority int
	Comment        string
	Active         bool
}

type GreylistingResponse struct {
	ID             uint64
	Account        string
	Priority       int
	Sender         string
	SenderPriority int
	Comment        string
	Active         bool
}

type LogRequest struct {
	Timestamp string
	Admin     string
	IP        string
	Domain    string
	Username  string
	Event     string
	Loglevel  string
	Msg       string
}

type LogResponse struct {
	ID        uint64
	Timestamp string
	Admin     string
	IP        string
	Domain    string
	Username  string
	Event     string
	Loglevel  string
	Msg       string
}

type ThrottleRequest struct {
	Account  string
	Kind     string
	Priority int
	Period   uint64
	MsgSize  int64
	MaxMsgs  int64
	MaxQuota int64
	MaxRcpts int64
}

type ThrottleResponse struct {
	ID       uint64
	Account  string
	Kind     string
	Priority int
	Period   uint64
	MsgSize  int64
	MaxMsgs  int64
	MaxQuota int64
	MaxRcpts int64
}

type WblistRequest struct {
	Sid uint64
	Wb  string
}

type WblistResponse struct {
	Rid uint64
	Sid uint64
	Wb  string
}

type QuotaRequest struct {
	Username string
	Bytes    uint64
	Messages uint64
	Domain   string
}

type QuotaResponse struct {
	Username string
	Bytes    uint64
	Messages uint64
	Domain   string
}

type RoundcubeUserRequest struct {
	Username           string
	MailHost           string
	Created            string
	LastLogin          *string
	FailedLogin        *string
	FailedLoginCounter *uint64
	Language           *string
	Preferences        *string
}

type RoundcubeUserResponse struct {
	UserID             uint64
	Username           string
	MailHost           string
	Created            string
	LastLogin          *string
	FailedLogin        *string
	FailedLoginCounter *uint64
	Language           *string
	Preferences        *string
}

type BannedResponse struct {
	IP string
}

type JailResponse struct {
	// Add fields as needed
} 