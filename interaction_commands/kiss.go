package interaction_commands

import (
	"fmt"

	embedutil "git.randomchars.net/Reviath/RemiliaScarlet/EmbedUtil"
	multiplexer "git.randomchars.net/Reviath/RemiliaScarlet/Multiplexer"
	"git.randomchars.net/Reviath/RemiliaScarlet/interactions"
	"git.randomchars.net/Reviath/RemiliaScarlet/sql"
	"github.com/bwmarrin/discordgo"
)

func KissCommand(session *discordgo.Session, interaction interactions.Interaction) interactions.InteractionResponse {
	if sql.CheckLanguage(interaction.GuildID) == "tr" {
		if sql.IsBlocked(interaction.GuildID, "kiss") == "true" {
			return multiplexer.CreateResponse("Bu komut bu sunucuda engellenmiş.")
		}

		u, err := session.User(multiplexer.GetUser(interaction.Data.Options[0].Value.(string)))
		if err == nil {
			embed := embedutil.NewEmbed().
				SetColor(0xff1000).
				SetDescription(fmt.Sprintf("<@%s>, <@%s> isimli kişiyi öptü 😘", interaction.Member.User.ID, u.ID)).
				SetImage("https://media.tenor.com/images/d68747a5865b12c465e5dff31c65d5c2/tenor.gif").MessageEmbed
			return multiplexer.CreateEmbedResponse(embed)
		}
		return multiplexer.CreateResponse("Bir üye belirtmelisin")
	}

	if sql.IsBlocked(interaction.GuildID, "kiss") == "true" {
		return multiplexer.CreateResponse("This command is blocked on this guild.")
	}

	u, err := session.User(multiplexer.GetUser(interaction.Data.Options[0].Value.(string)))
	if err == nil {
		embed := embedutil.NewEmbed().
			SetColor(0xff1000).
			SetDescription(fmt.Sprintf("<@%s> kisses <@%s>", interaction.Member.User.ID, u.ID)).
			SetImage("https://media.tenor.com/images/d68747a5865b12c465e5dff31c65d5c2/tenor.gif").MessageEmbed
		return multiplexer.CreateEmbedResponse(embed)
	}
	return multiplexer.CreateResponse("You need to specify the user.")
}
