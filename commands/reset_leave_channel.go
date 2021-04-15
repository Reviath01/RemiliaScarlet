package commands

import (
	CommandHandler "git.randomchars.net/Reviath/RemiliaScarlet/handler"
	"git.randomchars.net/Reviath/RemiliaScarlet/sql"
)

func ResetLeaveChannel(ctx CommandHandler.Context, _ []string) error {
	db := sql.Connect()

	type Tag struct {
		channelid string
	}

	var tag Tag

	if sql.CheckLanguage(ctx.Guild.ID) == "tr" {
		err := db.QueryRow("SELECT channelid FROM leavechannel WHERE guildid ='" + ctx.Guild.ID + "'").Scan(&tag.channelid)
		if err == nil {
			delete, err := db.Query("DELETE FROM leavechannel WHERE guildid ='" + ctx.Guild.ID + "'")
			if err != nil {
				ctx.Reply("Bir hata oluştu!")
				return nil
			}

			defer delete.Close()

			ctx.Reply("Başarıyla çıkış kanalı sıfırlandı.")

			return nil
		} else {
			ctx.Reply("Çıkış kanalı ayarlanmamış, yani sıfırlayamazsın.")

			return nil
		}
	}

	err := db.QueryRow("SELECT channelid FROM leavechannel WHERE guildid ='" + ctx.Guild.ID + "'").Scan(&tag.channelid)
	if err == nil {
		delete, err := db.Query("DELETE FROM leavechannel WHERE guildid ='" + ctx.Guild.ID + "'")
		if err != nil {
			ctx.Reply("An error occurred, please try again.")

			return nil
		}

		defer delete.Close()

		ctx.Reply("Successfully reset leave channel.")

		return nil
	} else {
		ctx.Reply("Leave channel is not existing, so you can't reset.")

		return nil
	}
}
