package commands

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	ctx "git.randomchars.net/Reviath/RemiliaScarlet/CommandHandler/Context"
	embedutil "git.randomchars.net/Reviath/RemiliaScarlet/EmbedUtil"
	"git.randomchars.net/Reviath/RemiliaScarlet/sql"
	"github.com/bwmarrin/discordgo"
)

type Author struct {
}

func (a Author) Execute(ctx ctx.Ctx, session *discordgo.Session) error {
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

	if sql.CheckLanguage(ctx.Guild().ID) == "tr" {
		if sql.IsBlocked(ctx.Guild().ID, "author") == "true" {
			_, _ = session.ChannelMessageSend(ctx.Channel().ID, "Bu komut bu sunucuda engellenmiş.")
			return nil
		}

		u, err := session.User(config.Owner)

		if err != nil {
			return nil
		}

		authorembed := embedutil.NewEmbed().
			SetColor(0x007bff).
			AddField("Sahibim:", "<@"+u.ID+"> (["+u.Username+"#"+u.Discriminator+"](https://discord.com/users/"+u.ID+"))").MessageEmbed
		_, _ = session.ChannelMessageSendEmbed(ctx.Channel().ID, authorembed)

		return nil
	}
	if sql.IsBlocked(ctx.Guild().ID, "author") == "true" {
		_, _ = session.ChannelMessageSend(ctx.Channel().ID, "This command is blocked on this guild.")
		return nil
	}

	u, err := session.User(config.Owner)

	if err != nil {
		return nil
	}

	authorembed := embedutil.NewEmbed().
		SetColor(0x007bff).
		AddField("My Author:", "<@"+u.ID+"> (["+u.Username+"#"+u.Discriminator+"](https://discord.com/users/"+u.ID+"))").MessageEmbed
	_, _ = session.ChannelMessageSendEmbed(ctx.Channel().ID, authorembed)

	return nil
}
