package commands

import (
	"runtime"
	"strconv"
	"time"

	embedutil "git.randomchars.net/Reviath/RemiliaScarlet/EmbedUtil"
	multiplexer "git.randomchars.net/Reviath/RemiliaScarlet/Multiplexer"
	CommandHandler "git.randomchars.net/Reviath/RemiliaScarlet/handler"
	"git.randomchars.net/Reviath/RemiliaScarlet/sql"
	"github.com/bwmarrin/discordgo"
)

var startTime = time.Now()

func StatsCommand(ctx CommandHandler.Context, _ []string) error {
	if sql.CheckLanguage(ctx.Guild.ID) == "tr" {
		if sql.IsBlocked(ctx.Guild.ID, "stats") == "true" {
			ctx.Reply("Bu komut bu sunucuda engellenmiş.")
			return nil
		}

		statembed := embedutil.NewEmbed().
			SetTitle("İstatistikler").
			SetColor(0x00ff03).
			AddField("Çalışma Süresi", multiplexer.GetDuration(time.Since(startTime))).
			AddField("Go Versiyonu", runtime.Version()).
			AddField("Discordgo Versiyonu", discordgo.VERSION).
			AddField("Sunucular", strconv.Itoa(len(ctx.Session.State.Guilds))).
			AddField("Goroutine'ler", strconv.Itoa(runtime.NumGoroutine())).MessageEmbed
		ctx.ReplyEmbed(statembed)

		return nil

	}

	if sql.IsBlocked(ctx.Guild.ID, "stats") == "true" {
		ctx.Reply("This command is blocked on this guild.")
		return nil
	}

	statembed := embedutil.NewEmbed().
		SetTitle("Stats").
		SetColor(0x00ff03).
		AddField("Uptime", multiplexer.GetDuration(time.Since(startTime))).
		AddField("Go version", runtime.Version()).
		AddField("Discordgo version", discordgo.VERSION).
		AddField("Server size", strconv.Itoa(len(ctx.Session.State.Guilds))).
		AddField("Goroutines", strconv.Itoa(runtime.NumGoroutine())).MessageEmbed
	ctx.ReplyEmbed(statembed)

	return nil
}
