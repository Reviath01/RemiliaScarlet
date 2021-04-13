package commands

import (
	ctx "git.randomchars.net/Reviath/RemiliaScarlet/CommandHandler/Context"
	embedutil "git.randomchars.net/Reviath/RemiliaScarlet/EmbedUtil"
	"git.randomchars.net/Reviath/RemiliaScarlet/sql"
	"github.com/bwmarrin/discordgo"
)

type Roles struct {
}

var roles string

func (r Roles) Execute(ctx ctx.Ctx, session *discordgo.Session) error {
	db := sql.Connect()

	type Tag struct {
		isblocked string
		lang      string
	}

	var tag Tag

	err := db.QueryRow("SELECT language FROM languages WHERE guildid ='" + ctx.Guild().ID + "'").Scan(&tag.lang)
	if err == nil && tag.lang == "tr" {
		err = db.QueryRow("SELECT isblocked FROM disabledcommands WHERE commandname ='roles' AND guildid ='" + ctx.Guild().ID + "'").Scan(&tag.isblocked)

		if err == nil && tag.isblocked == "True" {
			_, _ = session.ChannelMessageSend(ctx.Channel().ID, "Bu komut bu sunucuda engellenmiş.")
			return nil
		}

		for _, i := range ctx.Guild().Roles {
			roles += "<@&" + i.ID + ">" + ", \n"
		}

		embed := embedutil.NewEmbed().
			AddField("Roller:", roles).MessageEmbed

		_, _ = session.ChannelMessageSendEmbed(ctx.Channel().ID, embed)

		return nil
	}

	err = db.QueryRow("SELECT isblocked FROM disabledcommands WHERE commandname ='roles' AND guildid ='" + ctx.Guild().ID + "'").Scan(&tag.isblocked)

	if err == nil && tag.isblocked == "True" {
		_, _ = session.ChannelMessageSend(ctx.Channel().ID, "This command is blocked on this guild.")
		return nil
	}

	for _, i := range ctx.Guild().Roles {
		roles += "<@&" + i.ID + ">" + ", \n"
	}

	embed := embedutil.NewEmbed().
		AddField("Roles:", roles).MessageEmbed

	_, _ = session.ChannelMessageSendEmbed(ctx.Channel().ID, embed)

	return nil
}
