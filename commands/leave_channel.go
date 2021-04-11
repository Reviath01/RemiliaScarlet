package commands

import (
	"database/sql"
	"strings"

	ctx "git.randomchars.net/Reviath/RemiliaScarlet/CommandHandler/Context"
	"git.randomchars.net/Reviath/RemiliaScarlet/multiplexer"
	"github.com/bwmarrin/discordgo"
	_ "github.com/go-sql-driver/mysql"
)

type LeaveChannel struct {
}

func (l LeaveChannel) Execute(ctx ctx.Ctx, session *discordgo.Session) error {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/remilia")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	type Tag struct {
		channelid string
		lang      string
	}

	var tag Tag
	err = db.QueryRow("SELECT language FROM languages WHERE guildid ='" + ctx.Guild().ID + "'").Scan(&tag.lang)

	if err == nil && tag.lang == "tr" {
		perms, err := session.State.UserChannelPermissions(ctx.Author().ID, ctx.Channel().ID)
		if err == nil && !(int(perms)&discordgo.PermissionAdministrator == discordgo.PermissionAdministrator) {
			_, _ = session.ChannelMessageSend(ctx.Channel().ID, "Bu komutu kullanmak için yönetici yetkisine sahip olmalısın.")
			return nil
		}

		if len(strings.Join(ctx.Args(), " ")) < 1 {
			_, _ = session.ChannelMessageSend(ctx.Channel().ID, "Bir kanal belirtmelisin")

			return nil
		}
		c, err := session.Channel(multiplexer.GetChannel(ctx.Args()[0]))
		if err == nil {
			err = db.QueryRow("SELECT channelid FROM leavechannel WHERE guildid ='" + ctx.Guild().ID + "'").Scan(&tag.channelid)
			if err == nil {
				_, _ = session.ChannelMessageSend(ctx.Channel().ID, "Çıkış kanalı zaten ayarlanmış reset'lemek için reset_leave_channel komutunu kullanın.")
				return nil
			} else {
				insert, err := db.Query("INSERT INTO leavechannel (channelid, guildid) VALUES ('" + c.ID + "', '" + ctx.Guild().ID + "')")
				if err != nil {
					_, _ = session.ChannelMessageSend(ctx.Channel().ID, "Bir hata oluştu.")

					return nil
				}
				defer insert.Close()
				_, _ = session.ChannelMessageSend(ctx.Channel().ID, "Başarıyla ayarlandı.")
				return nil
			}
		} else {
			_, _ = session.ChannelMessageSend(ctx.Channel().ID, "Bir kanal belirtmelisin.")
			return nil
		}
	}

	perms, err := session.State.UserChannelPermissions(ctx.Author().ID, ctx.Channel().ID)
	if err == nil && !(int(perms)&discordgo.PermissionAdministrator == discordgo.PermissionAdministrator) {
		_, _ = session.ChannelMessageSend(ctx.Channel().ID, "You need administrator permission to run this command.")
		return nil
	}
	if len(strings.Join(ctx.Args(), " ")) < 1 {
		_, _ = session.ChannelMessageSend(ctx.Channel().ID, "You need to specify the channel.")

		return nil
	}

	c, err := session.Channel(multiplexer.GetChannel(ctx.Args()[0]))
	if err == nil {
		err = db.QueryRow("SELECT channelid FROM leavechannel WHERE guildid ='" + ctx.Guild().ID + "'").Scan(&tag.channelid)
		if err == nil {
			_, _ = session.ChannelMessageSend(ctx.Channel().ID, "Leave channel is already existing (to reset, use reset_leave_channel command).")

			return nil
		} else {
			insert, err := db.Query("INSERT INTO leavechannel (channelid, guildid) VALUES ('" + c.ID + "', '" + ctx.Guild().ID + "')")
			if err != nil {
				_, _ = session.ChannelMessageSend(ctx.Channel().ID, "An error occurred, please try again.")

				return nil
			}
			defer insert.Close()

			_, _ = session.ChannelMessageSend(ctx.Channel().ID, "Leave channel set successfully.")

			return nil
		}
	} else {
		_, _ = session.ChannelMessageSend(ctx.Channel().ID, "You need to specify the channel.")
		return nil
	}
}
