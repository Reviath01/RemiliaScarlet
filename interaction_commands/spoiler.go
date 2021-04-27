package interaction_commands

import (
	"fmt"

	embedutil "git.randomchars.net/Reviath/RemiliaScarlet/EmbedUtil"
	multiplexer "git.randomchars.net/Reviath/RemiliaScarlet/Multiplexer"
	"git.randomchars.net/Reviath/RemiliaScarlet/interactions"
	"git.randomchars.net/Reviath/RemiliaScarlet/sql"
	"github.com/bwmarrin/discordgo"
)

func SpoilerCommand(session *discordgo.Session, interaction interactions.Interaction) interactions.InteractionResponse {
	if sql.CheckLanguage(interaction.GuildID) == "tr" {
		if sql.IsBlocked(interaction.GuildID, "spoiler") == "true" {
			return multiplexer.CreateResponse("Bu komut bu sunucuda engellenmiş.")
		}

		spoilerembed := embedutil.NewEmbed().
			SetColor(0xe9ff00).
			SetDescription(fmt.Sprintf("|| %s ||", interaction.Data.Options[0].Value.(string))).MessageEmbed
		return multiplexer.CreateEmbedResponse(spoilerembed)
	}

	if sql.IsBlocked(interaction.GuildID, "spoiler") == "true" {
		return multiplexer.CreateResponse("This command is blocked on this guild.")
	}

	spoilerembed := embedutil.NewEmbed().
		SetColor(0xe9ff00).
		SetDescription(fmt.Sprintf("|| %s ||", interaction.Data.Options[0].Value.(string))).MessageEmbed
	return multiplexer.CreateEmbedResponse(spoilerembed)
}