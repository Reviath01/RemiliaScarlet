package commands

import (
	ctx "git.randomchars.net/Reviath/handlers/Context"
	"github.com/bwmarrin/discordgo"
    embedutil "git.randomchars.net/Reviath/embed-util"
    "strings"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)

type Kiss struct {

}

func (k Kiss) Execute(ctx ctx.Ctx, session *discordgo.Session) error {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/remilia")

    if err != nil {
        panic(err.Error())
    }

    defer db.Close()

    type Tag struct {
        isblocked string `json:"isblocked"`
    }

    var tag Tag

    err = db.QueryRow("SELECT isblocked FROM disabledcommands WHERE commandname ='kiss' AND guildid ='" + ctx.Guild().ID + "'").Scan(&tag.isblocked)

    if err == nil {
        if tag.isblocked == "True" {
            _, err = session.ChannelMessageSend(ctx.Channel().ID, "This command is blocked on this guild.")
            
            if err != nil {
                return nil
            }
            
            return err
        }
    }

    var args string
    if len(strings.Join(ctx.Args()," ")) < 1 {
        _, err := session.ChannelMessageSend(ctx.Channel().ID, "You need to specify the user.")
	    
        if err != nil {
            return nil
        }
        
        return err
    }
    args = ctx.Args()[0]

    if len(args) != 22 {
        _, err := session.ChannelMessageSend(ctx.Channel().ID, "You need to specify the user.")
	    
        if err != nil {
            return nil
        }
        
        return err
    }
        u, err := session.User(args[3:][:18])
        if err == nil {
        embed := embedutil.NewEmbed().
            SetColor(0xff1000).
            SetDescription("<@" + ctx.Author().ID + "> kisses <@" + u.ID + ">").
            SetImage("https://media.tenor.com/images/d68747a5865b12c465e5dff31c65d5c2/tenor.gif").MessageEmbed
	    _, err := session.ChannelMessageSendEmbed(ctx.Channel().ID, embed)
	    
        if err != nil {
            return nil
        }
        
        return err
        } else {
            _, err := session.ChannelMessageSend(ctx.Channel().ID, "You need to specify the user.")
	        
            if err != nil {
                return nil
            }
            
            return err
        }
    }
