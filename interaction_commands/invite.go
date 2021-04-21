package interaction_commands

import (
	"fmt"

	embedutil "git.randomchars.net/Reviath/RemiliaScarlet/EmbedUtil"
	multiplexer "git.randomchars.net/Reviath/RemiliaScarlet/Multiplexer"
	"git.randomchars.net/Reviath/RemiliaScarlet/interactions"
	"github.com/bwmarrin/discordgo"
)

func InviteCommand(session *discordgo.Session, interaction interactions.Interaction) interactions.InteractionResponse {
	embed := embedutil.NewEmbed().
		SetColor(0xc000ff).
		SetDescription(fmt.Sprintf("Click [here](https://discord.com/oauth2/authorize?client_id=%s&scope=applications.commands+bot&permissions=8) to invite me!", session.State.User.ID)).MessageEmbed
	return multiplexer.CreateEmbedResponse(embed)
}
