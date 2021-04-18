package commands

import (
	"strings"

	embedutil "git.randomchars.net/Reviath/RemiliaScarlet/EmbedUtil"
	multiplexer "git.randomchars.net/Reviath/RemiliaScarlet/Multiplexer"
	CommandHandler "git.randomchars.net/Reviath/RemiliaScarlet/handler"
	"git.randomchars.net/Reviath/RemiliaScarlet/sql"
)

func StartVoteCommand(ctx CommandHandler.Context, _ []string) error {
	if sql.CheckLanguage(ctx.Guild.ID) == "tr" {
		if !multiplexer.CheckManageMessagesPermission(ctx.Session, ctx.Message.Author.ID, ctx.Channel.ID) {
			ctx.Reply("Yeterli yetkiye sahip değilsin.")
			return nil
		}
		if sql.IsBlocked(ctx.Guild.ID, "start_vote") == "true" {
			ctx.Reply("Bu komut bu sunucuda engellenmiş.")
			return nil
		}

		if strings.Join(multiplexer.GetArgs(ctx.Message.Content, multiplexer.GetPrefix()), " ") == "" {
			ctx.Reply("Oylama başlatmak için bir mesaj belirtmelisin.")
			return nil
		}
		embed := embedutil.NewEmbed().
			SetTitle("Oylama Başladı!").
			SetColor(0xff9100).
			AddField("Oylama Sorusu:", strings.Join(multiplexer.GetArgs(ctx.Message.Content, multiplexer.GetPrefix()), " ")).MessageEmbed
		msg, err := ctx.ReplyEmbed(embed)
		if err != nil {
			return nil
		}
		ctx.Session.MessageReactionAdd(ctx.Channel.ID, msg.ID, "👍")
		ctx.Session.MessageReactionAdd(ctx.Channel.ID, msg.ID, "👎")

		return nil
	}

	if !multiplexer.CheckManageMessagesPermission(ctx.Session, ctx.Message.Author.ID, ctx.Channel.ID) {
		ctx.Reply("You don't have enough permission.")
		return nil
	}

	if sql.IsBlocked(ctx.Guild.ID, "start_vote") == "true" {
		ctx.Reply("This command is blocked on this guild.")
		return nil
	}

	if strings.Join(multiplexer.GetArgs(ctx.Message.Content, multiplexer.GetPrefix()), " ") == "" {
		ctx.Reply("You need to specify a message.")
		return nil
	}
	embed := embedutil.NewEmbed().
		SetTitle("Vote started!").
		SetColor(0xff9100).
		AddField("Vote question:", strings.Join(multiplexer.GetArgs(ctx.Message.Content, multiplexer.GetPrefix()), " ")).MessageEmbed
	msg, err := ctx.ReplyEmbed(embed)

	if err != nil {
		return nil
	}

	ctx.Session.MessageReactionAdd(ctx.Channel.ID, msg.ID, "👍")
	ctx.Session.MessageReactionAdd(ctx.Channel.ID, msg.ID, "👎")
	return nil
}
