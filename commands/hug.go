package commands

import (
	"fmt"
	"strings"

	multiplexer "git.randomchars.net/Reviath/RemiliaScarlet/Multiplexer"
	CommandHandler "git.randomchars.net/Reviath/RemiliaScarlet/handler"
	"git.randomchars.net/Reviath/RemiliaScarlet/sql"
	embedutil "git.randomchars.net/freenitori/embedutil"
)

// HugCommand is a handler for hug command
func HugCommand(ctx CommandHandler.Context, _ []string) error {
	switch sql.CheckLanguage(ctx.Guild.ID) {
	case "tr":
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
			embed := embedutil.New("", fmt.Sprintf("<@%s>, <@%s> isimli kişiye sarıldı 💖", ctx.Message.Author.ID, u.ID))
			embed.Color = 0xff1000
			embed.SetImage("https://i.pinimg.com/originals/4d/d7/49/4dd749423de10a319b5d9e8850bbace4.gif")
			ctx.ReplyEmbed(embed)
			return nil
		}
		ctx.Reply("Bir kişiyi etiketlemelisin.")
		return nil

	default:
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
			embed := embedutil.New("", fmt.Sprintf("<@%s> hugs <@%s>", ctx.Message.Author.ID, u.ID))
			embed.Color = 0xff1000
			embed.SetImage("https://i.pinimg.com/originals/4d/d7/49/4dd749423de10a319b5d9e8850bbace4.gif")
			ctx.ReplyEmbed(embed)

			return nil
		}
		ctx.Reply("You need to specify the user.")
		return nil
	}
}
