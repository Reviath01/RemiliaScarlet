package interactioncommands

import (
	"fmt"

	multiplexer "git.randomchars.net/Reviath/RemiliaScarlet/Multiplexer"
	"git.randomchars.net/Reviath/RemiliaScarlet/interactions"
	"git.randomchars.net/Reviath/RemiliaScarlet/sql"
	"github.com/bwmarrin/discordgo"
)

// LogCommand is log command for interactions
func LogCommand(session *discordgo.Session, interaction interactions.Interaction) interactions.InteractionResponse {
	db := sql.Connect()
	defer db.Close()

	type Tag struct {
		channelid string
	}

	var tag Tag

	switch sql.CheckLanguage(interaction.GuildID) {
	case "tr":
		if !multiplexer.CheckAdministratorPermission(session, interaction.Member.User.ID, interaction.GuildID) {
			return multiplexer.CreateResponse("Yeterli yetkiye sahip değilsin.")
		}
		c, err := session.Channel(multiplexer.GetChannel(interaction.Data.Options[0].Value.(string)))
		if err == nil {
			err = db.QueryRow(fmt.Sprintf("SELECT channelid FROM log WHERE guildid ='%s'", interaction.GuildID)).Scan(&tag.channelid)
			if err == nil {
				return multiplexer.CreateResponse("Log kanalı zaten ayarlanmış, sıfırlamak için reset_log komutunu kullan.")
			}
			insert, err := db.Query(fmt.Sprintf("INSERT INTO log (channelid, guildid) VALUES ('%s', '%s')", c.ID, interaction.GuildID))
			if err != nil {
				return multiplexer.CreateResponse("Bir hata oluştu.")
			}
			defer insert.Close()
			return multiplexer.CreateResponse("Log kanalı başarıyla ayarlandı!")
		}
		return multiplexer.CreateResponse("Log kanalını belirtmelisin.")
	default:

		if !multiplexer.CheckAdministratorPermission(session, interaction.Member.User.ID, interaction.GuildID) {
			return multiplexer.CreateResponse("You don't have enough permission.")
		}

		c, err := session.Channel(multiplexer.GetChannel(interaction.Data.Options[0].Value.(string)))
		if err == nil {

			err = db.QueryRow(fmt.Sprintf("SELECT channelid FROM log WHERE guildid ='%s'", interaction.GuildID)).Scan(&tag.channelid)
			if err == nil {
				return multiplexer.CreateResponse("Log is already existing (to reset, use reset_log command).")

			}
			insert, err := db.Query(fmt.Sprintf("INSERT INTO log (channelid, guildid) VALUES ('%s', '%s')", c.ID, interaction.GuildID))
			if err != nil {
				return multiplexer.CreateResponse("An error occurred, please try again.")
			}
			defer insert.Close()
			return multiplexer.CreateResponse("Logging channel set successfully.")
		}
		return multiplexer.CreateResponse("You need to specify the channel.")
	}
}
