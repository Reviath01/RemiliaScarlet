package interaction_commands

import (
	"fmt"

	multiplexer "git.randomchars.net/Reviath/RemiliaScarlet/Multiplexer"
	"git.randomchars.net/Reviath/RemiliaScarlet/interactions"
	"git.randomchars.net/Reviath/RemiliaScarlet/sql"
	"github.com/bwmarrin/discordgo"
)

func BanCommand(session *discordgo.Session, interaction interactions.Interaction) interactions.InteractionResponse {
	if multiplexer.CheckPermission(session, interaction.Member.User.ID, interaction.ChannelID) == false {
		response := interactions.InteractionResponse{
			Type: interactions.InteractionResponseTypeChannelMessageWithSource,
			Data: interactions.InteractionResponseData{
				TTS:     false,
				Content: "You don't have enough permission.",
			},
		}
		return response
	}
	if sql.CheckLanguage(interaction.GuildID) == "tr" {
		if sql.IsBlocked(interaction.GuildID, "ban") == "true" {
			response := interactions.InteractionResponse{
				Type: interactions.InteractionResponseTypeChannelMessageWithSource,
				Data: interactions.InteractionResponseData{
					TTS:     false,
					Content: "Bu komut bu sunucuda engellenmiş",
				},
			}
			return response
		}
		u, err := session.User(multiplexer.GetUser(interaction.Data.Options[0].Value.(string)))
		if err != nil {
			response := interactions.InteractionResponse{
				Type: interactions.InteractionResponseTypeChannelMessageWithSource,
				Data: interactions.InteractionResponseData{
					TTS:     false,
					Content: "Bir üye belirtmelisiniz",
				},
			}
			return response
		} else {
			err = session.GuildBanCreate(interaction.GuildID, u.ID, 0)
			if err != nil {
				response := interactions.InteractionResponse{
					Type: interactions.InteractionResponseTypeChannelMessageWithSource,
					Data: interactions.InteractionResponseData{
						TTS:     false,
						Content: "Yeterli yetkiye sahip değilim",
					},
				}
				return response
			}
		}
	}

	if sql.IsBlocked(interaction.GuildID, "ban") == "true" {
		response := interactions.InteractionResponse{
			Type: interactions.InteractionResponseTypeChannelMessageWithSource,
			Data: interactions.InteractionResponseData{
				TTS:     false,
				Content: "This command is blocked on this guild.",
			},
		}
		return response
	}

	u, err := session.User(multiplexer.GetUser(interaction.Data.Options[0].Value.(string)))
	if err != nil {
		response := interactions.InteractionResponse{
			Type: interactions.InteractionResponseTypeChannelMessageWithSource,
			Data: interactions.InteractionResponseData{
				TTS:     false,
				Content: "You need to specify a user.",
			},
		}
		return response
	}
	err = session.GuildBanCreate(interaction.GuildID, u.ID, 0)
	if err != nil {
		response := interactions.InteractionResponse{
			Type: interactions.InteractionResponseTypeChannelMessageWithSource,
			Data: interactions.InteractionResponseData{
				TTS:     false,
				Content: fmt.Sprintf("An error occurred: %s", err.Error()),
			},
		}
		return response
	}
	response := interactions.InteractionResponse{
		Type: interactions.InteractionResponseTypeChannelMessageWithSource,
		Data: interactions.InteractionResponseData{
			TTS:     false,
			Content: "Banned specified user.",
		},
	}
	return response
}
