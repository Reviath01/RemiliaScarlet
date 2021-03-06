package interactioncommands

import (
	multiplexer "git.randomchars.net/Reviath/RemiliaScarlet/Multiplexer"
	"git.randomchars.net/Reviath/RemiliaScarlet/interactions"
	"git.randomchars.net/Reviath/RemiliaScarlet/sql"
	"github.com/bwmarrin/discordgo"
)

// LanguageCommand is language command for interactions
func LanguageCommand(session *discordgo.Session, interaction interactions.Interaction) interactions.InteractionResponse {
	db := sql.Connect()
	defer db.Close()

	args := interaction.Data.Options[0].Value.(string)

	switch sql.CheckLanguage(interaction.GuildID) {
	case "tr":
		if !multiplexer.CheckAdministratorPermission(session, interaction.Member.User.ID, interaction.GuildID) {
			return multiplexer.CreateResponse("Yeterli yetkiye sahip değilsin.")
		}
		if args == "tr" {
			return multiplexer.CreateResponse("Dil zaten Türkçe ayarlı.")
		} else if args == "en" {
			delete, err := db.Query("DELETE FROM languages WHERE guildid ='" + interaction.GuildID + "'")
			if err != nil {
				return multiplexer.CreateResponse("Bir hata oluştu.")
			}
			defer delete.Close()
			return multiplexer.CreateResponse("Dili ingilizce olarak ayarlıyorum...")
		} else {
			return multiplexer.CreateResponse("Kullanılabilir diller: `tr`, `en`")
		}
	default:
		if !multiplexer.CheckAdministratorPermission(session, interaction.Member.User.ID, interaction.GuildID) {
			return multiplexer.CreateResponse("You don't have enough permission.")
		}
		if args == "tr" {
			insert, _ := db.Query("INSERT INTO languages (language, guildid) VALUES ('tr', '" + interaction.GuildID + "')")
			defer insert.Close()
			return multiplexer.CreateResponse("Dil Türkçe yapıldı!")
		} else if args == "en" {
			return multiplexer.CreateResponse("Language is en.")
		}
		return multiplexer.CreateResponse("Language options: `tr`, `en`")
	}
}
