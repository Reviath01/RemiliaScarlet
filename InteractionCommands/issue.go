package interactioncommands

import (
	"fmt"

	embedutil "git.randomchars.net/FreeNitori/EmbedUtil"
	multiplexer "git.randomchars.net/Reviath/RemiliaScarlet/Multiplexer"
	"git.randomchars.net/Reviath/RemiliaScarlet/config"
	"git.randomchars.net/Reviath/RemiliaScarlet/interactions"
	"git.randomchars.net/Reviath/RemiliaScarlet/sql"
	"github.com/bwmarrin/discordgo"
)

// IssueCommand is issue command for interactions
func IssueCommand(session *discordgo.Session, interaction interactions.Interaction) interactions.InteractionResponse {
	config.ReadConfig()
	u, _ := session.User(config.Owner)
	switch sql.CheckLanguage(interaction.GuildID) {
	case "tr":
		issueembed := embedutil.New("", "GitLab üzerinden bir issue oluşturmak için [buraya](https://git.randomchars.net/Reviath/RemiliaScarlet/-/issues/new) tıkla!")
		issueembed.Color = 0xffa935
		issueembed.AddField("Eğer GitLab kullanmayı bilmiyorsan,", fmt.Sprintf("Sahibime DM yoluyla ulaşabilirsiniz [%s#%s](https://discord.com/users/%s)", u.Username, u.Discriminator, u.ID), true)
		return multiplexer.CreateEmbedResponse(issueembed)
	default:
		issueembed := embedutil.New("", "Click [here](https://git.randomchars.net/Reviath/RemiliaScarlet/-/issues/new) to create an issue on GitLab")
		issueembed.Color = 0xffa935
		issueembed.AddField("If you don't know how to use GitLab,", fmt.Sprintf("You can write from DM to my owner [%s#%s](https://discord.com/users/%s)", u.Username, u.Discriminator, u.ID), true)
		return multiplexer.CreateEmbedResponse(issueembed)
	}
}
