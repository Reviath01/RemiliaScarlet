package interaction_commands

import (
	embedutil "git.randomchars.net/Reviath/RemiliaScarlet/EmbedUtil"
	multiplexer "git.randomchars.net/Reviath/RemiliaScarlet/Multiplexer"
	"git.randomchars.net/Reviath/RemiliaScarlet/interactions"
	"git.randomchars.net/Reviath/RemiliaScarlet/sql"
	"github.com/bwmarrin/discordgo"
)

func EmbedCommand(session *discordgo.Session, interaction interactions.Interaction) interactions.InteractionResponse {
	if sql.CheckLanguage(interaction.GuildID) == "tr" {
		if sql.IsBlocked(interaction.GuildID, "embed") == "true" {
			return multiplexer.CreateResponse("Bu komut bu sunucuda engellenmiş.")
		}

		embed := embedutil.NewEmbed().
			SetColor(0xc000ff).
			SetDescription(interaction.Data.Options[0].Value.(string)).MessageEmbed
		return multiplexer.CreateEmbedResponse(embed)
	}

	if sql.IsBlocked(interaction.GuildID, "embed") == "true" {
		return multiplexer.CreateResponse("This command is blocked on this guild.")
	}

	embed := embedutil.NewEmbed().
		SetColor(0xc000ff).
		SetDescription(interaction.Data.Options[0].Value.(string)).MessageEmbed
	return multiplexer.CreateEmbedResponse(embed)
}
