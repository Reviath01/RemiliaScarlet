package commands

import (
	"fmt"
	"strings"

	multiplexer "git.randomchars.net/Reviath/RemiliaScarlet/Multiplexer"
	CommandHandler "git.randomchars.net/Reviath/RemiliaScarlet/handler"
	"git.randomchars.net/Reviath/RemiliaScarlet/sql"
)

// LeaveMessageCommand is a handler for leave message command
func LeaveMessageCommand(ctx CommandHandler.Context, _ []string) error {
	db := sql.Connect()
	defer db.Close()
	type Tag struct {
		message string
	}

	var tag Tag

	switch sql.CheckLanguage(ctx.Guild.ID) {
	case "tr":
		if !multiplexer.CheckAdministratorPermission(ctx.Session, ctx.Message.Author.ID, ctx.Guild.ID) {
			ctx.Reply("Yeterli yetkiye sahip değilsin.")
			return nil
		}
		if len(strings.Join(multiplexer.GetArgs(ctx.Message.Content, multiplexer.GetPrefix()), " ")) < 1 || len(strings.Join(multiplexer.GetArgs(ctx.Message.Content, multiplexer.GetPrefix()), " ")) > 254 {
			ctx.Reply("Mesajınız 1 ve 255 karakter aralığında olmalıdır.")

			return nil
		}

		err := db.QueryRow(fmt.Sprintf("SELECT message FROM leavemessage WHERE guildid ='%s'", ctx.Guild.ID)).Scan(&tag.message)
		if err == nil {
			ctx.Reply("Çıkış mesajı zaten ayarlanmış, reset'lemek için reset_leave_message komutunu kullan.")

			return nil
		}
		insert, err := db.Query(fmt.Sprintf("INSERT INTO leavemessage (message, guildid) VALUES ('%s', '%s')", strings.Join(multiplexer.GetArgs(ctx.Message.Content, multiplexer.GetPrefix()), " "), ctx.Guild.ID))
		if err != nil {
			ctx.Reply("Bir hata oluştu.")
			return nil
		}
		defer insert.Close()
		ctx.Reply("Başarıyla çıkış mesajı ayarlandı.")
		return nil
	default:

		if !multiplexer.CheckAdministratorPermission(ctx.Session, ctx.Message.Author.ID, ctx.Guild.ID) {
			ctx.Reply("You don't have enough permission.")
			return nil
		}

		if len(strings.Join(multiplexer.GetArgs(ctx.Message.Content, multiplexer.GetPrefix()), " ")) < 1 || len(strings.Join(multiplexer.GetArgs(ctx.Message.Content, multiplexer.GetPrefix()), " ")) > 254 {
			ctx.Reply("Your message must be between 1 and 255 characters.")
			return nil
		}

		err := db.QueryRow(fmt.Sprintf("SELECT message FROM leavemessage WHERE guildid ='%s'", ctx.Guild.ID)).Scan(&tag.message)
		if err == nil {
			ctx.Reply("Leave message is already existing (to reset, use reset_leave_message command).")
			return nil
		}
		insert, err := db.Query(fmt.Sprintf("INSERT INTO leavemessage (message, guildid) VALUES ('%s', '%s')", strings.Join(multiplexer.GetArgs(ctx.Message.Content, multiplexer.GetPrefix()), " "), ctx.Guild.ID))
		if err != nil {
			ctx.Reply("An error occurred, please try again.")
			return nil
		}
		defer insert.Close()
		ctx.Reply("Leave message set successfully.")
		return nil
	}
}
