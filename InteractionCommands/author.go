package interactioncommands

import (
	"fmt"

	multiplexer "git.randomchars.net/Reviath/RemiliaScarlet/Multiplexer"
	"git.randomchars.net/Reviath/RemiliaScarlet/config"
	"git.randomchars.net/Reviath/RemiliaScarlet/interactions"
	"git.randomchars.net/Reviath/RemiliaScarlet/sql"
	embedutil "git.randomchars.net/freenitori/embedutil"
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
