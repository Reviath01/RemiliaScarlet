package interactioncommands

import (
	embedutil "git.randomchars.net/FreeNitori/EmbedUtil"
	multiplexer "git.randomchars.net/Reviath/RemiliaScarlet/Multiplexer"
	"git.randomchars.net/Reviath/RemiliaScarlet/interactions"
	"git.randomchars.net/Reviath/RemiliaScarlet/sql"
	"github.com/bwmarrin/discordgo"
)

// IssueCommand is issue command for interactions
func IssueCommand(session *discordgo.Session, interaction interactions.Interaction) interactions.InteractionResponse {
	switch sql.CheckLanguage(interaction.GuildID) {
	case "tr":
		issueembed := embedutil.New("", "GitLab üzerinden bir issue oluşturmak için [buraya](https://git.randomchars.net/Reviath/RemiliaScarlet/-/issues/new) tıkla!")
		issueembed.Color = 0xffa935
		issueembed.AddField("Eğer GitLab kullanmayı bilmiyorsan,", "[Sunucumuza](https://discord.gg/FshmaUh9eV) gelip sorunu belirtebilirsin.", true)
		return multiplexer.CreateEmbedResponse(issueembed)
	default:
		issueembed := embedutil.New("", "Click [here](https://git.randomchars.net/Reviath/RemiliaScarlet/-/issues/new) to create an issue on GitLab")
		issueembed.Color = 0xffa935
		issueembed.AddField("If you don't know how to use GitLab,", "You can come to our [guild](https://discord.gg/FshmaUh9eV) and specify the problem.", true)
		return multiplexer.CreateEmbedResponse(issueembed)
	}
}
