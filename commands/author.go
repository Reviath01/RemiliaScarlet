package commands

import (
	"fmt"

	"git.randomchars.net/Reviath/RemiliaScarlet/config"
	CommandHandler "git.randomchars.net/Reviath/RemiliaScarlet/handler"
	"git.randomchars.net/Reviath/RemiliaScarlet/sql"
	embedutil "git.randomchars.net/freenitori/embedutil"
)

// AuthorCommand is a handler for author command
func AuthorCommand(ctx CommandHandler.Context, _ []string) error {
	config.ReadConfig()

	switch sql.CheckLanguage(ctx.Guild.ID) {
	case "tr":
		switch sql.IsBlocked(ctx.Guild.ID, "author") {
		case "true":
			ctx.Reply("Bu komut bu sunucuda engellenmiş.")
			return nil
		}

		u, err := ctx.Session.User(config.Owner)

		if err != nil {
			return nil
		}

		authorembed := embedutil.New("", "")
		authorembed.Color = 0x007bff
		authorembed.AddField("Sahibim:", fmt.Sprintf("<@%s> ([%s#%s](https://discord.com/users/%s))", u.ID, u.Username, u.Discriminator, u.ID), true)
		ctx.ReplyEmbed(authorembed)

		return nil
	default:
		switch sql.IsBlocked(ctx.Guild.ID, "author") {
		case "true":
			ctx.Reply("This command is blocked on this guild.")
			return nil
		}

		u, err := ctx.Session.User(config.Owner)

		if err != nil {
			return nil
		}

		authorembed := embedutil.New("", "")
		authorembed.Color = 0x007bff
		authorembed.AddField("My Author:", fmt.Sprintf("<@%s> ([%s#%s](https://discord.com/users/%s))", u.ID, u.Username, u.Discriminator, u.ID), true)
		ctx.ReplyEmbed(authorembed)

		return nil
	}
}
