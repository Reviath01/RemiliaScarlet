package commands

import (
	"fmt"
	"strings"

	multiplexer "git.randomchars.net/Reviath/RemiliaScarlet/Multiplexer"
	CommandHandler "git.randomchars.net/Reviath/RemiliaScarlet/handler"
	"git.randomchars.net/Reviath/RemiliaScarlet/sql"
)

//Welcome channel command
func WelcomeChannelCommand(ctx CommandHandler.Context, _ []string) error {
	db := sql.Connect()

	type Tag struct {
		channelid string
	}

	var tag Tag

	switch sql.CheckLanguage(ctx.Guild.ID) {
	case "tr":
		if !multiplexer.CheckAdministratorPermission(ctx.Session, ctx.Message.Author.ID, ctx.Channel.ID) {
			ctx.Reply("Yeterli yetkiye sahip değilsin.")
			return nil
		}
		if len(strings.Join(multiplexer.GetArgs(ctx.Message.Content, multiplexer.GetPrefix()), " ")) < 1 {
			ctx.Reply("Bir kanal belirtmelisin.")

			return nil
		}
		c, err := ctx.Session.Channel(multiplexer.GetUser(multiplexer.GetArgs(ctx.Message.Content, multiplexer.GetPrefix())[0]))
		if err == nil {
			err = db.QueryRow(fmt.Sprintf("SELECT channelid FROM welcomechannel WHERE guildid ='%s'", ctx.Guild.ID)).Scan(&tag.channelid)
			if err == nil {
				ctx.Reply("Hoş geldin kanalı zaten ayarlanmış, sıfırlamak için reset_welcome_channel komutunu kullan.")

				return nil
			}
			insert, err := db.Query(fmt.Sprintf("INSERT INTO welcomechannel (channelid, guildid) VALUES ('%s', '%s')", c.ID, ctx.Guild.ID))
			if err != nil {
				ctx.Reply("Bir hata oluştu.")
				return nil
			}
			defer insert.Close()

			ctx.Reply("Hoş geldin kanalı başarıyla ayarlandı.")
			return nil
		}
		ctx.Reply("Bir kanal belirtmelisin.")
		return nil
	default:

		if !multiplexer.CheckAdministratorPermission(ctx.Session, ctx.Message.Author.ID, ctx.Channel.ID) {
			ctx.Reply("You don't have enough permission.")
			return nil
		}

		if len(strings.Join(multiplexer.GetArgs(ctx.Message.Content, multiplexer.GetPrefix()), " ")) < 1 {
			ctx.Reply("You need to specify the channel.")

			return nil
		}
		c, err := ctx.Session.Channel(multiplexer.GetUser(multiplexer.GetArgs(ctx.Message.Content, multiplexer.GetPrefix())[0]))
		if err == nil {
			err = db.QueryRow(fmt.Sprintf("SELECT channelid FROM welcomechannel WHERE guildid ='%s'", ctx.Guild.ID)).Scan(&tag.channelid)
			if err == nil {
				ctx.Reply("Welcome channel is already existing (to reset, use reset_welcome_channel command).")

				return nil
			}
			insert, err := db.Query(fmt.Sprintf("INSERT INTO welcomechannel (channelid, guildid) VALUES ('%s', '%s')", c.ID, ctx.Guild.ID))
			if err != nil {
				ctx.Reply("An error occurred, please try again.")
				return nil
			}
			defer insert.Close()

			ctx.Reply("Welcome channel set successfully.")
			return nil
		}
		ctx.Reply("You need to specify the channel.")
		return nil
	}
}
