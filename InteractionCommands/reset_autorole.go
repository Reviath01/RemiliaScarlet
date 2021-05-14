package interaction_commands

import (
	"fmt"

	multiplexer "git.randomchars.net/Reviath/RemiliaScarlet/Multiplexer"
	"git.randomchars.net/Reviath/RemiliaScarlet/interactions"
	"git.randomchars.net/Reviath/RemiliaScarlet/sql"
	"github.com/bwmarrin/discordgo"
)

// ResetAutoroleCommand is reset auto role command for interactions
func ResetAutoroleCommand(session *discordgo.Session, interaction interactions.Interaction) interactions.InteractionResponse {
	db := sql.Connect()
	defer db.Close()

	type Tag struct {
		roleid string
	}

	var tag Tag

	switch sql.CheckLanguage(interaction.GuildID) {
	case "tr":
		if !multiplexer.CheckAdministratorPermission(session, interaction.Member.User.ID, interaction.ChannelID) {
			return multiplexer.CreateResponse("Yeterli yetkiye sahip değilsin.")
		}

		var tag Tag

		err := db.QueryRow(fmt.Sprintf("SELECT roleid FROM autorole WHERE guildid ='%s'", interaction.GuildID)).Scan(&tag.roleid)
		if err == nil {
			delete, err := db.Query(fmt.Sprintf("DELETE FROM autorole WHERE guildid ='%s'", interaction.GuildID))
			if err != nil {
				return multiplexer.CreateResponse("Bir hata oluştu.")
			}
			defer delete.Close()

			return multiplexer.CreateResponse("Başarıyla otorol sıfırlandı.")

		}
		return multiplexer.CreateResponse("Otorol ayarlanmadığı için sıfırlanamaz.")
	default:

		if !multiplexer.CheckAdministratorPermission(session, interaction.Member.User.ID, interaction.ChannelID) {
			return multiplexer.CreateResponse("You don't have enough permission.")
		}

		err := db.QueryRow(fmt.Sprintf("SELECT roleid FROM autorole WHERE guildid ='%s'", interaction.GuildID)).Scan(&tag.roleid)
		if err == nil {
			delete, err := db.Query(fmt.Sprintf("DELETE FROM autorole WHERE guildid ='%s'", interaction.GuildID))
			if err != nil {
				return multiplexer.CreateResponse("An error occurred, please try again.")

			}

			defer delete.Close()

			return multiplexer.CreateResponse("Successfully reset auto role.")

		}
		return multiplexer.CreateResponse("Auto role is not existing, so you can't reset.")
	}
}
