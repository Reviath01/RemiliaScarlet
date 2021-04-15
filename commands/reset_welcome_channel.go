package commands

import (
	CommandHandler "git.randomchars.net/Reviath/RemiliaScarlet/handler"
	"git.randomchars.net/Reviath/RemiliaScarlet/sql"
)

func ResetWelcomeChannelCommand(ctx CommandHandler.Context, _ []string) error {
	db := sql.Connect()

	type Tag struct {
		channelid string
	}

	var tag Tag
	if sql.CheckLanguage(ctx.Guild.ID) == "tr" {
		err := db.QueryRow("SELECT channelid FROM welcomechannel WHERE guildid ='" + ctx.Guild.ID + "'").Scan(&tag.channelid)
		if err == nil {
			delete, err := db.Query("DELETE FROM welcomechannel WHERE guildid ='" + ctx.Guild.ID + "'")
			if err != nil {
				ctx.Reply("Bir hata oluştu.")
				return nil
			}

			defer delete.Close()

			ctx.Reply("Hoş geldin kanalı başarıyla sıfırlandı.")

			return nil
		} else {
			ctx.Reply("Hoş geldin kanalı ayarlanmamış, sıfırlayamazsın.")

			return nil
		}
	}
	err := db.QueryRow("SELECT channelid FROM welcomechannel WHERE guildid ='" + ctx.Guild.ID + "'").Scan(&tag.channelid)
	if err == nil {
		delete, err := db.Query("DELETE FROM welcomechannel WHERE guildid ='" + ctx.Guild.ID + "'")
		if err != nil {
			ctx.Reply("An error occurred, please try again.")

			return nil
		}

		defer delete.Close()

		ctx.Reply("Successfully reset welcome channel.")

		return nil
	} else {
		ctx.Reply("Welcome channel is not existing, so you can't reset.")

		return nil
	}
}
