package web

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

// GetClientUser returns to Bot's owner as *discordgo.User
func GetClientUser(session *discordgo.Session) *discordgo.User {
	cli, err := session.User("@me")
	if err != nil {
		log.Fatalln("Cannot get client user, exiting process.")
		return nil
	}
	return cli
}
