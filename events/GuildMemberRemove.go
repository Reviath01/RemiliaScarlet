package events

import (
	"github.com/bwmarrin/discordgo"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func GuildMemberRemove(s *discordgo.Session, event *discordgo.GuildMemberRemove) {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/remilia")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	type Tag struct {
		channelid string `json:"channelid"`
		leavemessage string `json:"message"`
	}

	var tag Tag
	var leavemessage string

	err = db.QueryRow("SELECT message FROM leavemessage WHERE guildid ='" + event.GuildID + "'").Scan(&tag.leavemessage)
	if err != nil {
		leavemessage = "<@" + event.User.ID + "> left the server!"
	} else {
		leavemessage = tag.leavemessage
	}


	err = db.QueryRow("SELECT channelid FROM leavechannel WHERE guildid ='" + event.GuildID + "'").Scan(&tag.channelid)
	if err != nil {
		return
	} else {
		channel, err := s.Channel(tag.channelid)
		if err != nil {
			return
		}
		s.ChannelMessageSend(channel.ID, leavemessage)
	}
}
