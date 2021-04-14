package events

import (
	"strconv"

	embedutil "git.randomchars.net/Reviath/RemiliaScarlet/EmbedUtil"
	"git.randomchars.net/Reviath/RemiliaScarlet/sql"
	"github.com/bwmarrin/discordgo"
)

func ChannelCreate(s *discordgo.Session, event *discordgo.ChannelCreate) {
	db := sql.Connect()

	type Tag struct {
		channelid string
	}

	var tag Tag

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

	if sql.CheckLanguage(event.GuildID) == "tr" {
		err := db.QueryRow("SELECT channelid FROM log WHERE guildid ='" + event.GuildID + "'").Scan(&tag.channelid)
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

	err := db.QueryRow("SELECT channelid FROM log WHERE guildid ='" + event.GuildID + "'").Scan(&tag.channelid)
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
