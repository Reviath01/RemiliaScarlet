package events

import (
	"database/sql"
	"go/types"

	"github.com/bwmarrin/discordgo"
	_ "github.com/go-sql-driver/mysql"
	embedutil "git.randomchars.net/Reviath/embed-util"
)

func ChannelCreate(s *discordgo.Session, event *discordgo.ChannelCreate) {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/remilia")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	type Tag struct {
		channelid string `json:"channelid"`
	}

	var tag Tag

	err = db.QueryRow("SELECT channelid FROM log WHERE guildid ='" + event.GuildID + "'").Scan(&tag.channelid)
	if err != nil {
		return
	} else {
		channel, err := discordgo.Channel(tag.channelid)
		if err != nil {
			return
		}

		embed := embedutil.NewEmbed().
		SetTitle("Channel Created!").
		AddField("Channel Name:", event.Channel.Name + " ( <#" + event.Channel.ID + "> )").
		AddField("Channel ID:", event.Channel.ID).
		AddField("Channel Type:", event.Channel.Type).MessageEmbed

		_, err = s.ChannelMessageSendEmbed(channel.ID, embed)
	}
}