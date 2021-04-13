package commands

import (
	"strings"

	ctx "git.randomchars.net/Reviath/RemiliaScarlet/CommandHandler/Context"
	"git.randomchars.net/Reviath/RemiliaScarlet/sql"
	"github.com/bwmarrin/discordgo"
)

type Language struct {
}

func (l Language) Execute(ctx ctx.Ctx, session *discordgo.Session) error {
	db := sql.Connect()

	type Tag struct {
		lang string
	}

	var tag Tag

	args := ctx.Args()

	err := db.QueryRow("SELECT language FROM languages WHERE guildid ='" + ctx.Guild().ID + "'").Scan(&tag.lang)
	if err == nil {
		if tag.lang == "tr" {
			if len(strings.Join(ctx.Args(), " ")) < 1 {
				_, _ = session.ChannelMessageSend(ctx.Channel().ID, "Yeni dili belirtmelisiniz.")

				return nil
			} else if args[0] == "tr" {
				_, _ = session.ChannelMessageSend(ctx.Channel().ID, "Dil zaten Türkçe ayarlı.")

				return nil
			} else if args[0] == "en" {
				_, _ = session.ChannelMessageSend(ctx.Channel().ID, "Dili ingilizce olarak ayarlıyorum...")

				delete, err := db.Query("DELETE FROM languages WHERE guildid ='" + ctx.Guild().ID + "'")
				if err != nil {
					_, _ = session.ChannelMessageSend(ctx.Channel().ID, "Bir hata oluştu.")

					return nil
				}

				defer delete.Close()
			} else {
				_, _ = session.ChannelMessageSend(ctx.Channel().ID, "Kullanılabilir diller: `tr`, `en`")

				return nil
			}
		} else {
			if len(strings.Join(ctx.Args(), " ")) < 1 {
				_, _ = session.ChannelMessageSend(ctx.Channel().ID, "You need to specify the language.")

				return nil
			} else if args[0] == "tr" {
				_, _ = session.ChannelMessageSend(ctx.Channel().ID, "Setting language as Turkish...")

				insert, err := db.Query("INSERT INTO languages (language, guildid) VALUES ('tr', '" + ctx.Guild().ID + "')")
				if err != nil {
					return nil
				}

				_, _ = session.ChannelMessageSend(ctx.Channel().ID, "Dil Türkçe yapıldı!")

				defer insert.Close()

				return nil

			} else if args[0] == "en" {
				_, _ = session.ChannelMessageSend(ctx.Channel().ID, "Language is en.")

				return nil
			} else {
				_, _ = session.ChannelMessageSend(ctx.Channel().ID, "Language options: `tr`, `en`")

				return nil
			}
		}
		return nil
	} else {
		if len(strings.Join(ctx.Args(), " ")) < 1 {
			_, _ = session.ChannelMessageSend(ctx.Channel().ID, "You need to specify the language.")

			return nil
		} else if args[0] == "tr" {
			_, _ = session.ChannelMessageSend(ctx.Channel().ID, "Setting language as Turkish...")

			insert, err := db.Query("INSERT INTO languages (language, guildid) VALUES ('tr', '" + ctx.Guild().ID + "')")

			if err != nil {
				return nil
			}

			_, _ = session.ChannelMessageSend(ctx.Channel().ID, "Dil Türkçe yapıldı!")

			defer insert.Close()

			return nil

		} else if args[0] == "en" {
			_, _ = session.ChannelMessageSend(ctx.Channel().ID, "Language is en.")

			return nil
		} else {
			_, _ = session.ChannelMessageSend(ctx.Channel().ID, "Language options: `tr`, `en`")

			return nil
		}
	}
}
