package commands

import (
	"database/sql"
	"strings"

	ctx "git.randomchars.net/Reviath/RemiliaScarlet/Context"
	"github.com/bwmarrin/discordgo"
	_ "github.com/go-sql-driver/mysql"
)

type Kick struct {
}

func (k Kick) Execute(ctx ctx.Ctx, session *discordgo.Session) error {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/remilia")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	type Tag struct {
		isblocked string
		lang      string
	}

	var tag Tag

	err = db.QueryRow("SELECT language FROM languages WHERE guildid ='" + ctx.Guild().ID + "'").Scan(&tag.lang)
	if err == nil && tag.lang == "tr" {
		err = db.QueryRow("SELECT isblocked FROM disabledcommands WHERE commandname ='kick' AND guildid ='" + ctx.Guild().ID + "'").Scan(&tag.isblocked)

		if err == nil {
			if tag.isblocked == "True" {
				_, _ = session.ChannelMessageSend(ctx.Channel().ID, "This command is blocked on this guild.")

				return nil
			}
		}

		perms, err := session.State.UserChannelPermissions(ctx.Author().ID, ctx.Channel().ID)
		if err == nil && !(int(perms)&discordgo.PermissionAdministrator == discordgo.PermissionAdministrator) {
			_, _ = session.ChannelMessageSend(ctx.Channel().ID, "Bu komutu kullanmak için üyeleri at yetkisine sahip olmalısın.")

			return nil
		}

		var args string
		if len(strings.Join(ctx.Args(), " ")) < 1 {
			_, _ = session.ChannelMessageSend(ctx.Channel().ID, "Bir üye belirtmelisin.")

			return nil
		}
		args = ctx.Args()[0]

		if len(args) < 19 {
			u, err := session.User(args)
			if err == nil {
				err = session.GuildMemberDelete(ctx.Guild().ID, u.ID)
				if err != nil {
					_, _ = session.ChannelMessageSend(ctx.Channel().ID, "Yeterli yetkim yok.")

					return nil
				}
				_, _ = session.ChannelMessageSend(ctx.Channel().ID, "Belirtilen kişi sunucudan atıldı.")

				return nil
			} else {
				_, _ = session.ChannelMessageSend(ctx.Channel().ID, "Bir üye belirtmelisin.")

				return nil
			}
		} else {
			if len(args) > 21 {
				u, err := session.User(args[3:][:18])
				if err == nil {
					err = session.GuildMemberDelete(ctx.Guild().ID, u.ID)
					if err != nil {
						_, _ = session.ChannelMessageSend(ctx.Channel().ID, "Yeterli yetkim yok.")

						return nil
					}
					_, _ = session.ChannelMessageSend(ctx.Channel().ID, "Belirtilen kişi sunucudan atıldı!")

					return nil
				} else {
					_, _ = session.ChannelMessageSend(ctx.Channel().ID, "Bir üye belirtmelisin.")

					return nil
				}
			} else {
				_, _ = session.ChannelMessageSend(ctx.Channel().ID, "Bir üye belirtmelisin.")

				return nil
			}
		}
	}
	err = db.QueryRow("SELECT isblocked FROM disabledcommands WHERE commandname ='kick' AND guildid ='" + ctx.Guild().ID + "'").Scan(&tag.isblocked)

	if err == nil {
		if tag.isblocked == "True" {
			_, _ = session.ChannelMessageSend(ctx.Channel().ID, "This command is blocked on this guild.")

			return nil
		}
	}

	perms, err := session.State.UserChannelPermissions(ctx.Author().ID, ctx.Channel().ID)
	if err == nil && !(int(perms)&discordgo.PermissionAdministrator == discordgo.PermissionAdministrator) {
		_, _ = session.ChannelMessageSend(ctx.Channel().ID, "You need kick members permission to run this command.")

		return nil
	}

	var args string
	if len(strings.Join(ctx.Args(), " ")) < 1 {
		_, _ = session.ChannelMessageSend(ctx.Channel().ID, "You need to specify the user.")

		return nil
	}
	args = ctx.Args()[0]

	if len(args) < 19 {
		u, err := session.User(args)
		if err == nil {
			err = session.GuildMemberDelete(ctx.Guild().ID, u.ID)
			if err != nil {
				_, _ = session.ChannelMessageSend(ctx.Channel().ID, "I do not have enough permission.")

				return nil
			}
			_, _ = session.ChannelMessageSend(ctx.Channel().ID, "Kicked specified user.")

			return nil
		} else {
			_, _ = session.ChannelMessageSend(ctx.Channel().ID, "You need to specify the user.")

			return nil
		}
	} else {
		if len(args) > 21 {
			u, err := session.User(args[3:][:18])
			if err == nil {
				err = session.GuildMemberDelete(ctx.Guild().ID, u.ID)
				if err != nil {
					_, _ = session.ChannelMessageSend(ctx.Channel().ID, "I do not have enough permission.")

					return nil
				}
				_, _ = session.ChannelMessageSend(ctx.Channel().ID, "Kicked specified user.")

				return nil
			} else {
				_, _ = session.ChannelMessageSend(ctx.Channel().ID, "You need to specify the user.")

				return nil
			}
		} else {
			_, _ = session.ChannelMessageSend(ctx.Channel().ID, "You need to specify the user.")

			return nil
		}
	}
}
