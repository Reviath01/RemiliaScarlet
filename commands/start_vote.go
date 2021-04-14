package commands

import (
	"strings"

	ctx "git.randomchars.net/Reviath/RemiliaScarlet/CommandHandler/Context"
	embedutil "git.randomchars.net/Reviath/RemiliaScarlet/EmbedUtil"
	"git.randomchars.net/Reviath/RemiliaScarlet/sql"
	"github.com/bwmarrin/discordgo"
)

type StartVote struct {
}

func (s StartVote) Execute(ctx ctx.Ctx, session *discordgo.Session) error {
	if sql.CheckLanguage(ctx.Guild().ID) == "tr" {
		if sql.IsBlocked(ctx.Guild().ID, "start_vote") == "true" {
			_, _ = session.ChannelMessageSend(ctx.Channel().ID, "Bu komut bu sunucuda engellenmiş.")
			return nil
		}

		perms, err := session.State.UserChannelPermissions(ctx.Author().ID, ctx.Channel().ID)
		if err == nil && !(int(perms)&discordgo.PermissionAdministrator == discordgo.PermissionAdministrator) {
			_, _ = session.ChannelMessageSend(ctx.Channel().ID, "Bu komutu kullanmak için mesajları yönet yetkisine sahip olmalısın.")

			return nil
		}
		if strings.Join(ctx.Args(), " ") == "" {
			_, _ = session.ChannelMessageSend(ctx.Channel().ID, "Oylama başlatmak için bir mesaj belirtmelisin.")

			return nil
		}
		embed := embedutil.NewEmbed().
			SetTitle("Oylama Başladı!").
			SetColor(0xff9100).
			AddField("Oylama Sorusu:", strings.Join(ctx.Args(), " ")).MessageEmbed
		msg, err := session.ChannelMessageSendEmbed(ctx.Channel().ID, embed)
		if err != nil {
			return nil
		}
		session.MessageReactionAdd(ctx.Channel().ID, msg.ID, "👍")
		session.MessageReactionAdd(ctx.Channel().ID, msg.ID, "👎")

		return nil
	}

	if sql.IsBlocked(ctx.Guild().ID, "start_vote") == "true" {
		_, _ = session.ChannelMessageSend(ctx.Channel().ID, "This command is blocked on this guild.")
		return nil
	}

	perms, err := session.State.UserChannelPermissions(ctx.Author().ID, ctx.Channel().ID)
	if err == nil && !(int(perms)&discordgo.PermissionAdministrator == discordgo.PermissionAdministrator) {
		_, _ = session.ChannelMessageSend(ctx.Channel().ID, "You need manage messages permission to run this command.")

		return nil
	}
	if strings.Join(ctx.Args(), " ") == "" {
		_, _ = session.ChannelMessageSend(ctx.Channel().ID, "You need to specify a message.")

		return nil
	}
	embed := embedutil.NewEmbed().
		SetTitle("Vote started!").
		SetColor(0xff9100).
		AddField("Vote question:", strings.Join(ctx.Args(), " ")).MessageEmbed
	msg, err := session.ChannelMessageSendEmbed(ctx.Channel().ID, embed)

	if err != nil {
		return nil
	}

	session.MessageReactionAdd(ctx.Channel().ID, msg.ID, "👍")
	session.MessageReactionAdd(ctx.Channel().ID, msg.ID, "👎")
	return nil
}
