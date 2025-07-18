package models

type Policy struct {
	ID                        uint64  `gorm:"primaryKey" json:"id"`
	PolicyName                string  `json:"policy_name"`
	VirusLover                string  `json:"virus_lover"`
	SpamLover                 string  `json:"spam_lover"`
	UncheckedLover            string  `json:"unchecked_lover"`
	BannedFilesLover          string  `json:"banned_files_lover"`
	BadHeaderLover            string  `json:"bad_header_lover"`
	BypassVirusChecks         string  `json:"bypass_virus_checks"`
	BypassSpamChecks          string  `json:"bypass_spam_checks"`
	BypassBannedChecks        string  `json:"bypass_banned_checks"`
	BypassHeaderChecks        string  `json:"bypass_header_checks"`
	VirusQuarantineTo         string  `json:"virus_quarantine_to"`
	SpamQuarantineTo          string  `json:"spam_quarantine_to"`
	BannedQuarantineTo        string  `json:"banned_quarantine_to"`
	UncheckedQuarantineTo     string  `json:"unchecked_quarantine_to"`
	BadHeaderQuarantineTo     string  `json:"bad_header_quarantine_to"`
	CleanQuarantineTo         string  `json:"clean_quarantine_to"`
	ArchiveQuarantineTo       string  `json:"archive_quarantine_to"`
	SpamTagLevel              float64 `json:"spam_tag_level"`
	SpamTag2Level             float64 `json:"spam_tag2_level"`
	SpamTag3Level             float64 `json:"spam_tag3_level"`
	SpamKillLevel             float64 `json:"spam_kill_level"`
	SpamDsnCutoffLevel        float64 `json:"spam_dsn_cutoff_level"`
	SpamQuarantineCutoffLevel float64 `json:"spam_quarantine_cutoff_level"`
	AddrExtensionVirus        string  `json:"addr_extension_virus"`
	AddrExtensionSpam         string  `json:"addr_extension_spam"`
	AddrExtensionBanned       string  `json:"addr_extension_banned"`
	AddrExtensionBadHeader    string  `json:"addr_extension_bad_header"`
	WarnVirusRecip            string  `json:"warnvirusrecip"`
	WarnBannedRecip           string  `json:"warnbannedrecip"`
	WarnBadhRecip             string  `json:"warnbadhrecip"`
	NewVirusAdmin             string  `json:"newvirus_admin"`
	VirusAdmin                string  `json:"virus_admin"`
	BannedAdmin               string  `json:"banned_admin"`
	BadHeaderAdmin            string  `json:"bad_header_admin"`
	SpamAdmin                 string  `json:"spam_admin"`
	SpamSubjectTag            string  `json:"spam_subject_tag"`
	SpamSubjectTag2           string  `json:"spam_subject_tag2"`
	SpamSubjectTag3           string  `json:"spam_subject_tag3"`
	MessageSizeLimit          int     `json:"message_size_limit"`
	BannedRulenames           string  `json:"banned_rulenames"`
	DisclaimerOptions         string  `json:"disclaimer_options"`
	ForwardMethod             string  `json:"forward_method"`
	SaUserconf                string  `json:"sa_userconf"`
	SaUsername                string  `json:"sa_username"`
} 