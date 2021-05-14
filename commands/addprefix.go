package commands

import (
	multiplexer "git.randomchars.net/Reviath/RemiliaScarlet/Multiplexer"
	CommandHandler "git.randomchars.net/Reviath/RemiliaScarlet/handler"
)

//Add prefix command
func AddPrefixCommand(ctx CommandHandler.Context, _ []string) error {
	ctx.Handler.AddPrefix(multiplexer.GetArgs(ctx.Message.Content, multiplexer.GetPrefix())[0])
	ctx.Reply("Successful")
	return nil
}
