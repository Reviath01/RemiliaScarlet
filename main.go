package main

import (
	"fmt"
	"time"

	"../config"
	Handler "git.randomchars.net/Reviath/remilia-scarlet-command-handler"
	"github.com/bwmarrin/discordgo"
)

func main() {
	fmt.Println("Starting!")

	client, err := discordgo.New("Bot " + config.Token)
	if err != nil {
		fmt.Printf("Creating a session failed: \"%s\".\n", err.Error())
		return
	}

	handler := Handler.New([]string{config.BotPrefix}, []string{config.Owner}, true, true, client.StateEnabled)
	client.AddHandler(handler.MessageHandler)

	handler.AddCommand("ping", "Check the bot's ping.", []string{"pong"}, false, false, discordgo.PermissionSendMessages, discordgo.PermissionSendMessages, Handler.CommandTypeEverywhere, pingCommand)

	handler.SetHelpCommand("help", []string{}, discordgo.PermissionSendMessages, discordgo.PermissionSendMessages, helpCommand)

	handler.SetOnErrorFunc(func(context Handler.Context, command *Handler.Command, content []string, err error) {
		fmt.Printf("An error occurred for command \"%s\": \"%s\".\n", command.Name, err.Error())
	})

	handler.SetPrerunFunc(func(context Handler.Context, command *Handler.Command, content []string) bool {
		return true
	})

	if err = client.Open(); err != nil {
		fmt.Printf("Opening the session failed: \"%s\".\n", err.Error())
		return
	}

	Handler.WaitForInterrupt()

	fmt.Println("Shutting down.")
	if err := client.Close(); err != nil {
		fmt.Printf("Closing the session failed. \"%s\". Ignoring.\n", err.Error())
	}
}

func pingCommand(ctx Handler.Context, _ []string) error {
	msg, err := ctx.Reply("Pong!")
	if err != nil {
		return err
	}

	sent, err := msg.Timestamp.Parse()
	if err != nil {
		return err
	}

	_, err = ctx.Session.ChannelMessageEdit(ctx.Message.ChannelID, msg.ID, fmt.Sprintf("Pong! Ping took **%dms**!", time.Now().Sub(sent).Milliseconds()))
	return err
}
