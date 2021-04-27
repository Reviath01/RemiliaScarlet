package events

import (
	"strconv"

	embedutil "git.randomchars.net/Reviath/RemiliaScarlet/EmbedUtil"
	"git.randomchars.net/Reviath/RemiliaScarlet/sql"
	"github.com/bwmarrin/discordgo"
)

func ChannelDelete(s *discordgo.Session, event *discordgo.ChannelDelete) {
	db := sql.Connect()
	defer db.Close()

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
		}
		embed := embedutil.NewEmbed().
			SetTitle("Kanal Silindi!").
			AddField("Kanalın İsmi:", event.Channel.Name).
			AddField("Kanalın İD'si:", event.Channel.ID).
			AddField("Kanal Tipi:", channeltype).
			SetColor(0xff1000).MessageEmbed

		_, _ = s.ChannelMessageSendEmbed(tag.channelid, embed)
		if err != nil {
			return
		}
		return
	}
	err := db.QueryRow("SELECT channelid FROM log WHERE guildid ='" + event.GuildID + "'").Scan(&tag.channelid)
	if err != nil {
		return
	}
	embed := embedutil.NewEmbed().
		SetTitle("Channel Deleted!").
		AddField("Channel Name:", event.Channel.Name).
		AddField("Channel ID:", event.Channel.ID).
		AddField("Channel Type:", channeltype).
		SetColor(0xff1000).MessageEmbed

	_, _ = s.ChannelMessageSendEmbed(tag.channelid, embed)
	if err != nil {
		return
	}
}
