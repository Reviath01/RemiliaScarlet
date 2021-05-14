package commands

import (
	"fmt"

	embedutil "git.randomchars.net/FreeNitori/EmbedUtil"
	CommandHandler "git.randomchars.net/Reviath/RemiliaScarlet/handler"
)

// VoteCommand is a handler for vote command
func VoteCommand(ctx CommandHandler.Context, _ []string) error {
	voteembed := embedutil.New("", fmt.Sprintf("Click [here](https://top.gg/bot/%s/vote) to vote me on Top.gg!", ctx.Session.State.User.ID))
	voteembed.Color = 0xc000ff
	ctx.ReplyEmbed(voteembed)
	return nil
}
