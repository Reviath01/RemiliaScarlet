package client

import (
	"fmt"

	"../config"
	"github.com/bwmarrin/discordgo"
	Handler "git.randomchars.net/Reviath/handlers"
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

	goBot.AddHandler(Handler.MessageHandler)

	BotID = u.ID
	BotUsername = u.Username
	BotDiscriminator = u.Discriminator

	err = goBot.Open()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Logging in as " + BotUsername + "#" + BotDiscriminator + " (" + BotID + ")")
}
