package commands

import (
	multiplexer "git.randomchars.net/Reviath/RemiliaScarlet/Multiplexer"
	CommandHandler "git.randomchars.net/Reviath/RemiliaScarlet/handler"
)

// RemovePrefixCommand is a handler for remove prefix command
func RemovePrefixCommand(ctx CommandHandler.Context, _ []string) error {
	ctx.Handler.RemovePrefix(multiplexer.GetArgs(ctx.Message.Content, multiplexer.GetPrefix())[0])
	ctx.Reply("Successful")
	return nil
}
