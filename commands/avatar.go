package commands

import (
	"database/sql"
	"strings"

	ctx "git.randomchars.net/Reviath/RemiliaScarlet/Context"
	embedutil "git.randomchars.net/Reviath/embed-util"
	"github.com/bwmarrin/discordgo"
	_ "github.com/go-sql-driver/mysql"
)

type Avatar struct {
}

func (a Avatar) Execute(ctx ctx.Ctx, session *discordgo.Session) error {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/remilia")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	type Tag struct {
		isblocked string `json:"isblocked"`
		lang      string `json:"language"`
	}

	var tag Tag

	err = db.QueryRow("SELECT language FROM languages WHERE guildid ='" + ctx.Guild().ID + "'").Scan(&tag.lang)
	if err == nil && tag.lang == "tr" {
		err = db.QueryRow("SELECT isblocked FROM disabledcommands WHERE commandname ='author' AND guildid ='" + ctx.Guild().ID + "'").Scan(&tag.isblocked)

		if err == nil {
			if tag.isblocked == "True" {
				_, err = session.ChannelMessageSend(ctx.Channel().ID, "Bu komut bu sunucuda engellenmiş.")
				return err
			}
		}

		if len(strings.Join(ctx.Args(), " ")) < 1 {
			avatarembed := embedutil.NewEmbed().
				SetColor(0xff1000).
				SetDescription(ctx.Author().Username + "#" + ctx.Author().Discriminator + " isimli kişinin profil fotoğrafı").
				SetImage(ctx.Author().AvatarURL("1024")).MessageEmbed
			_, err := session.ChannelMessageSendEmbed(ctx.Channel().ID, avatarembed)

			if err != nil {
				return nil
			}
			return err
		}

		var args string
		args = ctx.Args()[0]

		if len(args) < 19 {
			u, err := session.User(args)
			if err == nil {
				avatarembed := embedutil.NewEmbed().
					SetColor(0xff1000).
					SetDescription(u.Username + "#" + u.Discriminator + " isimli kişinin profil fotoğrafı").
					SetImage(u.AvatarURL("1024")).MessageEmbed
				_, err := session.ChannelMessageSendEmbed(ctx.Channel().ID, avatarembed)

				if err != nil {
					return nil
				}

				return err
			} else {
				avatarembed := embedutil.NewEmbed().
					SetColor(0xff1000).
					SetDescription(ctx.Author().Username + "#" + ctx.Author().Discriminator + " isimli kişinin profil fotoğrafı").
					SetImage(ctx.Author().AvatarURL("1024")).MessageEmbed
				_, err := session.ChannelMessageSendEmbed(ctx.Channel().ID, avatarembed)

				if err != nil {
					return nil
				}

				return err
			}
		} else {
			if len(args) > 21 {
				u, err := session.User(args[3:][:18])
				if err == nil {
					avatarembed := embedutil.NewEmbed().
						SetColor(0xff1000).
						SetDescription(u.Username + "#" + u.Discriminator + " isimli kişinin profil fotoğrafı").
						SetImage(u.AvatarURL("1024")).MessageEmbed
					_, err := session.ChannelMessageSendEmbed(ctx.Channel().ID, avatarembed)

					if err != nil {
						return nil
					}

					return err
				} else {
					avatarembed := embedutil.NewEmbed().
						SetColor(0xff1000).
						SetDescription(ctx.Author().Username + "#" + ctx.Author().Discriminator + " isimli kişinin profil fotoğrafı").
						SetImage(ctx.Author().AvatarURL("1024")).MessageEmbed
					_, err := session.ChannelMessageSendEmbed(ctx.Channel().ID, avatarembed)

					if err != nil {
						return nil
					}

					return err
				}
			} else {
				avatarembed := embedutil.NewEmbed().
					SetColor(0xff1000).
					SetImage(ctx.Author().AvatarURL("1024")).MessageEmbed
				_, err := session.ChannelMessageSendEmbed(ctx.Channel().ID, avatarembed)

				if err != nil {
					return nil
				}

				return err
			}
		}
	}

	err = db.QueryRow("SELECT isblocked FROM disabledcommands WHERE commandname ='author' AND guildid ='" + ctx.Guild().ID + "'").Scan(&tag.isblocked)

	if err == nil {
		if tag.isblocked == "True" {
			_, err = session.ChannelMessageSend(ctx.Channel().ID, "This command is blocked on this guild.")
			return err
		}
	}

	if len(strings.Join(ctx.Args(), " ")) < 1 {
		avatarembed := embedutil.NewEmbed().
			SetColor(0xff1000).
			SetDescription("Avatar of " + ctx.Author().Username + "#" + ctx.Author().Discriminator).
			SetImage(ctx.Author().AvatarURL("1024")).MessageEmbed
		_, err := session.ChannelMessageSendEmbed(ctx.Channel().ID, avatarembed)

		if err != nil {
			return nil
		}
		return err
	}

	var args string
	args = ctx.Args()[0]

	if len(args) < 19 {
		u, err := session.User(args)
		if err == nil {
			avatarembed := embedutil.NewEmbed().
				SetColor(0xff1000).
				SetDescription("Avatar of " + u.Username + "#" + u.Discriminator).
				SetImage(u.AvatarURL("1024")).MessageEmbed
			_, err := session.ChannelMessageSendEmbed(ctx.Channel().ID, avatarembed)

			if err != nil {
				return nil
			}

			return err
		} else {
			avatarembed := embedutil.NewEmbed().
				SetColor(0xff1000).
				SetDescription("Avatar of " + ctx.Author().Username + "#" + ctx.Author().Discriminator).
				SetImage(ctx.Author().AvatarURL("1024")).MessageEmbed
			_, err := session.ChannelMessageSendEmbed(ctx.Channel().ID, avatarembed)

			if err != nil {
				return nil
			}

			return err
		}
	} else {
		if len(args) > 21 {
			u, err := session.User(args[3:][:18])
			if err == nil {
				avatarembed := embedutil.NewEmbed().
					SetColor(0xff1000).
					SetDescription("Avatar of " + u.Username + "#" + u.Discriminator).
					SetImage(u.AvatarURL("1024")).MessageEmbed
				_, err := session.ChannelMessageSendEmbed(ctx.Channel().ID, avatarembed)

				if err != nil {
					return nil
				}

				return err
			} else {
				avatarembed := embedutil.NewEmbed().
					SetColor(0xff1000).
					SetDescription("Avatar of " + ctx.Author().Username + "#" + ctx.Author().Discriminator).
					SetImage(ctx.Author().AvatarURL("1024")).MessageEmbed
				_, err := session.ChannelMessageSendEmbed(ctx.Channel().ID, avatarembed)

				if err != nil {
					return nil
				}

				return err
			}
		} else {
			avatarembed := embedutil.NewEmbed().
				SetColor(0xff1000).
				SetImage(ctx.Author().AvatarURL("1024")).MessageEmbed
			_, err := session.ChannelMessageSendEmbed(ctx.Channel().ID, avatarembed)

			if err != nil {
				return nil
			}

			return err
		}
	}
}
