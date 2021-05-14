package commands

import (
	"strings"

	multiplexer "git.randomchars.net/Reviath/RemiliaScarlet/Multiplexer"
	CommandHandler "git.randomchars.net/Reviath/RemiliaScarlet/handler"
	"git.randomchars.net/Reviath/RemiliaScarlet/sql"
)

// LanguageCommand is a handler for language command
func LanguageCommand(ctx CommandHandler.Context, _ []string) error {
	db := sql.Connect()

	args := multiplexer.GetArgs(ctx.Message.Content, multiplexer.GetPrefix())

	if sql.CheckLanguage(ctx.Guild.ID) == "tr" {
		if !multiplexer.CheckAdministratorPermission(ctx.Session, ctx.Message.Author.ID, ctx.Channel.ID) {
			ctx.Reply("Yeterli yetkiye sahip değilsin.")
			return nil
		}
		if len(strings.Join(multiplexer.GetArgs(ctx.Message.Content, multiplexer.GetPrefix()), " ")) < 1 {
			ctx.Reply("Yeni dili belirtmelisiniz.")

			return nil
		} else if args[0] == "tr" {
			ctx.Reply("Dil zaten Türkçe ayarlı.")

			return nil
		} else if args[0] == "en" {
			ctx.Reply("Dili ingilizce olarak ayarlıyorum...")

			delete, err := db.Query("DELETE FROM languages WHERE guildid ='" + ctx.Guild.ID + "'")
			if err != nil {
				ctx.Reply("Bir hata oluştu.")

				return nil
			}

			defer delete.Close()
		} else {
			ctx.Reply("Kullanılabilir diller: `tr`, `en`")

			return nil
		}
	} else {
		if !multiplexer.CheckAdministratorPermission(ctx.Session, ctx.Message.Author.ID, ctx.Channel.ID) {
			ctx.Reply("You don't have enough permission.")
			return nil
		}
		if len(strings.Join(multiplexer.GetArgs(ctx.Message.Content, multiplexer.GetPrefix()), " ")) < 1 {
			ctx.Reply("You need to specify the language.")

			return nil
		} else if args[0] == "tr" {
			ctx.Reply("Setting language as Turkish...")

			insert, err := db.Query("INSERT INTO languages (language, guildid) VALUES ('tr', '" + ctx.Guild.ID + "')")
			if err != nil {
				return nil
			}

			ctx.Reply("Dil Türkçe yapıldı!")

			defer insert.Close()

			return nil

		} else if args[0] == "en" {
			ctx.Reply("Language is en.")

			return nil
		} else {
			ctx.Reply("Language options: `tr`, `en`")

			return nil
		}
	}
	return nil
}
