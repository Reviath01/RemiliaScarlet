package client

import (
	"fmt"

	"git.randomchars.net/Reviath/RemiliaScarlet/commands"
	"git.randomchars.net/Reviath/RemiliaScarlet/events"
	CommandHandler "git.randomchars.net/Reviath/RemiliaScarlet/handler"
	"github.com/bwmarrin/discordgo"
)

func Start() {
	client, err := discordgo.New("Bot " + Token)
	if err != nil {
		fmt.Printf("Creating a session failed: %s.\n", err.Error())
		return
	}

	u, err := client.User("@me")

	if err != nil {
		fmt.Println(err.Error())
	}

	BotID := u.ID
	BotUsername := u.Username
	BotDiscriminator := u.Discriminator

	fmt.Println("Logging in as " + BotUsername + "#" + BotDiscriminator + " (" + BotID + ")")

	handler := CommandHandler.New([]string{BotPrefix}, []string{Owner}, true, true, client.StateEnabled)
	client.AddHandler(handler.MessageHandler)
	handler.AddCommand("ping", "Check the bot's ping.", []string{"pong"}, false, false, discordgo.PermissionSendMessages, discordgo.PermissionSendMessages, CommandHandler.CommandTypeEverywhere, commands.PingCommand)
	handler.AddCommand("autorole", "Set auto role.", []string{"auto_role", "otorol"}, false, false, discordgo.PermissionManageRoles, discordgo.PermissionAdministrator, CommandHandler.CommandTypeGuild, commands.AutoRoleCommand)
	handler.AddCommand("afk", "Sets you as AFK.", []string{"set_afk"}, false, false, discordgo.PermissionSendMessages, discordgo.PermissionSendMessages, CommandHandler.CommandTypeGuild, commands.AfkCommand)
	handler.AddCommand("author", "Check bot's author.", []string{}, false, false, discordgo.PermissionSendMessages, discordgo.PermissionSendMessages, CommandHandler.CommandTypeEverywhere, commands.AuthorCommand)
	handler.AddCommand("avatar", "Fetch the profile photo of the user.", []string{"pfp", "profile", "pp", "avatar"}, false, false, discordgo.PermissionSendMessages, discordgo.PermissionSendMessages, CommandHandler.CommandTypeEverywhere, commands.AvatarCommand)
	handler.AddCommand("ban", "Bans the user.", []string{"yasakla"}, false, false, discordgo.PermissionBanMembers, discordgo.PermissionBanMembers, CommandHandler.CommandTypeGuild, commands.BanCommand)
	handler.AddCommand("disable", "Disables specified command.", []string{}, false, false, discordgo.PermissionSendMessages, discordgo.PermissionAdministrator, CommandHandler.CommandTypeGuild, commands.DisableCommand)
	handler.AddCommand("embed", "Sends your message as an embed.", []string{"send_embed"}, false, false, discordgo.PermissionSendMessages, discordgo.PermissionSendMessages, CommandHandler.CommandTypeEverywhere, commands.EmbedCommand)
	handler.AddCommand("enable", "Enables specified command.", []string{"enable_command"}, false, false, discordgo.PermissionSendMessages, discordgo.PermissionAdministrator, CommandHandler.CommandTypeGuild, commands.EnableCommand)
	handler.AddCommand("guildinfo", "Sends info about guild.", []string{"guild_İnfo"}, false, false, discordgo.PermissionSendMessages, discordgo.PermissionSendMessages, CommandHandler.CommandTypeGuild, commands.GuildInfoCommand)
	handler.AddCommand("hug", "Sends hug gif.", []string{"sarıl"}, false, false, discordgo.PermissionSendMessages, discordgo.PermissionSendMessages, CommandHandler.CommandTypeGuild, commands.HugCommand)
	handler.AddCommand("icon", "Fetch server icon.", []string{}, false, false, discordgo.PermissionSendMessages, discordgo.PermissionSendMessages, CommandHandler.CommandTypeGuild, commands.IconCommand)
	handler.AddCommand("invite", "Invite the bot.", []string{"davet"}, false, false, discordgo.PermissionSendMessages, discordgo.PermissionSendMessages, CommandHandler.CommandTypeEverywhere, commands.InviteCommand)
	handler.AddCommand("issue", "Create an issue.", []string{}, false, false, discordgo.PermissionSendMessages, discordgo.PermissionSendMessages, CommandHandler.CommandTypeEverywhere, commands.IssueCommand)
	handler.AddCommand("kick", "Kicks specified user.", []string{}, false, false, discordgo.PermissionSendMessages, discordgo.PermissionKickMembers, CommandHandler.CommandTypeGuild, commands.KickCommand)
	handler.SetHelpCommand("help", []string{}, discordgo.PermissionSendMessages, discordgo.PermissionSendMessages, commands.HelpCommand)

	handler.SetPrerunFunc(func(context CommandHandler.Context, command *CommandHandler.Command, content []string) bool {
		fmt.Printf("Command \"%s\" is being run by \"%s#%s\" (ID: %s).\n", command.Name, context.User.Username, context.User.Discriminator, context.User.ID)
		return true
	})

	if err = client.Open(); err != nil {
		fmt.Printf("Opening the session failed: \"%s\".\n", err.Error())
		return
	}

	client.AddHandler(events.ChannelCreate)
	client.AddHandler(events.ChannelDelete)
	client.AddHandler(events.GuildMemberAdd)
	client.AddHandler(events.GuildMemberRemove)
	client.AddHandler(events.GuildRoleCreate)
	client.AddHandler(events.GuildRoleDelete)

	CommandHandler.WaitForInterrupt()

	fmt.Println("Shutting down.")
	if err := client.Close(); err != nil {
		fmt.Printf("Closing the session failed. \"%s\". Ignoring.\n", err.Error())
	}
}
