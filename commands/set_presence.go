package commands

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	ctx "git.randomchars.net/Reviath/RemiliaScarlet/CommandHandler/Context"
	"github.com/bwmarrin/discordgo"
)

type SetPresence struct {
}

func (s SetPresence) Execute(ctx ctx.Ctx, session *discordgo.Session) error {
	file, err := ioutil.ReadFile("./config.json")

	if err != nil {
		fmt.Println(err.Error())
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

	u, err := session.User(config.Owner)
	if err != nil {
		return nil
	}

	if ctx.Author().ID != u.ID {
		_, _ = session.ChannelMessageSend(ctx.Channel().ID, "You do not own this bot.")

		return nil
	}
	err = session.UpdateGameStatus(0, strings.Join(ctx.Args(), " "))
	if err != nil {
		_, _ = session.ChannelMessageSend(ctx.Channel().ID, "An error occurred.")

		return nil
	}
	_, _ = session.ChannelMessageSend(ctx.Channel().ID, "Successfully updated presence.")

	return nil
}
