package commands

import (
	embedutil "git.randomchars.net/Reviath/RemiliaScarlet/EmbedUtil"
	CommandHandler "git.randomchars.net/Reviath/RemiliaScarlet/handler"
)

func InviteCommand(ctx CommandHandler.Context, _ []string) error {
	inviteembed := embedutil.NewEmbed().
		SetColor(0xc000ff).
		SetDescription("Click [here](https://discord.com/oauth2/authorize?client_id=" + ctx.Session.State.User.ID + "&scope=bot&permissions=8) to invite me!").MessageEmbed
	ctx.ReplyEmbed(inviteembed)

	return nil
}
