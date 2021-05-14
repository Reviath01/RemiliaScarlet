package interactioncommands

import (
	embedutil "git.randomchars.net/FreeNitori/EmbedUtil"
	multiplexer "git.randomchars.net/Reviath/RemiliaScarlet/Multiplexer"
	"git.randomchars.net/Reviath/RemiliaScarlet/interactions"
	"git.randomchars.net/Reviath/RemiliaScarlet/sql"
	"github.com/bwmarrin/discordgo"
)

// IconCommand is icon command for interactions
func IconCommand(session *discordgo.Session, interaction interactions.Interaction) interactions.InteractionResponse {
	guild, _ := session.Guild(string(interaction.GuildID))
	switch sql.CheckLanguage(interaction.GuildID) {
	case "tr":
		if sql.IsBlocked(interaction.GuildID, "icon") == "true" {
			return multiplexer.CreateResponse("Bu komut bu sunucuda engellenmiş.")
		}

		iconembed := embedutil.New("", "")
		iconembed.Color = 0x00f6ff
		iconembed.SetImage(guild.IconURL())
		return multiplexer.CreateEmbedResponse(iconembed)
	default:

		if sql.IsBlocked(interaction.GuildID, "icon") == "true" {
			return multiplexer.CreateResponse("This command is blocked on this guild.")
		}

		iconembed := embedutil.New("", "")
		iconembed.Color = 0x00f6ff
		iconembed.SetImage(guild.IconURL())
		return multiplexer.CreateEmbedResponse(iconembed)
	}
}
