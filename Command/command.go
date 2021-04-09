package command

import (
	ctx "git.randomchars.net/Reviath/RemiliaScarlet/Context"
	"github.com/bwmarrin/discordgo"
)

type Command interface {
	Execute(ctx ctx.Ctx,session *discordgo.Session) error
}
