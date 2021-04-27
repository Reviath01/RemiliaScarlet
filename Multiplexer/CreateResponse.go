package multiplexer

import (
	embedutil "git.randomchars.net/FreeNitori/EmbedUtil"
	"git.randomchars.net/Reviath/RemiliaScarlet/interactions"
	"github.com/bwmarrin/discordgo"
)

func CreateResponse(message string) interactions.InteractionResponse {
	response := interactions.InteractionResponse{
		Type: interactions.InteractionResponseTypeChannelMessageWithSource,
		Data: interactions.InteractionResponseData{
			TTS:     false,
			Content: message,
		},
	}
	return response
}

func CreateEmbedResponse(embed embedutil.Embed) interactions.InteractionResponse {
	response := interactions.InteractionResponse{
		Type: interactions.InteractionResponseTypeChannelMessageWithSource,
		Data: interactions.InteractionResponseData{
			TTS:    false,
			Embeds: []discordgo.MessageEmbed{*embed.MessageEmbed},
		},
	}
	return response
}
