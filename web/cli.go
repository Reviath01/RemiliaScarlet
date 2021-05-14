package web

import (
	"fmt"
	"os"
	"time"

	"git.randomchars.net/Reviath/RemiliaScarlet/config"
	"github.com/bwmarrin/discordgo"
)

//Fetching client user function.
func GetClientUser(session *discordgo.Session) *discordgo.User {
	cli, err := session.User(config.ClientID)
	if err != nil {
		fmt.Println("An error occurred while getting client user on web panel (please be sure that you wrote ClientID correct)")
		time.Sleep(1 * time.Second)
		fmt.Println(err.Error())
		os.Exit(1)
	}
	return cli
}
