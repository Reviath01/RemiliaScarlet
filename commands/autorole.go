package commands

import (
	"database/sql"
	ctx "git.randomchars.net/Reviath/handlers/Context"
	"github.com/bwmarrin/discordgo"
	_ "github.com/go-sql-driver/mysql"
)

type AutoRole struct {

}

func (a AutoRole) Execute(ctx ctx.Ctx, session *discordgo.Session) error {
    perms, err := session.State.UserChannelPermissions(ctx.Author().ID, ctx.Channel().ID)
	if err == nil && (int(perms)&discordgo.PermissionAdministrator == discordgo.PermissionAdministrator) == false {
        _, err := session.ChannelMessageSend(ctx.Channel().ID, "You need administrator permission to run this command.")
	return err
    }

    var args string

    args = ctx.Args()[0]

    if len(args) == 18 {
    db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/remilia")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	type Tag struct {
		roleid string `json:"roleid"`
	}

	var tag Tag

	err = db.QueryRow("SELECT roleid FROM autorole WHERE guildid ='" + ctx.Guild().ID + "'").Scan(&tag.roleid)
	if err == nil {
    _, err = session.ChannelMessageSend(ctx.Channel().ID, "Auto role is already existing (to reset, use reset_autorole command).")
    return err
    } else {
        insert, err := db.Query("INSERT INTO autorole (roleid, guildid) VALUES ('" + args + "', '" + ctx.Guild().ID + "')")
        if err != nil {
            _, err = session.ChannelMessageSend(ctx.Channel().ID, "An error occured, please try again.")
            
            if err != nil {
                return nil
            }

            return err
        }
        defer insert.Close()

        _, err = session.ChannelMessageSend(ctx.Channel().ID, "Auto role set successfully.")

        if err != nil {
            return nil
        }

		return err
	}
} else {
        if len(args) == 22 {
    db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/remilia")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	type Tag struct {
		roleid string `json:"roleid"`
	}

	var tag Tag

	err = db.QueryRow("SELECT roleid FROM autorole WHERE guildid ='" + ctx.Guild().ID + "'").Scan(&tag.roleid)
	if err == nil {
        _, err = session.ChannelMessageSend(ctx.Channel().ID, "Auto role is already existing (to reset, use reset_autorole command).")
        if err != nil {
            return nil
        }
        return err
    } else {
        insert, err := db.Query("INSERT INTO autorole (roleid, guildid) VALUES ('" + args[3:][:18] + "', '" + ctx.Guild().ID + "')")
        if err != nil {
            _, err = session.ChannelMessageSend(ctx.Channel().ID, "An error occured, please try again.")
            
            if err != nil {
                return nil
            }
            
            return err
        }
        defer insert.Close()

        _, err = session.ChannelMessageSend(ctx.Channel().ID, "Auto role set successfully.")
		
        if err != nil {
            return nil
        }

        return err
	    }
        } else {
            _, err = session.ChannelMessageSend(ctx.Channel().ID, "You need to specify the role.")
           
            if err != nil {
                return nil
            }

           return err
        }
	}
}