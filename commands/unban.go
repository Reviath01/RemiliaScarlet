package commands

import (
	"database/sql"
	"strings"

	ctx "git.randomchars.net/Reviath/RemiliaScarlet/CommandHandler/Context"
	"git.randomchars.net/Reviath/RemiliaScarlet/multiplexer"
	"github.com/bwmarrin/discordgo"
	_ "github.com/go-sql-driver/mysql"
)

type Unban struct {
}

func (q Unban) Execute(ctx ctx.Ctx, session *discordgo.Session) error {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/remilia")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	type Tag struct {
		isblocked string
		lang      string
	}

	var tag Tag

	err = db.QueryRow("SELECT language FROM languages WHERE guildid ='" + ctx.Guild().ID + "'").Scan(&tag.lang)
	if err == nil && tag.lang == "tr" {
		err = db.QueryRow("SELECT isblocked FROM disabledcommands WHERE commandname ='unban' AND guildid ='" + ctx.Guild().ID + "'").Scan(&tag.isblocked)

		if err == nil && tag.isblocked == "True" {
			_, _ = session.ChannelMessageSend(ctx.Channel().ID, "Bu komut bu sunucuda engellenmiş.")
			return nil
		}

		perms, err := session.State.UserChannelPermissions(ctx.Author().ID, ctx.Channel().ID)
		if err == nil && !(int(perms)&discordgo.PermissionAdministrator == discordgo.PermissionAdministrator) {
			_, _ = session.ChannelMessageSend(ctx.Channel().ID, "Bu komutu kullanabilmek için üyeleri yasakla yetkisine sahip olmalısın.")

			return nil
		}

		if len(strings.Join(ctx.Args(), " ")) < 1 {
			_, _ = session.ChannelMessageSend(ctx.Channel().ID, "Bir üye belirtmelisin.")

			return nil
		}

		u, err := session.User(multiplexer.GetUser(ctx.Args()[0]))
		if err == nil {
			err = session.GuildBanDelete(ctx.Guild().ID, u.ID)
			if err != nil {
				_, _ = session.ChannelMessageSend(ctx.Channel().ID, "Bu üye banlanmamış veya yeterli yetkiye sahip değilim.")
				return nil
			}
			_, _ = session.ChannelMessageSend(ctx.Channel().ID, "Belirtilen kişinin banı kaldırıldı.")
			return nil
		} else {
			_, _ = session.ChannelMessageSend(ctx.Channel().ID, "Bir üye belirtmelisin.")
			return nil
		}
	}

	err = db.QueryRow("SELECT isblocked FROM disabledcommands WHERE commandname ='unban' AND guildid ='" + ctx.Guild().ID + "'").Scan(&tag.isblocked)

	if err == nil && tag.isblocked == "True" {
		_, _ = session.ChannelMessageSend(ctx.Channel().ID, "This command is blocked on this guild.")
		return nil
	}

	perms, err := session.State.UserChannelPermissions(ctx.Author().ID, ctx.Channel().ID)
	if err == nil && !(int(perms)&discordgo.PermissionAdministrator == discordgo.PermissionAdministrator) {
		_, _ = session.ChannelMessageSend(ctx.Channel().ID, "You need ban members permission to run this command.")
		return nil
	}

	if len(strings.Join(ctx.Args(), " ")) < 1 {
		_, _ = session.ChannelMessageSend(ctx.Channel().ID, "You need to specify the user.")

		return nil
	}

	u, err := session.User(multiplexer.GetUser(ctx.Args()[0]))

	if err == nil {
		err = session.GuildBanDelete(ctx.Guild().ID, u.ID)
		if err != nil {
			_, _ = session.ChannelMessageSend(ctx.Channel().ID, "This user is not banned or I don't have enough permission.")
			return nil
		}
		_, _ = session.ChannelMessageSend(ctx.Channel().ID, "Unbanned specified user.")

		return nil
	} else {
		_, _ = session.ChannelMessageSend(ctx.Channel().ID, "You need to specify the user.")

		return nil
	}
}
