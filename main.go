package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"git.randomchars.net/Reviath/RemiliaScarlet/config"
	"git.randomchars.net/Reviath/RemiliaScarlet/executor"
	"git.randomchars.net/Reviath/RemiliaScarlet/web"
	"github.com/bwmarrin/discordgo"
)

func main() {
	config.ReadConfig()
	client, err := discordgo.New(fmt.Sprintf("Bot %s", config.Token))
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

	if config.LoadInteractions == "true" {
		executor.RunAllInteractions(client, BotID)
	}

	fmt.Printf("Logging in as %s#%s (%s)\n", BotUsername, BotDiscriminator, BotID)

	time.Sleep(1 * time.Second)

	go web.Listen()

	if err = client.Open(); err != nil {
		fmt.Printf("Opening the session failed: \"%s\".\n", err.Error())
		return
	}

	WaitForInterrupt()

	fmt.Println("Shutting down.")
	client.Close()
	os.Exit(1)
}

func WaitForInterrupt() {
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
}
