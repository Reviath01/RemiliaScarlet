package multiplexer

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

var (
	BotPrefix string        // BotPrefix is bot's prefix to cut prefix's length
	config    *configStruct // config structure
)

type configStruct struct {
	BotPrefix string `json:"BotPrefix"`
}

// GetPrefix returns to Bot's prefix from config.json file
func GetPrefix() string {
	file, err := ioutil.ReadFile("./config.json")

	if err != nil {
		fmt.Println(err.Error())
		return "err"
	}

	err = json.Unmarshal(file, &config)

	if err != nil {
		fmt.Println(err.Error())
		return "err"
	}

	BotPrefix = config.BotPrefix

	return BotPrefix
}
