package events

import (
	"strconv"

	"git.randomchars.net/Reviath/RemiliaScarlet/sql"
	embedutil "git.randomchars.net/freenitori/embedutil"
	"github.com/bwmarrin/discordgo"
)

// GuildRoleCreate is an event handler for guild role create event
func GuildRoleCreate(s *discordgo.Session, event *discordgo.GuildRoleCreate) {
	db := sql.Connect()
	defer db.Close()

	type Tag struct {
		channelid string
	}

	var tag Tag

	if sql.CheckLanguage(event.GuildID) == "tr" {
		err := db.QueryRow("SELECT channelid FROM log WHERE guildid ='" + event.GuildID + "'").Scan(&tag.channelid)
		if err != nil {
			return
		}
		embed := embedutil.New("Rol Oluşturuldu!", "")
		embed.AddField("Rolün İsmi:", event.Role.Name+" ( <@&"+event.Role.ID+"> )", true)
		embed.AddField("Rol İD'si:", event.Role.ID, true)
		embed.AddField("Rol Rengi:", strconv.Itoa(event.Role.Color), true)
		embed.Color = 0xefff00

		s.ChannelMessageSendEmbed(tag.channelid, embed.MessageEmbed)
		return
	}

	err := db.QueryRow("SELECT channelid FROM log WHERE guildid ='" + event.GuildID + "'").Scan(&tag.channelid)
	if err != nil {
		return
	}
	embed := embedutil.New("Role Created!", "")
	embed.AddField("Role Name:", event.Role.Name+" ( <@&"+event.Role.ID+"> )", true)
	embed.AddField("Role ID:", event.Role.ID, true)
	embed.AddField("Role Color:", strconv.Itoa(event.Role.Color), true)
	embed.Color = 0xefff00

	s.ChannelMessageSendEmbed(tag.channelid, embed.MessageEmbed)
}
