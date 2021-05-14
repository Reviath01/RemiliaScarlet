package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

var (
	Token            string        // Token is to run the bot with token
	BotPrefix        string        // BotPrefix is the prefix that the command handler will use
	Presence         string        // Presence is bot's playing part
	Owner            string        // Owner of the bot (needed on some commands)
	LoadInteractions string        // LoadInteractions is bot's slash commands
	Database         string        // Database name
	User             string        // User to connect MySQL database
	config           *configStruct // config structure
	Password         string        // Password to connect MySQL database
	Host             string        // Host to connect MySQL database
	WebURL           string        // WebURL to run website
	ClientID         string        // ClientID to use on website
	ClientSecret     string        // ClientSecret for OAuth2
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
	WebURL           string `json:"WebURL"`
	ClientID         string `json:"ClientID"`
	ClientSecret     string `json:"ClientSecret"`
}

// ReadConfig reads config.json file, returns error
func ReadConfig() error {
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
	WebURL = config.WebURL
	ClientID = config.ClientID
	ClientSecret = config.ClientSecret

	return nil
}
