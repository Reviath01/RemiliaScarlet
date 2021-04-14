package commands

import (
	ctx "git.randomchars.net/Reviath/RemiliaScarlet/CommandHandler/Context"
	"git.randomchars.net/Reviath/RemiliaScarlet/sql"
	"github.com/bwmarrin/discordgo"
)

type ResetLog struct {
}

func (r ResetLog) Execute(ctx ctx.Ctx, session *discordgo.Session) error {
	db := sql.Connect()

	type Tag struct {
		channelid string
	}

	var tag Tag

	if sql.CheckLanguage(ctx.Guild().ID) == "tr" {
		perms, err := session.State.UserChannelPermissions(ctx.Author().ID, ctx.Channel().ID)
		if err == nil && !(int(perms)&discordgo.PermissionAdministrator == discordgo.PermissionAdministrator) {
			_, _ = session.ChannelMessageSend(ctx.Channel().ID, "Bu komutu kullanmak için yönetici yetkisine sahip olmalısın.")

			return nil
		}

		err = db.QueryRow("SELECT channelid FROM log WHERE guildid ='" + ctx.Guild().ID + "'").Scan(&tag.channelid)
		if err == nil {
			delete, err := db.Query("DELETE FROM log WHERE guildid ='" + ctx.Guild().ID + "'")
			if err != nil {
				_, _ = session.ChannelMessageSend(ctx.Channel().ID, "Bir hata oluştu.")

				return nil
			}

			defer delete.Close()

			_, _ = session.ChannelMessageSend(ctx.Channel().ID, "Başarıyla log kanalı sıfırlandı.")

			return nil
		}
		_, _ = session.ChannelMessageSend(ctx.Channel().ID, "Log kanalı ayarlanmamış, sıfırlayamazsın.")
		return nil
	}

	perms, err := session.State.UserChannelPermissions(ctx.Author().ID, ctx.Channel().ID)
	if err == nil && !(int(perms)&discordgo.PermissionAdministrator == discordgo.PermissionAdministrator) {
		_, _ = session.ChannelMessageSend(ctx.Channel().ID, "You need administrator permission to run this command.")

		return nil
	}

	err = db.QueryRow("SELECT channelid FROM log WHERE guildid ='" + ctx.Guild().ID + "'").Scan(&tag.channelid)
	if err == nil {
		delete, err := db.Query("DELETE FROM log WHERE guildid ='" + ctx.Guild().ID + "'")
		if err != nil {
			_, _ = session.ChannelMessageSend(ctx.Channel().ID, "An error occurred, please try again.")

			return nil
		}

		defer delete.Close()

		_, _ = session.ChannelMessageSend(ctx.Channel().ID, "Successfully reset log.")

		return nil
	}
	_, _ = session.ChannelMessageSend(ctx.Channel().ID, "Log channel is not existing, so you can't reset.")
	return nil
}
