package commands

import (
	"database/sql"
	"strings"

	ctx "git.randomchars.net/Reviath/RemiliaScarlet/Context"
	embedutil "git.randomchars.net/Reviath/embed-util"
	"github.com/bwmarrin/discordgo"
	_ "github.com/go-sql-driver/mysql"
)

type Embed struct {
}

func (e Embed) Execute(ctx ctx.Ctx, session *discordgo.Session) error {
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
		err = db.QueryRow("SELECT isblocked FROM disabledcommands WHERE commandname ='embed' AND guildid ='" + ctx.Guild().ID + "'").Scan(&tag.isblocked)

		if err == nil {
			if tag.isblocked == "True" {
				_, err = session.ChannelMessageSend(ctx.Channel().ID, "Bu komut bu sunucuda engellenmiş.")

				if err != nil {
					return nil
				}

				return err
			}
		}

		if strings.Join(ctx.Args(), " ") == "" {
			_, err := session.ChannelMessageSend(ctx.Channel().ID, "Bir mesaj belirtmelisin.")

			if err != nil {
				return nil
			}

			return err
		}
		embed := embedutil.NewEmbed().
			SetColor(0xc000ff).
			SetDescription(strings.Join(ctx.Args(), " ")).MessageEmbed
		_, err = session.ChannelMessageSendEmbed(ctx.Channel().ID, embed)

		if err != nil {
			return nil
		}

		return err

	}

	err = db.QueryRow("SELECT isblocked FROM disabledcommands WHERE commandname ='embed' AND guildid ='" + ctx.Guild().ID + "'").Scan(&tag.isblocked)

	if err == nil {
		if tag.isblocked == "True" {
			_, err = session.ChannelMessageSend(ctx.Channel().ID, "This command is blocked on this guild.")

			if err != nil {
				return nil
			}

			return err
		}
	}

	if strings.Join(ctx.Args(), " ") == "" {
		_, err := session.ChannelMessageSend(ctx.Channel().ID, "You need to specify the message.")

		if err != nil {
			return nil
		}

		return err
	}
	embed := embedutil.NewEmbed().
		SetColor(0xc000ff).
		SetDescription(strings.Join(ctx.Args(), " ")).MessageEmbed
	_, err = session.ChannelMessageSendEmbed(ctx.Channel().ID, embed)

	if err != nil {
		return nil
	}

	return err
}
