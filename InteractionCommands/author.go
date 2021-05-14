package interaction_commands

import (
	"fmt"

	embedutil "git.randomchars.net/FreeNitori/EmbedUtil"
	multiplexer "git.randomchars.net/Reviath/RemiliaScarlet/Multiplexer"
	"git.randomchars.net/Reviath/RemiliaScarlet/config"
	"git.randomchars.net/Reviath/RemiliaScarlet/interactions"
	"git.randomchars.net/Reviath/RemiliaScarlet/sql"
	"github.com/bwmarrin/discordgo"
)

// AuthorCommand is author command for interactions
func AuthorCommand(session *discordgo.Session, interaction interactions.Interaction) interactions.InteractionResponse {
	config.ReadConfig()
	switch sql.CheckLanguage(interaction.GuildID) {
	case "tr":
		u, _ := session.User(config.Owner)
		embed := embedutil.New("", "")
		embed.Color = 0x007bff
		embed.AddField("Sahibim:", fmt.Sprintf("<@%s> ([%s#%s](https://discord.com/users/%s))", u.ID, u.Username, u.Discriminator, u.ID), true)
		return multiplexer.CreateEmbedResponse(embed)

	default:
		u, _ := session.User(config.Owner)
		embed := embedutil.New("", "")
		embed.Color = 0x007bff
		embed.AddField("My Author:", fmt.Sprintf("<@%s> ([%s#%s](https://discord.com/users/%s))", u.ID, u.Username, u.Discriminator, u.ID), true)
		return multiplexer.CreateEmbedResponse(embed)
	}
}
