package client

import (
	"fmt"

	"../starter"
	"../config"
	"github.com/bwmarrin/discordgo"
)

var BotUsername string
var BotDiscriminator string
var BotID string
var goBot *discordgo.Session

func Start() {
	goBot, err := discordgo.New("Bot " + config.Token)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	u, err := goBot.User("@me")

	if err != nil {
		fmt.Println(err.Error())
	}

	BotID = u.ID
	BotUsername = u.Username
	BotDiscriminator = u.Discriminator
	err = goBot.Open()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	starter.Run()

	fmt.Println("Logging in as " + BotUsername + "#" + BotDiscriminator + " (" + BotID + ")")
}
