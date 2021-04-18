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
			response := Commands.InteractionInviteCommand(session, interaction)
			err = interaction.Respond(session, response)
			if err != nil {
				return
			}
		}
		if interaction.Data.Name == "author" {
			response := Commands.InteractionAuthorCommand(session, interaction)
			err = interaction.Respond(session, response)
			if err != nil {
				return
			}
		}
		if interaction.Data.Name == "autorole" {
			response := Commands.InteractionAutoRoleCommand(session, interaction)
			err = interaction.Respond(session, response)
			if err != nil {
				return
			}
		}
		if interaction.Data.Name == "avatar" {
			response := Commands.InteractionAvatarCommand(session, interaction)
			err = interaction.Respond(session, response)
			if err != nil {
				return
			}
		}
	}
}
