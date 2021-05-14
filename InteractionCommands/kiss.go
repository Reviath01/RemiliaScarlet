package interaction_commands

import (
	"fmt"

	embedutil "git.randomchars.net/FreeNitori/EmbedUtil"
	multiplexer "git.randomchars.net/Reviath/RemiliaScarlet/Multiplexer"
	"git.randomchars.net/Reviath/RemiliaScarlet/interactions"
	"git.randomchars.net/Reviath/RemiliaScarlet/sql"
	"github.com/bwmarrin/discordgo"
)

// KissCommand is kiss command for interactions
func KissCommand(session *discordgo.Session, interaction interactions.Interaction) interactions.InteractionResponse {
	switch sql.CheckLanguage(interaction.GuildID) {
	case "tr":
		if sql.IsBlocked(interaction.GuildID, "kiss") == "true" {
			return multiplexer.CreateResponse("Bu komut bu sunucuda engellenmiş.")
		}

		u, err := session.User(multiplexer.GetUser(interaction.Data.Options[0].Value.(string)))
		if err == nil {
			embed := embedutil.New("", fmt.Sprintf("<@%s>, <@%s> isimli kişiyi öptü 😘", interaction.Member.User.ID, u.ID))
			embed.Color = 0xff1000
			embed.SetImage("https://media.tenor.com/images/d68747a5865b12c465e5dff31c65d5c2/tenor.gif")
			return multiplexer.CreateEmbedResponse(embed)
		}
		return multiplexer.CreateResponse("Bir üye belirtmelisin")
	default:

		if sql.IsBlocked(interaction.GuildID, "kiss") == "true" {
			return multiplexer.CreateResponse("This command is blocked on this guild.")
		}

		u, err := session.User(multiplexer.GetUser(interaction.Data.Options[0].Value.(string)))
		if err == nil {
			embed := embedutil.New("", fmt.Sprintf("<@%s> kisses <@%s>", interaction.Member.User.ID, u.ID))
			embed.Color = 0xff1000
			embed.SetImage("https://media.tenor.com/images/d68747a5865b12c465e5dff31c65d5c2/tenor.gif")
			return multiplexer.CreateEmbedResponse(embed)
		}
		return multiplexer.CreateResponse("You need to specify the user.")
	}
}
