package commands

import (
	"strconv"

	ctx "git.randomchars.net/Reviath/RemiliaScarlet/CommandHandler/Context"
	"git.randomchars.net/Reviath/RemiliaScarlet/sql"
	"github.com/bwmarrin/discordgo"
)

type Ping struct {
}

func (p Ping) Execute(ctx ctx.Ctx, session *discordgo.Session) error {
	if sql.CheckLanguage(ctx.Guild().ID) == "tr" {
		if sql.IsBlocked(ctx.Guild().ID, "ping") == "true" {
			_, _ = session.ChannelMessageSend(ctx.Channel().ID, "Bu komut bu sunucuda engellenmiş.")
			return nil
		}
		_, _ = session.ChannelMessageSend(ctx.Channel().ID, "Pong! "+strconv.Itoa(int(session.HeartbeatLatency().Milliseconds()))+"ms")
		return nil
	}

	if sql.IsBlocked(ctx.Guild().ID, "ping") == "true" {
		_, _ = session.ChannelMessageSend(ctx.Channel().ID, "This command is blocked on this guild.")
		return nil
	}

	_, _ = session.ChannelMessageSend(ctx.Channel().ID, "Pong! "+strconv.Itoa(int(session.HeartbeatLatency().Milliseconds()))+"ms")
	return nil
}
