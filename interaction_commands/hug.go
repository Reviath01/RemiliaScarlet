package interaction_commands

import (
	"fmt"

	embedutil "git.randomchars.net/FreeNitori/EmbedUtil"
	multiplexer "git.randomchars.net/Reviath/RemiliaScarlet/Multiplexer"
	"git.randomchars.net/Reviath/RemiliaScarlet/interactions"
	"git.randomchars.net/Reviath/RemiliaScarlet/sql"
	"github.com/bwmarrin/discordgo"
)

func HugCommand(session *discordgo.Session, interaction interactions.Interaction) interactions.InteractionResponse {
	switch sql.CheckLanguage(interaction.GuildID) {
	case "tr":
		if sql.IsBlocked(interaction.GuildID, "hug") == "true" {
			return multiplexer.CreateResponse("Bu komut bu sunucuda engellenmiş")
		}

		u, err := session.User(multiplexer.GetUser(interaction.Data.Options[0].Value.(string)))
		if err == nil {
			if u.ID == interaction.Member.User.ID {
				return multiplexer.CreateResponse("Kendine sarılamazsın")
			}
			embed := embedutil.New("", fmt.Sprintf("<@%s>, <@%s> isimli kişiye sarıldı 💖", interaction.Member.User.ID, u.ID))
			embed.Color = 0xff1000
			embed.SetImage("https://i.pinimg.com/originals/4d/d7/49/4dd749423de10a319b5d9e8850bbace4.gif")
			return multiplexer.CreateEmbedResponse(embed)
		}
		return multiplexer.CreateResponse("Bir kişiyi etiketlemelisin.")
	default:

		if sql.IsBlocked(interaction.GuildID, "hug") == "true" {
			return multiplexer.CreateResponse("This command is blocked on this guild.")
		}

		u, err := session.User(multiplexer.GetUser(interaction.Data.Options[0].Value.(string)))
		if err == nil {
			if u.ID == interaction.Member.User.ID {
				return multiplexer.CreateResponse("You can't hug yourself.")
			}
			embed := embedutil.New("", fmt.Sprintf("<@%s> hugs <@%s>", interaction.Member.User.ID, u.ID))
			embed.Color = 0xff1000
			embed.SetImage("https://i.pinimg.com/originals/4d/d7/49/4dd749423de10a319b5d9e8850bbace4.gif")
			return multiplexer.CreateEmbedResponse(embed)
		}
		return multiplexer.CreateResponse("You need to specify the user.")
	}
}
