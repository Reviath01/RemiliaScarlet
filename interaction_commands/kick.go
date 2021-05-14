package interaction_commands

import (
	multiplexer "git.randomchars.net/Reviath/RemiliaScarlet/Multiplexer"
	"git.randomchars.net/Reviath/RemiliaScarlet/interactions"
	"git.randomchars.net/Reviath/RemiliaScarlet/sql"
	"github.com/bwmarrin/discordgo"
)

//Kick slash command.
func KickCommand(session *discordgo.Session, interaction interactions.Interaction) interactions.InteractionResponse {
	switch sql.CheckLanguage(interaction.GuildID) {
	case "tr":
		if !multiplexer.CheckKickPermission(session, interaction.Member.User.ID, interaction.ChannelID) {
			return multiplexer.CreateResponse("Yeterli yetkiye sahip değilsin.")
		}
		if sql.IsBlocked(interaction.GuildID, "kick") == "true" {
			return multiplexer.CreateResponse("Bu komut bu sunucuda engellenmiş.")
		}

		u, err := session.User(multiplexer.GetUser(interaction.Data.Options[0].Value.(string)))
		if err == nil {
			err = session.GuildMemberDelete(interaction.GuildID, u.ID)
			if err != nil {
				return multiplexer.CreateResponse("Yeterli yetkim yok.")
			}
			return multiplexer.CreateResponse("Belirtilen kişi sunucudan atıldı.")

		}
		return multiplexer.CreateResponse("Bir üye belirtmelisin.")
	default:

		if !multiplexer.CheckKickPermission(session, interaction.Member.User.ID, interaction.ChannelID) {
			return multiplexer.CreateResponse("You don't have enough permission.")
		}

		if sql.IsBlocked(interaction.GuildID, "kick") == "true" {
			return multiplexer.CreateResponse("This command is blocked on this guild.")
		}

		u, err := session.User(multiplexer.GetUser(interaction.Data.Options[0].Value.(string)))
		if err == nil {
			err = session.GuildMemberDelete(interaction.GuildID, u.ID)
			if err != nil {
				return multiplexer.CreateResponse("I do not have enough permission.")
			}
			return multiplexer.CreateResponse("Kicked specified user.")
		}
		return multiplexer.CreateResponse("You need to specify the user.")
	}
}
