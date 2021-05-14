package commands

import (
	"fmt"

	embedutil "git.randomchars.net/FreeNitori/EmbedUtil"
	CommandHandler "git.randomchars.net/Reviath/RemiliaScarlet/handler"
	"git.randomchars.net/Reviath/RemiliaScarlet/sql"
)

//Roles command
func RolesCommand(ctx CommandHandler.Context, _ []string) error {
	var roles string

	switch sql.CheckLanguage(ctx.Guild.ID) {
	case "tr":
		if sql.IsBlocked(ctx.Guild.ID, "roles") == "true" {
			ctx.Reply("Bu komut bu sunucuda engellenmiş.")
			return nil
		}

		for _, i := range ctx.Guild.Roles {
			roles += fmt.Sprintf("<@&%s> ,\n", i.ID)
		}

		embed := embedutil.New("", "")
		embed.AddField("Roller:", roles, true)

		ctx.ReplyEmbed(embed)

		return nil
	default:
		if sql.IsBlocked(ctx.Guild.ID, "roles") == "true" {
			ctx.Reply("This command is blocked on this guild.")
			return nil
		}

		for _, i := range ctx.Guild.Roles {
			roles += fmt.Sprintf("<@&%s> ,\n", i.ID)
		}

		embed := embedutil.New("", "")
		embed.AddField("Roles:", roles, true)

		ctx.ReplyEmbed(embed)
		return nil
	}
}
