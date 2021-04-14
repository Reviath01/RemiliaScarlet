package commands

import (
	CommandHandler "git.randomchars.net/Reviath/RemiliaScarlet/handler"
	"git.randomchars.net/Reviath/RemiliaScarlet/sql"
)

func AfkCommand(ctx CommandHandler.Context, _ []string) error {
	db := sql.Connect()

	if sql.CheckLanguage(ctx.Guild.ID) == "tr" {

		if sql.IsBlocked(ctx.Guild.ID, "afk") == "true" {
			_, _ = ctx.Reply("Bu komut bu sunucuda kullanıma kapatılmış.")
			return nil
		}

		insert, err := db.Query("INSERT INTO afk (isafk, userid) VALUES ('true', '" + ctx.Message.Author.ID + "')")

		if err != nil {
			_, _ = ctx.Reply("Bir hata oluştu.")
			return nil
		}

		defer insert.Close()

		_, _ = ctx.Reply("Başarıyla AFK moduna geçtin.")

		return nil
	}

	if sql.IsBlocked(ctx.Guild.ID, "afk") == "true" {
		_, _ = ctx.Reply("This command is blocked on this guild.")
		return nil
	}

	insert, err := db.Query("INSERT INTO afk (isafk, userid) VALUES ('true', '" + ctx.Message.Author.ID + "')")

	if err != nil {
		ctx.Reply("An error occurred.")
		return nil
	}

	defer insert.Close()

	_, _ = ctx.Reply("I set you as AFK.")

	return nil
}
