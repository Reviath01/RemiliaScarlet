package executor

import (
	"fmt"

	"git.randomchars.net/Reviath/RemiliaScarlet/interactions"
	"github.com/bwmarrin/discordgo"
)

//Creates all interactions.
func RunAllInteractions(client *discordgo.Session, BotID string) {
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

	pingcommand := interactions.Command{
		Name:        "ping",
		Description: "Fetch bots ping.",
	}

	err = interactions.GlobalCommandCreate(client, BotID, pingcommand)
	if err != nil {
		print(err.Error())
	}

	reset_autorolecommand := interactions.Command{
		Name:        "reset_autorole",
		Description: "Reset auto role.",
	}

	err = interactions.GlobalCommandCreate(client, BotID, reset_autorolecommand)
	if err != nil {
		print(err.Error())
	}

	reset_leavechannelcommand := interactions.Command{
		Name:        "reset_leave_channel",
		Description: "Reset leave channel.",
	}

	err = interactions.GlobalCommandCreate(client, BotID, reset_leavechannelcommand)
	if err != nil {
		print(err.Error())
	}

	reset_leavemessagecommand := interactions.Command{
		Name:        "reset_leave_message",
		Description: "Reset leave message.",
	}

	err = interactions.GlobalCommandCreate(client, BotID, reset_leavemessagecommand)
	if err != nil {
		print(err.Error())
	}

	reset_logcommand := interactions.Command{
		Name:        "reset_log",
		Description: "Reset log channel.",
	}

	err = interactions.GlobalCommandCreate(client, BotID, reset_logcommand)
	if err != nil {
		print(err.Error())
	}

	reset_welcome_message_command := interactions.Command{
		Name:        "reset_welcome_message",
		Description: "Reset welcome message.",
	}

	err = interactions.GlobalCommandCreate(client, BotID, reset_welcome_message_command)
	if err != nil {
		print(err.Error())
	}

	reset_welcome_channel_command := interactions.Command{
		Name:        "reset_welcome_channel",
		Description: "Reset welcome channel.",
	}

	err = interactions.GlobalCommandCreate(client, BotID, reset_welcome_channel_command)
	if err != nil {
		print(err.Error())
	}

	roles_command := interactions.Command{
		Name:        "roles",
		Description: "Fetch all roles at your guild.",
	}

	err = interactions.GlobalCommandCreate(client, BotID, roles_command)
	if err != nil {
		print(err.Error())
	}

	settings_command := interactions.Command{
		Name:        "settings",
		Description: "Get server settings.",
	}

	err = interactions.GlobalCommandCreate(client, BotID, settings_command)
	if err != nil {
		print(err.Error())
	}

	slap_command := interactions.Command{
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

	err = interactions.GlobalCommandCreate(client, BotID, slap_command)
	if err != nil {
		print(err.Error())
	}

	spoiler_command := interactions.Command{
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

	err = interactions.GlobalCommandCreate(client, BotID, spoiler_command)
	if err != nil {
		print(err.Error())
	}

	stats_command := interactions.Command{
		Name:        "stats",
		Description: "Fetch bot's stats.",
	}

	err = interactions.GlobalCommandCreate(client, BotID, stats_command)
	if err != nil {
		print(err.Error())
	}

	unban_command := interactions.Command{
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

	err = interactions.GlobalCommandCreate(client, BotID, unban_command)
	if err != nil {
		print(err.Error())
	}
}
