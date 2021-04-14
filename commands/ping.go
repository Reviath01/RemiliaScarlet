package commands

import (
	"strconv"

	CommandHandler "git.randomchars.net/Reviath/RemiliaScarlet/handler"
	"git.randomchars.net/Reviath/RemiliaScarlet/sql"
)

func PingCommand(ctx CommandHandler.Context, _ []string) error {
	if sql.CheckLanguage(ctx.Guild.ID) == "tr" && sql.IsBlocked(ctx.Guild.ID, "ping") == "true" {
		_, _ = ctx.Reply("Bu komut bu sunucuda kullanıma kapatılmış.")
		return nil
	}
	if sql.IsBlocked(ctx.Guild.ID, "ping") == "true" {
		_, _ = ctx.Reply("This command is blocked on this guild.")
		return nil
	}
	_, _ = ctx.Reply("Pong! " + strconv.Itoa(int(ctx.Session.HeartbeatLatency().Milliseconds())) + "ms")
	return nil
}
