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
	if config.Token == "TOKEN" || config.ClientSecret == "CLIENT_SECRET" || config.WebPort == "WEB_PORT" || config.BaseURL == "BASE_URL" {
		fmt.Println("Edit config file before running.")
		os.Exit(1)
	}

	client, err := discordgo.New(fmt.Sprintf("Bot %s", config.Token))
	if err != nil {
		fmt.Printf("Creating a session failed: %s.\n", err.Error())
		return
	}

	u, err := client.User("@me")

	if err != nil {
		fmt.Printf("Fetching client user failed: %s.\n", err.Error())
	}

	client.Identify.Intents = discordgo.IntentsAll

	if config.LoadInteractions == "true" {
		executor.RunAllInteractions(client, u.ID)
	}

	executor.RunAllEvents(client)
	executor.RunAllCommands(client)

	fmt.Printf("Logging in as %s#%s (%s)\n", u.Username, u.Discriminator, u.ID)

	time.Sleep(1 * time.Second)

	if config.RunWeb == "true" {
		go web.Listen(client)
	}

	if err = client.Open(); err != nil {
		fmt.Printf("Opening the session failed: \"%s\".\n", err.Error())
		return
	}

	WaitForInterrupt()

	fmt.Println("Shutting down.")
	client.Close()
	os.Exit(1)
}

// WaitForInterrupt is a function that notifies when program is closed
func WaitForInterrupt() {
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
}
