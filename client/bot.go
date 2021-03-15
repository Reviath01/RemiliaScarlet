package client

import (
	"fmt"

	"../config"
	"github.com/bwmarrin/discordgo"
	Handler "git.randomchars.net/Reviath/remilia-scarlet-command-handler"
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

	handler := Handler.New([]string{config.BotPrefix}, []string{config.Owner, "808393372448325672"}, true, true, goBot.StateEnabled)

	goBot.AddHandler(handler.MessageHandler)

	handler.SetOnErrorFunc(func(context Handler.Context, command *Handler.Command, content []string, err error) {
		fmt.Printf("An error occurred for command \"%s\": \"%s\".\n", command.Name, err.Error())
	})

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

	fmt.Println("Logging in as " + BotUsername + "#" + BotDiscriminator + " (" + BotID + ")")
}
