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
        lang string `json:"language"`
    }

    var tag Tag

    err = db.QueryRow("SELECT language FROM languages WHERE guildid ='" + ctx.Guild().ID + "'").Scan(&tag.lang)
    if err == nil {
        if tag.lang == "tr" {
            err = db.QueryRow("SELECT isblocked FROM disabledcommands WHERE commandname ='kick' AND guildid ='" + ctx.Guild().ID + "'").Scan(&tag.isblocked)

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
            if err == nil && (int(perms)&discordgo.PermissionKickMembers == discordgo.PermissionKickMembers) == false {
                _, err := session.ChannelMessageSend(ctx.Channel().ID, "Bu komutu kullanmak için üyeleri at yetkisine sahip olmalısın.")
            
                if err != nil {
                    return nil
                }
            
                return err
            }
        
            var args string 
            if len(strings.Join(ctx.Args()," ")) < 1 {
                _, err = session.ChannelMessageSend(ctx.Channel().ID, "Bir üye belirtmelisin.")
                if err != nil {
                    return nil
                }
                return err
            }
            args = ctx.Args()[0]
        
            if len(args) < 19 {
            u, err := session.User(args)
            if err == nil {
                err = session.GuildMemberDelete(ctx.Guild().ID, u.ID)
                if err != nil {
                    _, err = session.ChannelMessageSend(ctx.Channel().ID, "Yeterli yetkim yok.")
                    
                    if err != nil {
                        return nil
                    }
                    
                    return err
                }
                _, err = session.ChannelMessageSend(ctx.Channel().ID, "Belirtilen kişi sunucudan atıldı.")
                
                if err != nil {
                    return nil
                }
                
                return err
            } else {
                _, err = session.ChannelMessageSend(ctx.Channel().ID, "Bir üye belirtmelisin.")
                
                if err != nil {
                    return nil
                }
                
                return err
            }
        } else {
            if len(args) > 21 {
             u, err := session.User(args[3:][:18])
                if err == nil {
                    err = session.GuildMemberDelete(ctx.Guild().ID, u.ID)
                    if err != nil {
                        _, err = session.ChannelMessageSend(ctx.Channel().ID, "Yeterli yetkim yok.")
                        
                        if err != nil {
                            return nil
                        }
                        
                        return err
                    }
                    _, err = session.ChannelMessageSend(ctx.Channel().ID, "Belirtilen kişi sunucudan atıldı!")
                    
                    if err != nil {
                        return nil
                    }
                    
                    return err
                } else {
                    _, err = session.ChannelMessageSend(ctx.Channel().ID, "Bir üye belirtmelisin.")
                   
                    if err != nil {
                        return nil
                    }
                   
                    return err
                }
            } else {
                    _, err := session.ChannelMessageSend(ctx.Channel().ID, "Bir üye belirtmelisin.")
                    
                    if err != nil {
                        return nil
                    }
                    
                    return err
                }
            }
        }
    }

    err = db.QueryRow("SELECT isblocked FROM disabledcommands WHERE commandname ='kick' AND guildid ='" + ctx.Guild().ID + "'").Scan(&tag.isblocked)

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
	if err == nil && (int(perms)&discordgo.PermissionKickMembers == discordgo.PermissionKickMembers) == false {
        _, err := session.ChannelMessageSend(ctx.Channel().ID, "You need kick members permission to run this command.")
	
        if err != nil {
            return nil
        }
    
        return err
    }

    var args string 
    if len(strings.Join(ctx.Args()," ")) < 1 {
        _, err = session.ChannelMessageSend(ctx.Channel().ID, "You need to specify the user.")
        if err != nil {
            return nil
        }
        return err
    }
    args = ctx.Args()[0]

    if len(args) < 19 {
    u, err := session.User(args)
    if err == nil {
        err = session.GuildMemberDelete(ctx.Guild().ID, u.ID)
        if err != nil {
            _, err = session.ChannelMessageSend(ctx.Channel().ID, "I do not have enough permission.")
            
            if err != nil {
                return nil
            }
            
            return err
        }
        _, err = session.ChannelMessageSend(ctx.Channel().ID, "Kicked specified user.")
		
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
    if len(args) > 21 {
     u, err := session.User(args[3:][:18])
        if err == nil {
            err = session.GuildMemberDelete(ctx.Guild().ID, u.ID)
            if err != nil {
                _, err = session.ChannelMessageSend(ctx.Channel().ID, "I do not have enough permission.")
                
                if err != nil {
                    return nil
                }
                
                return err
            }
            _, err = session.ChannelMessageSend(ctx.Channel().ID, "Kicked specified user.")
		    
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
