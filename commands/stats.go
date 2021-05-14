package commands

import (
	"runtime"
	"strconv"
	"time"

	embedutil "git.randomchars.net/FreeNitori/EmbedUtil"
	multiplexer "git.randomchars.net/Reviath/RemiliaScarlet/Multiplexer"
	CommandHandler "git.randomchars.net/Reviath/RemiliaScarlet/handler"
	"git.randomchars.net/Reviath/RemiliaScarlet/sql"
	"github.com/bwmarrin/discordgo"
)

var startTime = time.Now()

//Stats command
func StatsCommand(ctx CommandHandler.Context, _ []string) error {
	switch sql.CheckLanguage(ctx.Guild.ID) {
	case "tr":
		if sql.IsBlocked(ctx.Guild.ID, "stats") == "true" {
			ctx.Reply("Bu komut bu sunucuda engellenmiş.")
			return nil
		}

		statembed := embedutil.New("İstatistikler", "")
		statembed.Color = 0x00ff03
		statembed.AddField("Çalışma Süresi", multiplexer.GetDuration(time.Since(startTime)), true)
		statembed.AddField("Go Versiyonu", runtime.Version(), true)
		statembed.AddField("Discordgo Versiyonu", discordgo.VERSION, true)
		statembed.AddField("Sunucular", strconv.Itoa(len(ctx.Session.State.Guilds)), true)
		statembed.AddField("Goroutine'ler", strconv.Itoa(runtime.NumGoroutine()), true)
		ctx.ReplyEmbed(statembed)

		return nil
	default:

		if sql.IsBlocked(ctx.Guild.ID, "stats") == "true" {
			ctx.Reply("This command is blocked on this guild.")
			return nil
		}

		statembed := embedutil.New("Stats", "")
		statembed.Color = 0x00ff03
		statembed.AddField("Uptime", multiplexer.GetDuration(time.Since(startTime)), true)
		statembed.AddField("Go version", runtime.Version(), true)
		statembed.AddField("Discordgo version", discordgo.VERSION, true)
		statembed.AddField("Server size", strconv.Itoa(len(ctx.Session.State.Guilds)), true)
		statembed.AddField("Goroutines", strconv.Itoa(runtime.NumGoroutine()), true)
		ctx.ReplyEmbed(statembed)

		return nil
	}
}
