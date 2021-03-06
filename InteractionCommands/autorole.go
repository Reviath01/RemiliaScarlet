package interactioncommands

import (
	"fmt"

	multiplexer "git.randomchars.net/Reviath/RemiliaScarlet/Multiplexer"
	"git.randomchars.net/Reviath/RemiliaScarlet/interactions"
	"git.randomchars.net/Reviath/RemiliaScarlet/sql"
	"github.com/bwmarrin/discordgo"
)

// AutoRoleCommand is auto role command for interactions
func AutoRoleCommand(session *discordgo.Session, interaction interactions.Interaction) interactions.InteractionResponse {
	if !multiplexer.CheckAdministratorPermission(session, interaction.Member.User.ID, interaction.GuildID) {
		return multiplexer.CreateResponse("You don't have enough permission.")
	}

	db := sql.Connect()
	defer db.Close()

	type Tag struct {
		roleid string
	}

	var tag Tag

	err := db.QueryRow(fmt.Sprintf("SELECT roleid FROM autorole WHERE guildid ='%s'", interaction.GuildID)).Scan(&tag.roleid)
	if err == nil {
		return multiplexer.CreateResponse("Autorole is already existing.")
	}

	if multiplexer.GetRole(interaction.Data.Options[0].Value.(string)) == "Error" {
		return multiplexer.CreateResponse("You need to specify a role.")
	}

	insert, _ := db.Query(fmt.Sprintf("INSERT INTO autorole (roleid, guildid) VALUES ('%s', '%s')", multiplexer.GetRole(interaction.Data.Options[0].Value.(string)), interaction.GuildID))
	defer insert.Close()
	return multiplexer.CreateResponse("Successfully set autorole.")
}
