package web

type UserInfo struct {
	Name          string
	ID            string
	AvatarURL     string
	Discriminator string
	Bot           bool
}

type GuildInfo struct {
	Name []string
	ID   []string
	Icon []string
}
