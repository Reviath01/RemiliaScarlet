package commands

import (
	"fmt"

	embedutil "git.randomchars.net/Reviath/RemiliaScarlet/EmbedUtil"
	CommandHandler "git.randomchars.net/Reviath/RemiliaScarlet/handler"
	"git.randomchars.net/Reviath/RemiliaScarlet/interactions"
	"github.com/bwmarrin/discordgo"
)

func InviteCommand(ctx CommandHandler.Context, _ []string) error {
	inviteembed := embedutil.NewEmbed().
		SetColor(0xc000ff).
		SetDescription(fmt.Sprintf("Click [here](https://discord.com/oauth2/authorize?client_id=%s&scope=applications.commands+bot&permissions=8) to invite me!", ctx.Session.State.User.ID)).MessageEmbed
	ctx.ReplyEmbed(inviteembed)

	return nil
}

func InteractionInviteCommand(session *discordgo.Session, interaction interactions.Interaction) interactions.InteractionResponse {
	embed := embedutil.NewEmbed().
		SetColor(0xc000ff).
		SetDescription(fmt.Sprintf("Click [here](https://discord.com/oauth2/authorize?client_id=%s&scope=applications.commands+bot&permissions=8) to invite me!", session.State.User.ID)).MessageEmbed
	response := interactions.InteractionResponse{
		Type: interactions.InteractionResponseTypeChannelMessageWithSource,
		Data: interactions.InteractionResponseData{
			TTS:    false,
			Embeds: []discordgo.MessageEmbed{*embed},
		},
	}
	return response
}
