package commands

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	embedutil "git.randomchars.net/Reviath/RemiliaScarlet/EmbedUtil"
	CommandHandler "git.randomchars.net/Reviath/RemiliaScarlet/handler"
	"git.randomchars.net/Reviath/RemiliaScarlet/interactions"
	"git.randomchars.net/Reviath/RemiliaScarlet/sql"
	"github.com/bwmarrin/discordgo"
)

func AuthorCommand(ctx CommandHandler.Context, _ []string) error {
	file, err := ioutil.ReadFile("./config.json")

	if err != nil {
		return nil
	}

	type configStruct struct {
		Token     string `json:"Token"`
		BotPrefix string `json:"BotPrefix"`
		Presence  string `json:"Presence"`
		Owner     string `json:"Owner"`
	}

	var (
		config *configStruct
	)

	err = json.Unmarshal(file, &config)

	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	if sql.CheckLanguage(ctx.Guild.ID) == "tr" {
		if sql.IsBlocked(ctx.Guild.ID, "author") == "true" {
			_, _ = ctx.Reply("Bu komut bu sunucuda engellenmiş.")
			return nil
		}

		u, err := ctx.Session.User(config.Owner)

		if err != nil {
			return nil
		}

		authorembed := embedutil.NewEmbed().
			SetColor(0x007bff).
			AddField("Sahibim:", fmt.Sprintf("<@%s> ([%s#%s](https://discord.com/users/%s))", u.ID, u.Username, u.Discriminator, u.ID)).MessageEmbed
		_, _ = ctx.ReplyEmbed(authorembed)

		return nil
	}
	if sql.IsBlocked(ctx.Guild.ID, "author") == "true" {
		_, _ = ctx.Reply("This command is blocked on this guild.")
		return nil
	}

	u, err := ctx.Session.User(config.Owner)

	if err != nil {
		return nil
	}

	authorembed := embedutil.NewEmbed().
		SetColor(0x007bff).
		AddField("My Author:", fmt.Sprintf("<@%s> ([%s#%s](https://discord.com/users/%s))", u.ID, u.Username, u.Discriminator, u.ID)).MessageEmbed
	_, _ = ctx.ReplyEmbed(authorembed)

	return nil
}

func InteractionAuthorCommand(session *discordgo.Session, interaction interactions.Interaction) interactions.InteractionResponse {
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
		response := interactions.InteractionResponse{
			Type: interactions.InteractionResponseTypeChannelMessageWithSource,
			Data: interactions.InteractionResponseData{
				TTS:    false,
				Embeds: []discordgo.MessageEmbed{*embed},
			},
		}
		return response
	}
	u, _ := session.User(config.Owner)

	embed := embedutil.NewEmbed().
		SetColor(0x007bff).
		AddField("My Author:", fmt.Sprintf("<@%s> ([%s#%s](https://discord.com/users/%s))", u.ID, u.Username, u.Discriminator, u.ID)).MessageEmbed

	response := interactions.InteractionResponse{
		Type: interactions.InteractionResponseTypeChannelMessageWithSource,
		Data: interactions.InteractionResponseData{
			TTS:    false,
			Embeds: []discordgo.MessageEmbed{*embed},
		},
	}
	return response
}
