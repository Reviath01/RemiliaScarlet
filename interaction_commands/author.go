package interaction_commands

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	embedutil "git.randomchars.net/Reviath/RemiliaScarlet/EmbedUtil"
	multiplexer "git.randomchars.net/Reviath/RemiliaScarlet/Multiplexer"
	"git.randomchars.net/Reviath/RemiliaScarlet/interactions"
	"git.randomchars.net/Reviath/RemiliaScarlet/sql"
	"github.com/bwmarrin/discordgo"
)

func AuthorCommand(session *discordgo.Session, interaction interactions.Interaction) interactions.InteractionResponse {
	file, _ := ioutil.ReadFile("./config.json")

	type configStruct struct {
		Token     string `json:"Token"`
		BotPrefix string `json:"BotPrefix"`
		Presence  string `json:"Presence"`
		Owner     string `json:"Owner"`
	}

	var (
		config *configStruct
	)

	json.Unmarshal(file, &config)

	if sql.CheckLanguage(interaction.GuildID) == "tr" {
		u, _ := session.User(config.Owner)
		embed := embedutil.NewEmbed().
			SetColor(0x007bff).
			AddField("Sahibim:", fmt.Sprintf("<@%s> ([%s#%s](https://discord.com/users/%s))", u.ID, u.Username, u.Discriminator, u.ID)).MessageEmbed
		return multiplexer.CreateEmbedResponse(embed)
	}
	u, _ := session.User(config.Owner)

	embed := embedutil.NewEmbed().
		SetColor(0x007bff).
		AddField("My Author:", fmt.Sprintf("<@%s> ([%s#%s](https://discord.com/users/%s))", u.ID, u.Username, u.Discriminator, u.ID)).MessageEmbed
	return multiplexer.CreateEmbedResponse(embed)
}
