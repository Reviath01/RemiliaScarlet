package commands

import (
    "strings"
	ctx "git.randomchars.net/Reviath/handlers/Context"
	"github.com/bwmarrin/discordgo"
    embedutil "git.randomchars.net/Reviath/embed-util"
)

type Embed struct {

}

func (e Embed) Execute(ctx ctx.Ctx, session *discordgo.Session) error {
    if strings.Join(ctx.Args()," ") == "" {
    _, err := session.ChannelMessageSend(ctx.Channel().ID, "You need to specify the message.")
    return err
    }
      embed := embedutil.NewEmbed().
            SetColor(0xc000ff).
            SetDescription(strings.Join(ctx.Args()," ")).MessageEmbed
	_, err := session.ChannelMessageSendEmbed(ctx.Channel().ID, embed)
	return err
}
