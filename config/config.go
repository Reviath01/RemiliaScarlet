package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"flag"
)

var (
	Token     string
	BotPrefix string
	Owner string
	IgnoreBots string
	config *configStruct
	VersionStartup bool
)

type configStruct struct {
	Token     string `json:"Token"`
	BotPrefix string `json:"BotPrefix"`
	Owner string `json:"Owner"`
}

func flags() *types.Nil {
	flag.BoolVar(&VersionStartup, "v", false, "Display Version information and exit")
	flag.Parse()
	return nil
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

	return nil
}
