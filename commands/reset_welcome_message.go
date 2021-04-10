package commands

import (
	"database/sql"

	ctx "git.randomchars.net/Reviath/RemiliaScarlet/CommandHandler/Context"
	"github.com/bwmarrin/discordgo"
	_ "github.com/go-sql-driver/mysql"
)

type ResetWelcomeMessage struct {
}

func (r ResetWelcomeMessage) Execute(ctx ctx.Ctx, session *discordgo.Session) error {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/remilia")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	type Tag struct {
		message string
		lang    string
	}

	var tag Tag

	err = db.QueryRow("SELECT language FROM languages WHERE guildid ='" + ctx.Guild().ID + "'").Scan(&tag.lang)
	if err == nil && tag.lang == "tr" {
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
		} else {
			_, _ = session.ChannelMessageSend(ctx.Channel().ID, "Hoş geldin mesajı ayarlanmamış, sıfırlayamazsın.")

			return nil
		}
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
	} else {
		_, _ = session.ChannelMessageSend(ctx.Channel().ID, "Welcome message is not existing, so you can't reset.")

		return nil
	}
}
