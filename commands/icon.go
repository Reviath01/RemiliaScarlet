package commands

import (
	embedutil "git.randomchars.net/FreeNitori/EmbedUtil"
	CommandHandler "git.randomchars.net/Reviath/RemiliaScarlet/handler"
	"git.randomchars.net/Reviath/RemiliaScarlet/sql"
)

func IconCommand(ctx CommandHandler.Context, _ []string) error {
	if sql.CheckLanguage(ctx.Guild.ID) == "tr" {
		if sql.IsBlocked(ctx.Guild.ID, "icon") == "true" {
			ctx.Reply("Bu komut bu sunucuda engellenmiş.")
			return nil
		}

		iconembed := embedutil.New("", "")
		iconembed.Color = 0x00f6ff
		iconembed.SetImage(ctx.Guild.IconURL())
		ctx.ReplyEmbed(iconembed)

		return nil
	}

	if sql.IsBlocked(ctx.Guild.ID, "icon") == "true" {
		ctx.Reply("This command is blocked on this guild.")
		return nil
	}

	iconembed := embedutil.New("", "")
	iconembed.Color = 0x00f6ff
	iconembed.SetImage(ctx.Guild.IconURL())
	ctx.ReplyEmbed(iconembed)

	return nil
}
