package interaction_commands

import (
	"fmt"

	embedutil "git.randomchars.net/Reviath/RemiliaScarlet/EmbedUtil"
	multiplexer "git.randomchars.net/Reviath/RemiliaScarlet/Multiplexer"
	"git.randomchars.net/Reviath/RemiliaScarlet/interactions"
	"git.randomchars.net/Reviath/RemiliaScarlet/sql"
	"github.com/bwmarrin/discordgo"
)

func AvatarCommand(session *discordgo.Session, interaction interactions.Interaction) interactions.InteractionResponse {
	if sql.CheckLanguage(interaction.GuildID) == "tr" {
		if sql.IsBlocked(interaction.GuildID, "avatar") == "true" {
			response := interactions.InteractionResponse{
				Type: interactions.InteractionResponseTypeChannelMessageWithSource,
				Data: interactions.InteractionResponseData{
					TTS:     false,
					Content: "Bu komut bu sunucuda engellenmiş.",
				},
			}
			return response
		}
		u, err := session.User(multiplexer.GetUser(interaction.Data.Options[0].Value.(string)))
		if err == nil {
			avatarembed := embedutil.NewEmbed().
				SetColor(0xff1000).
				SetDescription(fmt.Sprintf("%s#%s isimli kişinin profil fotoğrafı", u.Username, u.Discriminator)).
				SetImage(u.AvatarURL("1024")).MessageEmbed
			response := interactions.InteractionResponse{
				Type: interactions.InteractionResponseTypeChannelMessageWithSource,
				Data: interactions.InteractionResponseData{
					TTS:    false,
					Embeds: []discordgo.MessageEmbed{*avatarembed},
				},
			}
			return response
		} else {
			avatarembed := embedutil.NewEmbed().
				SetColor(0xff1000).
				SetDescription(fmt.Sprintf("%s#%s isimli kişinin profil fotoğrafı", interaction.Member.User.Username, interaction.Member.User.Discriminator)).
				SetImage(interaction.Member.User.AvatarURL("1024")).MessageEmbed
			response := interactions.InteractionResponse{
				Type: interactions.InteractionResponseTypeChannelMessageWithSource,
				Data: interactions.InteractionResponseData{
					TTS:    false,
					Embeds: []discordgo.MessageEmbed{*avatarembed},
				},
			}
			return response
		}
	}

	if sql.IsBlocked(interaction.GuildID, "avatar") == "true" {
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
	if err == nil {
		avatarembed := embedutil.NewEmbed().
			SetColor(0xff1000).
			SetDescription(fmt.Sprintf("Avatar of %s#%s", u.Username, u.Discriminator)).
			SetImage(u.AvatarURL("1024")).MessageEmbed
		response := interactions.InteractionResponse{
			Type: interactions.InteractionResponseTypeChannelMessageWithSource,
			Data: interactions.InteractionResponseData{
				TTS:    false,
				Embeds: []discordgo.MessageEmbed{*avatarembed},
			},
		}
		return response
	} else {
		avatarembed := embedutil.NewEmbed().
			SetColor(0xff1000).
			SetDescription(fmt.Sprintf("Avatar of %s#%s", interaction.Member.User.Username, interaction.Member.User.Discriminator)).
			SetImage(interaction.Member.User.AvatarURL("1024")).MessageEmbed
		response := interactions.InteractionResponse{
			Type: interactions.InteractionResponseTypeChannelMessageWithSource,
			Data: interactions.InteractionResponseData{
				TTS:    false,
				Embeds: []discordgo.MessageEmbed{*avatarembed},
			},
		}
		return response
	}
}
