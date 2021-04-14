package commands

import (
	"strings"

	ctx "git.randomchars.net/Reviath/RemiliaScarlet/CommandHandler/Context"
	"git.randomchars.net/Reviath/RemiliaScarlet/sql"
	"github.com/bwmarrin/discordgo"
)

type WelcomeMessage struct {
}

func (w WelcomeMessage) Execute(ctx ctx.Ctx, session *discordgo.Session) error {
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
		if strings.Join(ctx.Args(), " ") == "" {
			_, _ = session.ChannelMessageSend(ctx.Channel().ID, "Mesajı belirtmelisin.")

			return nil
		}

		if len(strings.Join(ctx.Args(), " ")) > 254 {
			_, _ = session.ChannelMessageSend(ctx.Channel().ID, "Mesaj maksimum 255 karakter uzunluğunda olabilir.")

			return nil
		}

		err = db.QueryRow("SELECT message FROM welcomemessage WHERE guildid ='" + ctx.Guild().ID + "'").Scan(&tag.message)
		if err == nil {
			_, _ = session.ChannelMessageSend(ctx.Channel().ID, "Hoş geldin mesajı zaten ayarlanmış, sıfırlamak için reset_welcome_message komutunu kullan.")

			return nil
		}
		insert, err := db.Query("INSERT INTO welcomemessage (message, guildid) VALUES ('" + strings.Join(ctx.Args(), " ") + "', '" + ctx.Guild().ID + "')")
		if err != nil {
			_, _ = session.ChannelMessageSend(ctx.Channel().ID, "Bir hata oluştu.")

			return nil
		}
		defer insert.Close()

		_, _ = session.ChannelMessageSend(ctx.Channel().ID, "Hoş geldin mesajı başarıyla ayarlandı.")

		return nil
	}

	perms, err := session.State.UserChannelPermissions(ctx.Author().ID, ctx.Channel().ID)
	if err == nil && !(int(perms)&discordgo.PermissionAdministrator == discordgo.PermissionAdministrator) {
		_, _ = session.ChannelMessageSend(ctx.Channel().ID, "You need administrator permission to run this command.")

		return nil
	}
	if strings.Join(ctx.Args(), " ") == "" {
		_, _ = session.ChannelMessageSend(ctx.Channel().ID, "You need to specify the message.")

		return nil
	}

	if len(strings.Join(ctx.Args(), " ")) > 254 {
		_, _ = session.ChannelMessageSend(ctx.Channel().ID, "Your message can be up to 255 characters long.")

		return nil
	}

	err = db.QueryRow("SELECT message FROM welcomemessage WHERE guildid ='" + ctx.Guild().ID + "'").Scan(&tag.message)
	if err == nil {
		_, _ = session.ChannelMessageSend(ctx.Channel().ID, "Welcome message is already existing (to reset, use reset_welcome_message command).")

		return nil
	}
	insert, err := db.Query("INSERT INTO welcomemessage (message, guildid) VALUES ('" + strings.Join(ctx.Args(), " ") + "', '" + ctx.Guild().ID + "')")
	if err != nil {
		_, _ = session.ChannelMessageSend(ctx.Channel().ID, "An error occurred, please try again.")

		return nil
	}
	defer insert.Close()

	_, _ = session.ChannelMessageSend(ctx.Channel().ID, "Welcome message set successfully.")

	return nil
}
