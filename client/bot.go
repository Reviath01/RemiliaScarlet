package client

import (
	"fmt"

	"../config"
	"../embed"
	"github.com/bwmarrin/discordgo"
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

	u, err := goBot.User("@me")

	if err != nil {
		fmt.Println(err.Error())
	}

	BotID = u.ID
	BotUsername = u.Username
	BotDiscriminator = u.Discriminator

	goBot.AddHandler(messageHandler)

	err = goBot.Open()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Logging in as " + BotUsername + "#" + BotDiscriminator + " (" + BotID + ")")
}

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {

	if m.Author.ID == BotID {
		return
	}

	if m.Content == config.BotPrefix + "hug" {
		_, _ = s.ChannelMessageSendEmbed(m.ChannelID, embed.ImageEmbedWithColor("https://images-ext-1.discordapp.net/external/xT0jPvqo3SiH7lJG88IXCgxbrmd36Rt6-Hgvu5aw8Rg/https/i.pinimg.com/originals/4d/d7/49/4dd749423de10a319b5d9e8850bbace4.gif?width=596&height=596", 0xff0000))
	}
}
