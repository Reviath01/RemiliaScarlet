package client

import (
	"fmt"

	"../commands"
	"../config"
	"github.com/bwmarrin/discordgo"
	CommandHandler "git.randomchars.net/Reviath/handlers/CommandHandler"
	commandMap "git.randomchars.net/Reviath/handlers/CommandMap"
	"../events"
)

var goBot *discordgo.Session

func Start() {
	goBot, err := discordgo.New("Bot " + config.Token)
	
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	var prefix string
	prefix = config.BotPrefix

	handler := CommandHandler.New(goBot, commandMap.New(), prefix)
	handler.GetCommandMap().RegisterCommand("ping", commands.Ping{}, true)
	handler.GetCommandMap().RegisterCommand("stats", commands.Stats{}, true)
	handler.GetCommandMap().RegisterCommand("avatar", commands.Avatar{}, true)
	handler.GetCommandMap().RegisterCommand("invite", commands.Invite{}, true)
	handler.GetCommandMap().RegisterCommand("author", commands.Author{}, true)
	handler.GetCommandMap().RegisterCommand("issue", commands.Issue{}, true)
	handler.GetCommandMap().RegisterCommand("guild_info", commands.GuildInfo{}, true)
	handler.GetCommandMap().RegisterCommand("embed", commands.Embed{}, true)
	handler.GetCommandMap().RegisterCommand("icon", commands.Icon{}, true)
	handler.GetCommandMap().RegisterCommand("spoiler", commands.Spoiler{}, true)
	handler.GetCommandMap().RegisterCommand("hug", commands.Hug{}, true)
	handler.GetCommandMap().RegisterCommand("kiss", commands.Kiss{}, true)
	handler.GetCommandMap().RegisterCommand("slap", commands.Slap{}, true)
	handler.GetCommandMap().RegisterCommand("roles", commands.Roles{}, true)
	handler.GetCommandMap().RegisterCommand("ban", commands.Ban{}, true)
	handler.GetCommandMap().RegisterCommand("start_vote", commands.StartVote{}, true)
	handler.GetCommandMap().RegisterCommand("unban", commands.Unban{}, true)
	handler.GetCommandMap().RegisterCommand("kick", commands.Kick{}, true)
	handler.GetCommandMap().RegisterCommand("afk", commands.Afk{}, true)
	handler.GetCommandMap().RegisterCommand("help", commands.Help{}, true)
	handler.GetCommandMap().RegisterCommand("welcome_channel", commands.WelcomeChannel{}, true)
	handler.GetCommandMap().RegisterCommand("leave_channel", commands.LeaveChannel{}, true)
	handler.GetCommandMap().RegisterCommand("auto_role", commands.AutoRole{}, true)
	handler.GetCommandMap().RegisterCommand("reset_auto_role", commands.ResetAutorole{}, true)
	handler.GetCommandMap().RegisterCommand("settings", commands.Settings{}, true)
	handler.GetCommandMap().RegisterCommand("leave_message", commands.LeaveMessage{}, true)
	handler.GetCommandMap().RegisterCommand("welcome_message", commands.WelcomeMessage{}, true)
	handler.GetCommandMap().RegisterCommand("reset_leave_channel", commands.ResetLeaveChannel{}, true)
	handler.GetCommandMap().RegisterCommand("reset_welcome_channel", commands.ResetWelcomeChannel{}, true)
	handler.GetCommandMap().RegisterCommand("reset_welcome_message", commands.ResetWelcomeMessage{}, true)
	handler.GetCommandMap().RegisterCommand("reset_leave_message", commands.ResetLeaveMessage{}, true)
	handler.GetCommandMap().RegisterCommand("log", commands.Log{}, true)

	u, err := goBot.User("@me")

	if err != nil {
		fmt.Println(err.Error())
	}

	goBot.Identify.Intents = discordgo.IntentsAll

	goBot.AddHandler(events.Ready)
	goBot.AddHandler(events.GuildMemberAdd)
	goBot.AddHandler(events.GuildMemberRemove)
	goBot.AddHandler(events.ChannelCreate)
	goBot.AddHandler(events.ChannelDelete)

	BotID := u.ID
	BotUsername := u.Username
	BotDiscriminator := u.Discriminator

	err = goBot.Open()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Logging in as " + BotUsername + "#" + BotDiscriminator + " (" + BotID + ")")
}
