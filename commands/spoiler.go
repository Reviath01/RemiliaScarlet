package commands

import (
	"strings"

	ctx "git.randomchars.net/Reviath/RemiliaScarlet/CommandHandler/Context"
	embedutil "git.randomchars.net/Reviath/RemiliaScarlet/EmbedUtil"
	"git.randomchars.net/Reviath/RemiliaScarlet/sql"
	"github.com/bwmarrin/discordgo"
)

type Spoiler struct {
}

func (s Spoiler) Execute(ctx ctx.Ctx, session *discordgo.Session) error {
	db := sql.Connect()

	type Tag struct {
		isblocked string
	}

	var tag Tag

	if sql.CheckLanguage(ctx.Guild().ID) == "tr" {
		err := db.QueryRow("SELECT isblocked FROM disabledcommands WHERE commandname ='spoiler' AND guildid ='" + ctx.Guild().ID + "'").Scan(&tag.isblocked)

		if err == nil && tag.isblocked == "True" {
			_, _ = session.ChannelMessageSend(ctx.Channel().ID, "Bu komut bu sunucuda engellenmiş.")
			return nil
		}

		if strings.Join(ctx.Args(), " ") == "" {
			_, _ = session.ChannelMessageSend(ctx.Channel().ID, "Spoiler olarak gönderilecek mesajı belirtmelisin.")

			return nil
		}
		spoilerembed := embedutil.NewEmbed().
			SetColor(0xe9ff00).
			SetDescription("|| " + strings.Join(ctx.Args(), " ") + " ||").MessageEmbed
		_, _ = session.ChannelMessageSendEmbed(ctx.Channel().ID, spoilerembed)

		return nil

	}

	err := db.QueryRow("SELECT isblocked FROM disabledcommands WHERE commandname ='spoiler' AND guildid ='" + ctx.Guild().ID + "'").Scan(&tag.isblocked)

	if err == nil && tag.isblocked == "True" {
		_, _ = session.ChannelMessageSend(ctx.Channel().ID, "This command is blocked on this guild.")
		return nil
	}

	if strings.Join(ctx.Args(), " ") == "" {
		_, _ = session.ChannelMessageSend(ctx.Channel().ID, "You need to specify the message.")

		return nil
	}
	spoilerembed := embedutil.NewEmbed().
		SetColor(0xe9ff00).
		SetDescription("|| " + strings.Join(ctx.Args(), " ") + " ||").MessageEmbed
	_, _ = session.ChannelMessageSendEmbed(ctx.Channel().ID, spoilerembed)

	return nil
}
