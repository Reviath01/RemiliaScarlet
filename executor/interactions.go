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
}
