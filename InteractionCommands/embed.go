package interactioncommands

import (
	embedutil "git.randomchars.net/FreeNitori/EmbedUtil"
	multiplexer "git.randomchars.net/Reviath/RemiliaScarlet/Multiplexer"
	"git.randomchars.net/Reviath/RemiliaScarlet/interactions"
	"git.randomchars.net/Reviath/RemiliaScarlet/sql"
	"github.com/bwmarrin/discordgo"
)

// EmbedCommand is embed command for interactions
func EmbedCommand(session *discordgo.Session, interaction interactions.Interaction) interactions.InteractionResponse {
	switch sql.CheckLanguage(interaction.GuildID) {
	case "tr":
		if sql.IsBlocked(interaction.GuildID, "embed") == "true" {
			return multiplexer.CreateResponse("Bu komut bu sunucuda engellenmiş.")
		}

		embed := embedutil.New("", interaction.Data.Options[0].Value.(string))
		embed.Color = 0xc000ff
		return multiplexer.CreateEmbedResponse(embed)
	default:

		if sql.IsBlocked(interaction.GuildID, "embed") == "true" {
			return multiplexer.CreateResponse("This command is blocked on this guild.")
		}

		embed := embedutil.New("", interaction.Data.Options[0].Value.(string))
		embed.Color = 0xc000ff
		return multiplexer.CreateEmbedResponse(embed)
	}
}
