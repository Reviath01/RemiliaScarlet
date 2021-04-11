package events

import (
	"database/sql"
	"strconv"

	embedutil "git.randomchars.net/Reviath/RemiliaScarlet/EmbedUtil"
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
		channelid string
		lang      string
	}

	var tag Tag

	err = db.QueryRow("SELECT language FROM languages WHERE guildid ='" + event.GuildID + "'").Scan(&tag.lang)

	var channeltype string

	if strconv.Itoa(int(event.Channel.Type)) == "0" {
		channeltype = "Text"
	} else if strconv.Itoa(int(event.Channel.Type)) == "2" {
		channeltype = "Voice"
	} else if strconv.Itoa(int(event.Channel.Type)) == "6" {
		channeltype = "Store"
	} else if strconv.Itoa(int(event.Channel.Type)) == "13" {
		channeltype = "Stage"
	} else if strconv.Itoa(int(event.Channel.Type)) == "5" {
		channeltype = "News"
	} else {
		channeltype = "Unknown Type (Type ID: " + strconv.Itoa(int(event.Channel.Type))
	}

	if err == nil && tag.lang == "tr" {
		err = db.QueryRow("SELECT channelid FROM log WHERE guildid ='" + event.GuildID + "'").Scan(&tag.channelid)
		if err != nil {
			return
		} else {
			embed := embedutil.NewEmbed().
				SetTitle("Kanal Oluşturuldu!").
				AddField("Kanal İsmi:", event.Channel.Name+" ( <#"+event.Channel.ID+"> )").
				AddField("Kanalın İD'si:", event.Channel.ID).
				AddField("Kanal Tipi:", channeltype).
				SetColor(0xff1000).MessageEmbed

			_, _ = s.ChannelMessageSendEmbed(tag.channelid, embed)
			if err != nil {
				return
			}
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
			AddField("Channel Type:", channeltype).
			SetColor(0xff1000).MessageEmbed

		_, _ = s.ChannelMessageSendEmbed(tag.channelid, embed)
		if err != nil {
			return
		}
	}
}
