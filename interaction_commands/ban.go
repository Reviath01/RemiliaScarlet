package interaction_commands

import (
	"fmt"

	multiplexer "git.randomchars.net/Reviath/RemiliaScarlet/Multiplexer"
	"git.randomchars.net/Reviath/RemiliaScarlet/interactions"
	"git.randomchars.net/Reviath/RemiliaScarlet/sql"
	"github.com/bwmarrin/discordgo"
)

func BanCommand(session *discordgo.Session, interaction interactions.Interaction) interactions.InteractionResponse {
	if !multiplexer.CheckBanPermission(session, interaction.Member.User.ID, interaction.ChannelID) {
		return multiplexer.CreateResponse("You don't have enough permission.")
	}
	if sql.CheckLanguage(interaction.GuildID) == "tr" {
		if sql.IsBlocked(interaction.GuildID, "ban") == "true" {
			return multiplexer.CreateResponse("Bu komut bu sunucuda engellenmiş")
		}
		u, err := session.User(multiplexer.GetUser(interaction.Data.Options[0].Value.(string)))
		if err != nil {
			return multiplexer.CreateResponse("Bir üye belirtmelisin.")
		}

		err = session.GuildBanCreate(interaction.GuildID, u.ID, 0)
		if err != nil {
			return multiplexer.CreateResponse("Yeterli yetkiye sahip değilim.")
		}
	}

	if sql.IsBlocked(interaction.GuildID, "ban") == "true" {
		return multiplexer.CreateResponse("This command is blocked on this guild.")
	}

	u, err := session.User(multiplexer.GetUser(interaction.Data.Options[0].Value.(string)))
	if err != nil {
		return multiplexer.CreateResponse("You need to specify a user.")
	}
	err = session.GuildBanCreate(interaction.GuildID, u.ID, 0)
	if err != nil {
		return multiplexer.CreateResponse(fmt.Sprintf("An error occurred: %s", err.Error()))
	}
	return multiplexer.CreateResponse("Successfully performed ban on specified user.")
}
