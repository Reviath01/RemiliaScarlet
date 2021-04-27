package commands

import (
	"fmt"

	embedutil "git.randomchars.net/Reviath/RemiliaScarlet/EmbedUtil"
	CommandHandler "git.randomchars.net/Reviath/RemiliaScarlet/handler"
)

func VoteCommand(ctx CommandHandler.Context, _ []string) error {
	voteeembed := embedutil.NewEmbed().
		SetColor(0xc000ff).
		SetDescription(fmt.Sprintf("Click [here](https://top.gg/bot/%s/vote) to vote me on Top.gg!", ctx.Session.State.User.ID)).MessageEmbed
	ctx.ReplyEmbed(voteeembed)
	return nil
}
