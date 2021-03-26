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
    if len(strings.Join(ctx.Args()," ")) < 19 {
    u, err := session.User(strings.Join(ctx.Args()," "))
    if err == nil {
        avatarembed := embedutil.NewEmbed().
            SetColor(0xff1000).
            SetDescription("Avatar of " + u.Username + "#" + u.Discriminator).
            SetImage(u.AvatarURL("1024")).MessageEmbed
	    _, err := session.ChannelMessageSendEmbed(ctx.Channel().ID, avatarembed)
	return err 
    } else {
        avatarembed := embedutil.NewEmbed().
            SetColor(0xff1000).
            SetDescription("Avatar of " + ctx.Author().Username + "#" + ctx.Author().Discriminator).
            SetImage(ctx.Author().AvatarURL("1024")).MessageEmbed
	    _, err := session.ChannelMessageSendEmbed(ctx.Channel().ID, avatarembed)
	        return err
    }
} else {
    if len(strings.Join(ctx.Args()," ")) > 21 {
     u, err := session.User(strings.Join(ctx.Args()," ")[3:][:18])
        if err == nil {
        avatarembed := embedutil.NewEmbed().
            SetColor(0xff1000).
            SetDescription("Avatar of " + u.Username + "#" + u.Discriminator).
            SetImage(u.AvatarURL("1024")).MessageEmbed
	    _, err := session.ChannelMessageSendEmbed(ctx.Channel().ID, avatarembed)
	    return err
        } else {
        avatarembed := embedutil.NewEmbed().
            SetColor(0xff1000).
            SetDescription("Avatar of " + ctx.Author().Username + "#" + ctx.Author().Discriminator).
            SetImage(ctx.Author().AvatarURL("1024")).MessageEmbed
	    _, err := session.ChannelMessageSendEmbed(ctx.Channel().ID, avatarembed)
	        return err
        }
    } else {
        avatarembed := embedutil.NewEmbed().
            SetColor(0xff1000).
            SetImage(ctx.Author().AvatarURL("1024")).MessageEmbed
	    _, err := session.ChannelMessageSendEmbed(ctx.Channel().ID, avatarembed)
	        return err
    }
}
}
