package commands

import (
	"strings"

	ctx "git.randomchars.net/Reviath/RemiliaScarlet/CommandHandler/Context"
	embedutil "git.randomchars.net/Reviath/RemiliaScarlet/EmbedUtil"
	multiplexer "git.randomchars.net/Reviath/RemiliaScarlet/Multiplexer"
	"git.randomchars.net/Reviath/RemiliaScarlet/sql"
	"github.com/bwmarrin/discordgo"
)

type Kiss struct {
}

func (k Kiss) Execute(ctx ctx.Ctx, session *discordgo.Session) error {
	if sql.CheckLanguage(ctx.Guild().ID) == "tr" {
		if sql.IsBlocked(ctx.Guild().ID, "afk") == "kiss" {
			_, _ = session.ChannelMessageSend(ctx.Channel().ID, "Bu komut bu sunucuda engellenmiş.")
			return nil
		}

		if len(strings.Join(ctx.Args(), " ")) < 1 {
			_, _ = session.ChannelMessageSend(ctx.Channel().ID, "Bir üye belirtmelisin.")

			return nil
		}

		u, err := session.User(multiplexer.GetUser(ctx.Args()[0]))
		if err == nil {
			embed := embedutil.NewEmbed().
				SetColor(0xff1000).
				SetDescription("<@" + ctx.Author().ID + ">, <@" + u.ID + "> isimli kişiyi öptü 😘").
				SetImage("https://media.tenor.com/images/d68747a5865b12c465e5dff31c65d5c2/tenor.gif").MessageEmbed
			_, _ = session.ChannelMessageSendEmbed(ctx.Channel().ID, embed)

			return nil
		} else {
			_, _ = session.ChannelMessageSend(ctx.Channel().ID, "Bir üye belirtmelisin")
			return nil
		}
	}

	if sql.IsBlocked(ctx.Guild().ID, "kiss") == "true" {
		_, _ = session.ChannelMessageSend(ctx.Channel().ID, "This command is blocked on this guild.")
		return nil
	}

	if len(strings.Join(ctx.Args(), " ")) < 1 {
		_, _ = session.ChannelMessageSend(ctx.Channel().ID, "You need to specify the user.")

		return nil
	}

	u, err := session.User(multiplexer.GetUser(ctx.Args()[0]))
	if err == nil {
		embed := embedutil.NewEmbed().
			SetColor(0xff1000).
			SetDescription("<@" + ctx.Author().ID + "> kisses <@" + u.ID + ">").
			SetImage("https://media.tenor.com/images/d68747a5865b12c465e5dff31c65d5c2/tenor.gif").MessageEmbed
		_, _ = session.ChannelMessageSendEmbed(ctx.Channel().ID, embed)

		return nil
	} else {
		_, _ = session.ChannelMessageSend(ctx.Channel().ID, "You need to specify the user.")
		return nil
	}
}
