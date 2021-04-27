package commands

import (
	"strings"

	embedutil "git.randomchars.net/Reviath/RemiliaScarlet/EmbedUtil"
	multiplexer "git.randomchars.net/Reviath/RemiliaScarlet/Multiplexer"
	CommandHandler "git.randomchars.net/Reviath/RemiliaScarlet/handler"
	"git.randomchars.net/Reviath/RemiliaScarlet/sql"
)

func EmbedCommand(ctx CommandHandler.Context, _ []string) error {
	if sql.CheckLanguage(ctx.Guild.ID) == "tr" {
		if sql.IsBlocked(ctx.Guild.ID, "embed") == "true" {
			ctx.Reply("Bu komut bu sunucuda engellenmiş.")
			return nil
		}

		if strings.Join(multiplexer.GetArgs(ctx.Message.Content, multiplexer.GetPrefix()), " ") == "" {
			ctx.Reply("Bir mesaj belirtmelisin.")

			return nil
		}
		embed := embedutil.NewEmbed().
			SetColor(0xc000ff).
			SetDescription(strings.Join(multiplexer.GetArgs(ctx.Message.Content, multiplexer.GetPrefix()), " ")).MessageEmbed
		ctx.ReplyEmbed(embed)

		return nil

	}

	if sql.IsBlocked(ctx.Guild.ID, "embed") == "true" {
		ctx.Reply("This command is blocked on this guild.")
		return nil
	}

	if strings.Join(multiplexer.GetArgs(ctx.Message.Content, multiplexer.GetPrefix()), " ") == "" {
		ctx.Reply("You need to specify the message.")

		return nil
	}
	embed := embedutil.NewEmbed().
		SetColor(0xc000ff).
		SetDescription(strings.Join(multiplexer.GetArgs(ctx.Message.Content, multiplexer.GetPrefix()), " ")).MessageEmbed
	ctx.ReplyEmbed(embed)

	return nil
}
