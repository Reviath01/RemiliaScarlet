package client

import (
    "github.com/bwmarrin/discordgo"
    "../config"
)

func Ready(s *discordgo.Session, event *discordgo.Ready) {
	s.UpdateGameStatus(0, config.Presence)
}
