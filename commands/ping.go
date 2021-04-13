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
	db := sql.Connect()

	type Tag struct {
		isblocked string
		lang      string
	}

	var tag Tag

	err := db.QueryRow("SELECT language FROM languages WHERE guildid ='" + ctx.Guild().ID + "'").Scan(&tag.lang)
	if err == nil && tag.lang == "tr" {
		err = db.QueryRow("SELECT isblocked FROM disabledcommands WHERE commandname ='ping' AND guildid ='" + ctx.Guild().ID + "'").Scan(&tag.isblocked)

		if err == nil && tag.isblocked == "True" {
			_, _ = session.ChannelMessageSend(ctx.Channel().ID, "Bu komut bu sunucuda engellenmiş.")
			return nil
		}
		_, _ = session.ChannelMessageSend(ctx.Channel().ID, "Pong! "+strconv.Itoa(int(session.HeartbeatLatency().Milliseconds()))+"ms")
		return nil
	}

	err = db.QueryRow("SELECT isblocked FROM disabledcommands WHERE commandname ='ping' AND guildid ='" + ctx.Guild().ID + "'").Scan(&tag.isblocked)

	if err == nil && tag.isblocked == "True" {
		_, _ = session.ChannelMessageSend(ctx.Channel().ID, "This command is blocked on this guild.")
		return nil
	}

	_, _ = session.ChannelMessageSend(ctx.Channel().ID, "Pong! "+strconv.Itoa(int(session.HeartbeatLatency().Milliseconds()))+"ms")
	return nil
}
