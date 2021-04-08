package commands

import (
	ctx "git.randomchars.net/Reviath/handlers/Context"
	"github.com/bwmarrin/discordgo"
    embedutil "git.randomchars.net/Reviath/embed-util"
)

type Vote struct {

}

func (v Vote) Execute(ctx ctx.Ctx, session *discordgo.Session) error {
      voteeembed := embedutil.NewEmbed().
            SetColor(0xc000ff).
            SetDescription("Click [here](https://top.gg/bot/"+ session.State.User.ID + "/vote) to vote me on Top.gg!").MessageEmbed
	_, err := session.ChannelMessageSendEmbed(ctx.Channel().ID, voteeembed)

	if err != nil {
		return nil
	}
	
	return err
}
