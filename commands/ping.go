package commands

import (
	ctx "git.randomchars.net/Reviath/handlers/Context"
	"github.com/bwmarrin/discordgo"
	"strconv"
)

type Ping struct {

}

func (p Ping) Execute(ctx ctx.Ctx, session *discordgo.Session) error {
	_, err := session.ChannelMessageSend(ctx.Channel().ID, "Pong! " + strconv.Itoa(int(session.HeartbeatLatency().Milliseconds())) + "ms")
	return err
}
