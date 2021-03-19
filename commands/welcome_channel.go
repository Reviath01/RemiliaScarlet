package commands

import (
	ctx "git.randomchars.net/Reviath/handlers/Context"
	"github.com/bwmarrin/discordgo"
    "strings"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type WelcomeChannel struct {

}

func (w WelcomeChannel) Execute(ctx ctx.Ctx, session *discordgo.Session) error {
    perms, err := session.State.UserChannelPermissions(ctx.Author().ID, ctx.Channel().ID)
	if err == nil && (int(perms)&discordgo.PermissionAdministrator == discordgo.PermissionAdministrator) == false {
        _, err := session.ChannelMessageSend(ctx.Channel().ID, "You need administrator permission to run this command.")
	return err
    }

    if len(strings.Join(ctx.Args()," ")) < 19 {
    c, err := session.Channel(strings.Join(ctx.Args()," "))
    if err == nil {
    db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/remilia")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	type Tag struct {
		channelid string `json:"channelid"`
	}

	var tag Tag

	err = db.QueryRow("SELECT channelid FROM welcomechannel WHERE guildid ='" + ctx.Guild().ID + "'").Scan(&tag.channelid)
	if err == nil {
    _, err = session.ChannelMessageSend(ctx.Channel().ID, "Welcome channel is already existing (to reset, use reset_welcome_channel command).")
    return err
    } else {
        insert, err := db.Query("INSERT INTO welcomechannel (channelid, guildid) VALUES ('" + c.ID + "', '" + ctx.Guild().ID + "')")
        if err != nil {
            _, err = session.ChannelMessageSend(ctx.Channel().ID, "An error occured, please try again.")
            return err
        }
        defer insert.Close()

        _, err = session.ChannelMessageSend(ctx.Channel().ID, "Welcome channel set successfully.")
		return err
	}
        } else {
            _, err = session.ChannelMessageSend(ctx.Channel().ID, "You need to specify the channel.")
            return err
        }
    } else {
        if len(strings.Join(ctx.Args()," ")) > 20 {
        c, err := session.Channel(strings.Join(ctx.Args()," ")[2:][:18])
        if err == nil {
    db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/remilia")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	type Tag struct {
		channelid string `json:"channelid"`
	}

	var tag Tag

	err = db.QueryRow("SELECT channelid FROM welcomechannel WHERE guildid ='" + ctx.Guild().ID + "'").Scan(&tag.channelid)
	if err == nil {
    _, err = session.ChannelMessageSend(ctx.Channel().ID, "Welcome channel is already existing (to reset, use reset_welcome_channel command).")
    return err
    } else {
        insert, err := db.Query("INSERT INTO welcomechannel (channelid, guildid) VALUES ('" + c.ID + "', '" + ctx.Guild().ID + "')")
        if err != nil {
            _, err = session.ChannelMessageSend(ctx.Channel().ID, "An error occured, please try again.")
            return err
        }
        defer insert.Close()

        _, err = session.ChannelMessageSend(ctx.Channel().ID, "Welcome channel set successfully.")
		return err
	    }
        } else {
            _, err = session.ChannelMessageSend(ctx.Channel().ID, "You need to specify the channel.")
           return err
        }
    } else {
            _, err := session.ChannelMessageSend(ctx.Channel().ID, "You need to specify the channel.")
            return err
        }
    }
}
