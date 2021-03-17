package client

import (
	ctx "git.randomchars.net/Reviath/handlers/Context"
	"github.com/bwmarrin/discordgo"
    embedutil "git.randomchars.net/Reviath/embed-util"
    "strconv"
)

type GuildInfo struct {

}

func (g GuildInfo) Execute(ctx ctx.Ctx, session *discordgo.Session) error {
    embed := embedutil.NewEmbed().
    SetColor(0xefff00).
    AddField("Guild Name", ctx.GetGuild().Name).
    AddField("Member Count", strconv.Itoa(ctx.GetGuild().MemberCount)).
    AddField("Region", ctx.GetGuild().Region).
    AddField("ID:", ctx.GetGuild().ID).
    AddField("Locale", ctx.GetGuild().PreferredLocale).MessageEmbed
	_, err := session.ChannelMessageSendEmbed(ctx.GetChannel().ID, embed)
	return err
}
