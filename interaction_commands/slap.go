package interaction_commands

import (
	"fmt"

	embedutil "git.randomchars.net/FreeNitori/EmbedUtil"
	multiplexer "git.randomchars.net/Reviath/RemiliaScarlet/Multiplexer"
	"git.randomchars.net/Reviath/RemiliaScarlet/interactions"
	"git.randomchars.net/Reviath/RemiliaScarlet/sql"
	"github.com/bwmarrin/discordgo"
)

//Slap slash command.
func SlapCommand(session *discordgo.Session, interaction interactions.Interaction) interactions.InteractionResponse {
	switch sql.CheckLanguage(interaction.GuildID) {
	case "tr":
		if sql.IsBlocked(interaction.GuildID, "slap") == "true" {
			return multiplexer.CreateResponse("Bu komut bu sunucuda engellenmiş.")
		}

		u, err := session.User(multiplexer.GetUser(interaction.Data.Options[0].Value.(string)))
		if err == nil {
			embed := embedutil.New("", fmt.Sprintf("<@%s>, <@%s> isimli kişiye vurdu", interaction.Member.User.ID, u.ID))
			embed.Color = 0xff1000
			embed.SetImage("https://images-ext-1.discordapp.net/external/79sCWyD-TmmyjFxlaQIxAkAANAfV529d-LDHNkGDM0M/%3Fitemid%3D10426943/https/media1.tenor.com/images/b6d8a83eb652a30b95e87cf96a21e007/tenor.gif")
			return multiplexer.CreateEmbedResponse(embed)
		}
		return multiplexer.CreateResponse("Bir üye belirtmelisin.")
	default:

		if sql.IsBlocked(interaction.GuildID, "slap") == "true" {
			return multiplexer.CreateResponse("This command is blocked on this guild.")
		}

		u, err := session.User(multiplexer.GetUser(interaction.Data.Options[0].Value.(string)))
		if err == nil {
			embed := embedutil.New("", fmt.Sprintf("<@%s> slaps <@%s>", interaction.Member.User.ID, u.ID))
			embed.Color = 0xff1000
			embed.SetImage("https://images-ext-1.discordapp.net/external/79sCWyD-TmmyjFxlaQIxAkAANAfV529d-LDHNkGDM0M/%3Fitemid%3D10426943/https/media1.tenor.com/images/b6d8a83eb652a30b95e87cf96a21e007/tenor.gif")
			return multiplexer.CreateEmbedResponse(embed)
		}
		return multiplexer.CreateResponse("You need to specify the user.")
	}
}
