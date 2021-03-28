package commands

import (
	ctx "git.randomchars.net/Reviath/handlers/Context"
	"github.com/bwmarrin/discordgo"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)

type Afk struct {

}

func (a Afk) Execute(ctx ctx.Ctx, session *discordgo.Session) error {
    db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/remilia")

    if err != nil {
        panic(err.Error())
    }

    defer db.Close()

    type Tag struct {
        isblocked string `json:"isblocked"`
    }

    var tag Tag

    err = db.QueryRow("SELECT isblocked FROM disabledcommands WHERE commandname ='afk' AND guildid ='" + ctx.Guild().ID + "'").Scan(&tag.isblocked)

    if err == nil {
        if tag.isblocked == "True" {
            _, err = session.ChannelMessageSend(ctx.Channel().ID, "This command is blocked on this guild.")
            if err != nil {
                return nil
            }
            return err
        }
    }

    insert, err := db.Query("INSERT INTO afk (isafk, userid) VALUES ('true', '" + ctx.Author().ID + "')")

    if err != nil {
        _, err := session.ChannelMessageSend(ctx.Channel().ID, "An error occured.")
        return err
    }

       defer insert.Close()

        _, err = session.ChannelMessageSend(ctx.Channel().ID, "I set you as AFK.")
        
        if err != nil {
            return nil
        }
        
        return err
}
