package interaction_commands

import (
	"runtime"
	"strconv"
	"time"

	embedutil "git.randomchars.net/Reviath/RemiliaScarlet/EmbedUtil"
	multiplexer "git.randomchars.net/Reviath/RemiliaScarlet/Multiplexer"
	"git.randomchars.net/Reviath/RemiliaScarlet/interactions"
	"git.randomchars.net/Reviath/RemiliaScarlet/sql"
	"github.com/bwmarrin/discordgo"
)

var startTime = time.Now()

func StatsCommand(session *discordgo.Session, interaction interactions.Interaction) interactions.InteractionResponse {
	if sql.CheckLanguage(interaction.GuildID) == "tr" {
		if sql.IsBlocked(interaction.GuildID, "stats") == "true" {
			return multiplexer.CreateResponse("Bu komut bu sunucuda engellenmiş.")
		}

		statembed := embedutil.NewEmbed().
			SetTitle("İstatistikler").
			SetColor(0x00ff03).
			AddField("Çalışma Süresi", multiplexer.GetDuration(time.Since(startTime))).
			AddField("Go Versiyonu", runtime.Version()).
			AddField("Discordgo Versiyonu", discordgo.VERSION).
			AddField("Sunucular", strconv.Itoa(len(session.State.Guilds))).
			AddField("Goroutine'ler", strconv.Itoa(runtime.NumGoroutine())).MessageEmbed
		return multiplexer.CreateEmbedResponse(statembed)
	}

	if sql.IsBlocked(interaction.GuildID, "stats") == "true" {
		return multiplexer.CreateResponse("This command is blocked on this guild.")
	}

	statembed := embedutil.NewEmbed().
		SetTitle("Stats").
		SetColor(0x00ff03).
		AddField("Uptime", multiplexer.GetDuration(time.Since(startTime))).
		AddField("Go version", runtime.Version()).
		AddField("Discordgo version", discordgo.VERSION).
		AddField("Server size", strconv.Itoa(len(session.State.Guilds))).
		AddField("Goroutines", strconv.Itoa(runtime.NumGoroutine())).MessageEmbed
	return multiplexer.CreateEmbedResponse(statembed)

}
