package commands

import (
	"strings"

	embedutil "git.randomchars.net/Reviath/RemiliaScarlet/EmbedUtil"
	multiplexer "git.randomchars.net/Reviath/RemiliaScarlet/Multiplexer"
	CommandHandler "git.randomchars.net/Reviath/RemiliaScarlet/handler"
	"git.randomchars.net/Reviath/RemiliaScarlet/sql"
)

func AvatarCommand(ctx CommandHandler.Context, _ []string) error {
	if sql.CheckLanguage(ctx.Guild.ID) == "tr" {
		if sql.IsBlocked(ctx.Guild.ID, "avatar") == "true" {
			ctx.Reply("Bu komut bu sunucuda engellenmiş.")
			return nil
		}

		if len(strings.Join(multiplexer.GetArgs(ctx.Message.Content, multiplexer.GetPrefix()), " ")) < 1 {
			avatarembed := embedutil.NewEmbed().
				SetColor(0xff1000).
				SetDescription(ctx.Message.Author.Username + "#" + ctx.Message.Author.Discriminator + " isimli kişinin profil fotoğrafı").
				SetImage(ctx.Message.Author.AvatarURL("1024")).MessageEmbed
			ctx.ReplyEmbed(avatarembed)

			return nil
		}

		u, err := ctx.Session.User(multiplexer.GetUser(multiplexer.GetArgs(ctx.Message.Content, multiplexer.GetPrefix())[0]))
		if err == nil {
			avatarembed := embedutil.NewEmbed().
				SetColor(0xff1000).
				SetDescription(u.Username + "#" + u.Discriminator + " isimli kişinin profil fotoğrafı").
				SetImage(u.AvatarURL("1024")).MessageEmbed
			ctx.ReplyEmbed(avatarembed)

			return nil
		} else {
			avatarembed := embedutil.NewEmbed().
				SetColor(0xff1000).
				SetDescription(ctx.Message.Author.Username + "#" + ctx.Message.Author.Discriminator + " isimli kişinin profil fotoğrafı").
				SetImage(ctx.Message.Author.AvatarURL("1024")).MessageEmbed
			ctx.ReplyEmbed(avatarembed)

			return nil
		}
	}

	if sql.IsBlocked(ctx.Guild.ID, "avatar") == "true" {
		ctx.Reply("This command is blocked on this guild.")
		return nil
	}

	if len(strings.Join(multiplexer.GetArgs(ctx.Message.Content, multiplexer.GetPrefix()), " ")) < 1 {
		avatarembed := embedutil.NewEmbed().
			SetColor(0xff1000).
			SetDescription("Avatar of " + ctx.Message.Author.Username + "#" + ctx.Message.Author.Discriminator).
			SetImage(ctx.Message.Author.AvatarURL("1024")).MessageEmbed
		ctx.ReplyEmbed(avatarembed)

		return nil
	}

	u, err := ctx.Session.User(multiplexer.GetUser(multiplexer.GetArgs(ctx.Message.Content, multiplexer.GetPrefix())[0]))
	if err == nil {
		avatarembed := embedutil.NewEmbed().
			SetColor(0xff1000).
			SetDescription("Avatar of " + u.Username + "#" + u.Discriminator).
			SetImage(u.AvatarURL("1024")).MessageEmbed
		ctx.ReplyEmbed(avatarembed)

		return nil
	} else {
		avatarembed := embedutil.NewEmbed().
			SetColor(0xff1000).
			SetDescription("Avatar of " + ctx.Message.Author.Username + "#" + ctx.Message.Author.Discriminator).
			SetImage(ctx.Message.Author.AvatarURL("1024")).MessageEmbed
		ctx.ReplyEmbed(avatarembed)
		return nil
	}
}
