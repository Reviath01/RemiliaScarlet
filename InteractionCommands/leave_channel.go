package interactioncommands

import (
	"fmt"

	multiplexer "git.randomchars.net/Reviath/RemiliaScarlet/Multiplexer"
	"git.randomchars.net/Reviath/RemiliaScarlet/interactions"
	"git.randomchars.net/Reviath/RemiliaScarlet/sql"
	"github.com/bwmarrin/discordgo"
)

// LeaveChannelCommand is leave channel command for interactions
func LeaveChannelCommand(session *discordgo.Session, interaction interactions.Interaction) interactions.InteractionResponse {
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
			err = db.QueryRow(fmt.Sprintf("SELECT channelid FROM leavechannel WHERE guildid ='%s'", interaction.GuildID)).Scan(&tag.channelid)
			if err == nil {
				return multiplexer.CreateResponse("Çıkış kanalı zaten ayarlanmış reset'lemek için reset_leave_channel komutunu kullanın.")
			}
			insert, err := db.Query(fmt.Sprintf("INSERT INTO leavechannel (channelid, guildid) VALUES ('%s', '%s')", c.ID, interaction.GuildID))
			if err != nil {
				return multiplexer.CreateResponse("Bir hata oluştu.")
			}
			defer insert.Close()
			return multiplexer.CreateResponse("Başarıyla ayarlandı.")
		}
		return multiplexer.CreateResponse("Bir kanal belirtmelisin.")
	default:
		if !multiplexer.CheckAdministratorPermission(session, interaction.Member.User.ID, interaction.GuildID) {
			return multiplexer.CreateResponse("You don't have enough permission.")
		}

		c, err := session.Channel(multiplexer.GetChannel(interaction.Data.Options[0].Value.(string)))
		if err == nil {
			err = db.QueryRow(fmt.Sprintf("SELECT channelid FROM leavechannel WHERE guildid ='%s'", interaction.GuildID)).Scan(&tag.channelid)
			if err == nil {
				return multiplexer.CreateResponse("Leave channel is already existing (to reset, use reset_leave_channel command).")
			}
			insert, err := db.Query(fmt.Sprintf("INSERT INTO leavechannel (channelid, guildid) VALUES ('%s', '%s')", c.ID, interaction.GuildID))
			if err != nil {
				return multiplexer.CreateResponse("An error occurred, please try again.")
			}
			defer insert.Close()

			return multiplexer.CreateResponse("Leave channel set successfully.")
		}
		return multiplexer.CreateResponse("You need to specify the channel.")
	}
}
