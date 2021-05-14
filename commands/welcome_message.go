package commands

import (
	"fmt"
	"strings"

	multiplexer "git.randomchars.net/Reviath/RemiliaScarlet/Multiplexer"
	CommandHandler "git.randomchars.net/Reviath/RemiliaScarlet/handler"
	"git.randomchars.net/Reviath/RemiliaScarlet/sql"
)

// WelcomeMessageCommand is a handler for welcome message command
func WelcomeMessageCommand(ctx CommandHandler.Context, _ []string) error {
	db := sql.Connect()

	type Tag struct {
		message string
	}

	var tag Tag

	switch sql.CheckLanguage(ctx.Guild.ID) {
	case "tr":
		if strings.Join(multiplexer.GetArgs(ctx.Message.Content, multiplexer.GetPrefix()), " ") == "" {
			ctx.Reply("Mesajı belirtmelisin.")

			return nil
		}

		if len(strings.Join(multiplexer.GetArgs(ctx.Message.Content, multiplexer.GetPrefix()), " ")) > 254 {
			ctx.Reply("Mesaj maksimum 255 karakter uzunluğunda olabilir.")

			return nil
		}

		err := db.QueryRow(fmt.Sprintf("SELECT message FROM welcomemessage WHERE guildid ='%s'", ctx.Guild.ID)).Scan(&tag.message)
		if err == nil {
			ctx.Reply("Hoş geldin mesajı zaten ayarlanmış, sıfırlamak için reset_welcome_message komutunu kullan.")

			return nil
		}
		insert, err := db.Query(fmt.Sprintf("INSERT INTO welcomemessage (message, guildid) VALUES ('%s', '%s')", strings.Join(multiplexer.GetArgs(ctx.Message.Content, multiplexer.GetPrefix()), " "), ctx.Guild.ID))
		if err != nil {
			ctx.Reply("Bir hata oluştu.")

			return nil
		}
		defer insert.Close()
		ctx.Reply("Hoş geldin mesajı başarıyla ayarlandı.")
		return nil
	default:
		if strings.Join(multiplexer.GetArgs(ctx.Message.Content, multiplexer.GetPrefix()), " ") == "" {
			ctx.Reply("You need to specify the message.")

			return nil
		}

		if len(strings.Join(multiplexer.GetArgs(ctx.Message.Content, multiplexer.GetPrefix()), " ")) > 254 {
			ctx.Reply("Your message can be up to 255 characters long.")

			return nil
		}

		err := db.QueryRow(fmt.Sprintf("SELECT message FROM welcomemessage WHERE guildid ='%s'", ctx.Guild.ID)).Scan(&tag.message)
		if err == nil {
			ctx.Reply("Welcome message is already existing (to reset, use reset_welcome_message command).")
			return nil
		}
		insert, err := db.Query(fmt.Sprintf("INSERT INTO welcomemessage (message, guildid) VALUES ('%s', '%s')", strings.Join(multiplexer.GetArgs(ctx.Message.Content, multiplexer.GetPrefix()), " "), ctx.Guild.ID))
		if err != nil {
			ctx.Reply("An error occurred, please try again.")

			return nil
		}
		defer insert.Close()

		ctx.Reply("Welcome message set successfully.")

		return nil
	}
}
