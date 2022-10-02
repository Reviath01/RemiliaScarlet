package config

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/BurntSushi/toml"
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
	WebPort          string        // WebPort to run website at
	BaseURL          string        // BaseURL to run website at
	ClientSecret     string        // ClientSecret for OAuth2
	RunWeb           string        // RunWeb for checking if running server
)

type configStruct struct {
	// Discord stuff from remilia.conf
	Discord struct {
		Token            string `json:"Token"`            // Token from config file
		BotPrefix        string `json:"BotPrefix"`        // BotPrefix from config file
		Presence         string `json:"Presence"`         // Presence from config file
		Owner            string `json:"Owner"`            // Owner from config file
		LoadInteractions string `json:"LoadInteractions"` // LoadInteractions from config file
		ClientSecret     string `json:"ClientSecret"`     // ClientSecret from config file
	}
	// Database stuff from remilia.conf
	Database struct {
		Database string `json:"Database"` // DatabaseName from config file
		Host     string `json:"Host"`     // Host from config file
		User     string `json:"User"`     // User from config file
		Password string `json:"Password"` // Password from config file
	}
	// Web stuff from remilia.conf
	Web struct {
		WebPort string `json:"WebPort"` // WebURL from config file
		BaseURL string `json:"BaseURL"` // BaseURL from config file
		RunWeb  string `json:"RunWeb"`  // RunWeb from config file
	}
}

// ReadConfig reads config.json file, returns error
func ReadConfig() error {
	file, err := ioutil.ReadFile("./remilia.conf")

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	if _, err := toml.Decode(string(file), &config); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	Token = config.Discord.Token
	BotPrefix = config.Discord.BotPrefix
	Presence = config.Discord.Presence
	Owner = config.Discord.Owner
	Database = config.Database.Database
	Host = config.Database.Host
	User = config.Database.User
	Password = config.Database.Password
	LoadInteractions = config.Discord.LoadInteractions
	WebPort = config.Web.WebPort
	BaseURL = config.Web.BaseURL
	ClientSecret = config.Discord.ClientSecret
	RunWeb = config.Web.RunWeb

	return nil
}
