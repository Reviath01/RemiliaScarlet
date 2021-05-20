package webfuncs

// UserInfo structure to use on index.html
type UserInfo struct {
	Name          string
	ID            string
	AvatarURL     string
	Discriminator string
	Bot           bool
}

// Settings structure to use on guild.html
type Settings struct {
	WelcomeChannelID string
	LeaveChannelID   string
	RoleID           string
	WelcomeMessage   string
	LeaveMessage     string
	LogID            string
}

// ScanSQLData structure to use on web.go
type ScanSQLData struct {
	WelcomeChannelID string
	LeaveChannelID   string
	RoleID           string
	WelcomeMessage   string
	LeaveMessage     string
	LogID            string
}
