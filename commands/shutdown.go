package commands

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	ctx "git.randomchars.net/Reviath/handlers/Context"
	"github.com/bwmarrin/discordgo"
)

type ShutDown struct {

}

func (s ShutDown) Execute(ctx ctx.Ctx, session *discordgo.Session) error {
    file, err := ioutil.ReadFile("./config.json")

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

    type configStruct struct {
        Owner string `json:"Owner"`
    }
    
    var (
        config *configStruct
    ) 
    
    err = json.Unmarshal(file, &config)

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	if ctx.Author().ID != config.Owner {
		_, err = session.ChannelMessageSend(ctx.Channel().ID, "You do not own this bot.")
		if err != nil {
			return nil
		}
		return err
	}
	_, err = session.ChannelMessageSend(ctx.Channel().ID, "Performing shutdown.")
	if err != nil {
		return nil
	}
	os.Exit(1)
	return err
}