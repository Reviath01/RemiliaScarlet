package commands

import (
	ctx "git.randomchars.net/Reviath/RemiliaScarlet/CommandHandler/Context"
	"git.randomchars.net/Reviath/RemiliaScarlet/sql"
	"github.com/bwmarrin/discordgo"
)

type ResetWelcomeMessage struct {
}

func (r ResetWelcomeMessage) Execute(ctx ctx.Ctx, session *discordgo.Session) error {
	db := sql.Connect()

	type Tag struct {
		message string
	}

	var tag Tag

	if sql.CheckLanguage(ctx.Guild().ID) == "tr" {
		perms, err := session.State.UserChannelPermissions(ctx.Author().ID, ctx.Channel().ID)
		if err == nil && !(int(perms)&discordgo.PermissionAdministrator == discordgo.PermissionAdministrator) {
			_, _ = session.ChannelMessageSend(ctx.Channel().ID, "Bu komutu kullanmak için yönetici yetkisine sahip olmalısın.")

			return nil
		}

		err = db.QueryRow("SELECT message FROM welcomemessage WHERE guildid ='" + ctx.Guild().ID + "'").Scan(&tag.message)
		if err == nil {
			delete, err := db.Query("DELETE FROM welcomemessage WHERE guildid ='" + ctx.Guild().ID + "'")
			if err != nil {
				_, _ = session.ChannelMessageSend(ctx.Channel().ID, "Bir hata oluştu.")

				return nil
			}

			defer delete.Close()

			_, _ = session.ChannelMessageSend(ctx.Channel().ID, "Hoş geldin mesajı başarıyla sıfırlandı.")

			return nil
		}
		_, _ = session.ChannelMessageSend(ctx.Channel().ID, "Hoş geldin mesajı ayarlanmamış, sıfırlayamazsın.")
		return nil
	}

	perms, err := session.State.UserChannelPermissions(ctx.Author().ID, ctx.Channel().ID)
	if err == nil && !(int(perms)&discordgo.PermissionAdministrator == discordgo.PermissionAdministrator) {
		_, _ = session.ChannelMessageSend(ctx.Channel().ID, "You need administrator permission to run this command.")

		return nil
	}

	err = db.QueryRow("SELECT message FROM welcomemessage WHERE guildid ='" + ctx.Guild().ID + "'").Scan(&tag.message)
	if err == nil {
		delete, err := db.Query("DELETE FROM welcomemessage WHERE guildid ='" + ctx.Guild().ID + "'")
		if err != nil {
			_, _ = session.ChannelMessageSend(ctx.Channel().ID, "An error occurred, please try again.")

			return nil
		}

		defer delete.Close()

		_, _ = session.ChannelMessageSend(ctx.Channel().ID, "Successfully reset welcome message.")

		return nil
	}
	_, _ = session.ChannelMessageSend(ctx.Channel().ID, "Welcome message is not existing, so you can't reset.")
	return nil
}
