package commands

import (
	"strings"

	ctx "git.randomchars.net/Reviath/RemiliaScarlet/CommandHandler/Context"
	multiplexer "git.randomchars.net/Reviath/RemiliaScarlet/Multiplexer"
	"git.randomchars.net/Reviath/RemiliaScarlet/sql"
	"github.com/bwmarrin/discordgo"
)

type AutoRole struct {
}

func (a AutoRole) Execute(ctx ctx.Ctx, session *discordgo.Session) error {
	db := sql.Connect()

	type Tag struct {
		roleid string
	}

	var tag Tag

	if sql.CheckLanguage(ctx.Guild().ID) == "tr" {
		perms, err := session.State.UserChannelPermissions(ctx.Author().ID, ctx.Channel().ID)
		if err == nil && !(int(perms)&discordgo.PermissionAdministrator == discordgo.PermissionAdministrator) {
			_, _ = session.ChannelMessageSend(ctx.Channel().ID, "Bu komutu kullanmak için yönetici yetkisine sahip olmalısın.")
			return nil
		}
		if strings.Join(ctx.Args(), " ") == "" {
			_, _ = session.ChannelMessageSend(ctx.Channel().ID, "Rolü belirtmelisin.")
			return nil
		} else {
			args := multiplexer.GetRole(ctx.Args()[0])
			if args == "Error" {
				_, _ = session.ChannelMessageSend(ctx.Channel().ID, "Rolü belirtmelisin.")
				return nil
			}
			err = db.QueryRow("SELECT roleid FROM autorole WHERE guildid ='" + ctx.Guild().ID + "'").Scan(&tag.roleid)
			if err == nil {
				_, _ = session.ChannelMessageSend(ctx.Channel().ID, "Otorol zaten ayarlanmış, tekrar ayarlamak için reset_autorole komutunu kullan!")
				return nil
			} else {
				insert, err := db.Query("INSERT INTO autorole (roleid, guildid) VALUES ('" + args + "', '" + ctx.Guild().ID + "')")
				if err != nil {
					_, _ = session.ChannelMessageSend(ctx.Channel().ID, "Bir hata oluştu.")

					return nil
				}
				defer insert.Close()

				_, _ = session.ChannelMessageSend(ctx.Channel().ID, "Otorol başarıyla ayarlandı.")

				return nil
			}
		}
	}
	perms, err := session.State.UserChannelPermissions(ctx.Author().ID, ctx.Channel().ID)
	if err == nil && !(int(perms)&discordgo.PermissionAdministrator == discordgo.PermissionAdministrator) {
		_, _ = session.ChannelMessageSend(ctx.Channel().ID, "You need administrator permission to run this command.")
		return nil
	}

	if strings.Join(ctx.Args(), " ") == "" {
		_, _ = session.ChannelMessageSend(ctx.Channel().ID, "You need to specify the role.")
		return nil
	} else {
		args := multiplexer.GetRole(ctx.Args()[0])
		if args == "Error" {
			_, _ = session.ChannelMessageSend(ctx.Channel().ID, "You need to specify the role.")
			return nil
		}
		err = db.QueryRow("SELECT roleid FROM autorole WHERE guildid ='" + ctx.Guild().ID + "'").Scan(&tag.roleid)
		if err == nil {
			_, _ = session.ChannelMessageSend(ctx.Channel().ID, "Autorole is already existing, use reset_auto_role command to reset!")
			return nil
		} else {
			insert, err := db.Query("INSERT INTO autorole (roleid, guildid) VALUES ('" + args + "', '" + ctx.Guild().ID + "')")
			if err != nil {
				_, _ = session.ChannelMessageSend(ctx.Channel().ID, "An error occurred.")

				return nil
			}
			defer insert.Close()

			_, _ = session.ChannelMessageSend(ctx.Channel().ID, "Successfully set autorole.")

			return nil
		}
	}
}
