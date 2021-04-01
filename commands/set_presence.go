package commands

import (
	"strings"

	ctx "git.randomchars.net/Reviath/handlers/Context"
	"github.com/bwmarrin/discordgo"
	"../config"
)

type SetPresence struct {

}

func (s SetPresence) Execute(ctx ctx.Ctx, session *discordgo.Session) error {
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
