package executor

import (
	"fmt"

	"git.randomchars.net/Reviath/RemiliaScarlet/commands"
	"git.randomchars.net/Reviath/RemiliaScarlet/config"
	CommandHandler "git.randomchars.net/Reviath/RemiliaScarlet/handler"
	"github.com/bwmarrin/discordgo"
)

// RunAllCommands runs all commands and creates a CommandHandler
func RunAllCommands(client *discordgo.Session) {
	config.ReadConfig()
	handler := CommandHandler.New([]string{config.BotPrefix}, []string{config.Owner}, true, true, client.StateEnabled)
	client.AddHandler(handler.MessageHandler)
	handler.AddCommand("ping", "Check the bot's ping.", []string{"pong"}, false, false, discordgo.PermissionSendMessages, discordgo.PermissionSendMessages, CommandHandler.CommandTypeEverywhere, commands.PingCommand)
	handler.AddCommand("add_prefix", "Adds prefix to bot. (Owner-only).", []string{"addprefix", "add-prefix"}, true, false, discordgo.PermissionSendMessages, discordgo.PermissionSendMessages, CommandHandler.CommandTypeGuild, commands.AddPrefixCommand)
	handler.AddCommand("remove_prefix", "Removes prefix to bot. (Owner-only).", []string{"removeprefix", "remove-prefix"}, true, false, discordgo.PermissionSendMessages, discordgo.PermissionSendMessages, CommandHandler.CommandTypeGuild, commands.RemovePrefixCommand)
	handler.AddCommand("autorole", "Set auto role.", []string{"auto_role", "otorol"}, false, false, discordgo.PermissionSendMessages, discordgo.PermissionSendMessages, CommandHandler.CommandTypeGuild, commands.AutoRoleCommand)
	handler.AddCommand("author", "Check bot's author.", []string{}, false, false, discordgo.PermissionSendMessages, discordgo.PermissionSendMessages, CommandHandler.CommandTypeEverywhere, commands.AuthorCommand)
	handler.AddCommand("avatar", "Fetch the profile photo of the user.", []string{"pfp", "profile", "pp", "avatar"}, false, false, discordgo.PermissionSendMessages, discordgo.PermissionSendMessages, CommandHandler.CommandTypeEverywhere, commands.AvatarCommand)
	handler.AddCommand("ban", "Bans the user.", []string{"yasakla"}, false, false, discordgo.PermissionSendMessages, discordgo.PermissionSendMessages, CommandHandler.CommandTypeGuild, commands.BanCommand)
	handler.AddCommand("disable", "Disables specified command.", []string{}, false, false, discordgo.PermissionSendMessages, discordgo.PermissionSendMessages, CommandHandler.CommandTypeGuild, commands.DisableCommand)
	handler.AddCommand("embed", "Sends your message as an embed.", []string{"send_embed"}, false, false, discordgo.PermissionSendMessages, discordgo.PermissionSendMessages, CommandHandler.CommandTypeEverywhere, commands.EmbedCommand)
	handler.AddCommand("enable", "Enables specified command.", []string{"enable_command"}, false, false, discordgo.PermissionSendMessages, discordgo.PermissionSendMessages, CommandHandler.CommandTypeGuild, commands.EnableCommand)
	handler.AddCommand("guildinfo", "Sends info about guild.", []string{"guild_İnfo"}, false, false, discordgo.PermissionSendMessages, discordgo.PermissionSendMessages, CommandHandler.CommandTypeGuild, commands.GuildInfoCommand)
	handler.AddCommand("hug", "Sends hug gif.", []string{"sarıl"}, false, false, discordgo.PermissionSendMessages, discordgo.PermissionSendMessages, CommandHandler.CommandTypeGuild, commands.HugCommand)
	handler.AddCommand("icon", "Fetch server icon.", []string{}, false, false, discordgo.PermissionSendMessages, discordgo.PermissionSendMessages, CommandHandler.CommandTypeGuild, commands.IconCommand)
	handler.AddCommand("invite", "Invite the bot.", []string{"davet"}, false, false, discordgo.PermissionSendMessages, discordgo.PermissionSendMessages, CommandHandler.CommandTypeEverywhere, commands.InviteCommand)
	handler.AddCommand("issue", "Create an issue.", []string{}, false, false, discordgo.PermissionSendMessages, discordgo.PermissionSendMessages, CommandHandler.CommandTypeEverywhere, commands.IssueCommand)
	handler.AddCommand("kick", "Kicks specified user.", []string{}, false, false, discordgo.PermissionSendMessages, discordgo.PermissionSendMessages, CommandHandler.CommandTypeGuild, commands.KickCommand)
	handler.AddCommand("kiss", "Sends kiss gif.", []string{}, false, false, discordgo.PermissionSendMessages, discordgo.PermissionSendMessages, CommandHandler.CommandTypeGuild, commands.KissCommand)
	handler.AddCommand("language", "Change bot language.", []string{}, false, false, discordgo.PermissionSendMessages, discordgo.PermissionSendMessages, CommandHandler.CommandTypeGuild, commands.LanguageCommand)
	handler.AddCommand("leave_channel", "Set leave channel.", []string{"leavechannel", "leave-channel"}, false, false, discordgo.PermissionSendMessages, discordgo.PermissionSendMessages, CommandHandler.CommandTypeGuild, commands.LeaveChannelCommand)
	handler.AddCommand("leave_message", "Set leave message. \nPlaceholders: `{mention}`, `{username}`", []string{"leavemessage", "leave-message"}, false, false, discordgo.PermissionSendMessages, discordgo.PermissionSendMessages, CommandHandler.CommandTypeGuild, commands.LeaveMessageCommand)
	handler.AddCommand("log", "Set log channel.", []string{}, false, false, discordgo.PermissionSendMessages, discordgo.PermissionSendMessages, CommandHandler.CommandTypeGuild, commands.LogCommand)
	handler.AddCommand("reset_auto_role", "Reset auto role.", []string{"reset-auto-role", "resetautorole"}, false, false, discordgo.PermissionSendMessages, discordgo.PermissionSendMessages, CommandHandler.CommandTypeGuild, commands.ResetAutoroleCommand)
	handler.AddCommand("reset_leave_channel", "Reset leave channel.", []string{"reset-leave-channel"}, false, false, discordgo.PermissionSendMessages, discordgo.PermissionSendMessages, CommandHandler.CommandTypeGuild, commands.ResetLeaveChannel)
	handler.AddCommand("reset_leave_message", "Reset leave message.", []string{"reset-leave-message"}, false, false, discordgo.PermissionSendMessages, discordgo.PermissionSendMessages, CommandHandler.CommandTypeGuild, commands.ResetLeaveMessage)
	handler.AddCommand("reset_log", "Reset log channel.", []string{"reset-log"}, false, false, discordgo.PermissionSendMessages, discordgo.PermissionSendMessages, CommandHandler.CommandTypeGuild, commands.ResetLogCommand)
	handler.AddCommand("reset_welcome_channel", "Reset welcome channel.", []string{"reset-welcome-channel"}, false, false, discordgo.PermissionSendMessages, discordgo.PermissionSendMessages, CommandHandler.CommandTypeGuild, commands.ResetWelcomeChannelCommand)
	handler.AddCommand("reset_welcome_message", "Reset welcome message.", []string{"reset-welcome-message"}, false, false, discordgo.PermissionSendMessages, discordgo.PermissionSendMessages, CommandHandler.CommandTypeGuild, commands.ResetWelcomeMessageCommand)
	handler.AddCommand("roles", "Get all roles in guild.", []string{"roller"}, false, false, discordgo.PermissionSendMessages, discordgo.PermissionSendMessages, CommandHandler.CommandTypeGuild, commands.RolesCommand)
	handler.AddCommand("settings", "Shows all settings of guild.", []string{"ayarlar"}, false, false, discordgo.PermissionSendMessages, discordgo.PermissionSendMessages, CommandHandler.CommandTypeGuild, commands.SettingsCommand)
	handler.AddCommand("shutdown", "Shutdowns the bot (Owner-only).", []string{}, true, false, discordgo.PermissionSendMessages, discordgo.PermissionSendMessages, CommandHandler.CommandTypeGuild, commands.ShutdownCommand)
	handler.AddCommand("eval", "Test codes (Owner-only).", []string{}, true, false, discordgo.PermissionSendMessages, discordgo.PermissionSendMessages, CommandHandler.CommandTypeGuild, commands.EvalCommand)
	handler.AddCommand("slap", "Sends slap gif", []string{}, false, false, discordgo.PermissionSendMessages, discordgo.PermissionSendMessages, CommandHandler.CommandTypeGuild, commands.SlapCommand)
	handler.AddCommand("spoiler", "Sends your message as a spoiler.", []string{}, false, false, discordgo.PermissionSendMessages, discordgo.PermissionSendMessages, CommandHandler.CommandTypeGuild, commands.SpoilerCommand)
	handler.AddCommand("start_vote", "Start a vote.", []string{"start_vote"}, false, false, discordgo.PermissionSendMessages, discordgo.PermissionSendMessages, CommandHandler.CommandTypeGuild, commands.StartVoteCommand)
	handler.AddCommand("stats", "Shows bot stats.", []string{"istatistik"}, false, false, discordgo.PermissionSendMessages, discordgo.PermissionSendMessages, CommandHandler.CommandTypeGuild, commands.StatsCommand)
	handler.AddCommand("unban", "Unbans specified user.", []string{}, false, false, discordgo.PermissionSendMessages, discordgo.PermissionSendMessages, CommandHandler.CommandTypeGuild, commands.UnbanCommand)
	handler.AddCommand("vote", "Vote for the bot on Top.gg.", []string{"top.gg"}, false, false, discordgo.PermissionSendMessages, discordgo.PermissionSendMessages, CommandHandler.CommandTypeGuild, commands.VoteCommand)
	handler.AddCommand("welcome_channel", "Set welcome channel.", []string{"welcomechannel", "welcome-channel"}, false, false, discordgo.PermissionSendMessages, discordgo.PermissionSendMessages, CommandHandler.CommandTypeGuild, commands.WelcomeChannelCommand)
	handler.AddCommand("welcome_message", "Set welcome message. \nPlaceholders: `{mention}`, `{username}`", []string{"welcomemessage", "welcome-message"}, false, false, discordgo.PermissionSendMessages, discordgo.PermissionSendMessages, CommandHandler.CommandTypeGuild, commands.WelcomeMessageCommand)
	handler.SetHelpCommand("help", []string{"yardım"}, discordgo.PermissionSendMessages, discordgo.PermissionSendMessages, commands.HelpCommand)

	handler.SetPrerunFunc(func(context CommandHandler.Context, command *CommandHandler.Command, content []string) bool {
		fmt.Printf("Command \"%s\" is being run by \"%s#%s\" (ID: %s).\n", command.Name, context.User.Username, context.User.Discriminator, context.User.ID)
		return true
	})
}
