package interactioncommands

import (
	"fmt"

	multiplexer "git.randomchars.net/Reviath/RemiliaScarlet/Multiplexer"
	"git.randomchars.net/Reviath/RemiliaScarlet/interactions"
	"git.randomchars.net/Reviath/RemiliaScarlet/sql"
	embedutil "git.randomchars.net/freenitori/embedutil"
	"github.com/bwmarrin/discordgo"
)

// RolesCommand is roles command for interactions
func RolesCommand(session *discordgo.Session, interaction interactions.Interaction) interactions.InteractionResponse {
	var roles string

	Guild, _ := session.Guild(interaction.GuildID)
	switch sql.CheckLanguage(interaction.GuildID) {
	case "tr":
		if sql.IsBlocked(interaction.GuildID, "roles") == "true" {
			return multiplexer.CreateResponse("Bu komut bu sunucuda engellenmiş.")
		}
		for _, i := range Guild.Roles {
			roles += fmt.Sprintf("<@&%s> ,\n", i.ID)
		}
		embed := embedutil.New("", "")
		embed.AddField("Roller:", roles, true)
		return multiplexer.CreateEmbedResponse(embed)
	default:
		if sql.IsBlocked(interaction.GuildID, "roles") == "true" {
			return multiplexer.CreateResponse("This command is blocked on this guild.")
		}
		for _, i := range Guild.Roles {
			roles += fmt.Sprintf("<@&%s> ,\n", i.ID)
		}
		embed := embedutil.New("", "")
		embed.AddField("Roles:", roles, true)

		return multiplexer.CreateEmbedResponse(embed)
	}
}
