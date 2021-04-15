package commands

import (
	CommandHandler "git.randomchars.net/Reviath/RemiliaScarlet/handler"
	"git.randomchars.net/Reviath/RemiliaScarlet/sql"
)

func ResetAutoroleCommand(ctx CommandHandler.Context, _ []string) error {
	db := sql.Connect()

	type Tag struct {
		roleid string
	}

	var tag Tag

	if sql.CheckLanguage(ctx.Guild.ID) == "tr" {
		var tag Tag

		err := db.QueryRow("SELECT roleid FROM autorole WHERE guildid ='" + ctx.Guild.ID + "'").Scan(&tag.roleid)
		if err == nil {
			delete, err := db.Query("DELETE FROM autorole WHERE guildid ='" + ctx.Guild.ID + "'")
			if err != nil {
				ctx.Reply("Bir hata oluştu.")

				return nil
			}

			defer delete.Close()

			ctx.Reply("Başarıyla otorol sıfırlandı.")

			return nil
		} else {
			ctx.Reply("Otorol ayarlanmadığı için sıfırlanamaz.")

			return nil
		}
	}

	err := db.QueryRow("SELECT roleid FROM autorole WHERE guildid ='" + ctx.Guild.ID + "'").Scan(&tag.roleid)
	if err == nil {
		delete, err := db.Query("DELETE FROM autorole WHERE guildid ='" + ctx.Guild.ID + "'")
		if err != nil {
			ctx.Reply("An error occurred, please try again.")

			return nil
		}

		defer delete.Close()

		ctx.Reply("Successfully reset auto role.")

		return nil
	} else {
		ctx.Reply("Auto role is not existing, so you can't reset.")

		return nil
	}
}
