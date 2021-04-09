package commands

import (
	ctx "git.randomchars.net/Reviath/RemiliaScarlet/Context"
	"github.com/bwmarrin/discordgo"
	"strconv"
	"database/sql"
    _ "github.com/go-sql-driver/mysql"
)

type Ping struct {

}

func (p Ping) Execute(ctx ctx.Ctx, session *discordgo.Session) error {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/remilia")

    if err != nil {
        panic(err.Error())
    }

    defer db.Close()

    type Tag struct {
        isblocked string `json:"isblocked"`
        lang string `json:"language"`
    }

    var tag Tag

    err = db.QueryRow("").Scan(&tag.lang)
    if err == nil {
        if tag.lang == "tr" {
            err = db.QueryRow("SELECT isblocked FROM disabledcommands WHERE commandname ='ping' AND guildid ='" + ctx.Guild().ID + "'").Scan(&tag.isblocked)

            if err == nil {
                if tag.isblocked == "True" {
                    _, err = session.ChannelMessageSend(ctx.Channel().ID, "Bu komut bu sunucuda engellenmiş.")
                    
                    if err != nil {
                        return nil
                    }
                    
                    return err
                }
            }
        
            _, err = session.ChannelMessageSend(ctx.Channel().ID, "Pong! " + strconv.Itoa(int(session.HeartbeatLatency().Milliseconds())) + "ms")
            return err
                
        }
    }

    err = db.QueryRow("SELECT isblocked FROM disabledcommands WHERE commandname ='ping' AND guildid ='" + ctx.Guild().ID + "'").Scan(&tag.isblocked)

    if err == nil {
        if tag.isblocked == "True" {
            _, err = session.ChannelMessageSend(ctx.Channel().ID, "This command is blocked on this guild.")
            
			if err != nil {
				return nil
			}
			
			return err
        }
    }

	_, err = session.ChannelMessageSend(ctx.Channel().ID, "Pong! " + strconv.Itoa(int(session.HeartbeatLatency().Milliseconds())) + "ms")
	return err
}
