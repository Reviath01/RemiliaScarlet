package commands

import (
	"fmt"
	"strings"

	multiplexer "git.randomchars.net/Reviath/RemiliaScarlet/Multiplexer"
	CommandHandler "git.randomchars.net/Reviath/RemiliaScarlet/handler"
	"git.randomchars.net/Reviath/RemiliaScarlet/sql"
	embedutil "git.randomchars.net/freenitori/embedutil"
)

// SpoilerCommand is a handler for spoiler command
func SpoilerCommand(ctx CommandHandler.Context, _ []string) error {
	switch sql.CheckLanguage(ctx.Guild.ID) {
	case "tr":
		if sql.IsBlocked(ctx.Guild.ID, "spoiler") == "true" {
			ctx.Reply("Bu komut bu sunucuda engellenmiş.")
			return nil
		}

		if strings.Join(multiplexer.GetArgs(ctx.Message.Content, multiplexer.GetPrefix()), " ") == "" {
			ctx.Reply("Spoiler olarak gönderilecek mesajı belirtmelisin.")
			return nil
		}
		spoilerembed := embedutil.New("", fmt.Sprintf("|| %s ||", strings.Join(multiplexer.GetArgs(ctx.Message.Content, multiplexer.GetPrefix()), " ")))
		spoilerembed.Color = 0xe9ff00
		ctx.ReplyEmbed(spoilerembed)
		return nil
	default:

		if sql.IsBlocked(ctx.Guild.ID, "spoiler") == "true" {
			ctx.Reply("This command is blocked on this guild.")
			return nil
		}

		if strings.Join(multiplexer.GetArgs(ctx.Message.Content, multiplexer.GetPrefix()), " ") == "" {
			ctx.Reply("You need to specify the message.")

			return nil
		}
		spoilerembed := embedutil.New("", fmt.Sprintf("|| %s ||", strings.Join(multiplexer.GetArgs(ctx.Message.Content, multiplexer.GetPrefix()), " ")))
		spoilerembed.Color = 0xe9ff00
		ctx.ReplyEmbed(spoilerembed)

		return nil
	}
}
