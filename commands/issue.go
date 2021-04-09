package commands

import (
	"database/sql"

	ctx "git.randomchars.net/Reviath/RemiliaScarlet/Context"
	embedutil "git.randomchars.net/Reviath/embed-util"
	"github.com/bwmarrin/discordgo"
	_ "github.com/go-sql-driver/mysql"
)

type Issue struct {
}

func (i Issue) Execute(ctx ctx.Ctx, session *discordgo.Session) error {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/remilia")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	type Tag struct {
		isblocked string `json:"isblocked"`
		lang      string `json:"language"`
	}

	var tag Tag

	err = db.QueryRow("SELECT language FROM languages WHERE guildid ='" + ctx.Guild().ID + "'").Scan(&tag.lang)
	if err == nil && tag.lang == "tr" {
		issueembed := embedutil.NewEmbed().
			SetColor(0xffa935).
			SetDescription("GitLab üzerinden bir issue oluşturmak için [buraya](https://git.randomchars.net/Reviath/RemiliaScarlet/-/issues/new) tıkla!").
			AddField("Eğer GitLab kullanmayı bilmiyorsan,", "[Sunucumuza](https://discord.gg/xqsTvtM2hk) gelip sorunu belirtebilirsin.").MessageEmbed
		_, err = session.ChannelMessageSendEmbed(ctx.Channel().ID, issueembed)

		if err != nil {
			return nil
		}

		return err
	}
	issueembed := embedutil.NewEmbed().
		SetColor(0xffa935).
		SetDescription("Click [here](https://git.randomchars.net/Reviath/RemiliaScarlet/-/issues/new) to create an issue on GitLab").
		AddField("If you don't know how to use GitLab,", "You can come to our [guild](https://discord.gg/xqsTvtM2hk) and specify the problem.").MessageEmbed
	_, err = session.ChannelMessageSendEmbed(ctx.Channel().ID, issueembed)

	if err != nil {
		return nil
	}

	return err
}
