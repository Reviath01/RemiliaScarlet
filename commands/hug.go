package commands

import (
	ctx "git.randomchars.net/Reviath/handlers/Context"
	"github.com/bwmarrin/discordgo"
    embedutil "git.randomchars.net/Reviath/embed-util"
    "strings"
)

type Hug struct {

}

func (h Hug) Execute(ctx ctx.Ctx, session *discordgo.Session) error {
    if strings.Join(ctx.Args(), " ") == "" {
        _, err4 := session.ChannelMessageSend(ctx.Channel().ID, "You need to specify the user.")
	        return err4
    }
    if len(strings.Join(ctx.Args(), " ")) < 18 {
        _, err4 := session.ChannelMessageSend(ctx.Channel().ID, "You need to specify the user.")
	        return err4
    }
    u, err := session.User(strings.Join(ctx.Args()," "))
    if err == nil {
        embed := embedutil.NewEmbed().
            SetColor(0xff1000).
            SetDescription("<@" + ctx.Author().ID + "> hugs <@" + u.ID + ">").
            SetImage("https://i.pinimg.com/originals/4d/d7/49/4dd749423de10a319b5d9e8850bbace4.gif").MessageEmbed
	    _, err2 := session.ChannelMessageSendEmbed(ctx.Channel().ID, embed)
	return err2
    } else {	
        u2, err := session.User(strings.Join(ctx.Args()," ")[3:][:18])
        if err == nil {
        embed := embedutil.NewEmbed().
            SetColor(0xff1000).
            SetDescription("<@" + ctx.Author().ID + "> hugs <@" + u2.ID + ">").
            SetImage("https://i.pinimg.com/originals/4d/d7/49/4dd749423de10a319b5d9e8850bbace4.gif").MessageEmbed
	    _, err3 := session.ChannelMessageSendEmbed(ctx.Channel().ID, embed)
	    return err3
        } else {
            _, err4 := session.ChannelMessageSend(ctx.Channel().ID, "You need to specify the user.")
	        return err4
        }
    }
}
