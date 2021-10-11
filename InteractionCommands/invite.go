package interactioncommands

import (
	"fmt"

	multiplexer "git.randomchars.net/Reviath/RemiliaScarlet/Multiplexer"
	"git.randomchars.net/Reviath/RemiliaScarlet/interactions"
	embedutil "git.randomchars.net/freenitori/embedutil"
	"github.com/bwmarrin/discordgo"
)

// InviteCommand is invite command for interactions
func InviteCommand(session *discordgo.Session, interaction interactions.Interaction) interactions.InteractionResponse {
	embed := embedutil.New("", fmt.Sprintf("Click [here](https://discord.com/oauth2/authorize?client_id=%s&scope=applications.commands+bot&permissions=8) to invite me!", session.State.User.ID))
	embed.Color = 0xc000ff
	return multiplexer.CreateEmbedResponse(embed)
}
