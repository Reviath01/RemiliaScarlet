package interactions_handler

import (
	Commands "git.randomchars.net/Reviath/RemiliaScarlet/commands"
	"git.randomchars.net/Reviath/RemiliaScarlet/interactions"
	"github.com/bwmarrin/discordgo"
)

func InteractionHandler(session *discordgo.Session, event *discordgo.Event) {
	if event.Type == "INTERACTION_CREATE" {
		interaction, err := interactions.InteractionFromRaw(event.RawData)
		if err != nil {
			panic(err)
		}
		if interaction.Data.Name == "invite" {
			response := Commands.InteractiveInviteCommand(session)
			err = interaction.Respond(session, response)
			if err != nil {
				return
			}
		}
	}
}
