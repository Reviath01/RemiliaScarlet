package commands

import (
	ctx "git.randomchars.net/Reviath/handlers/Context"
	"github.com/bwmarrin/discordgo"
    embedutil "git.randomchars.net/Reviath/embed-util"
    "strings"
)

type StartVote struct {

}

func (s StartVote) Execute(ctx ctx.Ctx, session *discordgo.Session) error {
    perms, err := session.State.UserChannelPermissions(ctx.Author().ID, ctx.Channel().ID)
	if err == nil && (int(perms)&discordgo.PermissionManageMessages == discordgo.PermissionManageMessages) == false {
        _, err := session.ChannelMessageSend(ctx.Channel().ID, "You need manage messages permission to run this command.")
	return err
    }
    if strings.Join(ctx.Args()," ") == "" {
    _, err := session.ChannelMessageSend(ctx.Channel().ID, "You need to specify a message.")
    return err
    }
    embed := embedutil.NewEmbed().
    SetTitle("Vote started!").
    SetColor(0xff9100).
    AddField("Vote question:", strings.Join(ctx.Args()," ")).MessageEmbed
	msg, err := session.ChannelMessageSendEmbed(ctx.Channel().ID, embed)
    session.MessageReactionAdd(ctx.Channel().ID, msg.ID, "👍")
	session.MessageReactionAdd(ctx.Channel().ID, msg.ID, "👎")
	return err
}
