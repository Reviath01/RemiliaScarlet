package commands

import (
	ctx "git.randomchars.net/Reviath/RemiliaScarlet/Context"
	"github.com/bwmarrin/discordgo"
    embedutil "git.randomchars.net/Reviath/embed-util"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type Settings struct {

}

func (s Settings) Execute(ctx ctx.Ctx, session *discordgo.Session) error {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/remilia")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	type Tag struct {
		welcomechannelid string `json:"channelid"`
		leavechannelid string `json:"channelid"`
		roleid string `json:"roleid"`
		welcomemessage string `json:"message"`
		leavemessage string `json:"message"`
		logid string `json:"channelid"`
		isblocked string `json:"isblocked"`
		lang string `json:"language"`
	}

	var tag Tag
	var welcomechannel string
	var leavechannel string
	var autorole string
	var leavemsg string
	var welcomemsg string
	var logchannel string

	err = db.QueryRow("SELECT language FROM languages WHERE guildid ='" + ctx.Guild().ID + "'").Scan(&tag.lang)
    if err == nil {
        if tag.lang == "tr" {
			err = db.QueryRow("SELECT isblocked FROM disabledcommands WHERE commandname ='settings' AND guildid ='" + ctx.Guild().ID + "'").Scan(&tag.isblocked)

			if err == nil {
				if tag.isblocked == "True" {
					_, err = session.ChannelMessageSend(ctx.Channel().ID, "Bu komut bu sunucuda engellenmiş.")
					
					if err != nil {
						return nil
					}
					
					return err
				}
			}
		
			err = db.QueryRow("SELECT channelid FROM welcomechannel WHERE guildid ='" + ctx.Guild().ID + "'").Scan(&tag.welcomechannelid)
			if err == nil {
				welcomechannel = "<#" + tag.welcomechannelid + "> (" + tag.welcomechannelid + ")"
			} else {
				welcomechannel = "Ayarlanmamış."
			}
		
			err = db.QueryRow("SELECT channelid FROM leavechannel WHERE guildid ='" + ctx.Guild().ID + "'").Scan(&tag.leavechannelid)
			if err == nil {
				leavechannel = "<#" + tag.leavechannelid + "> (" + tag.leavechannelid + ")"
			} else {
				leavechannel = "Ayarlanmamış."
			}
		
			err = db.QueryRow("SELECT roleid FROM autorole WHERE guildid ='" + ctx.Guild().ID + "'").Scan(&tag.roleid)
			if err == nil {
				autorole = "<@&" + tag.roleid + "> (" + tag.roleid + ")"
			} else {
				autorole = "Ayarlanmamış."
			}
		
			err = db.QueryRow("SELECT message FROM welcomemessage WHERE guildid ='" + ctx.Guild().ID + "'").Scan(&tag.welcomemessage)
			if err == nil {
				welcomemsg = tag.welcomemessage
			} else {
				welcomemsg = "Ayarlanmamış."
			}
		
			err = db.QueryRow("SELECT message FROM leavemessage WHERE guildid ='" + ctx.Guild().ID + "'").Scan(&tag.leavemessage)
			if err == nil {
				leavemsg = tag.leavemessage
			} else {
				leavemsg = "Ayarlanmamış."
			}
		
			err = db.QueryRow("SELECT channelid FROM log WHERE guildid ='" + ctx.Guild().ID + "'").Scan(&tag.logid)
			if err == nil {
				logchannel = "<#" + tag.logid + "> (" + tag.logid + ")"
			} else {
				logchannel = "Ayarlanmamış."
			}
		
			embed := embedutil.NewEmbed().
				SetTitle(ctx.Guild().Name + " Settings").
				AddField("Hoş Geldin Kanalı:", welcomechannel).
				AddField("Çıkış Kanalı:", leavechannel).
				AddField("Otorol:", autorole).
				AddField("Çıkış Mesajı:", leavemsg).
				AddField("Hoş Geldin Mesajı:", welcomemsg).
				AddField("Log Kanalı:", logchannel).
				SetColor(0x00f0ff).MessageEmbed
				
			_, err = session.ChannelMessageSendEmbed(ctx.Channel().ID, embed)
			
			if err != nil {
				return nil
			}
			
			return err
		
		}
	}

	err = db.QueryRow("SELECT isblocked FROM disabledcommands WHERE commandname ='settings' AND guildid ='" + ctx.Guild().ID + "'").Scan(&tag.isblocked)

    if err == nil {
        if tag.isblocked == "True" {
            _, err = session.ChannelMessageSend(ctx.Channel().ID, "This command is blocked on this guild.")
            
			if err != nil {
				return nil
			}
			
			return err
        }
    }

	err = db.QueryRow("SELECT channelid FROM welcomechannel WHERE guildid ='" + ctx.Guild().ID + "'").Scan(&tag.welcomechannelid)
	if err == nil {
		welcomechannel = "<#" + tag.welcomechannelid + "> (" + tag.welcomechannelid + ")"
	} else {
		welcomechannel = "Not existing."
	}

	err = db.QueryRow("SELECT channelid FROM leavechannel WHERE guildid ='" + ctx.Guild().ID + "'").Scan(&tag.leavechannelid)
	if err == nil {
		leavechannel = "<#" + tag.leavechannelid + "> (" + tag.leavechannelid + ")"
	} else {
		leavechannel = "Not existing."
	}

	err = db.QueryRow("SELECT roleid FROM autorole WHERE guildid ='" + ctx.Guild().ID + "'").Scan(&tag.roleid)
	if err == nil {
		autorole = "<@&" + tag.roleid + "> (" + tag.roleid + ")"
	} else {
		autorole = "Not existing."
	}

	err = db.QueryRow("SELECT message FROM welcomemessage WHERE guildid ='" + ctx.Guild().ID + "'").Scan(&tag.welcomemessage)
	if err == nil {
		welcomemsg = tag.welcomemessage
	} else {
		welcomemsg = "Not existing."
	}

	err = db.QueryRow("SELECT message FROM leavemessage WHERE guildid ='" + ctx.Guild().ID + "'").Scan(&tag.leavemessage)
	if err == nil {
		leavemsg = tag.leavemessage
	} else {
		leavemsg = "Not existing."
	}

	err = db.QueryRow("SELECT channelid FROM log WHERE guildid ='" + ctx.Guild().ID + "'").Scan(&tag.logid)
	if err == nil {
		logchannel = "<#" + tag.logid + "> (" + tag.logid + ")"
	} else {
		logchannel = "Not existing."
	}

	embed := embedutil.NewEmbed().
        SetTitle(ctx.Guild().Name + " Settings").
		AddField("Welcome Channel:", welcomechannel).
		AddField("Leave Channel:", leavechannel).
		AddField("Autorole:", autorole).
		AddField("Leave Message:", leavemsg).
		AddField("Welcome Message:", welcomemsg).
		AddField("Log Channel:", logchannel).
		SetColor(0x00f0ff).MessageEmbed
        
	_, err = session.ChannelMessageSendEmbed(ctx.Channel().ID, embed)
	
	if err != nil {
		return nil
	}
	
	return err
}
