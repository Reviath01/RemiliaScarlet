package commands

import (
	ctx "git.randomchars.net/Reviath/handlers/Context"
	"github.com/bwmarrin/discordgo"
    embedutil "git.randomchars.net/Reviath/embed-util"
    "strings"
)

type Avatar struct {

}

func (a Avatar) Execute(ctx ctx.Ctx, session *discordgo.Session) error {
    u, err := session.User(strings.Join(ctx.Args()," "))
    if err == nil {
        avatarembed2 := embedutil.NewEmbed().
            SetColor(0xff1000).
            SetImage(u.AvatarURL("avatar")).MessageEmbed
	    _, err2 := session.ChannelMessageSendEmbed(ctx.Channel().ID, avatarembed2)
	return err2
    } else {
        avatarembed := embedutil.NewEmbed().
            SetColor(0xff1000).
            SetImage(ctx.Author().AvatarURL("avatar")).MessageEmbed
	    _, err := session.ChannelMessageSendEmbed(ctx.Channel().ID, avatarembed)
	    return err
    }
}
