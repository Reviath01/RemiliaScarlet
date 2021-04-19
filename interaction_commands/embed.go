package interaction_commands

import (
	embedutil "git.randomchars.net/Reviath/RemiliaScarlet/EmbedUtil"
	"git.randomchars.net/Reviath/RemiliaScarlet/interactions"
	"git.randomchars.net/Reviath/RemiliaScarlet/sql"
	"github.com/bwmarrin/discordgo"
)

func EmbedCommand(session *discordgo.Session, interaction interactions.Interaction) interactions.InteractionResponse {
	if sql.CheckLanguage(interaction.GuildID) == "tr" {
		if sql.IsBlocked(interaction.GuildID, "embed") == "true" {
			response := interactions.InteractionResponse{
				Type: interactions.InteractionResponseTypeChannelMessageWithSource,
				Data: interactions.InteractionResponseData{
					TTS:     false,
					Content: "Bu komut bu sunucuda engellenmiş.",
				},
			}
			return response
		}

		embed := embedutil.NewEmbed().
			SetColor(0xc000ff).
			SetDescription(interaction.Data.Options[0].Value.(string)).MessageEmbed
		response := interactions.InteractionResponse{
			Type: interactions.InteractionResponseTypeChannelMessageWithSource,
			Data: interactions.InteractionResponseData{
				TTS:    false,
				Embeds: []discordgo.MessageEmbed{*embed},
			},
		}
		return response
	}

	if sql.IsBlocked(interaction.GuildID, "embed") == "true" {
		response := interactions.InteractionResponse{
			Type: interactions.InteractionResponseTypeChannelMessageWithSource,
			Data: interactions.InteractionResponseData{
				TTS:     false,
				Content: "This command is blocked on this guild.",
			},
		}
		return response
	}

	embed := embedutil.NewEmbed().
		SetColor(0xc000ff).
		SetDescription(interaction.Data.Options[0].Value.(string)).MessageEmbed
	response := interactions.InteractionResponse{
		Type: interactions.InteractionResponseTypeChannelMessageWithSource,
		Data: interactions.InteractionResponseData{
			TTS:    false,
			Embeds: []discordgo.MessageEmbed{*embed},
		},
	}
	return response
}
