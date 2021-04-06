package commands

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	ctx "git.randomchars.net/Reviath/handlers/Context"
	"github.com/bwmarrin/discordgo"
)

type SetPresence struct {

}

func (s SetPresence) Execute(ctx ctx.Ctx, session *discordgo.Session) error {
    file, err := ioutil.ReadFile("../config.json")

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

    type configStruct struct {
        Token     string `json:"Token"`
        BotPrefix string `json:"BotPrefix"`
        Presence string `json:"Presence"`
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

	u, err := session.User(config.Owner)
	if err != nil {
		return nil
	}

	if ctx.Author().ID != u.ID {
	_, err := session.ChannelMessageSend(ctx.Channel().ID, "You do not own this bot.")
	if err != nil {
		return nil
	}
	return err
}
err = session.UpdateGameStatus(0, strings.Join(ctx.Args(), " "))
if err != nil {
	_, err = session.ChannelMessageSend(ctx.Channel().ID, "An error occured.")
	if err != nil {
		return nil 
	}
	return err
}
_, err = session.ChannelMessageSend(ctx.Channel().ID, "Successfully updated presence.")
if err != nil {
	return nil
}
return err
}
