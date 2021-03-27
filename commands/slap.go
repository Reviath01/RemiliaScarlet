package commands

import (
	ctx "git.randomchars.net/Reviath/handlers/Context"
	"github.com/bwmarrin/discordgo"
    embedutil "git.randomchars.net/Reviath/embed-util"
    "strings"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)

type Slap struct {

}

func (s Slap) Execute(ctx ctx.Ctx, session *discordgo.Session) error {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/remilia")

    if err != nil {
        panic(err.Error())
    }

    defer db.Close()

    type Tag struct {
        isblocked string `json:"isblocked"`
    }

    var tag Tag

    err = db.QueryRow("SELECT isblocked FROM disabledcommands WHERE commandname ='slap' AND guildid ='" + ctx.Guild().ID + "'").Scan(&tag.isblocked)

    if err == nil {
        if tag.isblocked == "True" {
            _, err = session.ChannelMessageSend(ctx.Channel().ID, "This command is blocked on this guild.")
            return err
        }
    }

    if len(strings.Join(ctx.Args(), " ")) != 22 {
        _, err := session.ChannelMessageSend(ctx.Channel().ID, "You need to specify the user.")
	        return err
    }
        u, err := session.User(strings.Join(ctx.Args()," ")[3:][:18])
        if err == nil {
        embed := embedutil.NewEmbed().
            SetColor(0xff1000).
            SetDescription("<@" + ctx.Author().ID + "> slaps <@" + u.ID + ">").
            SetImage("https://images-ext-1.discordapp.net/external/79sCWyD-TmmyjFxlaQIxAkAANAfV529d-LDHNkGDM0M/%3Fitemid%3D10426943/https/media1.tenor.com/images/b6d8a83eb652a30b95e87cf96a21e007/tenor.gif").MessageEmbed
	    _, err := session.ChannelMessageSendEmbed(ctx.Channel().ID, embed)
	    return err
        } else {
            _, err := session.ChannelMessageSend(ctx.Channel().ID, "You need to specify the user.")
	        return err
        }
    }
