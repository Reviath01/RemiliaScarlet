package commands

import (
	ctx "git.randomchars.net/Reviath/handlers/Context"
	"github.com/bwmarrin/discordgo"
    "strings"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)

type Kick struct {

}

func (k Kick) Execute(ctx ctx.Ctx, session *discordgo.Session) error {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/remilia")

    if err != nil {
        panic(err.Error())
    }

    defer db.Close()

    type Tag struct {
        isblocked string `json:"isblocked"`
    }

    var tag Tag

    err = db.QueryRow("SELECT isblocked FROM disabledcommands WHERE commandname ='kick' AND guildid ='" + ctx.Guild().ID + "'").Scan(&tag.isblocked)

    if err == nil {
        if tag.isblocked == "True" {
            _, err = session.ChannelMessageSend(ctx.Channel().ID, "This command is blocked on this guild.")
            return err
        }
    }

    perms, err := session.State.UserChannelPermissions(ctx.Author().ID, ctx.Channel().ID)
	if err == nil && (int(perms)&discordgo.PermissionKickMembers == discordgo.PermissionKickMembers) == false {
        _, err := session.ChannelMessageSend(ctx.Channel().ID, "You need kick members permission to run this command.")
	return err
    }
    if len(strings.Join(ctx.Args()," ")) < 19 {
    u, err := session.User(strings.Join(ctx.Args()," "))
    if err == nil {
        err = session.GuildMemberDelete(ctx.Guild().ID, u.ID)
        if err != nil {
            _, err = session.ChannelMessageSend(ctx.Channel().ID, "I do not have enough permission.")
            return err
        }
        _, err = session.ChannelMessageSend(ctx.Channel().ID, "Kicked specified user.")
		return err
	} else {
        _, err = session.ChannelMessageSend(ctx.Channel().ID, "You need to specify the user.")
        return err
    }
} else {
    if len(strings.Join(ctx.Args()," ")) > 21 {
     u, err := session.User(strings.Join(ctx.Args()," ")[3:][:18])
        if err == nil {
            err = session.GuildMemberDelete(ctx.Guild().ID, u.ID)
            if err != nil {
                _, err = session.ChannelMessageSend(ctx.Channel().ID, "I do not have enough permission.")
                return err
            }
            _, err = session.ChannelMessageSend(ctx.Channel().ID, "Kicked specified user.")
		    return err
        } else {
            _, err = session.ChannelMessageSend(ctx.Channel().ID, "You need to specify the user.")
           return err
        }
    } else {
            _, err := session.ChannelMessageSend(ctx.Channel().ID, "You need to specify the user.")
            return err
        }
    }
}
