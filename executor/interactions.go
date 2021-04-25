package executor

import (
	"fmt"
	"time"

	"git.randomchars.net/Reviath/RemiliaScarlet/interactions"
	"github.com/bwmarrin/discordgo"
)

func RunAllInteractions(client *discordgo.Session, BotID string) {
	time.Sleep(2 * time.Second)
	fmt.Print("Loading interactions. \n")
	invitecommand := interactions.Command{
		Name:        "invite",
		Description: "Invite the bot.",
	}

	err := interactions.GlobalCommandCreate(client, BotID, invitecommand)
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
	hugcommand := interactions.Command{
		Name:        "hug",
		Description: "Sends hug gif.",
		Options: []interactions.CommandOption{
			{
				Type:        interactions.OptionTypeString,
				Name:        "User",
				Required:    true,
				Description: "Specify a user to hug.",
			},
		},
	}

	err = interactions.GlobalCommandCreate(client, BotID, hugcommand)
	if err != nil {
		print(err.Error())
	}

	iconcommand := interactions.Command{
		Name:        "icon",
		Description: "Sends server icon.",
	}

	err = interactions.GlobalCommandCreate(client, BotID, iconcommand)
	if err != nil {
		print(err.Error())
	}

	guildinfocommand := interactions.Command{
		Name:        "guild_info",
		Description: "Get information about guild.",
	}

	err = interactions.GlobalCommandCreate(client, BotID, guildinfocommand)
	if err != nil {
		print(err.Error())
	}

	issuecommand := interactions.Command{
		Name:        "issue",
		Description: "Create an issue.",
	}

	err = interactions.GlobalCommandCreate(client, BotID, issuecommand)
	if err != nil {
		print(err.Error())
	}

	kickcommand := interactions.Command{
		Name:        "kick",
		Description: "Kick a user.",
		Options: []interactions.CommandOption{
			{
				Type:        interactions.OptionTypeString,
				Name:        "User",
				Required:    true,
				Description: "Specify a user to kick.",
			},
		},
	}

	err = interactions.GlobalCommandCreate(client, BotID, kickcommand)
	if err != nil {
		print(err.Error())
	}

	kisscommand := interactions.Command{
		Name:        "kiss",
		Description: "Sends kiss gif.",
		Options: []interactions.CommandOption{
			{
				Type:        interactions.OptionTypeString,
				Name:        "User",
				Required:    true,
				Description: "Specify a user to kiss.",
			},
		},
	}

	err = interactions.GlobalCommandCreate(client, BotID, kisscommand)
	if err != nil {
		print(err.Error())
	}

	languagecommand := interactions.Command{
		Name:        "language",
		Description: "Set bot language.",
		Options: []interactions.CommandOption{
			{
				Type:        interactions.OptionTypeString,
				Name:        "Language",
				Required:    true,
				Description: "New language of bot.",
			},
		},
	}

	err = interactions.GlobalCommandCreate(client, BotID, languagecommand)
	if err != nil {
		print(err.Error())
	}

	leavechannelcommand := interactions.Command{
		Name:        "leave_channel",
		Description: "Set leave channel.",
		Options: []interactions.CommandOption{
			{
				Type:        interactions.OptionTypeString,
				Name:        "Channel",
				Required:    true,
				Description: "Specify a channel.",
			},
		},
	}

	err = interactions.GlobalCommandCreate(client, BotID, leavechannelcommand)
	if err != nil {
		print(err.Error())
	}

	leavemessagecommand := interactions.Command{
		Name:        "leave_message",
		Description: "Set leave message.",
		Options: []interactions.CommandOption{
			{
				Type:        interactions.OptionTypeString,
				Name:        "Message",
				Required:    true,
				Description: "Specify a message.",
			},
		},
	}

	err = interactions.GlobalCommandCreate(client, BotID, leavemessagecommand)
	if err != nil {
		print(err.Error())
	}
	logcommand := interactions.Command{
		Name:        "log",
		Description: "Set log channel.",
		Options: []interactions.CommandOption{
			{
				Type:        interactions.OptionTypeString,
				Name:        "Channel",
				Required:    true,
				Description: "Specify a channel.",
			},
		},
	}

	err = interactions.GlobalCommandCreate(client, BotID, logcommand)
	if err != nil {
		print(err.Error())
	}
}
