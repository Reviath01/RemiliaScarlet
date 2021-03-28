package commands

import (
	ctx "git.randomchars.net/Reviath/handlers/Context"
	"github.com/bwmarrin/discordgo"
    embedutil "git.randomchars.net/Reviath/embed-util"
)

type Issue struct {

}

func (i Issue) Execute(ctx ctx.Ctx, session *discordgo.Session) error {
        issueembed := embedutil.NewEmbed().
            SetColor(0xffa935).
            SetDescription("Click [here](https://git.randomchars.net/Reviath/RemiliaScarlet/-/issues/new) to create an issue on GitLab").
            AddField("If you don't know how to use GitLab,","You can come to our [guild](https://discord.gg/xqsTvtM2hk) and specify the problem.").MessageEmbed
    _, err := session.ChannelMessageSendEmbed(ctx.Channel().ID, issueembed)

    if err != nil {
        return nil
    }

    return err
}
