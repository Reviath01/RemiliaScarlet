package events

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/bwmarrin/discordgo"
)

func Ready(s *discordgo.Session, event *discordgo.Ready) {
	file, err := ioutil.ReadFile("./config.json")

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	type configStruct struct {
		Presence string `json:"Presence"`
	}

	var (
		config *configStruct
	)

	err = json.Unmarshal(file, &config)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	s.UpdateStatus(0, config.Presence)
}
