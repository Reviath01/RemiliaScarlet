package commands

import (
	"fmt"
	"strconv"

	CommandHandler "git.randomchars.net/Reviath/RemiliaScarlet/handler"
	"git.randomchars.net/Reviath/RemiliaScarlet/sql"
)

// PingCommand is a handler for ping command
func PingCommand(ctx CommandHandler.Context, _ []string) error {
	if sql.CheckLanguage(ctx.Guild.ID) == "tr" && sql.IsBlocked(ctx.Guild.ID, "ping") == "true" {
		ctx.Reply("Bu komut bu sunucuda kullanıma kapatılmış.")
		return nil
	}
	if sql.IsBlocked(ctx.Guild.ID, "ping") == "true" {
		ctx.Reply("This command is blocked on this guild.")
		return nil
	}
	ctx.Reply(fmt.Sprintf("Pong! %sms", strconv.Itoa(int(ctx.Session.HeartbeatLatency().Milliseconds()))))
	return nil
}
