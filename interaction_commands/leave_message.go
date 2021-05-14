package interaction_commands

import (
	"fmt"

	multiplexer "git.randomchars.net/Reviath/RemiliaScarlet/Multiplexer"
	"git.randomchars.net/Reviath/RemiliaScarlet/interactions"
	"git.randomchars.net/Reviath/RemiliaScarlet/sql"
	"github.com/bwmarrin/discordgo"
)

//Leave message slash command.
func LeaveMessageCommand(session *discordgo.Session, interaction interactions.Interaction) interactions.InteractionResponse {
	db := sql.Connect()
	defer db.Close()

	type Tag struct {
		message string
	}

	var tag Tag

	switch sql.CheckLanguage(interaction.GuildID) {
	case "tr":
		if !multiplexer.CheckAdministratorPermission(session, interaction.Member.User.ID, interaction.ChannelID) {
			return multiplexer.CreateResponse("Yeterli yetkiye sahip değilsin.")
		}
		if len(interaction.Data.Options[0].Value.(string)) > 254 {
			return multiplexer.CreateResponse("Mesajınız 1 ve 255 karakter aralığında olmalıdır.")
		}

		err := db.QueryRow(fmt.Sprintf("SELECT message FROM leavemessage WHERE guildid ='%s'", interaction.GuildID)).Scan(&tag.message)
		if err == nil {
			return multiplexer.CreateResponse("Çıkış mesajı zaten ayarlanmış, reset'lemek için reset_leave_message komutunu kullan.")

		}
		insert, err := db.Query(fmt.Sprintf("INSERT INTO leavemessage (message, guildid) VALUES ('%s', '%s')", interaction.Data.Options[0].Value.(string), interaction.GuildID))
		if err != nil {
			return multiplexer.CreateResponse("Bir hata oluştu.")
		}
		defer insert.Close()
		return multiplexer.CreateResponse("Başarıyla çıkış mesajı ayarlandı.")
	default:

		if !multiplexer.CheckAdministratorPermission(session, interaction.Member.User.ID, interaction.ChannelID) {
			return multiplexer.CreateResponse("You don't have enough permission.")
		}

		if len(interaction.Data.Options[0].Value.(string)) > 254 {
			return multiplexer.CreateResponse("Your message must be between 1 and 255 characters.")
		}

		err := db.QueryRow(fmt.Sprintf("SELECT message FROM leavemessage WHERE guildid ='%s'", interaction.GuildID)).Scan(&tag.message)
		if err == nil {
			return multiplexer.CreateResponse("Leave message is already existing (to reset, use reset_leave_message command).")
		}
		insert, err := db.Query(fmt.Sprintf("INSERT INTO leavemessage (message, guildid) VALUES ('%s', '%s')", interaction.Data.Options[0].Value.(string), interaction.GuildID))
		if err != nil {
			return multiplexer.CreateResponse("An error occurred, please try again.")
		}
		defer insert.Close()
		return multiplexer.CreateResponse("Leave message set successfully.")
	}
}
