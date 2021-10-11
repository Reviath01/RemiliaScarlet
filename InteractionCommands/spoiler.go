package interactioncommands

import (
	"fmt"

	multiplexer "git.randomchars.net/Reviath/RemiliaScarlet/Multiplexer"
	"git.randomchars.net/Reviath/RemiliaScarlet/interactions"
	"git.randomchars.net/Reviath/RemiliaScarlet/sql"
	embedutil "git.randomchars.net/freenitori/embedutil"
	"github.com/bwmarrin/discordgo"
)

// SpoilerCommand is spoiler command for interactions
func SpoilerCommand(session *discordgo.Session, interaction interactions.Interaction) interactions.InteractionResponse {
	switch sql.CheckLanguage(interaction.GuildID) {
	case "tr":
		if sql.IsBlocked(interaction.GuildID, "spoiler") == "true" {
			return multiplexer.CreateResponse("Bu komut bu sunucuda engellenmiş.")
		}

		spoilerembed := embedutil.New("", fmt.Sprintf("|| %s ||", interaction.Data.Options[0].Value.(string)))
		spoilerembed.Color = 0xe9ff00
		return multiplexer.CreateEmbedResponse(spoilerembed)
	default:

		if sql.IsBlocked(interaction.GuildID, "spoiler") == "true" {
			return multiplexer.CreateResponse("This command is blocked on this guild.")
		}

		spoilerembed := embedutil.New("", fmt.Sprintf("|| %s ||", interaction.Data.Options[0].Value.(string)))
		spoilerembed.Color = 0xe9ff00
		return multiplexer.CreateEmbedResponse(spoilerembed)
	}
}
