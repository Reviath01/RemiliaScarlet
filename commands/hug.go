package commands

import (
	"fmt"
	"strings"

	embedutil "git.randomchars.net/Reviath/RemiliaScarlet/EmbedUtil"
	multiplexer "git.randomchars.net/Reviath/RemiliaScarlet/Multiplexer"
	CommandHandler "git.randomchars.net/Reviath/RemiliaScarlet/handler"
	"git.randomchars.net/Reviath/RemiliaScarlet/sql"
)

func HugCommand(ctx CommandHandler.Context, _ []string) error {
	if sql.CheckLanguage(ctx.Guild.ID) == "tr" {
		if sql.IsBlocked(ctx.Guild.ID, "hug") == "true" {
			ctx.Reply("Bu komut bu sunucuda engellenmiş.")
			return nil
		}

		if len(strings.Join(multiplexer.GetArgs(ctx.Message.Content, multiplexer.GetPrefix()), " ")) < 1 {
			ctx.Reply("Bir üye belirtmelisin.")
			return nil
		}

		u, err := ctx.Session.User(multiplexer.GetUser(multiplexer.GetArgs(ctx.Message.Content, multiplexer.GetPrefix())[0]))
		if err == nil {
			if u.ID == ctx.Message.Author.ID {
				ctx.Reply("Kendine sarılamazsın.")

				return nil
			}
			embed := embedutil.NewEmbed().
				SetColor(0xff1000).
				SetDescription(fmt.Sprintf("<@%s>, <@%s> isimli kişiye sarıldı 💖", ctx.Message.Author.ID, u.ID)).
				SetImage("https://i.pinimg.com/originals/4d/d7/49/4dd749423de10a319b5d9e8850bbace4.gif").MessageEmbed
			ctx.ReplyEmbed(embed)
			return nil
		} else {
			ctx.Reply("Bir kişiyi etiketlemelisin.")

			return nil
		}
	}

	if sql.IsBlocked(ctx.Guild.ID, "hug") == "true" {
		ctx.Reply("This command is blocked on this guild.")
		return nil
	}

	if len(strings.Join(multiplexer.GetArgs(ctx.Message.Content, multiplexer.GetPrefix()), " ")) < 1 {
		ctx.Reply("You need to specify the user.")

		return nil

	}

	u, err := ctx.Session.User(multiplexer.GetUser(multiplexer.GetArgs(ctx.Message.Content, multiplexer.GetPrefix())[0]))
	if err == nil {
		if u.ID == ctx.Message.Author.ID {
			ctx.Reply("You can't hug yourself.")

			return nil
		}
		embed := embedutil.NewEmbed().
			SetColor(0xff1000).
			SetDescription(fmt.Sprintf("<@%s> hugs <@%s>", ctx.Message.Author.ID, u.ID)).
			SetImage("https://i.pinimg.com/originals/4d/d7/49/4dd749423de10a319b5d9e8850bbace4.gif").MessageEmbed
		ctx.ReplyEmbed(embed)

		return nil
	} else {
		ctx.Reply("You need to specify the user.")

		return nil
	}
}
