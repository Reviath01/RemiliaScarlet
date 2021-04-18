package commands

import (
	"strings"

	multiplexer "git.randomchars.net/Reviath/RemiliaScarlet/Multiplexer"
	CommandHandler "git.randomchars.net/Reviath/RemiliaScarlet/handler"
	"git.randomchars.net/Reviath/RemiliaScarlet/sql"
)

func KickCommand(ctx CommandHandler.Context, _ []string) error {
	if sql.CheckLanguage(ctx.Guild.ID) == "tr" {
		if !multiplexer.CheckKickPermission(ctx.Session, ctx.Message.Author.ID, ctx.Channel.ID) {
			ctx.Reply("Yeterli yetkiye sahip değilsin.")
			return nil
		}
		if sql.IsBlocked(ctx.Guild.ID, "kick") == "true" {
			ctx.Reply("This command is blocked on this guild.")
			return nil
		}

		if len(strings.Join(multiplexer.GetArgs(ctx.Message.Content, multiplexer.GetPrefix()), " ")) < 1 {
			ctx.Reply("Bir üye belirtmelisin.")

			return nil
		}
		u, err := ctx.Session.User(multiplexer.GetUser(multiplexer.GetArgs(ctx.Message.Content, multiplexer.GetPrefix())[0]))
		if err == nil {
			err = ctx.Session.GuildMemberDelete(ctx.Guild.ID, u.ID)
			if err != nil {
				ctx.Reply("Yeterli yetkim yok.")

				return nil
			}
			ctx.Reply("Belirtilen kişi sunucudan atıldı.")

			return nil
		}
		ctx.Reply("Bir üye belirtmelisin.")
		return nil
	}

	if !multiplexer.CheckKickPermission(ctx.Session, ctx.Message.Author.ID, ctx.Channel.ID) {
		ctx.Reply("You don't have enough permission.")
		return nil
	}

	if sql.IsBlocked(ctx.Guild.ID, "kick") == "true" {
		ctx.Reply("This command is blocked on this guild.")
		return nil
	}

	if len(strings.Join(multiplexer.GetArgs(ctx.Message.Content, multiplexer.GetPrefix()), " ")) < 1 {
		ctx.Reply("You need to specify the user.")

		return nil
	}

	u, err := ctx.Session.User(multiplexer.GetUser(multiplexer.GetArgs(ctx.Message.Content, multiplexer.GetPrefix())[0]))
	if err == nil {
		err = ctx.Session.GuildMemberDelete(ctx.Guild.ID, u.ID)
		if err != nil {
			ctx.Reply("I do not have enough permission.")

			return nil
		}
		ctx.Reply("Kicked specified user.")

		return nil
	}
	ctx.Reply("You need to specify the user.")
	return nil
}
