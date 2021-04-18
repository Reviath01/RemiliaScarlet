package interaction_commands

import (
	"fmt"

	multiplexer "git.randomchars.net/Reviath/RemiliaScarlet/Multiplexer"
	"git.randomchars.net/Reviath/RemiliaScarlet/interactions"
	"git.randomchars.net/Reviath/RemiliaScarlet/sql"
	"github.com/bwmarrin/discordgo"
)

func AutoRoleCommand(session *discordgo.Session, interaction interactions.Interaction) interactions.InteractionResponse {
	if !multiplexer.CheckAdministratorPermission(session, interaction.Member.User.ID, interaction.ChannelID) {
		response := interactions.InteractionResponse{
			Type: interactions.InteractionResponseTypeChannelMessageWithSource,
			Data: interactions.InteractionResponseData{
				TTS:     false,
				Content: "You don't have enough permission.",
			},
		}
		return response
	}

	db := sql.Connect()

	type Tag struct {
		roleid string
	}

	var tag Tag

	err := db.QueryRow(fmt.Sprintf("SELECT roleid FROM autorole WHERE guildid ='%s'", interaction.GuildID)).Scan(&tag.roleid)
	if err == nil {
		response := interactions.InteractionResponse{
			Type: interactions.InteractionResponseTypeChannelMessageWithSource,
			Data: interactions.InteractionResponseData{
				TTS:     false,
				Content: "Autorole is already existing.",
			},
		}
		return response
	}

	if multiplexer.GetRole(interaction.Data.Options[0].Value.(string)) == "Error" {
		response := interactions.InteractionResponse{
			Type: interactions.InteractionResponseTypeChannelMessageWithSource,
			Data: interactions.InteractionResponseData{
				TTS:     false,
				Content: "You need to specify a role.",
			},
		}
		return response
	}

	insert, _ := db.Query(fmt.Sprintf("INSERT INTO autorole (roleid, guildid) VALUES ('%s', '%s')", multiplexer.GetRole(interaction.Data.Options[0].Value.(string)), interaction.GuildID))
	defer insert.Close()
	response := interactions.InteractionResponse{
		Type: interactions.InteractionResponseTypeChannelMessageWithSource,
		Data: interactions.InteractionResponseData{
			TTS:     false,
			Content: "Successfully set autorole.",
		},
	}
	return response
}
