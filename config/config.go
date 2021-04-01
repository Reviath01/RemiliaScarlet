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
	Owner string
	config *configStruct
)

type configStruct struct {
	Token     string `json:"Token"`
	BotPrefix string `json:"BotPrefix"`
	Presence string `json:"Presence"`
	Owner string `json:"Owner"`
}

func ReadConfig() error {
	fmt.Println("Getting data from config file.")

	file, err := ioutil.ReadFile("./config.json")

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	err = json.Unmarshal(file, &config)

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	Token = config.Token
	BotPrefix = config.BotPrefix
	Presence = config.Presence
	Owner = config.Owner

	return nil
}
