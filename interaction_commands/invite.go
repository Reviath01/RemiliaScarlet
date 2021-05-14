package interaction_commands

import (
	"fmt"

	embedutil "git.randomchars.net/FreeNitori/EmbedUtil"
	multiplexer "git.randomchars.net/Reviath/RemiliaScarlet/Multiplexer"
	"git.randomchars.net/Reviath/RemiliaScarlet/interactions"
	"github.com/bwmarrin/discordgo"
)

//Invite slash command.
func InviteCommand(session *discordgo.Session, interaction interactions.Interaction) interactions.InteractionResponse {
	embed := embedutil.New("", fmt.Sprintf("Click [here](https://discord.com/oauth2/authorize?client_id=%s&scope=applications.commands+bot&permissions=8) to invite me!", session.State.User.ID))
	embed.Color = 0xc000ff
	return multiplexer.CreateEmbedResponse(embed)
}
