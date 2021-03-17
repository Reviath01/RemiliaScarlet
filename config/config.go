package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

var (
	Token     string
	Prefix1 string
	Prefix2 string
	config *configStruct
)

type configStruct struct {
	Token     string `json:"Token"`
	Prefix1 string `json:"Prefix1"`
	Prefix2 string `json:"Prefix2"`
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
	Prefix1 = config.Prefix1
	Prefix2 = config.Prefix2

	return nil
}
