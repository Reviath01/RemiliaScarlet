package handler

import (
	"io"

	embedutil "git.randomchars.net/FreeNitori/EmbedUtil"
	"github.com/bwmarrin/discordgo"
)

//Sends reply.
func (c *Context) Reply(message string) (*discordgo.Message, error) {
	return c.Session.ChannelMessageSend(c.Channel.ID, message)
}

//Sends complex reply.
func (c *Context) ReplyComplex(message string, tts bool, embed *discordgo.MessageEmbed, files []*discordgo.File) (*discordgo.Message, error) {
	return c.Session.ChannelMessageSendComplex(c.Channel.ID, &discordgo.MessageSend{
		Content: message,
		Embed:   embed,
		TTS:     tts,
		Files:   files,
	})
}

//Reply to message channel with embedutil's embed.
func (c *Context) ReplyEmbed(embed embedutil.Embed) (*discordgo.Message, error) {
	return c.Session.ChannelMessageSendEmbed(c.Channel.ID, embed.MessageEmbed)
}

//Reply to message channel with discordgo's embed.
func (c *Context) ReplyDiscordgoEmbed(embed *discordgo.MessageEmbed) (*discordgo.Message, error) {
	return c.Session.ChannelMessageSendEmbed(c.Channel.ID, embed)
}

//Reply to message channel with file.
func (c *Context) ReplyFile(filename string, file io.Reader) (*discordgo.Message, error) {
	return c.Session.ChannelFileSend(c.Channel.ID, filename, file)
}
