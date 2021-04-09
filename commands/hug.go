package commands

import (
	"database/sql"
	"strings"

	ctx "git.randomchars.net/Reviath/RemiliaScarlet/Context"
	embedutil "git.randomchars.net/Reviath/embed-util"
	"github.com/bwmarrin/discordgo"
	_ "github.com/go-sql-driver/mysql"
)

type Hug struct {
}

func (h Hug) Execute(ctx ctx.Ctx, session *discordgo.Session) error {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/remilia")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	type Tag struct {
		isblocked string `json:"isblocked"`
		lang      string `json:"language"`
	}

	var tag Tag

	err = db.QueryRow("SELECT language FROM languages WHERE guildid ='" + ctx.Guild().ID + "'").Scan(&tag.lang)
	if err == nil && tag.lang == "tr" {
		err = db.QueryRow("SELECT isblocked FROM disabledcommands WHERE commandname ='hug' AND guildid ='" + ctx.Guild().ID + "'").Scan(&tag.isblocked)

		if err == nil {
			if tag.isblocked == "True" {
				_, err = session.ChannelMessageSend(ctx.Channel().ID, "Bu komut bu sunucuda engellenmiş.")

				if err != nil {
					return nil
				}

				return err
			}
		}

		var args string
		if len(strings.Join(ctx.Args(), " ")) < 1 {
			_, err := session.ChannelMessageSend(ctx.Channel().ID, "Bir üye belirtmelisin.")

			if err != nil {
				return nil
			}

			return err

		}

		args = ctx.Args()[0]

		if len(args) != 22 {
			_, err := session.ChannelMessageSend(ctx.Channel().ID, "Bir üye belirtmelisin.")

			if err != nil {
				return nil
			}

			return err
		}
		u, err := session.User(strings.Join(ctx.Args(), " ")[3:][:18])
		if err == nil {
			if u.ID == ctx.Author().ID {
				_, err = session.ChannelMessageSend(ctx.Channel().ID, "Kendine sarılamazsın.")

				if err != nil {
					return nil
				}

				return err
			}
			embed := embedutil.NewEmbed().
				SetColor(0xff1000).
				SetDescription("<@" + ctx.Author().ID + ">, <@" + u.ID + "> isimli kişiye sarıldı 💖").
				SetImage("https://i.pinimg.com/originals/4d/d7/49/4dd749423de10a319b5d9e8850bbace4.gif").MessageEmbed
			_, err := session.ChannelMessageSendEmbed(ctx.Channel().ID, embed)

			if err != nil {
				return nil
			}

			return err
		} else {
			_, err := session.ChannelMessageSend(ctx.Channel().ID, "Bir kişiyi etiketlemelisin.")

			if err != nil {
				return nil
			}

			return err
		}
	}

	err = db.QueryRow("SELECT isblocked FROM disabledcommands WHERE commandname ='hug' AND guildid ='" + ctx.Guild().ID + "'").Scan(&tag.isblocked)

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
	if len(strings.Join(ctx.Args(), " ")) < 1 {
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
	u, err := session.User(strings.Join(ctx.Args(), " ")[3:][:18])
	if err == nil {
		if u.ID == ctx.Author().ID {
			_, err = session.ChannelMessageSend(ctx.Channel().ID, "You can't hug yourself.")

			if err != nil {
				return nil
			}

			return err
		}
		embed := embedutil.NewEmbed().
			SetColor(0xff1000).
			SetDescription("<@" + ctx.Author().ID + "> hugs <@" + u.ID + ">").
			SetImage("https://i.pinimg.com/originals/4d/d7/49/4dd749423de10a319b5d9e8850bbace4.gif").MessageEmbed
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
