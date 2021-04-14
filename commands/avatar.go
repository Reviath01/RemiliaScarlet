package commands

import (
	"strings"

	ctx "git.randomchars.net/Reviath/RemiliaScarlet/CommandHandler/Context"
	embedutil "git.randomchars.net/Reviath/RemiliaScarlet/EmbedUtil"
	multiplexer "git.randomchars.net/Reviath/RemiliaScarlet/Multiplexer"
	"git.randomchars.net/Reviath/RemiliaScarlet/sql"
	"github.com/bwmarrin/discordgo"
)

type Avatar struct {
}

func (a Avatar) Execute(ctx ctx.Ctx, session *discordgo.Session) error {
	db := sql.Connect()

	type Tag struct {
		isblocked string
	}

	var tag Tag

	if sql.CheckLanguage(ctx.Guild().ID) == "tr" {
		err := db.QueryRow("SELECT isblocked FROM disabledcommands WHERE commandname ='author' AND guildid ='" + ctx.Guild().ID + "'").Scan(&tag.isblocked)

		if err == nil {
			if tag.isblocked == "True" {
				_, _ = session.ChannelMessageSend(ctx.Channel().ID, "Bu komut bu sunucuda engellenmiş.")
				return nil
			}
		}

		if len(strings.Join(ctx.Args(), " ")) < 1 {
			avatarembed := embedutil.NewEmbed().
				SetColor(0xff1000).
				SetDescription(ctx.Author().Username + "#" + ctx.Author().Discriminator + " isimli kişinin profil fotoğrafı").
				SetImage(ctx.Author().AvatarURL("1024")).MessageEmbed
			_, _ = session.ChannelMessageSendEmbed(ctx.Channel().ID, avatarembed)

			return nil
		}

		u, err := session.User(multiplexer.GetUser(ctx.Args()[0]))
		if err == nil {
			avatarembed := embedutil.NewEmbed().
				SetColor(0xff1000).
				SetDescription(u.Username + "#" + u.Discriminator + " isimli kişinin profil fotoğrafı").
				SetImage(u.AvatarURL("1024")).MessageEmbed
			_, _ = session.ChannelMessageSendEmbed(ctx.Channel().ID, avatarembed)

			return nil
		} else {
			avatarembed := embedutil.NewEmbed().
				SetColor(0xff1000).
				SetDescription(ctx.Author().Username + "#" + ctx.Author().Discriminator + " isimli kişinin profil fotoğrafı").
				SetImage(ctx.Author().AvatarURL("1024")).MessageEmbed
			_, _ = session.ChannelMessageSendEmbed(ctx.Channel().ID, avatarembed)

			return nil
		}
	}

	err := db.QueryRow("SELECT isblocked FROM disabledcommands WHERE commandname ='author' AND guildid ='" + ctx.Guild().ID + "'").Scan(&tag.isblocked)

	if err == nil && tag.isblocked == "True" {
		_, _ = session.ChannelMessageSend(ctx.Channel().ID, "This command is blocked on this guild.")
		return nil
	}

	if len(strings.Join(ctx.Args(), " ")) < 1 {
		avatarembed := embedutil.NewEmbed().
			SetColor(0xff1000).
			SetDescription("Avatar of " + ctx.Author().Username + "#" + ctx.Author().Discriminator).
			SetImage(ctx.Author().AvatarURL("1024")).MessageEmbed
		_, _ = session.ChannelMessageSendEmbed(ctx.Channel().ID, avatarembed)

		return nil
	}

	u, err := session.User(multiplexer.GetUser(ctx.Args()[0]))
	if err == nil {
		avatarembed := embedutil.NewEmbed().
			SetColor(0xff1000).
			SetDescription("Avatar of " + u.Username + "#" + u.Discriminator).
			SetImage(u.AvatarURL("1024")).MessageEmbed
		_, _ = session.ChannelMessageSendEmbed(ctx.Channel().ID, avatarembed)

		return nil
	} else {
		avatarembed := embedutil.NewEmbed().
			SetColor(0xff1000).
			SetDescription("Avatar of " + ctx.Author().Username + "#" + ctx.Author().Discriminator).
			SetImage(ctx.Author().AvatarURL("1024")).MessageEmbed
		_, _ = session.ChannelMessageSendEmbed(ctx.Channel().ID, avatarembed)
		return nil
	}
}
