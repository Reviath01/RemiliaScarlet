package interaction_commands

import (
	embedutil "git.randomchars.net/Reviath/RemiliaScarlet/EmbedUtil"
	multiplexer "git.randomchars.net/Reviath/RemiliaScarlet/Multiplexer"
	"git.randomchars.net/Reviath/RemiliaScarlet/interactions"
	"git.randomchars.net/Reviath/RemiliaScarlet/sql"
	"github.com/bwmarrin/discordgo"
)

func IconCommand(session *discordgo.Session, interaction interactions.Interaction) interactions.InteractionResponse {
	guild, _ := session.Guild(string(interaction.GuildID))
	if sql.CheckLanguage(interaction.GuildID) == "tr" {
		if sql.IsBlocked(interaction.GuildID, "icon") == "true" {
			return multiplexer.CreateResponse("Bu komut bu sunucuda engellenmiş.")
		}

		iconembed := embedutil.NewEmbed().
			SetColor(0x00f6ff).
			SetImage(guild.IconURL()).MessageEmbed
		return multiplexer.CreateEmbedResponse(iconembed)
	}

	if sql.IsBlocked(interaction.GuildID, "icon") == "true" {
		return multiplexer.CreateResponse("This command is blocked on this guild.")
	}

	iconembed := embedutil.NewEmbed().
		SetColor(0x00f6ff).
		SetImage(guild.IconURL()).MessageEmbed
	return multiplexer.CreateEmbedResponse(iconembed)
}