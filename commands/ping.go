package commands

import (
	"strconv"

	CommandHandler "git.randomchars.net/Reviath/RemiliaScarlet/handler"
)

func PingCommand(ctx CommandHandler.Context, _ []string) error {
	_, _ = ctx.Reply("Pong! " + strconv.Itoa(int(ctx.Session.HeartbeatLatency().Milliseconds())) + "ms")
	return nil
}
