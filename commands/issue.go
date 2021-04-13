package commands

import (
	ctx "git.randomchars.net/Reviath/RemiliaScarlet/CommandHandler/Context"
	embedutil "git.randomchars.net/Reviath/RemiliaScarlet/EmbedUtil"
	"git.randomchars.net/Reviath/RemiliaScarlet/sql"
	"github.com/bwmarrin/discordgo"
)

type Issue struct {
}

func (i Issue) Execute(ctx ctx.Ctx, session *discordgo.Session) error {
	db := sql.Connect()

	type Tag struct {
		lang string
	}

	var tag Tag

	err := db.QueryRow("SELECT language FROM languages WHERE guildid ='" + ctx.Guild().ID + "'").Scan(&tag.lang)
	if err == nil && tag.lang == "tr" {
		issueembed := embedutil.NewEmbed().
			SetColor(0xffa935).
			SetDescription("GitLab üzerinden bir issue oluşturmak için [buraya](https://git.randomchars.net/Reviath/RemiliaScarlet/-/issues/new) tıkla!").
			AddField("Eğer GitLab kullanmayı bilmiyorsan,", "[Sunucumuza](https://discord.gg/xqsTvtM2hk) gelip sorunu belirtebilirsin.").MessageEmbed
		_, _ = session.ChannelMessageSendEmbed(ctx.Channel().ID, issueembed)

		return nil
	}
	issueembed := embedutil.NewEmbed().
		SetColor(0xffa935).
		SetDescription("Click [here](https://git.randomchars.net/Reviath/RemiliaScarlet/-/issues/new) to create an issue on GitLab").
		AddField("If you don't know how to use GitLab,", "You can come to our [guild](https://discord.gg/xqsTvtM2hk) and specify the problem.").MessageEmbed
	_, _ = session.ChannelMessageSendEmbed(ctx.Channel().ID, issueembed)

	return nil
}
