package events

import (
	embedutil "git.randomchars.net/Reviath/RemiliaScarlet/EmbedUtil"
	"git.randomchars.net/Reviath/RemiliaScarlet/sql"
	"github.com/bwmarrin/discordgo"
)

func GuildRoleDelete(s *discordgo.Session, event *discordgo.GuildRoleDelete) {
	db := sql.Connect()

	type Tag struct {
		channelid string
	}

	var tag Tag

	if sql.CheckLanguage(event.GuildID) == "tr" {
		err := db.QueryRow("SELECT channelid FROM log WHERE guildid ='" + event.GuildID + "'").Scan(&tag.channelid)
		if err != nil {
			return
		} else {
			embed := embedutil.NewEmbed().
				SetTitle("Rol Silindi!").
				AddField("Rol İD'si:", event.RoleID).
				SetColor(0xefff00).MessageEmbed

			_, _ = s.ChannelMessageSendEmbed(tag.channelid, embed)
			if err != nil {
				return
			}
			return
		}
	}

	err := db.QueryRow("SELECT channelid FROM log WHERE guildid ='" + event.GuildID + "'").Scan(&tag.channelid)
	if err != nil {
		return
	} else {
		embed := embedutil.NewEmbed().
			SetTitle("Role Deleted!").
			AddField("Role ID:", event.RoleID).
			SetColor(0xefff00).MessageEmbed

		_, _ = s.ChannelMessageSendEmbed(tag.channelid, embed)
		if err != nil {
			return
		}
	}
}
