package commands

import (
	ctx "git.randomchars.net/Reviath/handlers/Context"
	"github.com/bwmarrin/discordgo"
    embedutil "git.randomchars.net/Reviath/embed-util"
    "strconv"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)

type GuildInfo struct {

}

func (g GuildInfo) Execute(ctx ctx.Ctx, session *discordgo.Session) error {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/remilia")

    if err != nil {
        panic(err.Error())
    }

    defer db.Close()

    type Tag struct {
        isblocked string `json:"isblocked"`
    }

    var tag Tag

    err = db.QueryRow("SELECT isblocked FROM disabledcommands WHERE commandname ='guild_info' AND guildid ='" + ctx.Guild().ID + "'").Scan(&tag.isblocked)

    if err == nil {
        if tag.isblocked == "True" {
            _, err = session.ChannelMessageSend(ctx.Channel().ID, "This command is blocked on this guild.")
            
            if err != nil {
                return nil
            }

            return err
        }
    }

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
	_, err = session.ChannelMessageSendEmbed(ctx.Channel().ID, embed)
	
    if err != nil {
        return nil
    }

    return err
}
