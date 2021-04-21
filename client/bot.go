package client

import (
	"fmt"
	"time"

	"git.randomchars.net/Reviath/RemiliaScarlet/executor"
	CommandHandler "git.randomchars.net/Reviath/RemiliaScarlet/handler"
	"github.com/bwmarrin/discordgo"
)

func Start() {
	client, err := discordgo.New(fmt.Sprintf("Bot %s", Token))
	if err != nil {
		fmt.Printf("Creating a session failed: %s.\n", err.Error())
		return
	}

	u, err := client.User("@me")

	if err != nil {
		fmt.Printf("Fetching client user failed: %s.\n", err.Error())
	}

	BotID := u.ID
	BotUsername := u.Username
	BotDiscriminator := u.Discriminator

	executor.RunAllCommands(client)
	executor.RunAllEvents(client)
	executor.RunAllInteractions(client, BotID)

	time.Sleep(1 * time.Second)

	fmt.Printf("Logging in as %s#%s (%s)\n", BotUsername, BotDiscriminator, BotID)

	if err = client.Open(); err != nil {
		fmt.Printf("Opening the session failed: \"%s\".\n", err.Error())
		return
	}

	CommandHandler.WaitForInterrupt()

	fmt.Println("Shutting down.")
	if err := client.Close(); err != nil {
		fmt.Printf("Closing the session failed. \"%s\". Ignoring.\n", err.Error())
	}
}
