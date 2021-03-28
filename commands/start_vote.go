package commands

import (
	ctx "git.randomchars.net/Reviath/handlers/Context"
	"github.com/bwmarrin/discordgo"
    embedutil "git.randomchars.net/Reviath/embed-util"
    "strings"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)

type StartVote struct {

}

func (s StartVote) Execute(ctx ctx.Ctx, session *discordgo.Session) error {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/remilia")

    if err != nil {
        panic(err.Error())
    }

    defer db.Close()

    type Tag struct {
        isblocked string `json:"isblocked"`
    }

    var tag Tag

    err = db.QueryRow("SELECT isblocked FROM disabledcommands WHERE commandname ='start_vote' AND guildid ='" + ctx.Guild().ID + "'").Scan(&tag.isblocked)

    if err == nil {
        if tag.isblocked == "True" {
            _, err = session.ChannelMessageSend(ctx.Channel().ID, "This command is blocked on this guild.")
            
            if err != nil {
                return nil
            }

            return err
        }
    }

    perms, err := session.State.UserChannelPermissions(ctx.Author().ID, ctx.Channel().ID)
	if err == nil && (int(perms)&discordgo.PermissionManageMessages == discordgo.PermissionManageMessages) == false {
        _, err := session.ChannelMessageSend(ctx.Channel().ID, "You need manage messages permission to run this command.")
	
        if err != nil {
            return nil
        }

        return err
    }
    if strings.Join(ctx.Args()," ") == "" {
    _, err := session.ChannelMessageSend(ctx.Channel().ID, "You need to specify a message.")
    
    if err != nil {
        return nil
    }

    return err
    }
    embed := embedutil.NewEmbed().
    SetTitle("Vote started!").
    SetColor(0xff9100).
    AddField("Vote question:", strings.Join(ctx.Args()," ")).MessageEmbed
	msg, err := session.ChannelMessageSendEmbed(ctx.Channel().ID, embed)
    session.MessageReactionAdd(ctx.Channel().ID, msg.ID, "👍")
	session.MessageReactionAdd(ctx.Channel().ID, msg.ID, "👎")
	
    if err != nil {
        return nil
    }

    return err
}
