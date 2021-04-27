package commands

import (
	embedutil "git.randomchars.net/FreeNitori/EmbedUtil"
	CommandHandler "git.randomchars.net/Reviath/RemiliaScarlet/handler"
	"git.randomchars.net/Reviath/RemiliaScarlet/sql"
)

func IssueCommand(ctx CommandHandler.Context, _ []string) error {
	if sql.CheckLanguage(ctx.Guild.ID) == "tr" {
		issueembed := embedutil.New("", "GitLab üzerinden bir issue oluşturmak için [buraya](https://git.randomchars.net/Reviath/RemiliaScarlet/-/issues/new) tıkla!")
		issueembed.Color = 0xffa935
		issueembed.AddField("Eğer GitLab kullanmayı bilmiyorsan,", "[Sunucumuza](https://discord.gg/FshmaUh9eV) gelip sorunu belirtebilirsin.", true)
		ctx.ReplyEmbed(issueembed)
		return nil
	}

	issueembed := embedutil.New("", "Click [here](https://git.randomchars.net/Reviath/RemiliaScarlet/-/issues/new) to create an issue on GitLab")
	issueembed.Color = 0xffa935
	issueembed.AddField("If you don't know how to use GitLab,", "You can come to our [guild](https://discord.gg/FshmaUh9eV) and specify the problem.", true)
	ctx.ReplyEmbed(issueembed)
	return nil
}
