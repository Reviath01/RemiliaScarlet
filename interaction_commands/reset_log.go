package interaction_commands

import (
	"fmt"

	multiplexer "git.randomchars.net/Reviath/RemiliaScarlet/Multiplexer"
	"git.randomchars.net/Reviath/RemiliaScarlet/interactions"
	"git.randomchars.net/Reviath/RemiliaScarlet/sql"
	"github.com/bwmarrin/discordgo"
)

func ResetLogCommand(session *discordgo.Session, interaction interactions.Interaction) interactions.InteractionResponse {
	db := sql.Connect()

	type Tag struct {
		channelid string
	}

	var tag Tag

	if sql.CheckLanguage(interaction.GuildID) == "tr" {
		if !multiplexer.CheckAdministratorPermission(session, interaction.Member.User.ID, interaction.ChannelID) {
			return multiplexer.CreateResponse("Yeterli yetkiye sahip değilsin.")
		}
		err := db.QueryRow(fmt.Sprintf("SELECT channelid FROM log WHERE guildid ='%s'", interaction.GuildID)).Scan(&tag.channelid)
		if err == nil {
			delete, err := db.Query(fmt.Sprintf("DELETE FROM log WHERE guildid ='%s'", interaction.GuildID))
			if err != nil {
				return multiplexer.CreateResponse("Bir hata oluştu.")
			}

			defer delete.Close()

			return multiplexer.CreateResponse("Başarıyla log kanalı sıfırlandı.")

		}
		return multiplexer.CreateResponse("Log kanalı ayarlanmamış, sıfırlayamazsın.")
	}

	if !multiplexer.CheckAdministratorPermission(session, interaction.Member.User.ID, interaction.ChannelID) {
		return multiplexer.CreateResponse("You don't have enough permission.")
	}

	err := db.QueryRow(fmt.Sprintf("SELECT channelid FROM log WHERE guildid ='%s'", interaction.GuildID)).Scan(&tag.channelid)
	if err == nil {
		delete, err := db.Query(fmt.Sprintf("DELETE FROM log WHERE guildid ='%s'", interaction.GuildID))
		if err != nil {
			return multiplexer.CreateResponse("An error occurred, please try again.")
		}

		defer delete.Close()

		return multiplexer.CreateResponse("Successfully reset log.")

	}
	return multiplexer.CreateResponse("Log channel is not existing, so you can't reset.")
}
