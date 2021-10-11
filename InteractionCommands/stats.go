package interactioncommands

import (
	"runtime"
	"strconv"
	"time"

	multiplexer "git.randomchars.net/Reviath/RemiliaScarlet/Multiplexer"
	"git.randomchars.net/Reviath/RemiliaScarlet/interactions"
	"git.randomchars.net/Reviath/RemiliaScarlet/sql"
	embedutil "git.randomchars.net/freenitori/embedutil"
	"github.com/bwmarrin/discordgo"
)

var startTime = time.Now()

// StatsCommand is stats command for interactions
func StatsCommand(session *discordgo.Session, interaction interactions.Interaction) interactions.InteractionResponse {
	switch sql.CheckLanguage(interaction.GuildID) {
	case "tr":
		if sql.IsBlocked(interaction.GuildID, "stats") == "true" {
			return multiplexer.CreateResponse("Bu komut bu sunucuda engellenmiş.")
		}

		statembed := embedutil.New("İstatistikler", "")
		statembed.Color = 0x00ff03
		statembed.AddField("Çalışma Süresi", multiplexer.GetDuration(time.Since(startTime)), true)
		statembed.AddField("Go Versiyonu", runtime.Version(), true)
		statembed.AddField("Discordgo Versiyonu", discordgo.VERSION, true)
		statembed.AddField("Sunucular", strconv.Itoa(len(session.State.Guilds)), true)
		statembed.AddField("Goroutine'ler", strconv.Itoa(runtime.NumGoroutine()), true)
		return multiplexer.CreateEmbedResponse(statembed)
	default:

		if sql.IsBlocked(interaction.GuildID, "stats") == "true" {
			return multiplexer.CreateResponse("This command is blocked on this guild.")
		}

		statembed := embedutil.New("Stats", "")
		statembed.Color = 0x00ff03
		statembed.AddField("Uptime", multiplexer.GetDuration(time.Since(startTime)), true)
		statembed.AddField("Go version", runtime.Version(), true)
		statembed.AddField("Discordgo version", discordgo.VERSION, true)
		statembed.AddField("Server size", strconv.Itoa(len(session.State.Guilds)), true)
		statembed.AddField("Goroutines", strconv.Itoa(runtime.NumGoroutine()), true)
		return multiplexer.CreateEmbedResponse(statembed)
	}
}
