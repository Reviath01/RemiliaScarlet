package interactions

import (
	"fmt"

	embedutil "git.randomchars.net/Reviath/RemiliaScarlet/EmbedUtil"
	"github.com/bwmarrin/discordgo"
)

func InteractionHandler(session *discordgo.Session, event *discordgo.Event) {
	if event.Type == "INTERACTION_CREATE" {
		interaction, err := InteractionFromRaw(event.RawData)
		if err != nil {
			panic(err)
		}
		if interaction.Data.Name == "invite" {
			if err != nil {
				panic(err)
			}
			embed := embedutil.NewEmbed().
				SetColor(0xc000ff).
				SetDescription(fmt.Sprintf("Click [here](https://discord.com/oauth2/authorize?client_id=%s&scope=applications.commands+bot&permissions=8) to invite me!", session.State.User.ID)).MessageEmbed
			response := InteractionResponse{
				Type: InteractionResponseTypeChannelMessageWithSource,
				Data: InteractionResponseData{
					TTS:    false,
					Embeds: []discordgo.MessageEmbed{*embed},
				},
			}
			err = interaction.Respond(session, response)
			if err != nil {
				return
			}
		}
	}
}
