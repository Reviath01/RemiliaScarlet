package multiplexer

import (
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

func CreateEmbedResponse(embed *discordgo.MessageEmbed) interactions.InteractionResponse {
	response := interactions.InteractionResponse{
		Type: interactions.InteractionResponseTypeChannelMessageWithSource,
		Data: interactions.InteractionResponseData{
			TTS:    false,
			Embeds: []discordgo.MessageEmbed{*embed},
		},
	}
	return response
}
