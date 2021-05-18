package interactioncommands

import (
	multiplexer "git.randomchars.net/Reviath/RemiliaScarlet/Multiplexer"
	"git.randomchars.net/Reviath/RemiliaScarlet/interactions"
	"git.randomchars.net/Reviath/RemiliaScarlet/sql"
	"github.com/bwmarrin/discordgo"
)

// UnbanCommand is unban command for interactions
func UnbanCommand(session *discordgo.Session, interaction interactions.Interaction) interactions.InteractionResponse {
	switch sql.CheckLanguage(interaction.GuildID) {
	case "tr":
		if !multiplexer.CheckBanPermission(session, interaction.Member.User.ID, interaction.GuildID) {
			return multiplexer.CreateResponse("Yeterli yetkiye sahip değilsin.")
		}
		if sql.IsBlocked(interaction.GuildID, "unban") == "true" {
			return multiplexer.CreateResponse("Bu komut bu sunucuda engellenmiş.")
		}

		u, err := session.User(multiplexer.GetUser(interaction.Data.Options[0].Value.(string)))
		if err == nil {
			err = session.GuildBanDelete(interaction.GuildID, u.ID)
			if err != nil {
				return multiplexer.CreateResponse("Bu üye banlanmamış veya yeterli yetkiye sahip değilim.")

			}
			return multiplexer.CreateResponse("Belirtilen kişinin banı kaldırıldı.")

		}
		return multiplexer.CreateResponse("Bir üye belirtmelisin.")

	default:

		if !multiplexer.CheckBanPermission(session, interaction.Member.User.ID, interaction.ChannelID) {
			return multiplexer.CreateResponse("You don't have enough permission.")

		}

		if sql.IsBlocked(interaction.GuildID, "unban") == "true" {
			return multiplexer.CreateResponse("This command is blocked on this guild.")

		}

		u, err := session.User(multiplexer.GetUser(interaction.Data.Options[0].Value.(string)))

		if err == nil {
			err = session.GuildBanDelete(interaction.GuildID, u.ID)
			if err != nil {
				return multiplexer.CreateResponse("This user is not banned or I don't have enough permission.")

			}
			return multiplexer.CreateResponse("Unbanned specified user.")

		}
		return multiplexer.CreateResponse("You need to specify the user.")
	}
}
