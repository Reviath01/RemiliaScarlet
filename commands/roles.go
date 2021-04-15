package commands

import (
	embedutil "git.randomchars.net/Reviath/RemiliaScarlet/EmbedUtil"
	CommandHandler "git.randomchars.net/Reviath/RemiliaScarlet/handler"
	"git.randomchars.net/Reviath/RemiliaScarlet/sql"
)

func RolesCommand(ctx CommandHandler.Context, _ []string) error {
	var roles string

	if sql.CheckLanguage(ctx.Guild.ID) == "tr" {
		if sql.IsBlocked(ctx.Guild.ID, "roles") == "true" {
			ctx.Reply("Bu komut bu sunucuda engellenmiş.")
			return nil
		}

		for _, i := range ctx.Guild.Roles {
			roles += "<@&" + i.ID + ">" + ", \n"
		}

		embed := embedutil.NewEmbed().
			AddField("Roller:", roles).MessageEmbed

		ctx.ReplyEmbed(embed)

		return nil
	}

	if sql.IsBlocked(ctx.Guild.ID, "roles") == "true" {
		ctx.Reply("This command is blocked on this guild.")
		return nil
	}

	for _, i := range ctx.Guild.Roles {
		roles += "<@&" + i.ID + ">" + ", \n"
	}

	embed := embedutil.NewEmbed().
		AddField("Roles:", roles).MessageEmbed

	ctx.ReplyEmbed(embed)

	return nil
}
