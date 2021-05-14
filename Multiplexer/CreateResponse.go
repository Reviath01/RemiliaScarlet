package multiplexer

import (
	embedutil "git.randomchars.net/FreeNitori/EmbedUtil"
	"git.randomchars.net/Reviath/RemiliaScarlet/interactions"
	"github.com/bwmarrin/discordgo"
)

// CreateResponse returns an interaction response
// This is for only string data
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

// CreateEmbedResponse returns an interaction response
// This is for only embed
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
