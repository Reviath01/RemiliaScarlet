package client

import (
	"fmt"

	"../config"
	"github.com/bwmarrin/discordgo"
	CommandHandler "git.randomchars.net/Reviath/handlers/CommandHandler"
	commandMap "git.randomchars.net/Reviath/handlers/CommandMap"
)

var BotUsername string
var BotDiscriminator string
var BotID string
var goBot *discordgo.Session

func Start() {
	goBot, err := discordgo.New("Bot " + config.Token)
	
	if err != nil {
		fmt.Println(err.Error())
		return
	}

var prefixes = []string{"r!", "r?"}

	handler := CommandHandler.New(goBot, commandMap.New(), prefixes)
	handler.GetCommandMap().RegisterCommand("ping", Ping{}, true)
	handler.GetCommandMap().RegisterCommand("stats", Stats{}, true)
	handler.GetCommandMap().RegisterCommand("avatar", Avatar{}, true)
	handler.GetCommandMap().RegisterCommand("invite", Invite{}, true)
	handler.GetCommandMap().RegisterCommand("author", Author{}, true)
	handler.GetCommandMap().RegisterCommand("issue", Issue{}, true)
	handler.GetCommandMap().RegisterCommand("guild_info", GuildInfo{}, true)

	u, err := goBot.User("@me")

	if err != nil {
		fmt.Println(err.Error())
	}

	BotID = u.ID
	BotUsername = u.Username
	BotDiscriminator = u.Discriminator

	err = goBot.Open()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Logging in as " + BotUsername + "#" + BotDiscriminator + " (" + BotID + ")")
}
