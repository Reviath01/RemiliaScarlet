package starter

import (
    "../config"
    "fmt"
)

var prefix string
var owner string
var ignoreBots string

func Run() CommandHandler {
	prefix = config.BotPrefix
	owner = config.Owner
	ignoreBots = config.IgnoreBots

	return CommandHandler{
		enabled:          true,
		prefixes:         prefix,
		owners:           owner,
		ignoreBots:       ignoreBots,
	}
}
