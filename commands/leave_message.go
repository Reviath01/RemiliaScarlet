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
		message string `json:"message"`
		lang    string `json:"language"`
	}

	var tag Tag

	err = db.QueryRow("SELECT language FROM languages WHERE guildid ='" + ctx.Guild().ID + "'").Scan(&tag.lang)
	if err == nil && tag.lang == "tr" {
		perms, err := session.State.UserChannelPermissions(ctx.Author().ID, ctx.Channel().ID)
		if err == nil && (int(perms)&discordgo.PermissionAdministrator == discordgo.PermissionAdministrator) == false {
			_, err := session.ChannelMessageSend(ctx.Channel().ID, "Bu komutu kullanmak için yönetici yetkisine sahip olmalısın.")

			if err != nil {
				return nil
			}

			return err
		}
		if len(strings.Join(ctx.Args(), " ")) < 1 {
			_, err := session.ChannelMessageSend(ctx.Channel().ID, "Yeni mesajı belirtmelisin.")

			if err != nil {
				return nil
			}

			return err
		}

		if len(strings.Join(ctx.Args(), " ")) > 254 {
			_, err = session.ChannelMessageSend(ctx.Channel().ID, "Mesaj maksimum 255 karakter uzunluğunda olabilir.")

			if err != nil {
				return nil
			}

			return err
		}

		err = db.QueryRow("SELECT message FROM leavemessage WHERE guildid ='" + ctx.Guild().ID + "'").Scan(&tag.message)
		if err == nil {
			_, err = session.ChannelMessageSend(ctx.Channel().ID, "Çıkış mesajı zaten ayarlanmış, reset'lemek için reset_leave_message komutunu kullan.")

			if err != nil {
				return nil
			}

			return err
		} else {
			insert, err := db.Query("INSERT INTO leavemessage (message, guildid) VALUES ('" + strings.Join(ctx.Args(), " ") + "', '" + ctx.Guild().ID + "')")
			if err != nil {
				_, err = session.ChannelMessageSend(ctx.Channel().ID, "Bir hata oluştu.")

				if err != nil {
					return nil
				}

				return err
			}
			defer insert.Close()

			_, err = session.ChannelMessageSend(ctx.Channel().ID, "Başarıyla çıkış mesajı ayarlandı.")

			if err != nil {
				return nil
			}

			return err
		}
	}

	perms, err := session.State.UserChannelPermissions(ctx.Author().ID, ctx.Channel().ID)
	if err == nil && (int(perms)&discordgo.PermissionAdministrator == discordgo.PermissionAdministrator) == false {
		_, err := session.ChannelMessageSend(ctx.Channel().ID, "You need administrator permission to run this command.")

		if err != nil {
			return nil
		}

		return err
	}
	if len(strings.Join(ctx.Args(), " ")) < 1 {
		_, err := session.ChannelMessageSend(ctx.Channel().ID, "You need to specify the message.")

		if err != nil {
			return nil
		}

		return err
	}

	if len(strings.Join(ctx.Args(), " ")) > 254 {
		_, err = session.ChannelMessageSend(ctx.Channel().ID, "Your message can be up to 255 characters long.")

		if err != nil {
			return nil
		}

		return err
	}

	err = db.QueryRow("SELECT message FROM leavemessage WHERE guildid ='" + ctx.Guild().ID + "'").Scan(&tag.message)
	if err == nil {
		_, err = session.ChannelMessageSend(ctx.Channel().ID, "Leave message is already existing (to reset, use reset_leave_message command).")

		if err != nil {
			return nil
		}

		return err
	} else {
		insert, err := db.Query("INSERT INTO leavemessage (message, guildid) VALUES ('" + strings.Join(ctx.Args(), " ") + "', '" + ctx.Guild().ID + "')")
		if err != nil {
			_, err = session.ChannelMessageSend(ctx.Channel().ID, "An error occurred, please try again.")

			if err != nil {
				return nil
			}

			return err
		}
		defer insert.Close()

		_, err = session.ChannelMessageSend(ctx.Channel().ID, "Leave message set successfully.")

		if err != nil {
			return nil
		}

		return err
	}
}
