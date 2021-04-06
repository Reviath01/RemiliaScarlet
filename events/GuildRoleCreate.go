package events

import (
	"database/sql"

	embedutil "git.randomchars.net/Reviath/embed-util"
	"github.com/bwmarrin/discordgo"
	_ "github.com/go-sql-driver/mysql"
	"strconv"
)

func GuildRoleCreate(s *discordgo.Session, event *discordgo.GuildRoleCreate) {
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

	err = db.QueryRow("SELECT language FROM languages WHERE guildid ='" + event.GuildID + "'").Scan(&tag.lang)

	if err == nil {
		if tag.lang == "tr" {
			err = db.QueryRow("SELECT channelid FROM log WHERE guildid ='" + event.GuildID + "'").Scan(&tag.channelid)
			if err != nil {
				return
			} else {
				embed := embedutil.NewEmbed().
					SetTitle("Rol Oluşturuldu!").
					AddField("Rolün İsmi:", event.Role.Name + " ( <@&" + event.Role.ID + "> )").
					AddField("Rol İD'si:", event.Role.ID).
					AddField("Rol Rengi:", strconv.Itoa(event.Role.Color)).
					SetColor(0xefff00).MessageEmbed
		
				_, err = s.ChannelMessageSendEmbed(tag.channelid, embed)
				if err != nil {
					return
				}
			}
			return
		}
		return
	}

	err = db.QueryRow("SELECT channelid FROM log WHERE guildid ='" + event.GuildID + "'").Scan(&tag.channelid)
	if err != nil {
		return
	} else {
		embed := embedutil.NewEmbed().
			SetTitle("Role Created!").
			AddField("Role Name:", event.Role.Name + " ( <@&" + event.Role.ID + "> )").
			AddField("Role ID:", event.Role.ID).
			AddField("Role Color:", strconv.Itoa(event.Role.Color)).
			SetColor(0xefff00).MessageEmbed

		_, err = s.ChannelMessageSendEmbed(tag.channelid, embed)
		if err != nil {
			return
		}
	}
}
