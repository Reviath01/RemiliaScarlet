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

	handler := Handler.New([]string{config.BotPrefix}, []string{config.Owner, "808393372448325672"}, true, true, client.StateEnabled)

	goBot.AddHandler(handler.MessageHandler)

	handler.SetOnErrorFunc(func(context Handler.Context, command *Handler.Command, content []string, err error) {
		fmt.Printf("An error occurred for command \"%s\": \"%s\".\n", command.Name, err.Error())
	})

	handler.SetHelpCommand("help", []string{}, discordgo.PermissionSendMessages, discordgo.PermissionSendMessages, helpCommand)

	handler.AddCommand("ping", "Check the bot's ping.", []string{"pong"}, false, false, discordgo.PermissionSendMessages, discordgo.PermissionSendMessages, Handler.CommandTypeEverywhere, pingCommand)


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
