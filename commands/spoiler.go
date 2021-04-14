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
	if sql.CheckLanguage(ctx.Guild().ID) == "tr" {
		if sql.IsBlocked(ctx.Guild().ID, "spoiler") == "true" {
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

	if sql.IsBlocked(ctx.Guild().ID, "spoiler") == "true" {
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
