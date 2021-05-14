package commands

import (
	"strings"

	multiplexer "git.randomchars.net/Reviath/RemiliaScarlet/Multiplexer"
	CommandHandler "git.randomchars.net/Reviath/RemiliaScarlet/handler"
	"git.randomchars.net/Reviath/RemiliaScarlet/sql"
)

// UnbanCommand is a handler for unban command
func UnbanCommand(ctx CommandHandler.Context, _ []string) error {
	switch sql.CheckLanguage(ctx.Guild.ID) {
	case "tr":
		if !multiplexer.CheckBanPermission(ctx.Session, ctx.Message.Author.ID, ctx.Channel.ID) {
			ctx.Reply("Yeterli yetkiye sahip değilsin.")
			return nil
		}
		if sql.IsBlocked(ctx.Guild.ID, "unban") == "true" {
			ctx.Reply("Bu komut bu sunucuda engellenmiş.")
			return nil
		}

		if len(strings.Join(multiplexer.GetArgs(ctx.Message.Content, multiplexer.GetPrefix()), " ")) < 1 {
			ctx.Reply("Bir üye belirtmelisin.")
			return nil
		}

		u, err := ctx.Session.User(multiplexer.GetUser(multiplexer.GetArgs(ctx.Message.Content, multiplexer.GetPrefix())[0]))
		if err == nil {
			err = ctx.Session.GuildBanDelete(ctx.Guild.ID, u.ID)
			if err != nil {
				ctx.Reply("Bu üye banlanmamış veya yeterli yetkiye sahip değilim.")
				return nil
			}
			ctx.Reply("Belirtilen kişinin banı kaldırıldı.")
			return nil
		}
		ctx.Reply("Bir üye belirtmelisin.")
		return nil
	default:

		if !multiplexer.CheckBanPermission(ctx.Session, ctx.Message.Author.ID, ctx.Channel.ID) {
			ctx.Reply("You don't have enough permission.")
			return nil
		}

		if sql.IsBlocked(ctx.Guild.ID, "unban") == "true" {
			ctx.Reply("This command is blocked on this guild.")
			return nil
		}

		if len(strings.Join(multiplexer.GetArgs(ctx.Message.Content, multiplexer.GetPrefix()), " ")) < 1 {
			ctx.Reply("You need to specify the user.")

			return nil
		}

		u, err := ctx.Session.User(multiplexer.GetUser(multiplexer.GetArgs(ctx.Message.Content, multiplexer.GetPrefix())[0]))

		if err == nil {
			err = ctx.Session.GuildBanDelete(ctx.Guild.ID, u.ID)
			if err != nil {
				ctx.Reply("This user is not banned or I don't have enough permission.")
				return nil
			}
			ctx.Reply("Unbanned specified user.")

			return nil
		}
		ctx.Reply("You need to specify the user.")
		return nil
	}
}
