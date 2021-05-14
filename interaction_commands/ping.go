package interaction_commands

import (
	"fmt"
	"strconv"

	multiplexer "git.randomchars.net/Reviath/RemiliaScarlet/Multiplexer"
	"git.randomchars.net/Reviath/RemiliaScarlet/interactions"
	"git.randomchars.net/Reviath/RemiliaScarlet/sql"
	"github.com/bwmarrin/discordgo"
)

//Ping slash command.
func PingCommand(session *discordgo.Session, interaction interactions.Interaction) interactions.InteractionResponse {
	if sql.CheckLanguage(interaction.GuildID) == "tr" && sql.IsBlocked(interaction.GuildID, "ping") == "true" {
		return multiplexer.CreateResponse("Bu komut bu sunucuda kullanıma kapatılmış.")
	}
	if sql.IsBlocked(interaction.GuildID, "ping") == "true" {
		return multiplexer.CreateResponse("This command is blocked on this guild.")
	}
	return multiplexer.CreateResponse(fmt.Sprintf("Pong! %sms", strconv.Itoa(int(session.HeartbeatLatency().Milliseconds()))))
}
