package interactioncommands

import (
	"fmt"

	multiplexer "git.randomchars.net/Reviath/RemiliaScarlet/Multiplexer"
	"git.randomchars.net/Reviath/RemiliaScarlet/interactions"
	"git.randomchars.net/Reviath/RemiliaScarlet/sql"
	"github.com/bwmarrin/discordgo"
)

// ResetWelcomeMessageCommand is reset welcome message command for interactions
func ResetWelcomeMessageCommand(session *discordgo.Session, interaction interactions.Interaction) interactions.InteractionResponse {
	db := sql.Connect()
	defer db.Close()

	type Tag struct {
		message string
	}

	var tag Tag

	switch sql.CheckLanguage(interaction.GuildID) {
	case "tr":
		if !multiplexer.CheckAdministratorPermission(session, interaction.Member.User.ID, interaction.GuildID) {
			return multiplexer.CreateResponse("Yeterli yetkiye sahip değilsin.")

		}

		err := db.QueryRow(fmt.Sprintf("SELECT message FROM welcomemessage WHERE guildid ='%s'", interaction.GuildID)).Scan(&tag.message)
		if err == nil {
			delete, err := db.Query(fmt.Sprintf("DELETE FROM welcomemessage WHERE guildid ='%s'", interaction.GuildID))
			if err != nil {
				return multiplexer.CreateResponse("Bir hata oluştu.")
			}
			defer delete.Close()
			return multiplexer.CreateResponse("Hoş geldin mesajı başarıyla sıfırlandı.")
		}
		return multiplexer.CreateResponse("Hoş geldin mesajı ayarlanmamış, sıfırlayamazsın.")
	default:

		if !multiplexer.CheckAdministratorPermission(session, interaction.Member.User.ID, interaction.GuildID) {
			return multiplexer.CreateResponse("You don't have enough permission.")

		}

		err := db.QueryRow(fmt.Sprintf("SELECT message FROM welcomemessage WHERE guildid ='%s'", interaction.GuildID)).Scan(&tag.message)
		if err == nil {
			delete, err := db.Query(fmt.Sprintf("DELETE FROM welcomemessage WHERE guildid ='%s'", interaction.GuildID))
			if err != nil {
				return multiplexer.CreateResponse("An error occurred, please try again.")
			}
			defer delete.Close()
			return multiplexer.CreateResponse("Successfully reset welcome message.")
		}
		return multiplexer.CreateResponse("Welcome message is not existing, so you can't reset.")
	}
}
