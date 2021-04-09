package commands

import (
	ctx "git.randomchars.net/Reviath/RemiliaScarlet/Context"
	embedutil "git.randomchars.net/Reviath/embed-util"
	"github.com/bwmarrin/discordgo"
)

type Invite struct {
}

func (i Invite) Execute(ctx ctx.Ctx, session *discordgo.Session) error {
	inviteembed := embedutil.NewEmbed().
		SetColor(0xc000ff).
		SetDescription("Click [here](https://discord.com/oauth2/authorize?client_id=" + session.State.User.ID + "&scope=bot&permissions=8) to invite me!").MessageEmbed
	_, err := session.ChannelMessageSendEmbed(ctx.Channel().ID, inviteembed)

	if err != nil {
		return nil
	}

	return err
}
