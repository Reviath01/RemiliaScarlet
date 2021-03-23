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
    if len(strings.Join(ctx.Args(), " ")) != 22 {
        _, err := session.ChannelMessageSend(ctx.Channel().ID, "You need to specify the user.")
	        return err
    }
        u, err := session.User(strings.Join(ctx.Args()," ")[3:][:18])
        if u.ID == ctx.Author().ID {
            _, err = session.ChannelMessageSend(ctx.Channel().ID, "You can't hug yourself.")
            return err
        }
        if err == nil {
        embed := embedutil.NewEmbed().
            SetColor(0xff1000).
            SetDescription("<@" + ctx.Author().ID + "> hugs <@" + u.ID + ">").
            SetImage("https://i.pinimg.com/originals/4d/d7/49/4dd749423de10a319b5d9e8850bbace4.gif").MessageEmbed
	    _, err := session.ChannelMessageSendEmbed(ctx.Channel().ID, embed)
	    return err
        } else {
            _, err := session.ChannelMessageSend(ctx.Channel().ID, "You need to specify the user.")
	        return err
        }
    }
