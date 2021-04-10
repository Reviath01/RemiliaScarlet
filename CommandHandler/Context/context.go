package ctx

import "github.com/bwmarrin/discordgo"

type Ctx interface {
	Args() []string
	Author() *discordgo.User
	Channel() *discordgo.Channel
	Guild() *discordgo.Guild
	Message() *discordgo.MessageCreate
}

type BaseCtx struct {
	args    []string
	author  *discordgo.User
	channel *discordgo.Channel
	guild   *discordgo.Guild
	message *discordgo.MessageCreate
}

func (b BaseCtx) Args() []string {
	return b.args
}

func (b BaseCtx) Author() *discordgo.User {
	return b.author
}

func (b BaseCtx) Channel() *discordgo.Channel {
	return b.channel
}

func (b BaseCtx) Guild() *discordgo.Guild {
	return b.guild
}

func (b BaseCtx) Message() *discordgo.MessageCreate {
	return b.message
}

func New(args []string, msg *discordgo.MessageCreate, session *discordgo.Session) *BaseCtx {
	ctx := &BaseCtx{args: args}
	ctx.author = msg.Author
	ctx.channel, _ = session.Channel(msg.ChannelID)
	ctx.guild, _ = session.Guild(msg.GuildID)
	ctx.message = msg
	return ctx
}
