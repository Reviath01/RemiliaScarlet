package events

import (
	"strings"

	"git.randomchars.net/Reviath/RemiliaScarlet/sql"
	"github.com/bwmarrin/discordgo"
)

func GuildMemberAdd(s *discordgo.Session, event *discordgo.GuildMemberAdd) {
	db := sql.Connect()

	type Tag struct {
		channelid      string
		roleid         string
		welcomemessage string
		lang           string
	}

	var tag Tag
	var welcomemsg string

	err := db.QueryRow("SELECT language FROM languages WHERE guildid ='" + event.GuildID + "'").Scan(&tag.lang)

	if err == nil && tag.lang == "tr" {
		err = db.QueryRow("SELECT message FROM welcomemessage WHERE guildid ='" + event.GuildID + "'").Scan(&tag.welcomemessage)
		if err != nil {
			welcomemsg = "Sunucuya hoş geldin <@" + event.User.ID + ">!"
		} else {
			welcomemsg = strings.NewReplacer("{mention}", "<@"+event.User.ID+">", "{username}", event.User.Username).Replace(tag.welcomemessage)
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
		return
	}

	err = db.QueryRow("SELECT message FROM welcomemessage WHERE guildid ='" + event.GuildID + "'").Scan(&tag.welcomemessage)
	if err != nil {
		welcomemsg = "Welcome to server <@" + event.User.ID + ">"
	} else {
		welcomemsg = strings.NewReplacer("{mention}", "<@"+event.User.ID+">", "{username}", event.User.Username).Replace(tag.welcomemessage)
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
