package commands

import (
	"fmt"

	CommandHandler "git.randomchars.net/Reviath/RemiliaScarlet/handler"
	embedutil "git.randomchars.net/freenitori/embedutil"
)

// InviteCommand is a handler for invite command
func InviteCommand(ctx CommandHandler.Context, _ []string) error {
	inviteembed := embedutil.New("", fmt.Sprintf("Click [here](https://discord.com/oauth2/authorize?client_id=%s&scope=applications.commands+bot&permissions=8) to invite me!", ctx.Session.State.User.ID))
	inviteembed.Color = 0xc000ff
	ctx.ReplyEmbed(inviteembed)

	return nil
}
