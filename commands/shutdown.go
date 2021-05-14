package commands

import (
	"os"

	CommandHandler "git.randomchars.net/Reviath/RemiliaScarlet/handler"
)

// ShutdownCommand is a handler for shutdown command
func ShutdownCommand(ctx CommandHandler.Context, _ []string) error {
	ctx.Reply("Performing shutdown.")

	ctx.Session.Close()
	os.Exit(1)
	return nil
}
