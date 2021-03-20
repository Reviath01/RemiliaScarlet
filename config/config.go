package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

var (
	Token     string
	BotPrefix string
	Presence string
	config *configStruct
)

type configStruct struct {
	Token     string `json:"Token"`
	BotPrefix string `json:"BotPrefix"`
	Presence string `json:"Presence"`
}

func ReadConfig() error {
	fmt.Println("Getting data from config file.")

	file, err := ioutil.ReadFile("./config.json")

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	if config.Token == "INSERT_TOKEN_HERE" || config.BotPrefix == "BOTS_PREFIX" || config.Presence == "YOUR_BOTS_GAME_ACTIVITY" {
		fmt.Println("Please edit the configuration file (config.json) before using.")
	}

	err = json.Unmarshal(file, &config)

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	Token = config.Token
	BotPrefix = config.BotPrefix
	Presence = config.Presence

	return nil
}
