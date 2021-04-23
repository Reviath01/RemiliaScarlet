package interactions_handler

import (
	"git.randomchars.net/Reviath/RemiliaScarlet/interaction_commands"
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
			response := interaction_commands.InviteCommand(session, interaction)
			err = interaction.Respond(session, response)
			if err != nil {
				return
			}
		}
		if interaction.Data.Name == "author" {
			response := interaction_commands.AuthorCommand(session, interaction)
			err = interaction.Respond(session, response)
			if err != nil {
				return
			}
		}
		if interaction.Data.Name == "autorole" {
			response := interaction_commands.AutoRoleCommand(session, interaction)
			err = interaction.Respond(session, response)
			if err != nil {
				return
			}
		}
		if interaction.Data.Name == "avatar" {
			response := interaction_commands.AvatarCommand(session, interaction)
			err = interaction.Respond(session, response)
			if err != nil {
				return
			}
		}
		if interaction.Data.Name == "ban" {
			response := interaction_commands.BanCommand(session, interaction)
			err = interaction.Respond(session, response)
			if err != nil {
				return
			}
		}
		if interaction.Data.Name == "disable" {
			response := interaction_commands.DisableCommand(session, interaction)
			err = interaction.Respond(session, response)
			if err != nil {
				return
			}
		}
		if interaction.Data.Name == "embed" {
			response := interaction_commands.EmbedCommand(session, interaction)
			err = interaction.Respond(session, response)
			if err != nil {
				return
			}
		}
		if interaction.Data.Name == "hug" {
			response := interaction_commands.HugCommand(session, interaction)
			err = interaction.Respond(session, response)
			if err != nil {
				return
			}
		}
	}
}
