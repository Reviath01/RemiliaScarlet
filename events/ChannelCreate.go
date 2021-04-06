package events

import (
	"database/sql"

	embedutil "git.randomchars.net/Reviath/embed-util"
	"github.com/bwmarrin/discordgo"
	_ "github.com/go-sql-driver/mysql"
)

func ChannelCreate(s *discordgo.Session, event *discordgo.ChannelCreate) {
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
			SetTitle("Kanal Oluşturuldu!").
			AddField("Kanal İsmi:", event.Channel.Name+" ( <#"+event.Channel.ID+"> )").
			AddField("Kanalın İD'si:", event.Channel.ID).
			AddField("Kanal Tipi:", string(rune(event.Channel.Type))).
			SetColor(0xff1000).MessageEmbed

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
			SetTitle("Channel Created!").
			AddField("Channel Name:", event.Channel.Name+" ( <#"+event.Channel.ID+"> )").
			AddField("Channel ID:", event.Channel.ID).
			AddField("Channel Type:", string(event.Channel.Type)).
			SetColor(0xff1000).MessageEmbed

		_, err = s.ChannelMessageSendEmbed(tag.channelid, embed)
		if err != nil {
			return
		}
	}
}
