package commands

import (
	"database/sql"

	ctx "git.randomchars.net/Reviath/handlers/Context"
	"github.com/bwmarrin/discordgo"
	_ "github.com/go-sql-driver/mysql"
)

type ResetAutorole struct {

}

func (r ResetAutorole) Execute(ctx ctx.Ctx, session *discordgo.Session) error {
    perms, err := session.State.UserChannelPermissions(ctx.Author().ID, ctx.Channel().ID)
	if err == nil && (int(perms)&discordgo.PermissionAdministrator == discordgo.PermissionAdministrator) == false {
        _, err := session.ChannelMessageSend(ctx.Channel().ID, "You need administrator permission to run this command.")
	
        if err != nil {
            return nil
        }

        return err
    }

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
        delete, err := db.Query("DELETE FROM autorole WHERE guildid ='" + ctx.Guild().ID + "'")
        if err != nil {
            _, err = session.ChannelMessageSend(ctx.Channel().ID, "An error occured, please try again.")
            
            if err != nil {
                return nil
            }

            return err
        }

        defer delete.Close()

        _, err = session.ChannelMessageSend(ctx.Channel().ID, "Successfully reset auto role.")
        
        if err != nil {
            return nil
        }

        return err
    } else {
        _, err = session.ChannelMessageSend(ctx.Channel().ID, "Auto role is not existing, so you can't reset.")
        
        if err != nil {
            return nil
        }

        return err
    }
}
