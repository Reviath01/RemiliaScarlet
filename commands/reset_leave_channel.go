package commands

import (
	ctx "git.randomchars.net/Reviath/RemiliaScarlet/Context"
	"github.com/bwmarrin/discordgo"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type ResetLeaveChannel struct {

}

func (r ResetLeaveChannel) Execute(ctx ctx.Ctx, session *discordgo.Session) error {
    db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/remilia")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	type Tag struct {
		channelid string `json:"channelid"`
        lang string `json:"language"`
    }

	var tag Tag

    err = db.QueryRow("SELECT language FROM languages WHERE guildid ='" + ctx.Guild().ID + "'").Scan(&tag.lang)
    if err == nil {
        if tag.lang == "tr" {
            perms, err := session.State.UserChannelPermissions(ctx.Author().ID, ctx.Channel().ID)
            if err == nil && (int(perms)&discordgo.PermissionAdministrator == discordgo.PermissionAdministrator) == false {
                _, err := session.ChannelMessageSend(ctx.Channel().ID, "Bu komutu kullanmak için yönetici yetkisine sahip olmalısın.")
            
                if err != nil {
                    return nil
                }
        
                return err
            }
        
            err = db.QueryRow("SELECT channelid FROM leavechannel WHERE guildid ='" + ctx.Guild().ID + "'").Scan(&tag.channelid)
            if err == nil {
                delete, err := db.Query("DELETE FROM leavechannel WHERE guildid ='" + ctx.Guild().ID + "'")
                if err != nil {
                    _, err = session.ChannelMessageSend(ctx.Channel().ID, "Bir hata oluştu!")
                    
                    if err != nil {
                        return nil
                    }
        
                    return err
                }
        
                defer delete.Close()
        
                _, err = session.ChannelMessageSend(ctx.Channel().ID, "Başarıyla çıkış kanalı sıfırlandı.")
                
                if err != nil {
                    return nil
                }
        
                return err
            } else {
                _, err = session.ChannelMessageSend(ctx.Channel().ID, "Çıkış kanalı ayarlanmamış, yani sıfırlayamazsın.")
                
                if err != nil {
                    return nil
                }
        
                return err
            }
        
        }
    }

    perms, err := session.State.UserChannelPermissions(ctx.Author().ID, ctx.Channel().ID)
	if err == nil && (int(perms)&discordgo.PermissionAdministrator == discordgo.PermissionAdministrator) == false {
        _, err := session.ChannelMessageSend(ctx.Channel().ID, "You need administrator permission to run this command.")
	
        if err != nil {
            return nil
        }

        return err
    }

	err = db.QueryRow("SELECT channelid FROM leavechannel WHERE guildid ='" + ctx.Guild().ID + "'").Scan(&tag.channelid)
    if err == nil {
        delete, err := db.Query("DELETE FROM leavechannel WHERE guildid ='" + ctx.Guild().ID + "'")
        if err != nil {
            _, err = session.ChannelMessageSend(ctx.Channel().ID, "An error occured, please try again.")
            
            if err != nil {
                return nil
            }

            return err
        }

        defer delete.Close()

        _, err = session.ChannelMessageSend(ctx.Channel().ID, "Successfully reset leave channel.")
        
        if err != nil {
            return nil
        }

        return err
    } else {
        _, err = session.ChannelMessageSend(ctx.Channel().ID, "Leave channel is not existing, so you can't reset.")
        
        if err != nil {
            return nil
        }

        return err
    }
}
