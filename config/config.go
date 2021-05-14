package config

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
	WebURL           string
	ClientID         string
	ClientSecret     string
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

//Reads config file.
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
