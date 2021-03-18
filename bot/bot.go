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

var prefixes = []string{config.Prefix1, config.Prefix2}

	handler := CommandHandler.New(goBot, commandMap.New(), prefixes)
	handler.GetCommandMap().RegisterCommand("ping", Commands.Ping{}, true)
	handler.GetCommandMap().RegisterCommand("stats", Commands.Stats{}, true)
	handler.GetCommandMap().RegisterCommand("avatar", Commands.Avatar{}, true)
	handler.GetCommandMap().RegisterCommand("invite", Commands.Invite{}, true)
	handler.GetCommandMap().RegisterCommand("author", Commands.Author{}, true)
	handler.GetCommandMap().RegisterCommand("issue", Commands.Issue{}, true)
	handler.GetCommandMap().RegisterCommand("guild_info", Commands.GuildInfo{}, true)
	handler.GetCommandMap().RegisterCommand("embed", Commands.Embed{}, true)

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
