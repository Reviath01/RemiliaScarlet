package commands

import (
	ctx "git.randomchars.net/Reviath/handlers/Context"
	"github.com/bwmarrin/discordgo"
    embedutil "git.randomchars.net/Reviath/embed-util"
    "strings"
)

type Kiss struct {

}

func (k Kiss) Execute(ctx ctx.Ctx, session *discordgo.Session) error {
    if len(strings.Join(ctx.Args(), " ")) != 22 {
        _, err := session.ChannelMessageSend(ctx.Channel().ID, "You need to specify the user.")
	        return err
    }
        u, err := session.User(strings.Join(ctx.Args()," ")[3:][:18])
        if err == nil {
        embed := embedutil.NewEmbed().
            SetColor(0xff1000).
            SetDescription("<@" + ctx.Author().ID + "> kisses <@" + u.ID + ">").
            SetImage("https://media.tenor.com/images/d68747a5865b12c465e5dff31c65d5c2/tenor.gif").MessageEmbed
	    _, err := session.ChannelMessageSendEmbed(ctx.Channel().ID, embed)
	    return err
        } else {
            _, err := session.ChannelMessageSend(ctx.Channel().ID, "You need to specify the user.")
	        return err
        }
    }
