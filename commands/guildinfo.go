package commands

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
    AddField("Guild Name", ctx.Guild().Name).
    AddField("Member Count", strconv.Itoa(ctx.Guild().MemberCount)).
    AddField("Region", ctx.Guild().Region).
    AddField("Guild Owner", "<@" + ctx.Guild().OwnerID + ">").
    AddField("Guild Owner ID", ctx.Guild().OwnerID).
    AddField("Afk Timeout", strconv.Itoa(ctx.Guild().AfkTimeout)).
    AddField("ID:", ctx.Guild().ID).
    AddField("Locale", ctx.Guild().PreferredLocale).MessageEmbed
	_, err := session.ChannelMessageSendEmbed(ctx.Channel().ID, embed)
	return err
}
