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
	ClientSecret     string        // ClientSecret for OAuth2
)

type configStruct struct {
	Token            string `json:"Token"`            // Token from config file
	BotPrefix        string `json:"BotPrefix"`        // BotPrefix from config file
	Presence         string `json:"Presence"`         // Presence from config file
	Owner            string `json:"Owner"`            // Owner from config file
	LoadInteractions string `json:"LoadInteractions"` // LoadInteractions from config file
	Database         string `json:"Database"`         // Database from config file
	Host             string `json:"Host"`             // Host from config file
	User             string `json:"User"`             // User from config file
	Password         string `json:"Password"`         // Password from config file
	WebURL           string `json:"WebURL"`           // WebURL from config file
	ClientSecret     string `json:"ClientSecret"`     // ClientSecret from config file
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
	ClientSecret = config.ClientSecret

	return nil
}
