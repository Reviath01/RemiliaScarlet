package commands

import (
	ctx "git.randomchars.net/Reviath/handlers/Context"
	"github.com/bwmarrin/discordgo"
    embedutil "git.randomchars.net/Reviath/embed-util"
)

type Roles struct {

}

var roles string

func (r Roles) Execute(ctx ctx.Ctx, session *discordgo.Session) error {
    for _, i := range ctx.Guild().Roles {
        roles += "<@&" + i.ID + ">" + ", \n"
    }

    embed := embedutil.NewEmbed().
        AddField("Roles:", roles).MessageEmbed
        
	_, err := session.ChannelMessageSendEmbed(ctx.Channel().ID, embed)
	return err
}
