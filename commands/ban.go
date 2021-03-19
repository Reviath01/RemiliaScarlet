package commands

import (
	ctx "git.randomchars.net/Reviath/handlers/Context"
	"github.com/bwmarrin/discordgo"
    "strings"
)

type Ban struct {

}

func (b Ban) Execute(ctx ctx.Ctx, session *discordgo.Session) error {
    perms, err := session.State.UserChannelPermissions(ctx.Author().ID, ctx.Channel().ID)
	if err == nil && (int(perms)&discordgo.PermissionBanMembers == discordgo.PermissionBanMembers) == false {
        _, err := session.ChannelMessageSend(ctx.Channel().ID, "You need ban members permission to run this command.")
	return err
    }

    if len(strings.Join(ctx.Args()," ")) < 19 {
    u, err := session.User(strings.Join(ctx.Args()," "))
    if err == nil {
        err = session.GuildBanCreate(ctx.Guild().ID, u.ID, 0)
        if err != nil {
            _, err = session.ChannelMessageSend(ctx.Channel().ID, "I do not have enough permission.")
            return err
        }
        _, err = session.ChannelMessageSend(ctx.Channel().ID, "Banned specified user.")
		return err
	} else {
        _, err = session.ChannelMessageSend(ctx.Channel().ID, "You need to specify the user.")
        return err
    }
} else {
    if len(strings.Join(ctx.Args()," ")) > 21 {
     u, err := session.User(strings.Join(ctx.Args()," ")[3:][:18])
        if err == nil {
            err = session.GuildBanCreate(ctx.Guild().ID, u.ID, 0)
            if err != nil {
                _, err = session.ChannelMessageSend(ctx.Channel().ID, "I do not have enough permission.")
                return err
            }
            _, err = session.ChannelMessageSend(ctx.Channel().ID, "Banned specified user.")
		    return err
        } else {
            _, err = session.ChannelMessageSend(ctx.Channel().ID, "You need to specify the user.")
           return err
        }
    } else {
            _, err := session.ChannelMessageSend(ctx.Channel().ID, "You need to specify the user.")
            return err
        }
    }
}
