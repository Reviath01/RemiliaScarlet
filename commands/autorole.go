package commands

import (
	"strings"

	CommandHandler "git.randomchars.net/Reviath/RemiliaScarlet/handler"
	multiplexer "git.randomchars.net/Reviath/RemiliaScarlet/multiplexer"
	"git.randomchars.net/Reviath/RemiliaScarlet/sql"
	"github.com/bwmarrin/discordgo"
)

func AutoRoleCommand(ctx CommandHandler.Context, _ []string) error {
	db := sql.Connect()

	type Tag struct {
		roleid string
	}

	var tag Tag

	if sql.CheckLanguage(ctx.Guild.ID) == "tr" {
		perms, err := ctx.Session.State.UserChannelPermissions(ctx.Message.Author.ID, ctx.Channel.ID)
		if err == nil && !(int(perms)&discordgo.PermissionAdministrator == discordgo.PermissionAdministrator) {
			ctx.Reply("Bu komutu kullanmak için yönetici yetkisine sahip olmalısın.")
			return nil
		}
		if strings.Join(multiplexer.GetArgs(ctx.Message.Content, multiplexer.GetPrefix()), " ") == "" {
			ctx.Reply("Rolü belirtmelisin.")
			return nil
		} else {
			args := multiplexer.GetRole(multiplexer.GetArgs(ctx.Message.Content, multiplexer.GetPrefix())[0])
			if args == "Error" {
				ctx.Reply("Rolü belirtmelisin.")
				return nil
			}
			err = db.QueryRow("SELECT roleid FROM autorole WHERE guildid ='" + ctx.Guild.ID + "'").Scan(&tag.roleid)
			if err == nil {
				ctx.Reply("Otorol zaten ayarlanmış, tekrar ayarlamak için reset_autorole komutunu kullan!")
				return nil
			} else {
				insert, err := db.Query("INSERT INTO autorole (roleid, guildid) VALUES ('" + args + "', '" + ctx.Guild.ID + "')")
				if err != nil {
					ctx.Reply("Bir hata oluştu.")

					return nil
				}
				defer insert.Close()

				ctx.Reply("Otorol başarıyla ayarlandı.")

				return nil
			}
		}
	}
	perms, err := ctx.Session.State.UserChannelPermissions(ctx.Message.Author.ID, ctx.Channel.ID)
	if err == nil && !(int(perms)&discordgo.PermissionAdministrator == discordgo.PermissionAdministrator) {
		ctx.Reply("You need administrator permission to run this command.")
		return nil
	}

	if strings.Join(multiplexer.GetArgs(ctx.Message.Content, multiplexer.GetPrefix()), " ") == "" {
		ctx.Reply("You need to specify the role.")
		return nil
	} else {
		args := multiplexer.GetRole(multiplexer.GetArgs(ctx.Message.Content, multiplexer.GetPrefix())[0])
		if args == "Error" {
			ctx.Reply("You need to specify the role.")
			return nil
		}
		err = db.QueryRow("SELECT roleid FROM autorole WHERE guildid ='" + ctx.Guild.ID + "'").Scan(&tag.roleid)
		if err == nil {
			ctx.Reply("Autorole is already existing, use reset_auto_role command to reset!")
			return nil
		} else {
			insert, err := db.Query("INSERT INTO autorole (roleid, guildid) VALUES ('" + args + "', '" + ctx.Guild.ID + "')")
			if err != nil {
				ctx.Reply("An error occurred.")

				return nil
			}
			defer insert.Close()

			ctx.Reply("Successfully set autorole.")

			return nil
		}
	}
}
