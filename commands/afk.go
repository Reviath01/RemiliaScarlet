package commands

import (
	ctx "git.randomchars.net/Reviath/RemiliaScarlet/CommandHandler/Context"
	"git.randomchars.net/Reviath/RemiliaScarlet/sql"
	"github.com/bwmarrin/discordgo"
)

type Afk struct {
}

func (a Afk) Execute(ctx ctx.Ctx, session *discordgo.Session) error {
	db := sql.Connect()

	if sql.CheckLanguage(ctx.Guild().ID) == "tr" {

		if sql.IsBlocked(ctx.Guild().ID, "afk") == "true" {
			_, _ = session.ChannelMessageSend(ctx.Channel().ID, "Bu komut bu sunucuda engellenmiş.")
			return nil
		}

		insert, err := db.Query("INSERT INTO afk (isafk, userid) VALUES ('true', '" + ctx.Author().ID + "')")

		if err != nil {
			_, _ = session.ChannelMessageSend(ctx.Channel().ID, "Bir hata oluştu.")
			return nil
		}

		defer insert.Close()

		_, _ = session.ChannelMessageSend(ctx.Channel().ID, "Başarıyla AFK moduna geçtin.")

		return nil
	}

	if sql.IsBlocked(ctx.Guild().ID, "afk") == "true" {
		_, _ = session.ChannelMessageSend(ctx.Channel().ID, "This command is blocked on this guild.")
		return nil
	}

	insert, err := db.Query("INSERT INTO afk (isafk, userid) VALUES ('true', '" + ctx.Author().ID + "')")

	if err != nil {
		_, _ = session.ChannelMessageSend(ctx.Channel().ID, "An error occurred.")
		return nil
	}

	defer insert.Close()

	_, _ = session.ChannelMessageSend(ctx.Channel().ID, "I set you as AFK.")

	return nil
}
