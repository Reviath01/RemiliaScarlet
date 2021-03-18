package commands

import (
	ctx "git.randomchars.net/Reviath/handlers/Context"
	"github.com/bwmarrin/discordgo"
    embedutil "git.randomchars.net/Reviath/embed-util"
)

type Icon struct {

}

func (i Icon) Execute(ctx ctx.Ctx, session *discordgo.Session) error {
      iconeembed := embedutil.NewEmbed().
            SetColor(0x00f6ff).
            SetImage(ctx.Guild().IconURL()).MessageEmbed
	_, err := session.ChannelMessageSendEmbed(ctx.Channel().ID, iconeembed)
	return err
}
