package client

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

var (
	Token            string
	BotPrefix        string
	Presence         string
	Owner            string
	LoadInteractions string
	Database         string
	User             string
	config           *configStruct
	Password         string
	Host             string
)

type configStruct struct {
	Token            string `json:"Token"`
	BotPrefix        string `json:"BotPrefix"`
	Presence         string `json:"Presence"`
	Owner            string `json:"Owner"`
	LoadInteractions string `json:"LoadInteractions"`
	Database         string `json:"Database"`
	Host             string `json:"Host"`
	User             string `json:"User"`
	Password         string `json:"Password"`
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
	Database = config.Database
	Host = config.Host
	User = config.User
	Password = config.Password
	LoadInteractions = config.LoadInteractions

	return nil
}
