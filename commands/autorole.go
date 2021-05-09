package commands

import (
	"fmt"
	"strings"

	multiplexer "git.randomchars.net/Reviath/RemiliaScarlet/Multiplexer"
	CommandHandler "git.randomchars.net/Reviath/RemiliaScarlet/handler"
	"git.randomchars.net/Reviath/RemiliaScarlet/sql"
)

func AutoRoleCommand(ctx CommandHandler.Context, _ []string) error {
	db := sql.Connect()

	type Tag struct {
		roleid string
	}

	var tag Tag

	switch sql.CheckLanguage(ctx.Guild.ID) {
	case "tr":
		if !multiplexer.CheckAdministratorPermission(ctx.Session, ctx.Message.Author.ID, ctx.Channel.ID) {
			ctx.Reply("Yeterli yetkiye sahip değilsin.")
			return nil
		}
		if strings.Join(multiplexer.GetArgs(ctx.Message.Content, multiplexer.GetPrefix()), " ") == "" {
			ctx.Reply("Rolü belirtmelisin.")
			return nil
		}
		args := multiplexer.GetRole(multiplexer.GetArgs(ctx.Message.Content, multiplexer.GetPrefix())[0])
		if args == "Error" {
			ctx.Reply("Rolü belirtmelisin.")
			return nil
		}
		err := db.QueryRow(fmt.Sprintf("SELECT roleid FROM autorole WHERE guildid ='%s'", ctx.Guild.ID)).Scan(&tag.roleid)
		if err == nil {
			ctx.Reply("Otorol zaten ayarlanmış, tekrar ayarlamak için reset_autorole komutunu kullan!")
			return nil
		}
		insert, err := db.Query(fmt.Sprintf("INSERT INTO autorole (roleid, guildid) VALUES ('%s', '%s')", args, ctx.Guild.ID))
		if err != nil {
			ctx.Reply("Bir hata oluştu.")
			return nil
		}
		defer insert.Close()

		ctx.Reply("Otorol başarıyla ayarlandı.")

		return nil
	default:
		if !multiplexer.CheckAdministratorPermission(ctx.Session, ctx.Message.Author.ID, ctx.Channel.ID) {
			ctx.Reply("You don't have enough permission.")
			return nil
		}

		if strings.Join(multiplexer.GetArgs(ctx.Message.Content, multiplexer.GetPrefix()), " ") == "" {
			ctx.Reply("You need to specify the role.")
			return nil
		}
		args := multiplexer.GetRole(multiplexer.GetArgs(ctx.Message.Content, multiplexer.GetPrefix())[0])
		if args == "Error" {
			ctx.Reply("You need to specify the role.")
			return nil
		}
		err := db.QueryRow(fmt.Sprintf("SELECT roleid FROM autorole WHERE guildid ='%s'", ctx.Guild.ID)).Scan(&tag.roleid)
		if err == nil {
			ctx.Reply("Autorole is already existing, use reset_auto_role command to reset!")
			return nil
		}
		insert, err := db.Query(fmt.Sprintf("INSERT INTO autorole (roleid, guildid) VALUES ('%s', '%s')", args, ctx.Guild.ID))
		if err != nil {
			ctx.Reply("An error occurred.")
			return nil
		}
		defer insert.Close()
		ctx.Reply("Successfully set autorole.")
		return nil
	}
}
