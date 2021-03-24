package events

import (
	"database/sql"

	embedutil "git.randomchars.net/Reviath/embed-util"
	"github.com/bwmarrin/discordgo"
	_ "github.com/go-sql-driver/mysql"
)

func GuildRoleDelete(s *discordgo.Session, event *discordgo.GuildRoleDelete) {
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
		embed := embedutil.NewEmbed().
			SetTitle("Role Deleted!").
			AddField("Role ID:", event.RoleID).
			SetColor(0xefff00).MessageEmbed

		_, err = s.ChannelMessageSendEmbed(tag.channelid, embed)
		if err != nil {
			return
		}
	}
}
