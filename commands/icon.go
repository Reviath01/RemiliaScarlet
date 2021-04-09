package commands

import (
	ctx "git.randomchars.net/Reviath/RemiliaScarlet/Context"
	"github.com/bwmarrin/discordgo"
    embedutil "git.randomchars.net/Reviath/embed-util"
	"database/sql"
    _ "github.com/go-sql-driver/mysql"
)

type Icon struct {

}

func (i Icon) Execute(ctx ctx.Ctx, session *discordgo.Session) error {
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

    err = db.QueryRow("SELECT language FROM languages WHERE guildid ='" + ctx.Guild().ID + "'").Scan(&tag.lang)
    if err == nil {
        if tag.lang == "tr" {
            err = db.QueryRow("SELECT isblocked FROM disabledcommands WHERE commandname ='icon' AND guildid ='" + ctx.Guild().ID + "'").Scan(&tag.isblocked)

    if err == nil {
        if tag.isblocked == "True" {
            _, err = session.ChannelMessageSend(ctx.Channel().ID, "Bu komut bu sunucuda engellenmiş.")
            
			if err != nil {
				return nil
			}
			
			return err
        }
    }

      iconeembed := embedutil.NewEmbed().
            SetColor(0x00f6ff).
            SetImage(ctx.Guild().IconURL()).MessageEmbed
	_, err = session.ChannelMessageSendEmbed(ctx.Channel().ID, iconeembed)
	
	if err != nil {
		return nil
	}
	
	return err

        }
    }

    err = db.QueryRow("SELECT isblocked FROM disabledcommands WHERE commandname ='icon' AND guildid ='" + ctx.Guild().ID + "'").Scan(&tag.isblocked)

    if err == nil {
        if tag.isblocked == "True" {
            _, err = session.ChannelMessageSend(ctx.Channel().ID, "This command is blocked on this guild.")
            
			if err != nil {
				return nil
			}
			
			return err
        }
    }

      iconeembed := embedutil.NewEmbed().
            SetColor(0x00f6ff).
            SetImage(ctx.Guild().IconURL()).MessageEmbed
	_, err = session.ChannelMessageSendEmbed(ctx.Channel().ID, iconeembed)
	
	if err != nil {
		return nil
	}
	
	return err
}
