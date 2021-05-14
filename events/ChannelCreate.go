package events

import (
	"fmt"
	"strconv"

	embedutil "git.randomchars.net/FreeNitori/EmbedUtil"
	"git.randomchars.net/Reviath/RemiliaScarlet/sql"
	"github.com/bwmarrin/discordgo"
)

// ChannelCreate is an event handler for channel create event
func ChannelCreate(s *discordgo.Session, event *discordgo.ChannelCreate) {
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
	} else if strconv.Itoa(int(event.Channel.Type)) == "4" {
		channeltype = "Category"
	} else {
		channeltype = fmt.Sprintf("Unknown Type (Type ID: %s)", strconv.Itoa(int(event.Channel.Type)))
	}

	if sql.CheckLanguage(event.GuildID) == "tr" {
		err := db.QueryRow("SELECT channelid FROM log WHERE guildid ='" + event.GuildID + "'").Scan(&tag.channelid)
		if err != nil {
			return
		}
		embed := embedutil.New("Kanal Oluşturuldu!", "")
		embed.AddField("Kanal İsmi:", event.Channel.Name+" ( <#"+event.Channel.ID+"> )", true)
		embed.AddField("Kanalın İD'si:", event.Channel.ID, true)
		embed.AddField("Kanal Tipi:", channeltype, true)
		embed.Color = 0xff1000

		_, _ = s.ChannelMessageSendEmbed(tag.channelid, embed.MessageEmbed)
		if err != nil {
			return
		}
		return
	}

	err := db.QueryRow("SELECT channelid FROM log WHERE guildid ='" + event.GuildID + "'").Scan(&tag.channelid)
	if err != nil {
		return
	}
	embed := embedutil.New("Channel Created!", "")
	embed.AddField("Channel Name:", event.Channel.Name+" ( <#"+event.Channel.ID+"> )", true)
	embed.AddField("Channel ID:", event.Channel.ID, true)
	embed.AddField("Channel Type:", channeltype, true)
	embed.Color = 0xff1000

	_, _ = s.ChannelMessageSendEmbed(tag.channelid, embed.MessageEmbed)
	if err != nil {
		return
	}
}
