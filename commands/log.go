package commands

import (
	"strings"

	multiplexer "git.randomchars.net/Reviath/RemiliaScarlet/Multiplexer"
	CommandHandler "git.randomchars.net/Reviath/RemiliaScarlet/handler"
	"git.randomchars.net/Reviath/RemiliaScarlet/sql"
)

func LogCommand(ctx CommandHandler.Context, _ []string) error {
	db := sql.Connect()

	type Tag struct {
		channelid string
	}

	var tag Tag

	if sql.CheckLanguage(ctx.Guild.ID) == "tr" {
		if len(strings.Join(multiplexer.GetArgs(ctx.Message.Content, multiplexer.GetPrefix()), " ")) < 1 {
			ctx.Reply("Log kanalını belirtmelisin.")
			return nil
		}
		c, err := ctx.Session.Channel(multiplexer.GetChannel(multiplexer.GetArgs(ctx.Message.Content, multiplexer.GetPrefix())[0]))
		if err == nil {
			err = db.QueryRow("SELECT channelid FROM log WHERE guildid ='" + ctx.Guild.ID + "'").Scan(&tag.channelid)
			if err == nil {
				ctx.Reply("Log kanalı zaten ayarlanmış, sıfırlamak için reset_log komutunu kullan.")
				return nil
			} else {
				insert, err := db.Query("INSERT INTO log (channelid, guildid) VALUES ('" + c.ID + "', '" + ctx.Guild.ID + "')")
				if err != nil {
					ctx.Reply("Bir hata oluştu.")
					return nil
				}
				defer insert.Close()

				ctx.Reply("Log kanalı başarıyla ayarlandı!")

				return nil
			}
		} else {
			ctx.Reply("Log kanalını belirtmelisin.")
			return nil
		}
	}

	if len(strings.Join(multiplexer.GetArgs(ctx.Message.Content, multiplexer.GetPrefix()), " ")) < 1 {
		ctx.Reply("You need to specify the channel.")
		return nil
	}

	c, err := ctx.Session.Channel(multiplexer.GetChannel(multiplexer.GetArgs(ctx.Message.Content, multiplexer.GetPrefix())[0]))
	if err == nil {

		err = db.QueryRow("SELECT channelid FROM log WHERE guildid ='" + ctx.Guild.ID + "'").Scan(&tag.channelid)
		if err == nil {
			ctx.Reply("Log is already existing (to reset, use reset_log command).")

			return nil
		} else {
			insert, err := db.Query("INSERT INTO log (channelid, guildid) VALUES ('" + c.ID + "', '" + ctx.Guild.ID + "')")
			if err != nil {
				ctx.Reply("An error occurred, please try again.")

				return nil
			}
			defer insert.Close()

			ctx.Reply("Logging channel set successfully.")

			return nil
		}
	} else {
		ctx.Reply("You need to specify the channel.")

		return nil
	}
}
