package commands

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	ctx "git.randomchars.net/Reviath/RemiliaScarlet/Context"
	"github.com/bwmarrin/discordgo"
)

type ShutDown struct {
}

func (s ShutDown) Execute(ctx ctx.Ctx, session *discordgo.Session) error {
	file, err := ioutil.ReadFile("./config.json")

	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	type configStruct struct {
		Owner string `json:"Owner"`
	}

	var (
		config *configStruct
	)

	err = json.Unmarshal(file, &config)

	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	if ctx.Author().ID != config.Owner {
		_, _ = session.ChannelMessageSend(ctx.Channel().ID, "You do not own this bot.")

		return nil
	}
	_, _ = session.ChannelMessageSend(ctx.Channel().ID, "Performing shutdown.")

	os.Exit(1)
	return nil
}
