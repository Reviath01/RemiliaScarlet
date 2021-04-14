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
	if sql.CheckLanguage(ctx.Guild().ID) == "tr" {
		if sql.IsBlocked(ctx.Guild().ID, "icon") == "true" {
			_, _ = session.ChannelMessageSend(ctx.Channel().ID, "Bu komut bu sunucuda engellenmiş.")
			return nil
		}

		iconeembed := embedutil.NewEmbed().
			SetColor(0x00f6ff).
			SetImage(ctx.Guild().IconURL()).MessageEmbed
		_, _ = session.ChannelMessageSendEmbed(ctx.Channel().ID, iconeembed)

		return nil
	}

	if sql.IsBlocked(ctx.Guild().ID, "icon") == "true" {
		_, _ = session.ChannelMessageSend(ctx.Channel().ID, "This command is blocked on this guild.")
		return nil
	}

	iconeembed := embedutil.NewEmbed().
		SetColor(0x00f6ff).
		SetImage(ctx.Guild().IconURL()).MessageEmbed
	_, _ = session.ChannelMessageSendEmbed(ctx.Channel().ID, iconeembed)

	return nil
}
