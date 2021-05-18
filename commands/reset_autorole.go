package commands

import (
	"fmt"

	multiplexer "git.randomchars.net/Reviath/RemiliaScarlet/Multiplexer"
	CommandHandler "git.randomchars.net/Reviath/RemiliaScarlet/handler"
	"git.randomchars.net/Reviath/RemiliaScarlet/sql"
)

// ResetAutoroleCommand is a handler for reset auto role command
func ResetAutoroleCommand(ctx CommandHandler.Context, _ []string) error {
	db := sql.Connect()

	type Tag struct {
		roleid string
	}

	var tag Tag

	switch sql.CheckLanguage(ctx.Guild.ID) {
	case "tr":
		if !multiplexer.CheckAdministratorPermission(ctx.Session, ctx.Message.Author.ID, ctx.Guild.ID) {
			ctx.Reply("Yeterli yetkiye sahip değilsin.")
			return nil
		}
		var tag Tag

		err := db.QueryRow(fmt.Sprintf("SELECT roleid FROM autorole WHERE guildid ='%s'", ctx.Guild.ID)).Scan(&tag.roleid)
		if err == nil {
			delete, err := db.Query(fmt.Sprintf("DELETE FROM autorole WHERE guildid ='%s'", ctx.Guild.ID))
			if err != nil {
				ctx.Reply("Bir hata oluştu.")

				return nil
			}

			defer delete.Close()

			ctx.Reply("Başarıyla otorol sıfırlandı.")

			return nil
		}
		ctx.Reply("Otorol ayarlanmadığı için sıfırlanamaz.")
		return nil
	default:

		if !multiplexer.CheckAdministratorPermission(ctx.Session, ctx.Message.Author.ID, ctx.Guild.ID) {
			ctx.Reply("You don't have enough permission.")
			return nil
		}

		err := db.QueryRow(fmt.Sprintf("SELECT roleid FROM autorole WHERE guildid ='%s'", ctx.Guild.ID)).Scan(&tag.roleid)
		if err == nil {
			delete, err := db.Query(fmt.Sprintf("DELETE FROM autorole WHERE guildid ='%s'", ctx.Guild.ID))
			if err != nil {
				ctx.Reply("An error occurred, please try again.")

				return nil
			}

			defer delete.Close()

			ctx.Reply("Successfully reset auto role.")

			return nil
		}
		ctx.Reply("Auto role is not existing, so you can't reset.")
		return nil
	}
}
