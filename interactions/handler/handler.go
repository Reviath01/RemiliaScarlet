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
		if interaction.Data.Name == "icon" {
			response := interaction_commands.IconCommand(session, interaction)
			err = interaction.Respond(session, response)
			if err != nil {
				return
			}
		}
		if interaction.Data.Name == "guild_info" {
			response := interaction_commands.GuildInfoCommand(session, interaction)
			err = interaction.Respond(session, response)
			if err != nil {
				return
			}
		}
		if interaction.Data.Name == "issue" {
			response := interaction_commands.IssueCommand(session, interaction)
			err = interaction.Respond(session, response)
			if err != nil {
				return
			}
		}
		if interaction.Data.Name == "kick" {
			response := interaction_commands.KickCommand(session, interaction)
			err = interaction.Respond(session, response)
			if err != nil {
				return
			}
		}
		if interaction.Data.Name == "kiss" {
			response := interaction_commands.KissCommand(session, interaction)
			err = interaction.Respond(session, response)
			if err != nil {
				return
			}
		}
		if interaction.Data.Name == "language" {
			response := interaction_commands.LanguageCommand(session, interaction)
			err = interaction.Respond(session, response)
			if err != nil {
				return
			}
		}
		if interaction.Data.Name == "leave_channel" {
			response := interaction_commands.LeaveChannelCommand(session, interaction)
			err = interaction.Respond(session, response)
			if err != nil {
				return
			}
		}
		if interaction.Data.Name == "leave_message" {
			response := interaction_commands.LeaveMessageCommand(session, interaction)
			err = interaction.Respond(session, response)
			if err != nil {
				return
			}
		}
		if interaction.Data.Name == "log" {
			response := interaction_commands.LogCommand(session, interaction)
			err = interaction.Respond(session, response)
			if err != nil {
				return
			}
		}
		if interaction.Data.Name == "ping" {
			response := interaction_commands.PingCommand(session, interaction)
			err = interaction.Respond(session, response)
			if err != nil {
				return
			}
		}
		if interaction.Data.Name == "reset_autorole" {
			response := interaction_commands.ResetAutoroleCommand(session, interaction)
			err = interaction.Respond(session, response)
			if err != nil {
				return
			}
		}
		if interaction.Data.Name == "reset_leave_channel" {
			response := interaction_commands.ResetLeaveChannel(session, interaction)
			err = interaction.Respond(session, response)
			if err != nil {
				return
			}
		}
		if interaction.Data.Name == "reset_leave_message" {
			response := interaction_commands.ResetLeaveMessage(session, interaction)
			err = interaction.Respond(session, response)
			if err != nil {
				return
			}
		}
		if interaction.Data.Name == "reset_log" {
			response := interaction_commands.ResetLogCommand(session, interaction)
			err = interaction.Respond(session, response)
			if err != nil {
				return
			}
		}
		if interaction.Data.Name == "reset_welcome_channel" {
			response := interaction_commands.ResetWelcomeChannelCommand(session, interaction)
			err = interaction.Respond(session, response)
			if err != nil {
				return
			}
		}
		if interaction.Data.Name == "reset_welcome_message" {
			response := interaction_commands.ResetWelcomeMessageCommand(session, interaction)
			err = interaction.Respond(session, response)
			if err != nil {
				return
			}
		}
	}
}
