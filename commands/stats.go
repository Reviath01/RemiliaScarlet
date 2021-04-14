package commands

import (
	"runtime"
	"strconv"
	"time"

	ctx "git.randomchars.net/Reviath/RemiliaScarlet/CommandHandler/Context"
	embedutil "git.randomchars.net/Reviath/RemiliaScarlet/EmbedUtil"
	multiplexer "git.randomchars.net/Reviath/RemiliaScarlet/Multiplexer"
	"git.randomchars.net/Reviath/RemiliaScarlet/sql"
	"github.com/bwmarrin/discordgo"
)

type Stats struct {
}

var startTime = time.Now()

func (s Stats) Execute(ctx ctx.Ctx, session *discordgo.Session) error {
	db := sql.Connect()

	type Tag struct {
		isblocked string
	}

	var tag Tag

	if sql.CheckLanguage(ctx.Guild().ID) == "tr" {
		err := db.QueryRow("SELECT isblocked FROM disabledcommands WHERE commandname ='stats' AND guildid ='" + ctx.Guild().ID + "'").Scan(&tag.isblocked)

		if err == nil && tag.isblocked == "True" {
			_, _ = session.ChannelMessageSend(ctx.Channel().ID, "Bu komut bu sunucuda engellenmiş.")
			return nil
		}

		statembed := embedutil.NewEmbed().
			SetTitle("İstatistikler").
			SetColor(0x00ff03).
			AddField("Çalışma Süresi", multiplexer.GetDuration(time.Since(startTime))).
			AddField("Go Versiyonu", runtime.Version()).
			AddField("Discordgo Versiyonu", discordgo.VERSION).
			AddField("Sunucular", strconv.Itoa(len(session.State.Guilds))).
			AddField("Goroutine'ler", strconv.Itoa(runtime.NumGoroutine())).MessageEmbed
		_, _ = session.ChannelMessageSendEmbed(ctx.Channel().ID, statembed)

		return nil

	}

	err := db.QueryRow("SELECT isblocked FROM disabledcommands WHERE commandname ='stats' AND guildid ='" + ctx.Guild().ID + "'").Scan(&tag.isblocked)

	if err == nil && tag.isblocked == "True" {
		_, _ = session.ChannelMessageSend(ctx.Channel().ID, "This command is blocked on this guild.")
		return nil
	}

	statembed := embedutil.NewEmbed().
		SetTitle("Stats").
		SetColor(0x00ff03).
		AddField("Uptime", multiplexer.GetDuration(time.Since(startTime))).
		AddField("Go version", runtime.Version()).
		AddField("Discordgo version", discordgo.VERSION).
		AddField("Server size", strconv.Itoa(len(session.State.Guilds))).
		AddField("Goroutines", strconv.Itoa(runtime.NumGoroutine())).MessageEmbed
	_, _ = session.ChannelMessageSendEmbed(ctx.Channel().ID, statembed)

	return nil
}
