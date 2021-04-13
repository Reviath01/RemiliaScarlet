package events

import (
	"strings"

	"git.randomchars.net/Reviath/RemiliaScarlet/sql"
	"github.com/bwmarrin/discordgo"
)

func GuildMemberRemove(s *discordgo.Session, event *discordgo.GuildMemberRemove) {
	db := sql.Connect()

	type Tag struct {
		channelid    string
		leavemessage string
	}

	var tag Tag
	var leavemessage string

	err := db.QueryRow("SELECT message FROM leavemessage WHERE guildid ='" + event.GuildID + "'").Scan(&tag.leavemessage)
	if err != nil {
		leavemessage = "<@" + event.User.ID + "> left the server!"
	} else {
		leavemessage = strings.NewReplacer("{mention}", "<@"+event.User.ID+">", "{username}", event.User.Username).Replace(tag.leavemessage)
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
