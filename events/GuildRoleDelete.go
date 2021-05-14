package events

import (
	embedutil "git.randomchars.net/FreeNitori/EmbedUtil"
	"git.randomchars.net/Reviath/RemiliaScarlet/sql"
	"github.com/bwmarrin/discordgo"
)

// GuildRoleDelete is an event handler for guild role delete event
func GuildRoleDelete(s *discordgo.Session, event *discordgo.GuildRoleDelete) {
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
		embed := embedutil.New("Rol Silindi!", "")
		embed.AddField("Rol İD'si:", event.RoleID, true)
		embed.Color = 0xefff00

		_, _ = s.ChannelMessageSendEmbed(tag.channelid, embed.MessageEmbed)
		return
	}

	err := db.QueryRow("SELECT channelid FROM log WHERE guildid ='" + event.GuildID + "'").Scan(&tag.channelid)
	if err != nil {
		return
	}
	embed := embedutil.New("Role Deleted!", "")
	embed.AddField("Role ID:", event.RoleID, true)
	embed.Color = 0xefff00

	_, _ = s.ChannelMessageSendEmbed(tag.channelid, embed.MessageEmbed)
}
