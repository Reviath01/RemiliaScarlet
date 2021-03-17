package client

import (
	ctx "git.randomchars.net/Reviath/handlers/Context"
	"github.com/bwmarrin/discordgo"
	embedutil "git.randomchars.net/Reviath/embed-util"
)

type Author struct {

}

func (a Author) Execute(ctx ctx.Ctx, session *discordgo.Session) error {
    authorembed := embedutil.NewEmbed().
    SetColor(0x007bff).
    AddField("My Author:", "<@770218429096656917> ([Reviath#0001](https://discord.com/users/770218429096656917))").MessageEmbed
	_, err := session.ChannelMessageSendEmbed(ctx.GetChannel().ID, authorembed)
	return err
}
