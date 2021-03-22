package commands

import (
	ctx "git.randomchars.net/Reviath/handlers/Context"
	"github.com/bwmarrin/discordgo"
	embedutil "git.randomchars.net/Reviath/embed-util"
    "../config"
)

type Help struct {

}

func (h Help) Execute(ctx ctx.Ctx, session *discordgo.Session) error {
    u, err := session.User("@me")
    embed := embedutil.NewEmbed().
        SetTitle(u.Username + " Help Menu!").
        SetColor(0x2ecc71).
        AddField("User Commands", config.BotPrefix + "afk: `Sets you as AFK.` \n" + config.BotPrefix + "avatar: `Fetch the profile picture of a user.` \n" + config.BotPrefix + "ping: `Sends latency of the bot.` \n" + config.BotPrefix + "invite: `Sends invite link.` \n" + config.BotPrefix + "roles: `Shows list of roles.` \n" + config.BotPrefix + "guild_info: `Gives information about guild.` \n" + config.BotPrefix + "stats: `Shows stats for bot.` \n" + config.BotPrefix + "settings: `Shows server settings.`").
        AddField("Moderation Commands", config.BotPrefix + "ban: `Bans the user.` \n" + config.BotPrefix + "kick: `Kicks the user.` \n" + config.BotPrefix + "unban: `Unbans the user.` \n" + config.BotPrefix + "start_vote: `Starts a vote.` \n" + config.BotPrefix + "welcome_channel: `When someone joins your server, the bot sends a message to the channel you set.` \n" + config.BotPrefix + "leave_channel: `When someone leaves your server, the bot sends a message to the channel you set.` \n" + config.BotPrefix + "reset_welcome_channel: `Resets welcome channel.` \n" + config.BotPrefix + "reset_leave_channel: `Resets leave channel.` \n" + config.BotPrefix + "auto_role: `Sets autorole.` \n" + config.BotPrefix + "reset_auto_role: `Resets autorole.` \n" + config.BotPrefix + "welcome_message: `Sets welcome message.` \n" + config.BotPrefix + "leave_message: `Sets leave message.`").
        AddField("Fun Commands", config.BotPrefix + "embed: `Sends your message as an embed.` \n" + config.BotPrefix + "hug: `Allows you to hug someone.` \n" + config.BotPrefix + "icon: `Sends guild icon.` \n" + config.BotPrefix + "kiss: `Allows you to kiss someone.` \n" + config.BotPrefix + "slap: `Sends slap gif.` \n" + config.BotPrefix + "spoiler: `Sends your message as a spoiler.`").MessageEmbed
    _, err = session.ChannelMessageSendEmbed(ctx.Channel().ID, embed)
	return err
}
