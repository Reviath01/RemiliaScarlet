package commands

import (
	"database/sql"
	"strings"

	ctx "git.randomchars.net/Reviath/RemiliaScarlet/Context"
	"github.com/bwmarrin/discordgo"
	_ "github.com/go-sql-driver/mysql"
)

type LeaveMessage struct {
}

func (l LeaveMessage) Execute(ctx ctx.Ctx, session *discordgo.Session) error {
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
		if len(strings.Join(ctx.Args(), " ")) < 1 || len(strings.Join(ctx.Args(), " ")) > 254 {
			_, _ = session.ChannelMessageSend(ctx.Channel().ID, "Mesajınız 1 ve 255 karakter aralığında olmalıdır.")

			return nil
		}

		err = db.QueryRow("SELECT message FROM leavemessage WHERE guildid ='" + ctx.Guild().ID + "'").Scan(&tag.message)
		if err == nil {
			_, _ = session.ChannelMessageSend(ctx.Channel().ID, "Çıkış mesajı zaten ayarlanmış, reset'lemek için reset_leave_message komutunu kullan.")

			return nil
		} else {
			insert, err := db.Query("INSERT INTO leavemessage (message, guildid) VALUES ('" + strings.Join(ctx.Args(), " ") + "', '" + ctx.Guild().ID + "')")
			if err != nil {
				_, _ = session.ChannelMessageSend(ctx.Channel().ID, "Bir hata oluştu.")
				return nil
			}
			defer insert.Close()

			_, _ = session.ChannelMessageSend(ctx.Channel().ID, "Başarıyla çıkış mesajı ayarlandı.")

			return nil
		}
	}

	perms, err := session.State.UserChannelPermissions(ctx.Author().ID, ctx.Channel().ID)
	if err == nil && !(int(perms)&discordgo.PermissionAdministrator == discordgo.PermissionAdministrator) {
		_, _ = session.ChannelMessageSend(ctx.Channel().ID, "You need administrator permission to run this command.")

		return nil
	}
	if len(strings.Join(ctx.Args(), " ")) < 1 || len(strings.Join(ctx.Args(), " ")) > 254 {
		_, _ = session.ChannelMessageSend(ctx.Channel().ID, "Your message must be between 1 and 255 characters.")
		return nil
	}

	err = db.QueryRow("SELECT message FROM leavemessage WHERE guildid ='" + ctx.Guild().ID + "'").Scan(&tag.message)
	if err == nil {
		_, _ = session.ChannelMessageSend(ctx.Channel().ID, "Leave message is already existing (to reset, use reset_leave_message command).")
		return nil
	} else {
		insert, err := db.Query("INSERT INTO leavemessage (message, guildid) VALUES ('" + strings.Join(ctx.Args(), " ") + "', '" + ctx.Guild().ID + "')")
		if err != nil {
			_, _ = session.ChannelMessageSend(ctx.Channel().ID, "An error occurred, please try again.")
			return nil
		}
		defer insert.Close()

		_, _ = session.ChannelMessageSend(ctx.Channel().ID, "Leave message set successfully.")

		return nil
	}
}
