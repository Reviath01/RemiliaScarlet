package client

import (
	ctx "git.randomchars.net/Reviath/handlers/Context"
	"github.com/bwmarrin/discordgo"
    embedutil "git.randomchars.net/Reviath/embed-util"
)

type Avatar struct {

}

func (a Avatar) Execute(ctx ctx.Ctx, session *discordgo.Session) error {
    avatarembed := embedutil.NewEmbed().
    SetColor(0xff1000).
    SetImage(ctx.GetAuthor().AvatarURL("avatar")).MessageEmbed
	_, err := session.ChannelMessageSendEmbed(ctx.GetChannel().ID, avatarembed)
	return err
}
