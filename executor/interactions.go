package executor

import (
	"fmt"

	"git.randomchars.net/Reviath/RemiliaScarlet/interactions"
	"github.com/bwmarrin/discordgo"
)

// RunAllInteractions creates global interaction commands for bot
func RunAllInteractions(client *discordgo.Session, BotID string) {
	fmt.Print("Loading interactions. \n")
	invitecommand := interactions.Command{
		Name:        "invite",
		Description: "Invite the bot.",
	}

	interactions.GlobalCommandCreate(client, BotID, invitecommand)

	authorcommand := interactions.Command{
		Name:        "author",
		Description: "Check bot's author.",
	}

	interactions.GlobalCommandCreate(client, BotID, authorcommand)

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

	interactions.GlobalCommandCreate(client, BotID, autorolecommand)

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

	interactions.GlobalCommandCreate(client, BotID, disablecommand)

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

	interactions.GlobalCommandCreate(client, BotID, avatarcommand)

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

	interactions.GlobalCommandCreate(client, BotID, embedcommand)

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

	interactions.GlobalCommandCreate(client, BotID, bancommand)
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

	interactions.GlobalCommandCreate(client, BotID, hugcommand)

	iconcommand := interactions.Command{
		Name:        "icon",
		Description: "Sends server icon.",
	}

	interactions.GlobalCommandCreate(client, BotID, iconcommand)

	guildinfocommand := interactions.Command{
		Name:        "guild_info",
		Description: "Get information about guild.",
	}

	interactions.GlobalCommandCreate(client, BotID, guildinfocommand)

	issuecommand := interactions.Command{
		Name:        "issue",
		Description: "Create an issue.",
	}

	interactions.GlobalCommandCreate(client, BotID, issuecommand)

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

	interactions.GlobalCommandCreate(client, BotID, kickcommand)

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

	interactions.GlobalCommandCreate(client, BotID, kisscommand)

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

	interactions.GlobalCommandCreate(client, BotID, languagecommand)

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

	interactions.GlobalCommandCreate(client, BotID, leavechannelcommand)

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

	interactions.GlobalCommandCreate(client, BotID, leavemessagecommand)

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

	interactions.GlobalCommandCreate(client, BotID, logcommand)

	pingcommand := interactions.Command{
		Name:        "ping",
		Description: "Fetch bots ping.",
	}

	interactions.GlobalCommandCreate(client, BotID, pingcommand)

	resetautorolecommand := interactions.Command{
		Name:        "reset_autorole",
		Description: "Reset auto role.",
	}

	interactions.GlobalCommandCreate(client, BotID, resetautorolecommand)

	resetleavechannelcommand := interactions.Command{
		Name:        "reset_leave_channel",
		Description: "Reset leave channel.",
	}

	interactions.GlobalCommandCreate(client, BotID, resetleavechannelcommand)

	resetleavemessagecommand := interactions.Command{
		Name:        "reset_leave_message",
		Description: "Reset leave message.",
	}

	interactions.GlobalCommandCreate(client, BotID, resetleavemessagecommand)

	resetlogcommand := interactions.Command{
		Name:        "reset_log",
		Description: "Reset log channel.",
	}

	interactions.GlobalCommandCreate(client, BotID, resetlogcommand)

	resetwelcomemessagecommand := interactions.Command{
		Name:        "reset_welcome_message",
		Description: "Reset welcome message.",
	}

	interactions.GlobalCommandCreate(client, BotID, resetwelcomemessagecommand)

	resetwelcomechannelcommand := interactions.Command{
		Name:        "reset_welcome_channel",
		Description: "Reset welcome channel.",
	}

	interactions.GlobalCommandCreate(client, BotID, resetwelcomechannelcommand)

	rolescommand := interactions.Command{
		Name:        "roles",
		Description: "Fetch all roles at your guild.",
	}

	interactions.GlobalCommandCreate(client, BotID, rolescommand)

	settingscommand := interactions.Command{
		Name:        "settings",
		Description: "Get server settings.",
	}

	interactions.GlobalCommandCreate(client, BotID, settingscommand)

	slapcommand := interactions.Command{
		Name:        "slap",
		Description: "Sends slap gif.",
		Options: []interactions.CommandOption{
			{
				Type:        interactions.OptionTypeString,
				Name:        "User",
				Required:    true,
				Description: "Specify a user.",
			},
		},
	}

	interactions.GlobalCommandCreate(client, BotID, slapcommand)

	spoilercommand := interactions.Command{
		Name:        "spoiler",
		Description: "Sends your message as spoiler.",
		Options: []interactions.CommandOption{
			{
				Type:        interactions.OptionTypeString,
				Name:        "Message",
				Required:    true,
				Description: "Specify a message.",
			},
		},
	}

	interactions.GlobalCommandCreate(client, BotID, spoilercommand)

	statscommand := interactions.Command{
		Name:        "stats",
		Description: "Fetch bot's stats.",
	}

	interactions.GlobalCommandCreate(client, BotID, statscommand)

	unbancommand := interactions.Command{
		Name:        "unban",
		Description: "Unban a user.",
		Options: []interactions.CommandOption{
			{
				Type:        interactions.OptionTypeString,
				Name:        "User",
				Required:    true,
				Description: "Specify a user.",
			},
		},
	}

	interactions.GlobalCommandCreate(client, BotID, unbancommand)
}
