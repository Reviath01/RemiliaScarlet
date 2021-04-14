package commands

import (
	"strings"

	ctx "git.randomchars.net/Reviath/RemiliaScarlet/CommandHandler/Context"
	embedutil "git.randomchars.net/Reviath/RemiliaScarlet/EmbedUtil"
	multiplexer "git.randomchars.net/Reviath/RemiliaScarlet/Multiplexer"
	"git.randomchars.net/Reviath/RemiliaScarlet/sql"
	"github.com/bwmarrin/discordgo"
)

type Hug struct {
}

func (h Hug) Execute(ctx ctx.Ctx, session *discordgo.Session) error {
	if sql.CheckLanguage(ctx.Guild().ID) == "tr" {
		if sql.IsBlocked(ctx.Guild().ID, "hug") == "true" {
			_, _ = session.ChannelMessageSend(ctx.Channel().ID, "Bu komut bu sunucuda engellenmiş.")
			return nil
		}

		if len(strings.Join(ctx.Args(), " ")) < 1 {
			_, _ = session.ChannelMessageSend(ctx.Channel().ID, "Bir üye belirtmelisin.")
			return nil
		}

		u, err := session.User(multiplexer.GetUser(ctx.Args()[0]))
		if err == nil {
			if u.ID == ctx.Author().ID {
				_, _ = session.ChannelMessageSend(ctx.Channel().ID, "Kendine sarılamazsın.")

				return nil
			}
			embed := embedutil.NewEmbed().
				SetColor(0xff1000).
				SetDescription("<@" + ctx.Author().ID + ">, <@" + u.ID + "> isimli kişiye sarıldı 💖").
				SetImage("https://i.pinimg.com/originals/4d/d7/49/4dd749423de10a319b5d9e8850bbace4.gif").MessageEmbed
			_, _ = session.ChannelMessageSendEmbed(ctx.Channel().ID, embed)
			return nil
		} else {
			_, _ = session.ChannelMessageSend(ctx.Channel().ID, "Bir kişiyi etiketlemelisin.")

			return nil
		}
	}

	if sql.IsBlocked(ctx.Guild().ID, "hug") == "true" {
		_, _ = session.ChannelMessageSend(ctx.Channel().ID, "This command is blocked on this guild.")
		return nil
	}

	if len(strings.Join(ctx.Args(), " ")) < 1 {
		_, _ = session.ChannelMessageSend(ctx.Channel().ID, "You need to specify the user.")

		return nil

	}

	u, err := session.User(multiplexer.GetUser(ctx.Args()[0]))
	if err == nil {
		if u.ID == ctx.Author().ID {
			_, _ = session.ChannelMessageSend(ctx.Channel().ID, "You can't hug yourself.")

			return nil
		}
		embed := embedutil.NewEmbed().
			SetColor(0xff1000).
			SetDescription("<@" + ctx.Author().ID + "> hugs <@" + u.ID + ">").
			SetImage("https://i.pinimg.com/originals/4d/d7/49/4dd749423de10a319b5d9e8850bbace4.gif").MessageEmbed
		_, _ = session.ChannelMessageSendEmbed(ctx.Channel().ID, embed)

		return nil
	} else {
		_, _ = session.ChannelMessageSend(ctx.Channel().ID, "You need to specify the user.")

		return nil
	}
}
