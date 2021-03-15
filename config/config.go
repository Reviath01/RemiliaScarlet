package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

var (
	Token     string
	BotPrefix string
	Owner string
	IgnoreBots string
	config *configStruct
)

type configStruct struct {
	Token     string `json:"Token"`
	BotPrefix string `json:"BotPrefix"`
	IgnoreBots string `json:"IgnoreBots"`
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
	Owner = config.Owner
	IgnoreBots = config.IgnoreBots

	return nil
}
