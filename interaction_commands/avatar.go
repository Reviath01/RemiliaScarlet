package interaction_commands

import (
	"fmt"

	embedutil "git.randomchars.net/FreeNitori/EmbedUtil"
	multiplexer "git.randomchars.net/Reviath/RemiliaScarlet/Multiplexer"
	"git.randomchars.net/Reviath/RemiliaScarlet/interactions"
	"git.randomchars.net/Reviath/RemiliaScarlet/sql"
	"github.com/bwmarrin/discordgo"
)

//Avatar slash command.
func AvatarCommand(session *discordgo.Session, interaction interactions.Interaction) interactions.InteractionResponse {
	switch sql.CheckLanguage(interaction.GuildID) {
	case "tr":
		if sql.IsBlocked(interaction.GuildID, "avatar") == "true" {
			return multiplexer.CreateResponse("Bu komut bu sunucuda engellenmiş.")
		}
		u, err := session.User(multiplexer.GetUser(interaction.Data.Options[0].Value.(string)))
		if err == nil {
			avatarembed := embedutil.New("", fmt.Sprintf("%s#%s isimli kişinin profil fotoğrafı", u.Username, u.Discriminator))
			avatarembed.Color = 0xff1000
			avatarembed.SetImage(u.AvatarURL("1024"))
			return multiplexer.CreateEmbedResponse(avatarembed)
		}
		avatarembed := embedutil.New("", fmt.Sprintf("%s#%s isimli kişinin profil fotoğrafı", interaction.Member.User.Username, interaction.Member.User.Discriminator))
		avatarembed.Color = 0xff1000
		avatarembed.SetImage(interaction.Member.User.AvatarURL("1024"))
		return multiplexer.CreateEmbedResponse(avatarembed)
	default:

		if sql.IsBlocked(interaction.GuildID, "avatar") == "true" {
			return multiplexer.CreateResponse("This command is blocked on this guild.")
		}

		u, err := session.User(multiplexer.GetUser(interaction.Data.Options[0].Value.(string)))
		if err == nil {
			avatarembed := embedutil.New("", fmt.Sprintf("Avatar of %s#%s", u.Username, u.Discriminator))
			avatarembed.Color = 0xff1000
			avatarembed.SetImage(u.AvatarURL("1024"))
			return multiplexer.CreateEmbedResponse(avatarembed)
		}
		avatarembed := embedutil.New("", fmt.Sprintf("Avatar of %s#%s", interaction.Member.User.Username, interaction.Member.User.Discriminator))
		avatarembed.Color = 0xff1000
		avatarembed.SetImage(interaction.Member.User.AvatarURL("1024"))
		return multiplexer.CreateEmbedResponse(avatarembed)
	}
}
