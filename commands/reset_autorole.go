package commands

import (
	ctx "git.randomchars.net/Reviath/RemiliaScarlet/CommandHandler/Context"
	"git.randomchars.net/Reviath/RemiliaScarlet/sql"
	"github.com/bwmarrin/discordgo"
)

type ResetAutorole struct {
}

func (r ResetAutorole) Execute(ctx ctx.Ctx, session *discordgo.Session) error {
	db := sql.Connect()

	type Tag struct {
		roleid string
		lang   string
	}

	var tag Tag

	err := db.QueryRow("SELECT language FROM languages WHERE guildid ='" + ctx.Guild().ID + "'").Scan(&tag.lang)
	if err == nil && tag.lang == "tr" {
		perms, err := session.State.UserChannelPermissions(ctx.Author().ID, ctx.Channel().ID)
		if err == nil && !(int(perms)&discordgo.PermissionAdministrator == discordgo.PermissionAdministrator) {
			_, _ = session.ChannelMessageSend(ctx.Channel().ID, "Bu komutu kullanmak için yönetici yetkisine sahip olmalısınız.")

			return nil
		}

		var tag Tag

		err = db.QueryRow("SELECT roleid FROM autorole WHERE guildid ='" + ctx.Guild().ID + "'").Scan(&tag.roleid)
		if err == nil {
			delete, err := db.Query("DELETE FROM autorole WHERE guildid ='" + ctx.Guild().ID + "'")
			if err != nil {
				_, _ = session.ChannelMessageSend(ctx.Channel().ID, "Bir hata oluştu.")

				return nil
			}

			defer delete.Close()

			_, _ = session.ChannelMessageSend(ctx.Channel().ID, "Başarıyla otorol sıfırlandı.")

			return nil
		} else {
			_, _ = session.ChannelMessageSend(ctx.Channel().ID, "Otorol ayarlanmadığı için sıfırlanamaz.")

			return nil
		}
	}

	perms, err := session.State.UserChannelPermissions(ctx.Author().ID, ctx.Channel().ID)
	if err == nil && !(int(perms)&discordgo.PermissionAdministrator == discordgo.PermissionAdministrator) {
		_, _ = session.ChannelMessageSend(ctx.Channel().ID, "You need administrator permission to run this command.")

		return nil
	}

	err = db.QueryRow("SELECT roleid FROM autorole WHERE guildid ='" + ctx.Guild().ID + "'").Scan(&tag.roleid)
	if err == nil {
		delete, err := db.Query("DELETE FROM autorole WHERE guildid ='" + ctx.Guild().ID + "'")
		if err != nil {
			_, _ = session.ChannelMessageSend(ctx.Channel().ID, "An error occurred, please try again.")

			return nil
		}

		defer delete.Close()

		_, _ = session.ChannelMessageSend(ctx.Channel().ID, "Successfully reset auto role.")

		return nil
	} else {
		_, _ = session.ChannelMessageSend(ctx.Channel().ID, "Auto role is not existing, so you can't reset.")

		return nil
	}
}
