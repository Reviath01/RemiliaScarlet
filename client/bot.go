package main

import (
	"fmt"
	"time"

	"../config"
	Handler "git.randomchars.net/Reviath/remilia-scarlet-command-handler"
	"github.com/bwmarrin/discordgo"
	embed "git.randomchars.net/Reviath/remilia-scarlet-embed-util"
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

	handler.AddCommand("ping", "Check the bot's ping.", []string{"pong"}, false, false, discordgo.PermissionSendMessages, discordgo.PermissionSendMessages, Handler.CommandTypeEverywhere, pingCommand) //Adding ping command to bot.

	handler.AddCommand("hug", "Sends hug gif.", []string{"hug_gif"}, false, false, discordgo.PermissionSendMessages, discordgo.PermissionSendMessages, Handler.CommandTypeEverywhere, HugCommand) //Adding hug command to bot.

	handler.SetOnErrorFunc(func(context Handler.Context, command *Handler.Command, content []string, err error) {
		fmt.Printf("An error occurred for command \"%s\": \"%s\".\n", command.Name, err.Error())
	}) //If error on a command, print the error.

	handler.SetPrerunFunc(func(context Handler.Context, command *Handler.Command, content []string) bool {
		return true
	})

	if err = client.Open(); err != nil { //If there is an error, logging the error
		fmt.Printf("Opening the session failed: \"%s\".\n", err.Error())
		return
	}

	Handler.WaitForInterrupt()

	fmt.Println("Shutting down.")
	if err := client.Close(); err != nil { //If there is an error, logging the error
		fmt.Printf("Closing the session failed. \"%s\". Ignoring.\n", err.Error())
	}
}

//Ping commands code
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

//Hug commands code
func HugCommand(ctx Handler.Context _ []string) error {
	msg, err := ctx.ReplyEmbed(embed.ImageEmbed("https://images-ext-1.discordapp.net/external/xT0jPvqo3SiH7lJG88IXCgxbrmd36Rt6-Hgvu5aw8Rg/https/i.pinimg.com/originals/4d/d7/49/4dd749423de10a319b5d9e8850bbace4.gif?width=596&height=596"))
	if err != nil {
		return err
	}
}
