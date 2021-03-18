package client

import (
	"fmt"

	"../commands"
	"../config"
	"github.com/bwmarrin/discordgo"
	CommandHandler "git.randomchars.net/Reviath/handlers/CommandHandler"
	commandMap "git.randomchars.net/Reviath/handlers/CommandMap"
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

	u, err := goBot.User("@me")

	if err != nil {
		fmt.Println(err.Error())
	}

	goBot.Identify.Intents = discordgo.IntentsAll

	goBot.AddHandler(Ready)

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

func Ready(s *discordgo.Session, event *discordgo.Ready) {
	s.UpdateGameStatus(0, config.Presence)
}
