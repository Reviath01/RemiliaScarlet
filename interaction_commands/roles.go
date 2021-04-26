package interaction_commands

import (
	"fmt"

	embedutil "git.randomchars.net/Reviath/RemiliaScarlet/EmbedUtil"
	multiplexer "git.randomchars.net/Reviath/RemiliaScarlet/Multiplexer"
	"git.randomchars.net/Reviath/RemiliaScarlet/interactions"
	"git.randomchars.net/Reviath/RemiliaScarlet/sql"
	"github.com/bwmarrin/discordgo"
)

func RolesCommand(session *discordgo.Session, interaction interactions.Interaction) interactions.InteractionResponse {
	var roles string

	Guild, _ := session.Guild(interaction.GuildID)

	if sql.CheckLanguage(interaction.GuildID) == "tr" {
		if sql.IsBlocked(interaction.GuildID, "roles") == "true" {
			return multiplexer.CreateResponse("Bu komut bu sunucuda engellenmiş.")
		}
		for _, i := range Guild.Roles {
			roles += fmt.Sprintf("<@&%s> ,\n", i.ID)
		}
		embed := embedutil.NewEmbed().
			AddField("Roller:", roles).MessageEmbed
		return multiplexer.CreateEmbedResponse(embed)
	}
	if sql.IsBlocked(interaction.GuildID, "roles") == "true" {
		return multiplexer.CreateResponse("This command is blocked on this guild.")
	}
	for _, i := range Guild.Roles {
		roles += fmt.Sprintf("<@&%s> ,\n", i.ID)
	}
	embed := embedutil.NewEmbed().
		AddField("Roles:", roles).MessageEmbed

	return multiplexer.CreateEmbedResponse(embed)
}
