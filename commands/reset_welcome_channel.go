package commands

import (
	"fmt"

	multiplexer "git.randomchars.net/Reviath/RemiliaScarlet/Multiplexer"
	CommandHandler "git.randomchars.net/Reviath/RemiliaScarlet/handler"
	"git.randomchars.net/Reviath/RemiliaScarlet/sql"
)

func ResetWelcomeChannelCommand(ctx CommandHandler.Context, _ []string) error {
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
		err := db.QueryRow(fmt.Sprintf("SELECT channelid FROM welcomechannel WHERE guildid ='%s'", ctx.Guild.ID)).Scan(&tag.channelid)
		if err == nil {
			delete, err := db.Query(fmt.Sprintf("DELETE FROM welcomechannel WHERE guildid ='%s'", ctx.Guild.ID))
			if err != nil {
				ctx.Reply("Bir hata oluştu.")
				return nil
			}

			defer delete.Close()

			ctx.Reply("Hoş geldin kanalı başarıyla sıfırlandı.")

			return nil
		}
		ctx.Reply("Hoş geldin kanalı ayarlanmamış, sıfırlayamazsın.")
		return nil
	default:

		if !multiplexer.CheckAdministratorPermission(ctx.Session, ctx.Message.Author.ID, ctx.Channel.ID) {
			ctx.Reply("You don't have enough permission.")
			return nil
		}

		err := db.QueryRow(fmt.Sprintf("SELECT channelid FROM welcomechannel WHERE guildid ='%s'", ctx.Guild.ID)).Scan(&tag.channelid)
		if err == nil {
			delete, err := db.Query(fmt.Sprintf("DELETE FROM welcomechannel WHERE guildid ='%s'", ctx.Guild.ID))
			if err != nil {
				ctx.Reply("An error occurred, please try again.")

				return nil
			}

			defer delete.Close()

			ctx.Reply("Successfully reset welcome channel.")

			return nil
		}
		ctx.Reply("Welcome channel is not existing, so you can't reset.")
		return nil
	}
}
