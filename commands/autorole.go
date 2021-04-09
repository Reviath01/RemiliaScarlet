package commands

import (
	"database/sql"
	"strings"

	ctx "git.randomchars.net/Reviath/RemiliaScarlet/Context"
	"github.com/bwmarrin/discordgo"
	_ "github.com/go-sql-driver/mysql"
)

type AutoRole struct {
}

func (a AutoRole) Execute(ctx ctx.Ctx, session *discordgo.Session) error {

	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/remilia")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	type Tag struct {
		roleid string
		lang   string
	}

	var tag Tag

	err = db.QueryRow("SELECT language FROM languages WHERE guildid ='" + ctx.Guild().ID + "'").Scan(&tag.lang)
	if err == nil && tag.lang == "tr" {
		perms, err := session.State.UserChannelPermissions(ctx.Author().ID, ctx.Channel().ID)
		if err == nil && !(int(perms)&discordgo.PermissionAdministrator == discordgo.PermissionAdministrator) {
			_, _ = session.ChannelMessageSend(ctx.Channel().ID, "Bu komutu kullanmak için yönetici yetkisine sahip olmalısın.")
			return nil
		}
		if strings.Join(ctx.Args(), " ") == "" {
			_, _ = session.ChannelMessageSend(ctx.Channel().ID, "Rolü belirtmelisin")

			return nil
		} else {
			args := ctx.Args()[0]

			if len(args) == 18 {

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
			} else {
				if len(args) == 22 {

					err = db.QueryRow("SELECT roleid FROM autorole WHERE guildid ='" + ctx.Guild().ID + "'").Scan(&tag.roleid)
					if err == nil {
						_, _ = session.ChannelMessageSend(ctx.Channel().ID, "Otorol zaten ayarlanmış, tekrar ayarlamak için reset_autorole komutunu kullan!")
						return nil
					} else {
						insert, err := db.Query("INSERT INTO autorole (roleid, guildid) VALUES ('" + args[3:][:18] + "', '" + ctx.Guild().ID + "')")
						if err != nil {
							_, _ = session.ChannelMessageSend(ctx.Channel().ID, "Bir hata oluştu.")

							return nil
						}
						defer insert.Close()

						_, _ = session.ChannelMessageSend(ctx.Channel().ID, "Otorol başarıyla ayarlandı.")

						return nil
					}
				} else {
					_, _ = session.ChannelMessageSend(ctx.Channel().ID, "Bir rol belirtmelisin.")

					return nil
				}
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
		args := ctx.Args()[0]

		if len(args) == 18 {

			err = db.QueryRow("SELECT roleid FROM autorole WHERE guildid ='" + ctx.Guild().ID + "'").Scan(&tag.roleid)
			if err == nil {
				_, _ = session.ChannelMessageSend(ctx.Channel().ID, "Auto role is already existing (to reset, use reset_autorole command).")
				return nil
			} else {
				insert, err := db.Query("INSERT INTO autorole (roleid, guildid) VALUES ('" + args + "', '" + ctx.Guild().ID + "')")
				if err != nil {
					_, _ = session.ChannelMessageSend(ctx.Channel().ID, "An error occurred, please try again.")
					return nil
				}
				defer insert.Close()

				_, _ = session.ChannelMessageSend(ctx.Channel().ID, "Auto role set successfully.")
				return nil
			}
		} else {
			if len(args) == 22 {
				db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/remilia")

				if err != nil {
					panic(err.Error())
				}

				defer db.Close()

				type Tag struct {
					roleid string
				}

				var tag Tag

				err = db.QueryRow("SELECT roleid FROM autorole WHERE guildid ='" + ctx.Guild().ID + "'").Scan(&tag.roleid)
				if err == nil {
					_, _ = session.ChannelMessageSend(ctx.Channel().ID, "Auto role is already existing (to reset, use reset_autorole command).")
					return nil
				} else {
					insert, err := db.Query("INSERT INTO autorole (roleid, guildid) VALUES ('" + args[3:][:18] + "', '" + ctx.Guild().ID + "')")
					if err != nil {
						_, _ = session.ChannelMessageSend(ctx.Channel().ID, "An error occurred, please try again.")

						return nil
					}
					defer insert.Close()

					_, _ = session.ChannelMessageSend(ctx.Channel().ID, "Auto role set successfully.")

					return nil
				}
			} else {
				_, _ = session.ChannelMessageSend(ctx.Channel().ID, "You need to specify the role.")

				return nil
			}
		}
	}
}
