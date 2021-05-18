package web

import (
	"github.com/bwmarrin/discordgo"
)

// GetClientUser returns to Bot's owner as *discordgo.User
func GetClientUser(session *discordgo.Session) *discordgo.User {
	cli, _ := session.User("@me")
	return cli
}
