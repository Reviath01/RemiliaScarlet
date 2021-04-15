package commands

import (
	"strings"

	embedutil "git.randomchars.net/Reviath/RemiliaScarlet/EmbedUtil"
	multiplexer "git.randomchars.net/Reviath/RemiliaScarlet/Multiplexer"
	CommandHandler "git.randomchars.net/Reviath/RemiliaScarlet/handler"
	"git.randomchars.net/Reviath/RemiliaScarlet/sql"
)

func SpoilerCommand(ctx CommandHandler.Context, _ []string) error {
	if sql.CheckLanguage(ctx.Guild.ID) == "tr" {
		if sql.IsBlocked(ctx.Guild.ID, "spoiler") == "true" {
			ctx.Reply("Bu komut bu sunucuda engellenmiş.")
			return nil
		}

		if strings.Join(multiplexer.GetArgs(ctx.Message.Content, multiplexer.GetPrefix()), " ") == "" {
			ctx.Reply("Spoiler olarak gönderilecek mesajı belirtmelisin.")

			return nil
		}
		spoilerembed := embedutil.NewEmbed().
			SetColor(0xe9ff00).
			SetDescription("|| " + strings.Join(multiplexer.GetArgs(ctx.Message.Content, multiplexer.GetPrefix()), " ") + " ||").MessageEmbed
		ctx.ReplyEmbed(spoilerembed)
		return nil
	}

	if sql.IsBlocked(ctx.Guild.ID, "spoiler") == "true" {
		ctx.Reply("This command is blocked on this guild.")
		return nil
	}

	if strings.Join(multiplexer.GetArgs(ctx.Message.Content, multiplexer.GetPrefix()), " ") == "" {
		ctx.Reply("You need to specify the message.")

		return nil
	}
	spoilerembed := embedutil.NewEmbed().
		SetColor(0xe9ff00).
		SetDescription("|| " + strings.Join(multiplexer.GetArgs(ctx.Message.Content, multiplexer.GetPrefix()), " ") + " ||").MessageEmbed
	ctx.ReplyEmbed(spoilerembed)

	return nil
}
