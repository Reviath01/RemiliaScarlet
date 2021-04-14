package events

import (
	"strconv"

	embedutil "git.randomchars.net/Reviath/RemiliaScarlet/EmbedUtil"
	"git.randomchars.net/Reviath/RemiliaScarlet/sql"
	"github.com/bwmarrin/discordgo"
)

func GuildRoleCreate(s *discordgo.Session, event *discordgo.GuildRoleCreate) {
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
				SetTitle("Rol Oluşturuldu!").
				AddField("Rolün İsmi:", event.Role.Name+" ( <@&"+event.Role.ID+"> )").
				AddField("Rol İD'si:", event.Role.ID).
				AddField("Rol Rengi:", strconv.Itoa(event.Role.Color)).
				SetColor(0xefff00).MessageEmbed

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
			SetTitle("Role Created!").
			AddField("Role Name:", event.Role.Name+" ( <@&"+event.Role.ID+"> )").
			AddField("Role ID:", event.Role.ID).
			AddField("Role Color:", strconv.Itoa(event.Role.Color)).
			SetColor(0xefff00).MessageEmbed

		_, _ = s.ChannelMessageSendEmbed(tag.channelid, embed)
		if err != nil {
			return
		}
	}
}
