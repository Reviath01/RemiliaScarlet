package events

import (
	"git.randomchars.net/Reviath/RemiliaScarlet/config"
	"github.com/bwmarrin/discordgo"
)

func Ready(s *discordgo.Session, event *discordgo.Ready) {
	config.ReadConfig()

	s.UpdateGameStatus(0, config.Presence)
}
