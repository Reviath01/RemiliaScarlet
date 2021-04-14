package commands

import (
	"strings"

	ctx "git.randomchars.net/Reviath/RemiliaScarlet/CommandHandler/Context"
	multiplexer "git.randomchars.net/Reviath/RemiliaScarlet/Multiplexer"
	"git.randomchars.net/Reviath/RemiliaScarlet/sql"
	"github.com/bwmarrin/discordgo"
)

type Kick struct {
}

func (k Kick) Execute(ctx ctx.Ctx, session *discordgo.Session) error {
	if sql.CheckLanguage(ctx.Guild().ID) == "tr" {
		if sql.IsBlocked(ctx.Guild().ID, "kick") == "true" {
			_, _ = session.ChannelMessageSend(ctx.Channel().ID, "This command is blocked on this guild.")
			return nil
		}

		perms, err := session.State.UserChannelPermissions(ctx.Author().ID, ctx.Channel().ID)
		if err == nil && !(int(perms)&discordgo.PermissionAdministrator == discordgo.PermissionAdministrator) {
			_, _ = session.ChannelMessageSend(ctx.Channel().ID, "Bu komutu kullanmak için üyeleri at yetkisine sahip olmalısın.")
			return nil
		}

		if len(strings.Join(ctx.Args(), " ")) < 1 {
			_, _ = session.ChannelMessageSend(ctx.Channel().ID, "Bir üye belirtmelisin.")

			return nil
		}
		u, err := session.User(multiplexer.GetUser(ctx.Args()[0]))
		if err == nil {
			err = session.GuildMemberDelete(ctx.Guild().ID, u.ID)
			if err != nil {
				_, _ = session.ChannelMessageSend(ctx.Channel().ID, "Yeterli yetkim yok.")

				return nil
			}
			_, _ = session.ChannelMessageSend(ctx.Channel().ID, "Belirtilen kişi sunucudan atıldı.")

			return nil
		} else {
			_, _ = session.ChannelMessageSend(ctx.Channel().ID, "Bir üye belirtmelisin.")

			return nil
		}
	}

	if sql.IsBlocked(ctx.Guild().ID, "kick") == "true" {
		_, _ = session.ChannelMessageSend(ctx.Channel().ID, "This command is blocked on this guild.")
		return nil
	}

	perms, err := session.State.UserChannelPermissions(ctx.Author().ID, ctx.Channel().ID)
	if err == nil && !(int(perms)&discordgo.PermissionAdministrator == discordgo.PermissionAdministrator) {
		_, _ = session.ChannelMessageSend(ctx.Channel().ID, "You need kick members permission to run this command.")

		return nil
	}

	if len(strings.Join(ctx.Args(), " ")) < 1 {
		_, _ = session.ChannelMessageSend(ctx.Channel().ID, "You need to specify the user.")

		return nil
	}

	u, err := session.User(multiplexer.GetUser(ctx.Args()[0]))
	if err == nil {
		err = session.GuildMemberDelete(ctx.Guild().ID, u.ID)
		if err != nil {
			_, _ = session.ChannelMessageSend(ctx.Channel().ID, "I do not have enough permission.")

			return nil
		}
		_, _ = session.ChannelMessageSend(ctx.Channel().ID, "Kicked specified user.")

		return nil
	} else {
		_, _ = session.ChannelMessageSend(ctx.Channel().ID, "You need to specify the user.")

		return nil
	}
}
