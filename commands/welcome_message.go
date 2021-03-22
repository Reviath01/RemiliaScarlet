package commands

import (
	ctx "git.randomchars.net/Reviath/handlers/Context"
	"github.com/bwmarrin/discordgo"
    "strings"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type WelcomeMessage struct {

}

func (w WelcomeMessage) Execute(ctx ctx.Ctx, session *discordgo.Session) error {
    perms, err := session.State.UserChannelPermissions(ctx.Author().ID, ctx.Channel().ID)
	if err == nil && (int(perms)&discordgo.PermissionAdministrator == discordgo.PermissionAdministrator) == false {
        _, err := session.ChannelMessageSend(ctx.Channel().ID, "You need administrator permission to run this command.")
	return err
    }
	if strings.Join(ctx.Args(), " ") == "" {
		_, err = session.ChannelMessageSend(ctx.Channel().ID, "You need to specify the message.")
		return err
	}

	if len(strings.Join(ctx.Args(), " ")) > 254 {
		_, err = session.ChannelMessageSend(ctx.Channel().ID, "Your message can be up to 255 characters long.")
		return err
	}

    db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/remilia")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	type Tag struct {
		message string `json:"message"`
	}

	var tag Tag

	err = db.QueryRow("SELECT message FROM welcomemessage WHERE guildid ='" + ctx.Guild().ID + "'").Scan(&tag.message)
	if err == nil {
    _, err = session.ChannelMessageSend(ctx.Channel().ID, "Welcome message is already existing (to reset, use reset_welcome_message command).")
    return err
    } else {
        insert, err := db.Query("INSERT INTO welcomemessage (message, guildid) VALUES ('" + strings.Join(ctx.Args(), " ") + "', '" + ctx.Guild().ID + "')")
        if err != nil {
            _, err = session.ChannelMessageSend(ctx.Channel().ID, "An error occured, please try again.")
            return err
        }
        defer insert.Close()

        _, err = session.ChannelMessageSend(ctx.Channel().ID, "Welcome message set successfully.")
		return err
	    }
	}