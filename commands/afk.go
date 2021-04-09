package commands

import (
	"database/sql"

	ctx "git.randomchars.net/Reviath/RemiliaScarlet/Context"
	"github.com/bwmarrin/discordgo"
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
		isblocked string
		lang      string
	}

	var tag Tag

	err = db.QueryRow("SELECT language FROM languages WHERE guildid ='" + ctx.Guild().ID + "'").Scan(&tag.lang)

	if err == nil && tag.lang == "tr" {

		err = db.QueryRow("SELECT isblocked FROM disabledcommands WHERE commandname ='afk' AND guildid ='" + ctx.Guild().ID + "'").Scan(&tag.isblocked)

		if err == nil {
			if tag.isblocked == "True" {
				_, _ = session.ChannelMessageSend(ctx.Channel().ID, "Bu komut bu sunucuda engellenmiş.")
				return nil
			}
		}

		insert, err := db.Query("INSERT INTO afk (isafk, userid) VALUES ('true', '" + ctx.Author().ID + "')")

		if err != nil {
			_, _ = session.ChannelMessageSend(ctx.Channel().ID, "Bir hata oluştu.")
			return nil
		}

		defer insert.Close()

		_, _ = session.ChannelMessageSend(ctx.Channel().ID, "Başarıyla AFK moduna geçtin.")

		return nil
	}

	err = db.QueryRow("SELECT isblocked FROM disabledcommands WHERE commandname ='afk' AND guildid ='" + ctx.Guild().ID + "'").Scan(&tag.isblocked)

	if err == nil {
		if tag.isblocked == "True" {
			_, _ = session.ChannelMessageSend(ctx.Channel().ID, "This command is blocked on this guild.")
			return nil
		}
	}

	insert, err := db.Query("INSERT INTO afk (isafk, userid) VALUES ('true', '" + ctx.Author().ID + "')")

	if err != nil {
		_, _ = session.ChannelMessageSend(ctx.Channel().ID, "An error occurred.")
		return nil
	}

	defer insert.Close()

	_, _ = session.ChannelMessageSend(ctx.Channel().ID, "I set you as AFK.")

	return nil
}
