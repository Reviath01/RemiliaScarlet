package commands

import (
	"fmt"

	embedutil "git.randomchars.net/Reviath/RemiliaScarlet/EmbedUtil"
	CommandHandler "git.randomchars.net/Reviath/RemiliaScarlet/handler"
)

func InviteCommand(ctx CommandHandler.Context, _ []string) error {
	inviteembed := embedutil.NewEmbed().
		SetColor(0xc000ff).
		SetDescription(fmt.Sprintf("Click [here](https://discord.com/oauth2/authorize?client_id=%s&scope=applications.commands+bot&permissions=8) to invite me!", ctx.Session.State.User.ID)).MessageEmbed
	ctx.ReplyEmbed(inviteembed)

	return nil
}
