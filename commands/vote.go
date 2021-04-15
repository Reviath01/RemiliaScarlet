package commands

import (
	embedutil "git.randomchars.net/Reviath/RemiliaScarlet/EmbedUtil"
	CommandHandler "git.randomchars.net/Reviath/RemiliaScarlet/handler"
)

func VoteCommand(ctx CommandHandler.Context, _ []string) error {
	voteeembed := embedutil.NewEmbed().
		SetColor(0xc000ff).
		SetDescription("Click [here](https://top.gg/bot/" + ctx.Session.State.User.ID + "/vote) to vote me on Top.gg!").MessageEmbed
	ctx.ReplyEmbed(voteeembed)
	return nil
}
