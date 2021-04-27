package commands

import (
	embedutil "git.randomchars.net/Reviath/RemiliaScarlet/EmbedUtil"
	CommandHandler "git.randomchars.net/Reviath/RemiliaScarlet/handler"
	"git.randomchars.net/Reviath/RemiliaScarlet/sql"
)

func IconCommand(ctx CommandHandler.Context, _ []string) error {
	if sql.CheckLanguage(ctx.Guild.ID) == "tr" {
		if sql.IsBlocked(ctx.Guild.ID, "icon") == "true" {
			ctx.Reply("Bu komut bu sunucuda engellenmiş.")
			return nil
		}

		iconeembed := embedutil.NewEmbed().
			SetColor(0x00f6ff).
			SetImage(ctx.Guild.IconURL()).MessageEmbed
		ctx.ReplyEmbed(iconeembed)

		return nil
	}

	if sql.IsBlocked(ctx.Guild.ID, "icon") == "true" {
		ctx.Reply("This command is blocked on this guild.")
		return nil
	}

	iconeembed := embedutil.NewEmbed().
		SetColor(0x00f6ff).
		SetImage(ctx.Guild.IconURL()).MessageEmbed
	ctx.ReplyEmbed(iconeembed)

	return nil
}
