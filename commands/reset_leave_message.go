package commands

import (
	"fmt"

	CommandHandler "git.randomchars.net/Reviath/RemiliaScarlet/handler"
	"git.randomchars.net/Reviath/RemiliaScarlet/sql"
)

func ResetLeaveMessage(ctx CommandHandler.Context, _ []string) error {
	db := sql.Connect()

	type Tag struct {
		message string
	}

	var tag Tag

	if sql.CheckLanguage(ctx.Guild.ID) == "tr" {
		err := db.QueryRow(fmt.Sprintf("SELECT message FROM leavemessage WHERE guildid ='%s'", ctx.Guild.ID)).Scan(&tag.message)
		if err == nil {
			delete, err := db.Query(fmt.Sprintf("DELETE FROM leavemessage WHERE guildid ='%s'", ctx.Guild.ID))
			if err != nil {
				ctx.Reply("Bir hata oluştu.")
				return nil
			}

			defer delete.Close()

			ctx.Reply("Başarıyla çıkış mesajı sıfırlandı.")

			return nil
		} else {
			ctx.Reply("Çıkış mesajı ayarlanmamış, sıfırlayamazsın.")

			return nil
		}
	}

	err := db.QueryRow(fmt.Sprintf("SELECT message FROM leavemessage WHERE guildid ='%s'", ctx.Guild.ID)).Scan(&tag.message)
	if err == nil {
		delete, err := db.Query(fmt.Sprintf("DELETE FROM leavemessage WHERE guildid ='%s'", ctx.Guild.ID))
		if err != nil {
			ctx.Reply("An error occurred, please try again.")

			return nil
		}

		defer delete.Close()

		ctx.Reply("Successfully reset leave message.")

		return nil
	} else {
		ctx.Reply("Leave message is not existing, so you can't reset.")

		return nil
	}
}
