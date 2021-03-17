package Commands

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
    SetImage(ctx.Author().AvatarURL("avatar")).MessageEmbed
	_, err := session.ChannelMessageSendEmbed(ctx.Channel().ID, avatarembed)
	return err
}
