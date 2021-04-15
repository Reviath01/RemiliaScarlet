package commands

import (
	CommandHandler "git.randomchars.net/Reviath/RemiliaScarlet/handler"
	"git.randomchars.net/Reviath/RemiliaScarlet/sql"
)

func ResetWelcomeMessageCommand(ctx CommandHandler.Context, _ []string) error {
	db := sql.Connect()

	type Tag struct {
		message string
	}

	var tag Tag

	if sql.CheckLanguage(ctx.Guild.ID) == "tr" {
		err := db.QueryRow("SELECT message FROM welcomemessage WHERE guildid ='" + ctx.Guild.ID + "'").Scan(&tag.message)
		if err == nil {
			delete, err := db.Query("DELETE FROM welcomemessage WHERE guildid ='" + ctx.Guild.ID + "'")
			if err != nil {
				ctx.Reply("Bir hata oluştu.")
				return nil
			}

			defer delete.Close()

			ctx.Reply("Hoş geldin mesajı başarıyla sıfırlandı.")

			return nil
		} else {
			ctx.Reply("Hoş geldin mesajı ayarlanmamış, sıfırlayamazsın.")

			return nil
		}
	}

	err := db.QueryRow("SELECT message FROM welcomemessage WHERE guildid ='" + ctx.Guild.ID + "'").Scan(&tag.message)
	if err == nil {
		delete, err := db.Query("DELETE FROM welcomemessage WHERE guildid ='" + ctx.Guild.ID + "'")
		if err != nil {
			ctx.Reply("An error occurred, please try again.")
			return nil
		}
		defer delete.Close()

		ctx.Reply("Successfully reset welcome message.")
		return nil
	} else {
		ctx.Reply("Welcome message is not existing, so you can't reset.")
		return nil
	}
}
