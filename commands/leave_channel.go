package commands

import (
	"fmt"
	"strings"

	multiplexer "git.randomchars.net/Reviath/RemiliaScarlet/Multiplexer"
	CommandHandler "git.randomchars.net/Reviath/RemiliaScarlet/handler"
	"git.randomchars.net/Reviath/RemiliaScarlet/sql"
)

func LeaveChannelCommand(ctx CommandHandler.Context, _ []string) error {
	db := sql.Connect()

	type Tag struct {
		channelid string
	}

	var tag Tag
	if sql.CheckLanguage(ctx.Guild.ID) == "tr" {
		if len(strings.Join(multiplexer.GetArgs(ctx.Message.Content, multiplexer.GetPrefix()), " ")) < 1 {
			ctx.Reply("Bir kanal belirtmelisin")

			return nil
		}
		c, err := ctx.Session.Channel(multiplexer.GetChannel(multiplexer.GetArgs(ctx.Message.Content, multiplexer.GetPrefix())[0]))
		if err == nil {
			err = db.QueryRow(fmt.Sprintf("SELECT channelid FROM leavechannel WHERE guildid ='%s'", ctx.Guild.ID)).Scan(&tag.channelid)
			if err == nil {
				ctx.Reply("Çıkış kanalı zaten ayarlanmış reset'lemek için reset_leave_channel komutunu kullanın.")
				return nil
			} else {
				insert, err := db.Query(fmt.Sprintf("INSERT INTO leavechannel (channelid, guildid) VALUES ('%s', '%s')", c.ID, ctx.Guild.ID))
				if err != nil {
					ctx.Reply("Bir hata oluştu.")

					return nil
				}
				defer insert.Close()
				ctx.Reply("Başarıyla ayarlandı.")
				return nil
			}
		} else {
			ctx.Reply("Bir kanal belirtmelisin.")
			return nil
		}
	}

	if len(strings.Join(multiplexer.GetArgs(ctx.Message.Content, multiplexer.GetPrefix()), " ")) < 1 {
		ctx.Reply("You need to specify the channel.")

		return nil
	}

	c, err := ctx.Session.Channel(multiplexer.GetChannel(multiplexer.GetArgs(ctx.Message.Content, multiplexer.GetPrefix())[0]))
	if err == nil {
		err = db.QueryRow(fmt.Sprintf("SELECT channelid FROM leavechannel WHERE guildid ='%s'", ctx.Guild.ID)).Scan(&tag.channelid)
		if err == nil {
			ctx.Reply("Leave channel is already existing (to reset, use reset_leave_channel command).")

			return nil
		} else {
			insert, err := db.Query(fmt.Sprintf("INSERT INTO leavechannel (channelid, guildid) VALUES ('%s', '%s')", c.ID, ctx.Guild.ID))
			if err != nil {
				ctx.Reply("An error occurred, please try again.")

				return nil
			}
			defer insert.Close()

			ctx.Reply("Leave channel set successfully.")

			return nil
		}
	} else {
		ctx.Reply("You need to specify the channel.")
		return nil
	}
}
