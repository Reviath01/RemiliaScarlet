package commands

import (
    "strings"
	ctx "git.randomchars.net/Reviath/handlers/Context"
	"github.com/bwmarrin/discordgo"
    embedutil "git.randomchars.net/Reviath/embed-util"
)

type Spoiler struct {

}

func (e Spoiler) Execute(ctx ctx.Ctx, session *discordgo.Session) error {
    if strings.Join(ctx.Args()," ") == "" {
    _, err := session.ChannelMessageSend(ctx.Channel().ID, "You need to specify the message.")
    return err
    }
      spoilerembed := embedutil.NewEmbed().
            SetColor(0xe9ff00).
            SetDescription("|| " + strings.Join(ctx.Args()," ") + " ||").MessageEmbed
	_, err := session.ChannelMessageSendEmbed(ctx.Channel().ID, spoilerembed)
	return err
}
