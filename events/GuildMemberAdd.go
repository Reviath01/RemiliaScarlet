package events

import (
	"github.com/bwmarrin/discordgo"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"strings"
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
		welcomemessage string `json:"message"`
	}

	var tag Tag
	var welcomemsg string

	err = db.QueryRow("SELECT message FROM welcomemessage WHERE guildid ='" + event.GuildID + "'").Scan(&tag.welcomemessage)
	if err != nil {
		welcomemsg = "Welcome to server <@" + event.User.ID + ">"
	} else {
		welcomemsg = strings.NewReplacer("{mention}", "<@" + event.User.ID + ">", "{username}", event.User.Username).Replace(tag.welcomemessage)
	}

	err = db.QueryRow("SELECT channelid FROM welcomechannel WHERE guildid ='" + event.GuildID + "'").Scan(&tag.channelid)
	if err != nil {
		return
	} else {
		channel, err := s.Channel(tag.channelid)
		if err != nil {
			return
		}
		s.ChannelMessageSend(channel.ID, welcomemsg)
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
