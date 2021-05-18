package commands

import (
	"fmt"

	embedutil "git.randomchars.net/FreeNitori/EmbedUtil"
	"git.randomchars.net/Reviath/RemiliaScarlet/config"
	CommandHandler "git.randomchars.net/Reviath/RemiliaScarlet/handler"
	"git.randomchars.net/Reviath/RemiliaScarlet/sql"
)

// IssueCommand is a handler for issue command
func IssueCommand(ctx CommandHandler.Context, _ []string) error {
	u, _ := ctx.Session.User(config.Owner)
	switch sql.CheckLanguage(ctx.Guild.ID) {
	case "tr":
		issueembed := embedutil.New("", "GitLab üzerinden bir issue oluşturmak için [buraya](https://git.randomchars.net/Reviath/RemiliaScarlet/-/issues/new) tıkla!")
		issueembed.Color = 0xffa935
		issueembed.AddField("Eğer GitLab kullanmayı bilmiyorsan,", fmt.Sprintf("Sahibime DM yoluyla ulaşabilirsiniz [%s#%s](https://discord.com/users/%s)", u.Username, u.Discriminator, u.ID), true)
		ctx.ReplyEmbed(issueembed)
		return nil

	default:

		issueembed := embedutil.New("", "Click [here](https://git.randomchars.net/Reviath/RemiliaScarlet/-/issues/new) to create an issue on GitLab")
		issueembed.Color = 0xffa935
		issueembed.AddField("If you don't know how to use GitLab,", fmt.Sprintf("You can write from DM to my owner [%s#%s](https://discord.com/users/%s)", u.Username, u.Discriminator, u.ID), true)
		ctx.ReplyEmbed(issueembed)
		return nil
	}
}
