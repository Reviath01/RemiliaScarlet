package commands

import (
	ctx "git.randomchars.net/Reviath/handlers/Context"
	"github.com/bwmarrin/discordgo"
    "strings"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)

type Ban struct {

}

func (b Ban) Execute(ctx ctx.Ctx, session *discordgo.Session) error {
    db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/remilia")

    if err != nil {
        panic(err.Error())
    }

    defer db.Close()

    type Tag struct {
        isblocked string `json:"isblocked"`
    }

    var tag Tag

    err = db.QueryRow("SELECT isblocked FROM disabledcommands WHERE commandname ='ban' AND guildid ='" + ctx.Guild().ID + "'").Scan(&tag.isblocked)

    if err == nil {
        if tag.isblocked == "True" {
            _, err = session.ChannelMessageSend(ctx.Channel().ID, "This command is blocked on this guild.")

            if err != nil {
                return nil
            }

            return err
        }
    }

    perms, err := session.State.UserChannelPermissions(ctx.Author().ID, ctx.Channel().ID)
	if err == nil && (int(perms)&discordgo.PermissionBanMembers == discordgo.PermissionBanMembers) == false {
        _, err := session.ChannelMessageSend(ctx.Channel().ID, "You need ban members permission to run this command.")
	return err
    }

    if len(strings.Join(ctx.Args()," ")) < 19 {
    u, err := session.User(strings.Join(ctx.Args()," "))
    if err == nil {
        err = session.GuildBanCreate(ctx.Guild().ID, u.ID, 0)
        if err != nil {
            _, err = session.ChannelMessageSend(ctx.Channel().ID, "I do not have enough permission.")

            if err != nil {
                return nil
            }

            return err
        }
        _, err = session.ChannelMessageSend(ctx.Channel().ID, "Banned specified user.")
		
        if err != nil {
            return nil
        }

        return err
	} else {
        _, err = session.ChannelMessageSend(ctx.Channel().ID, "You need to specify the user.")
        
        if err != nil {
            return nil
        }

        return err
    }
} else {
    if len(strings.Join(ctx.Args()," ")) > 21 {
     u, err := session.User(strings.Join(ctx.Args()," ")[3:][:18])
        if err == nil {
            err = session.GuildBanCreate(ctx.Guild().ID, u.ID, 0)
            if err != nil {
                _, err = session.ChannelMessageSend(ctx.Channel().ID, "I do not have enough permission.")
                
                if err != nil {
                    return nil
                }

                return err
            }
            _, err = session.ChannelMessageSend(ctx.Channel().ID, "Banned specified user.")
		    
            if err != nil {
                return nil
            }

            return err
        } else {
            _, err = session.ChannelMessageSend(ctx.Channel().ID, "You need to specify the user.")

            if err != nil {
                return nil
            }

           return err
        }
    } else {
            _, err := session.ChannelMessageSend(ctx.Channel().ID, "You need to specify the user.")
            
            if err != nil {
                return nil
            }

            return err
        }
    }
}
