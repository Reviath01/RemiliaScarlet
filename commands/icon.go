package commands

import (
	ctx "git.randomchars.net/Reviath/RemiliaScarlet/CommandHandler/Context"
	embedutil "git.randomchars.net/Reviath/RemiliaScarlet/EmbedUtil"
	"git.randomchars.net/Reviath/RemiliaScarlet/sql"
	"github.com/bwmarrin/discordgo"
)

type Icon struct {
}

func (i Icon) Execute(ctx ctx.Ctx, session *discordgo.Session) error {
	db := sql.Connect()

	type Tag struct {
		isblocked string
	}

	var tag Tag

	if sql.CheckLanguage(ctx.Guild().ID) == "tr" {
		err := db.QueryRow("SELECT isblocked FROM disabledcommands WHERE commandname ='icon' AND guildid ='" + ctx.Guild().ID + "'").Scan(&tag.isblocked)

		if err == nil && tag.isblocked == "True" {
			_, _ = session.ChannelMessageSend(ctx.Channel().ID, "Bu komut bu sunucuda engellenmiş.")
			return nil
		}

		iconeembed := embedutil.NewEmbed().
			SetColor(0x00f6ff).
			SetImage(ctx.Guild().IconURL()).MessageEmbed
		_, _ = session.ChannelMessageSendEmbed(ctx.Channel().ID, iconeembed)

		return nil
	}

	err := db.QueryRow("SELECT isblocked FROM disabledcommands WHERE commandname ='icon' AND guildid ='" + ctx.Guild().ID + "'").Scan(&tag.isblocked)

	if err == nil && tag.isblocked == "True" {
		_, _ = session.ChannelMessageSend(ctx.Channel().ID, "This command is blocked on this guild.")
		return nil
	}

	iconeembed := embedutil.NewEmbed().
		SetColor(0x00f6ff).
		SetImage(ctx.Guild().IconURL()).MessageEmbed
	_, _ = session.ChannelMessageSendEmbed(ctx.Channel().ID, iconeembed)

	return nil
}
