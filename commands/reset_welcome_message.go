package commands

import (
	ctx "git.randomchars.net/Reviath/handlers/Context"
	"github.com/bwmarrin/discordgo"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type ResetWelcomeMessage struct {

}

func (r ResetWelcomeMessage) Execute(ctx ctx.Ctx, session *discordgo.Session) error {
    db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/remilia")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	type Tag struct {
		message string `json:"message"`
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
        
            err = db.QueryRow("SELECT message FROM welcomemessage WHERE guildid ='" + ctx.Guild().ID + "'").Scan(&tag.message)
            if err == nil {
                delete, err := db.Query("DELETE FROM welcomemessage WHERE guildid ='" + ctx.Guild().ID + "'")
                if err != nil {
                    _, err = session.ChannelMessageSend(ctx.Channel().ID, "Bir hata oluştu.")
                    
                    if err != nil {
                        return nil
                    }
                    
                    return err
                }
        
                defer delete.Close()
        
                _, err = session.ChannelMessageSend(ctx.Channel().ID, "Hoş geldin mesajı başarıyla sıfırlandı.")
                
                if err != nil {
                    return nil
                }
                
                return err
            } else {
                _, err = session.ChannelMessageSend(ctx.Channel().ID, "Hoş geldin mesajı ayarlanmamış, sıfırlayamazsın.")
                
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

	err = db.QueryRow("SELECT message FROM welcomemessage WHERE guildid ='" + ctx.Guild().ID + "'").Scan(&tag.message)
    if err == nil {
        delete, err := db.Query("DELETE FROM welcomemessage WHERE guildid ='" + ctx.Guild().ID + "'")
        if err != nil {
            _, err = session.ChannelMessageSend(ctx.Channel().ID, "An error occured, please try again.")
            
            if err != nil {
                return nil
            }
            
            return err
        }

        defer delete.Close()

        _, err = session.ChannelMessageSend(ctx.Channel().ID, "Successfully reset welcome message.")
        
        if err != nil {
            return nil
        }
        
        return err
    } else {
        _, err = session.ChannelMessageSend(ctx.Channel().ID, "Welcome message is not existing, so you can't reset.")
        
        if err != nil {
            return nil
        }
        
        return err
    }
}
