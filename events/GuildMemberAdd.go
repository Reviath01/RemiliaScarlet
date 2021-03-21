package events

import (
	"github.com/bwmarrin/discordgo"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func GuildMemberAdd(s *discordgo.Session, event *discordgo.GuildMemberAdd) {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/remilia")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	type Tag struct {
		channelid string `json:"channelid"`
		roleid string `json:"roleid"`
	}

	var tag Tag

	err = db.QueryRow("SELECT channelid FROM welcomechannel WHERE guildid ='" + event.GuildID + "'").Scan(&tag.channelid)
	if err != nil {
		return
	} else {
		channel, err := s.Channel(tag.channelid)
		if err != nil {
			return
		}
		s.ChannelMessageSend(channel.ID, "Welcome to server <@" + event.User.ID + ">")
	}
	err = db.QueryRow("SELECT roleid FROM autorole WHERE guildid ='" + event.GuildID + "'").Scan(&tag.roleid)
	if err != nil {
		return
	} else {
		err = s.GuildMemberRoleAdd(event.GuildID, event.User.ID, tag.roleid)
		if err != nil {
			return
		}
	}
}
