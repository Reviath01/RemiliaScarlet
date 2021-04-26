package executor

import (
	"git.randomchars.net/Reviath/RemiliaScarlet/events"
	interactions_handler "git.randomchars.net/Reviath/RemiliaScarlet/interactions/handler"
	"github.com/bwmarrin/discordgo"
)

func RunAllEvents(client *discordgo.Session) {
	client.AddHandler(events.ChannelCreate)
	client.AddHandler(events.ChannelDelete)
	client.AddHandler(events.GuildMemberAdd)
	client.AddHandler(events.GuildMemberRemove)
	client.AddHandler(events.GuildRoleCreate)
	client.AddHandler(events.GuildRoleDelete)
	client.AddHandler(events.Ready)
	client.AddHandler(interactions_handler.InteractionHandler)
}