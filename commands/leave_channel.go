package commands

import (
	ctx "git.randomchars.net/Reviath/handlers/Context"
	"github.com/bwmarrin/discordgo"
    "strings"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type LeaveChannel struct {

}

func (l LeaveChannel) Execute(ctx ctx.Ctx, session *discordgo.Session) error {
    perms, err := session.State.UserChannelPermissions(ctx.Author().ID, ctx.Channel().ID)
	if err == nil && (int(perms)&discordgo.PermissionAdministrator == discordgo.PermissionAdministrator) == false {
        _, err := session.ChannelMessageSend(ctx.Channel().ID, "You need administrator permission to run this command.")
	
        if err != nil {
            return nil
        }
    
        return err
    }

    var args string
    if len(strings.Join(ctx.Args(), " ")) < 1 {
        _, err := session.ChannelMessageSend(ctx.Channel().ID, "You need to specify the channel.")
	    
        if err != nil {
            return nil
        }
        
        return err
    }
    args = ctx.Args()[0]

    if len(args) < 19 {
    c, err := session.Channel(args)
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

	err = db.QueryRow("SELECT channelid FROM leavechannel WHERE guildid ='" + ctx.Guild().ID + "'").Scan(&tag.channelid)
    if err == nil {
    _, err = session.ChannelMessageSend(ctx.Channel().ID, "Leave channel is already existing (to reset, use reset_leave_channel command).")
    
    if err != nil {
        return nil
    }
    
    return err
    } else {
        insert, err := db.Query("INSERT INTO leavechannel (channelid, guildid) VALUES ('" + c.ID + "', '" + ctx.Guild().ID + "')")
        if err != nil {
            _, err = session.ChannelMessageSend(ctx.Channel().ID, "An error occured, please try again.")
            
            if err != nil {
                return nil
            }
            
            return err
        }
        defer insert.Close()

        _, err = session.ChannelMessageSend(ctx.Channel().ID, "Leave channel set successfully.")
		
        if err != nil {
            return nil
        }
        
        return err
	}
        } else {
            _, err = session.ChannelMessageSend(ctx.Channel().ID, "You need to specify the channel.")
            
            if err != nil {
                return nil
            }
            
            return err
        }
    } else {
        if len(args) > 20 {
        c, err := session.Channel(args[2:][:18])
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

	err = db.QueryRow("SELECT channelid FROM leavechannel WHERE guildid ='" + ctx.Guild().ID + "'").Scan(&tag.channelid)
	if err == nil {
    _, err = session.ChannelMessageSend(ctx.Channel().ID, "Leave channel is already existing (to reset, use reset_leave_channel command).")
    
    if err != nil {
        return nil
    }
    
    return err
    } else {
        insert, err := db.Query("INSERT INTO leavechannel (channelid, guildid) VALUES ('" + c.ID + "', '" + ctx.Guild().ID + "')")
        if err != nil {
            _, err = session.ChannelMessageSend(ctx.Channel().ID, "An error occured, please try again.")
            
            if err != nil {
                return nil
            }
            
            return err
        }
        defer insert.Close()

        _, err = session.ChannelMessageSend(ctx.Channel().ID, "Leave channel set successfully.")
		
        if err != nil {
            return nil
        }
        
        return err
	    }
        } else {
            _, err = session.ChannelMessageSend(ctx.Channel().ID, "You need to specify the channel.")
           
            if err != nil {
                return nil
            }
           
            return err
        }
    } else {
            _, err := session.ChannelMessageSend(ctx.Channel().ID, "You need to specify the channel.")
            
            if err != nil {
                return nil
            }
            
            return err
        }
    }
}
