package commands

import (
	"fmt"
	"strings"

	embedutil "git.randomchars.net/FreeNitori/EmbedUtil"
	multiplexer "git.randomchars.net/Reviath/RemiliaScarlet/Multiplexer"
	CommandHandler "git.randomchars.net/Reviath/RemiliaScarlet/handler"
	"git.randomchars.net/Reviath/RemiliaScarlet/sql"
)

func SlapCommand(ctx CommandHandler.Context, _ []string) error {
	if sql.CheckLanguage(ctx.Guild.ID) == "tr" {
		if sql.IsBlocked(ctx.Guild.ID, "slap") == "true" {
			ctx.Reply("Bu komut bu sunucuda engellenmiş.")
			return nil
		}

		if len(strings.Join(multiplexer.GetArgs(ctx.Message.Content, multiplexer.GetPrefix()), " ")) < 1 {
			ctx.Reply("Bir üye belirtmelisin.")
			return nil
		}

		u, err := ctx.Session.User(multiplexer.GetUser(multiplexer.GetArgs(ctx.Message.Content, multiplexer.GetPrefix())[0]))
		if err == nil {
			embed := embedutil.New("", fmt.Sprintf("<@%s>, <@%s> isimli kişiye vurdu", ctx.Message.Author.ID, u.ID))
			embed.Color = 0xff1000
			embed.SetImage("https://images-ext-1.discordapp.net/external/79sCWyD-TmmyjFxlaQIxAkAANAfV529d-LDHNkGDM0M/%3Fitemid%3D10426943/https/media1.tenor.com/images/b6d8a83eb652a30b95e87cf96a21e007/tenor.gif")
			ctx.ReplyEmbed(embed)
			return nil
		}
		ctx.Reply("Bir üye belirtmelisin.")
		return nil
	}

	if sql.IsBlocked(ctx.Guild.ID, "slap") == "true" {
		ctx.Reply("This command is blocked on this guild.")
		return nil
	}

	if len(strings.Join(multiplexer.GetArgs(ctx.Message.Content, multiplexer.GetPrefix()), " ")) < 1 {
		ctx.Reply("You need to specify the user.")
		return nil
	}

	u, err := ctx.Session.User(multiplexer.GetUser(multiplexer.GetArgs(ctx.Message.Content, multiplexer.GetPrefix())[0]))
	if err == nil {
		embed := embedutil.New("", fmt.Sprintf("<@%s> slaps <@%s>", ctx.Message.Author.ID, u.ID))
		embed.Color = 0xff1000
		embed.SetImage("https://images-ext-1.discordapp.net/external/79sCWyD-TmmyjFxlaQIxAkAANAfV529d-LDHNkGDM0M/%3Fitemid%3D10426943/https/media1.tenor.com/images/b6d8a83eb652a30b95e87cf96a21e007/tenor.gif")
		ctx.ReplyEmbed(embed)
		return nil
	}
	ctx.Reply("You need to specify the user.")
	return nil
}
