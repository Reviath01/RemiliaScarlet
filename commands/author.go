package commands

import (
	ctx "git.randomchars.net/Reviath/handlers/Context"
	"github.com/bwmarrin/discordgo"
	embedutil "git.randomchars.net/Reviath/embed-util"
	"database/sql"
    _ "github.com/go-sql-driver/mysql"
)

type Author struct {

}

func (a Author) Execute(ctx ctx.Ctx, session *discordgo.Session) error {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/remilia")

    if err != nil {
        panic(err.Error())
    }

    defer db.Close()

    type Tag struct {
        isblocked string `json:"isblocked"`
    }

    var tag Tag

    err = db.QueryRow("SELECT isblocked FROM disabledcommands WHERE commandname ='author' AND guildid ='" + ctx.Guild().ID + "'").Scan(&tag.isblocked)

    if err == nil {
        if tag.isblocked == "True" {
            _, err = session.ChannelMessageSend(ctx.Channel().ID, "This command is blocked on this guild.")
            return err
        }
    }

    authorembed := embedutil.NewEmbed().
    SetColor(0x007bff).
    AddField("My Author:", "<@770218429096656917> ([Reviath#0001](https://discord.com/users/770218429096656917))").MessageEmbed
	_, err = session.ChannelMessageSendEmbed(ctx.Channel().ID, authorembed)
	
    if err != nil {
        return nil
    }

    return err
}