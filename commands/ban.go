package commands

import (
	"fmt"
	"strings"

	multiplexer "git.randomchars.net/Reviath/RemiliaScarlet/Multiplexer"
	CommandHandler "git.randomchars.net/Reviath/RemiliaScarlet/handler"
	"git.randomchars.net/Reviath/RemiliaScarlet/sql"
)

func BanCommand(ctx CommandHandler.Context, _ []string) error {
	if sql.CheckLanguage(ctx.Guild.ID) == "tr" {
		if sql.IsBlocked(ctx.Guild.ID, "ban") == "true" {
			ctx.Reply("Bu komut bu sunucuda engellenmiş.")
			return nil
		}

		if strings.Join(multiplexer.GetArgs(ctx.Message.Content, multiplexer.GetPrefix()), " ") == "" {
			ctx.Reply("Bir üye belirtmelisiniz.")
			return nil
		} else {
			u, err := ctx.Session.User(multiplexer.GetUser(multiplexer.GetArgs(ctx.Message.Content, multiplexer.GetPrefix())[0]))
			if err != nil {
				ctx.Reply("Bir üye belirtmelisiniz.")
			} else {
				err = ctx.Session.GuildBanCreate(ctx.Guild.ID, u.ID, 0)
				if err != nil {
					ctx.Reply("Yeterli yetkiye sahip değilim.")
				}
			}
		}
	}

	if sql.IsBlocked(ctx.Guild.ID, "ban") == "true" {
		ctx.Reply("This command is blocked on this guild.")
		return nil
	}

	if strings.Join(multiplexer.GetArgs(ctx.Message.Content, multiplexer.GetPrefix()), " ") == "" {
		ctx.Reply("You need to specify the user.")
		return nil
	} else {
		u, err := ctx.Session.User(multiplexer.GetUser(multiplexer.GetArgs(ctx.Message.Content, multiplexer.GetPrefix())[0]))
		if err != nil {
			ctx.Reply("You need to specify the user.")
		} else {
			err = ctx.Session.GuildBanCreate(ctx.Guild.ID, u.ID, 0)
			if err != nil {
				ctx.Reply(fmt.Sprintf("An error occurred %s", err.Error()))
				return nil
			}
			ctx.Reply("Successfully banned user.")
		}
	}
	return nil
}
