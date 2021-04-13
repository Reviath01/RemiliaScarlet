package commands

import (
	"strings"

	ctx "git.randomchars.net/Reviath/RemiliaScarlet/CommandHandler/Context"
	multiplexer "git.randomchars.net/Reviath/RemiliaScarlet/Multiplexer"
	"git.randomchars.net/Reviath/RemiliaScarlet/sql"
	"github.com/bwmarrin/discordgo"
)

type Ban struct {
}

func (b Ban) Execute(ctx ctx.Ctx, session *discordgo.Session) error {
	db := sql.Connect()

	type Tag struct {
		isblocked string
		lang      string
	}

	var tag Tag

	err := db.QueryRow("SELECT language FROM languages WHERE guildid ='" + ctx.Guild().ID + "'").Scan(&tag.lang)
	if err == nil && tag.lang == "tr" {
		err = db.QueryRow("SELECT isblocked FROM disabledcommands WHERE commandname ='ban' AND guildid ='" + ctx.Guild().ID + "'").Scan(&tag.isblocked)

		if err == nil && tag.isblocked == "True" {
			_, _ = session.ChannelMessageSend(ctx.Channel().ID, "Bu komut bu sunucuda engellenmiş.")
			return nil
		}

		perms, err := session.State.UserChannelPermissions(ctx.Author().ID, ctx.Channel().ID)
		if err == nil && !(int(perms)&discordgo.PermissionAdministrator == discordgo.PermissionAdministrator) {
			_, _ = session.ChannelMessageSend(ctx.Channel().ID, "Bu komutu kullanmak için üyeleri yasakla yetkisine sahip olmalısın.")
			return nil
		}

		if strings.Join(ctx.Args(), " ") == "" {
			_, _ = session.ChannelMessageSend(ctx.Channel().ID, "Bir üye belirtmelisiniz.")
			return nil
		} else {
			u, err := session.User(multiplexer.GetUser(ctx.Args()[0]))
			if err != nil {
				_, _ = session.ChannelMessageSend(ctx.Channel().ID, "Bir üye belirtmelisiniz.")
			} else {
				err = session.GuildBanCreate(ctx.Guild().ID, u.ID, 0)
				if err != nil {
					_, _ = session.ChannelMessageSend(ctx.Channel().ID, "Yeterli yetkiye sahip değilim.")
				}
			}
		}
	}

	err = db.QueryRow("SELECT isblocked FROM disabledcommands WHERE commandname ='ban' AND guildid ='" + ctx.Guild().ID + "'").Scan(&tag.isblocked)

	if err == nil && tag.isblocked == "True" {
		_, _ = session.ChannelMessageSend(ctx.Channel().ID, "This command is blocked on this guild.")
		return nil
	}

	perms, err := session.State.UserChannelPermissions(ctx.Author().ID, ctx.Channel().ID)
	if err == nil && !(int(perms)&discordgo.PermissionAdministrator == discordgo.PermissionAdministrator) {
		_, _ = session.ChannelMessageSend(ctx.Channel().ID, "You need ban members permission to run this command.")
		return nil
	}

	if strings.Join(ctx.Args(), " ") == "" {
		_, _ = session.ChannelMessageSend(ctx.Channel().ID, "You need to specify the user.")
		return nil
	} else {
		u, err := session.User(multiplexer.GetUser(ctx.Args()[0]))
		if err != nil {
			_, _ = session.ChannelMessageSend(ctx.Channel().ID, "You need to specify the user.")
		} else {
			err = session.GuildBanCreate(ctx.Guild().ID, u.ID, 0)
			if err != nil {
				_, _ = session.ChannelMessageSend(ctx.Channel().ID, "Cannot ban that user.")
			}
		}
	}
	return nil
}
