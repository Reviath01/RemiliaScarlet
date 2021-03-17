package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

var (
	Token     string
	config *configStruct
)

type configStruct struct {
	Token     string `json:"Token"`
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

	return nil
}
