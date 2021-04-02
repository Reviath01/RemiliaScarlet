package commands

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/bwmarrin/discordgo"
	ctx "git.randomchars.net/Reviath/handlers/Context"
	"strings"
)

type Language struct {

}

func (l Language) Execute(ctx ctx.Ctx, session *discordgo.Session) error {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/remilia")

    if err != nil {
        panic(err.Error())
    }

    defer db.Close()

	type Tag struct {
		lang string `json:"language"`
	}

	var tag Tag

	args := ctx.Args()

	err = db.QueryRow("SELECT language FROM languages WHERE guildid ='" + ctx.Guild().ID + "'").Scan(&tag.lang)
	if err == nil {
		if tag.lang == "tr" {
			if len(strings.Join(ctx.Args(), " ")) < 1 {
				_, err = session.ChannelMessageSend(ctx.Channel().ID, "Yeni dili belirtmelisiniz.")
				if err != nil {
					return nil
				}
				return err
			} else if args[0] == "tr" {
				_, err = session.ChannelMessageSend(ctx.Channel().ID, "Dil zaten Türkçe ayarlı.")
				if err != nil {
					return nil
				}
				return err
			} else if args[0] == "en" {
				_, err = session.ChannelMessageSend(ctx.Channel().ID, "Dili ingilizce olarak ayarlıyorum...")
				delete, err := db.Query("DELETE FROM languages WHERE guildid ='" + ctx.Guild().ID + "'")
				if err != nil {
					_, err = session.ChannelMessageSend(ctx.Channel().ID, "Bir hata oluştu.")
					if err != nil {
						return nil
					}
					return err 
				}
				return err

				defer delete.Close()
			} else {
				_, err = session.ChannelMessageSend(ctx.Channel().ID, "Kullanılabilir diller: `tr`, `en`")
				if err != nil {
					return nil
				}
				return err
			}
		} else {
			if len(strings.Join(ctx.Args(), " ")) < 1 {
				_, err = session.ChannelMessageSend(ctx.Channel().ID, "You need to specify the language.")
				if err != nil {
					return nil
				}
				return err
			} else if args[0] == "tr" {
				_,err = session.ChannelMessageSend(ctx.Channel().ID, "Setting language as Turkish...")
					
				if err != nil {
					return nil
				}
	
				insert, err := db.Query("INSERT INTO languages (language, guildid) VALUES ('tr', '" + ctx.Guild().ID + "')")
	
				if err != nil {
					return nil
				}
	
				defer insert.Close()
	
			} else if args[0] == "en" {
				_, err = session.ChannelMessageSend(ctx.Channel().ID, "Language is already en.")
				if err != nil {
					return nil
				}
				return err
			} else {
				_, err = session.ChannelMessageSend(ctx.Channel().ID, "Language options: `tr`, `en`")
				if err != nil {
					return nil
				}
				return err
			}
		}
		return nil
	} else {
		if len(strings.Join(ctx.Args(), " ")) < 1 {
			_, err = session.ChannelMessageSend(ctx.Channel().ID, "You need to specify the language.")
			
			if err != nil {
				return nil
			}

			return err
		} else if args[0] == "tr" {
			_,err = session.ChannelMessageSend(ctx.Channel().ID, "Setting language as Turkish...")
				
			if err != nil {
				return nil
			}

			insert, err := db.Query("INSERT INTO languages (language, guildid) VALUES ('tr', '" + ctx.Guild().ID + "')")

			if err != nil {
				return nil
			}

			defer insert.Close()

		} else if args[0] == "en" {
			_, err = session.ChannelMessageSend(ctx.Channel().ID, "Language is already en.")
			if err != nil {
				return nil
			}
			return err
		} else {
			_, err = session.ChannelMessageSend(ctx.Channel().ID, "Language options: `tr`, `en`")
			if err != nil {
				return nil
			}
			return err
		}
	}
	return nil
}