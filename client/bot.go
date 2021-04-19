package client

import (
	"fmt"

	"git.randomchars.net/Reviath/RemiliaScarlet/command_executer"
	"git.randomchars.net/Reviath/RemiliaScarlet/events"
	CommandHandler "git.randomchars.net/Reviath/RemiliaScarlet/handler"
	"git.randomchars.net/Reviath/RemiliaScarlet/interactions"
	interactions_handler "git.randomchars.net/Reviath/RemiliaScarlet/interactions/handler"
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

	fmt.Printf("Logging in as %s#%s (%s)\n", BotUsername, BotDiscriminator, BotID)

	command_executer.RunAllCommands(client)

	client.AddHandler(events.ChannelCreate)
	client.AddHandler(events.ChannelDelete)
	client.AddHandler(events.GuildMemberAdd)
	client.AddHandler(events.GuildMemberRemove)
	client.AddHandler(events.GuildRoleCreate)
	client.AddHandler(events.GuildRoleDelete)
	client.AddHandler(events.Ready)
	client.AddHandler(interactions_handler.InteractionHandler)

	invitecommand := interactions.Command{
		Name:        "invite",
		Description: "Invite the bot.",
	}

	err = interactions.GlobalCommandCreate(client, BotID, invitecommand)
	if err != nil {
		return
	}

	authorcommand := interactions.Command{
		Name:        "author",
		Description: "Check bot's author.",
	}

	err = interactions.GlobalCommandCreate(client, BotID, authorcommand)
	if err != nil {
		return
	}

	autorolecommand := interactions.Command{
		Name:        "autorole",
		Description: "Set auto role.",
		Options: []interactions.CommandOption{
			{
				Type:        interactions.OptionTypeString,
				Name:        "Role",
				Required:    true,
				Description: "This is the role you want to set as auto role",
			},
		},
	}

	err = interactions.GlobalCommandCreate(client, BotID, autorolecommand)
	if err != nil {
		return
	}

	disablecommand := interactions.Command{
		Name:        "disable",
		Description: "Disable a command.",
		Options: []interactions.CommandOption{
			{
				Type:        interactions.OptionTypeString,
				Name:        "Command",
				Required:    true,
				Description: "Specify a command to disable.",
			},
		},
	}

	err = interactions.GlobalCommandCreate(client, BotID, disablecommand)
	if err != nil {
		return
	}

	avatarcommand := interactions.Command{
		Name:        "avatar",
		Description: "Fetch a user's avatar.",
		Options: []interactions.CommandOption{
			{
				Type:        interactions.OptionTypeString,
				Name:        "User",
				Required:    true,
				Description: "Mention a user or give an ID",
			},
		},
	}

	err = interactions.GlobalCommandCreate(client, BotID, avatarcommand)
	if err != nil {
		print(err.Error())
	}

	embedcommand := interactions.Command{
		Name:        "embed",
		Description: "Sends your message as an embed.",
		Options: []interactions.CommandOption{
			{
				Type:        interactions.OptionTypeString,
				Name:        "Message",
				Required:    true,
				Description: "Specify a message to send as an embed.",
			},
		},
	}

	err = interactions.GlobalCommandCreate(client, BotID, embedcommand)
	if err != nil {
		print(err.Error())
	}

	bancommand := interactions.Command{
		Name:        "ban",
		Description: "Ban a user.",
		Options: []interactions.CommandOption{
			{
				Type:        interactions.OptionTypeString,
				Name:        "User",
				Required:    true,
				Description: "Specify a user to ban.",
			},
		},
	}

	err = interactions.GlobalCommandCreate(client, BotID, bancommand)
	if err != nil {
		print(err.Error())
	}

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
