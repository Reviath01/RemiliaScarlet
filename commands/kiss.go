package commands

import (
	"fmt"
	"strings"

	embedutil "git.randomchars.net/FreeNitori/EmbedUtil"
	multiplexer "git.randomchars.net/Reviath/RemiliaScarlet/Multiplexer"
	CommandHandler "git.randomchars.net/Reviath/RemiliaScarlet/handler"
	"git.randomchars.net/Reviath/RemiliaScarlet/sql"
)

// KissCommand is a handler for kiss command
func KissCommand(ctx CommandHandler.Context, _ []string) error {
	switch sql.CheckLanguage(ctx.Guild.ID) {
	case "tr":
		if sql.IsBlocked(ctx.Guild.ID, "kiss") == "true" {
			ctx.Reply("Bu komut bu sunucuda engellenmiş.")
			return nil
		}

		if len(strings.Join(multiplexer.GetArgs(ctx.Message.Content, multiplexer.GetPrefix()), " ")) < 1 {
			ctx.Reply("Bir üye belirtmelisin.")

			return nil
		}

		u, err := ctx.Session.User(multiplexer.GetUser(multiplexer.GetArgs(ctx.Message.Content, multiplexer.GetPrefix())[0]))
		if err == nil {
			embed := embedutil.New("", fmt.Sprintf("<@%s>, <@%s> isimli kişiyi öptü 😘", ctx.Message.Author.ID, u.ID))
			embed.Color = 0xff1000
			embed.SetImage("https://media.tenor.com/images/d68747a5865b12c465e5dff31c65d5c2/tenor.gif")
			ctx.ReplyEmbed(embed)
			return nil
		}
		ctx.Reply("Bir üye belirtmelisin")
		return nil

	default:
		if sql.IsBlocked(ctx.Guild.ID, "kiss") == "true" {
			ctx.Reply("This command is blocked on this guild.")
			return nil
		}

		if len(strings.Join(multiplexer.GetArgs(ctx.Message.Content, multiplexer.GetPrefix()), " ")) < 1 {
			ctx.Reply("You need to specify the user.")

			return nil
		}

		u, err := ctx.Session.User(multiplexer.GetUser(multiplexer.GetArgs(ctx.Message.Content, multiplexer.GetPrefix())[0]))
		if err == nil {
			embed := embedutil.New("", fmt.Sprintf("<@%s> kisses <@%s>", ctx.Message.Author.ID, u.ID))
			embed.Color = 0xff1000
			embed.SetImage("https://media.tenor.com/images/d68747a5865b12c465e5dff31c65d5c2/tenor.gif")
			ctx.ReplyEmbed(embed)

			return nil
		}
		ctx.Reply("You need to specify the user.")
		return nil
	}
}
