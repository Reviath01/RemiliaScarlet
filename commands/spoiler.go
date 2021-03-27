package commands

import (
    "strings"
	ctx "git.randomchars.net/Reviath/handlers/Context"
	"github.com/bwmarrin/discordgo"
    embedutil "git.randomchars.net/Reviath/embed-util"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)

type Spoiler struct {

}

func (s Spoiler) Execute(ctx ctx.Ctx, session *discordgo.Session) error {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/remilia")

    if err != nil {
        panic(err.Error())
    }

    defer db.Close()

    type Tag struct {
        isblocked string `json:"isblocked"`
    }

    var tag Tag

    err = db.QueryRow("SELECT isblocked FROM disabledcommands WHERE commandname ='spoiler' AND guildid ='" + ctx.Guild().ID + "'").Scan(&tag.isblocked)

    if err == nil {
        if tag.isblocked == "True" {
            _, err = session.ChannelMessageSend(ctx.Channel().ID, "This command is blocked on this guild.")
            return err
        }
    }

    if strings.Join(ctx.Args()," ") == "" {
    _, err := session.ChannelMessageSend(ctx.Channel().ID, "You need to specify the message.")
    return err
    }
      spoilerembed := embedutil.NewEmbed().
            SetColor(0xe9ff00).
            SetDescription("|| " + strings.Join(ctx.Args()," ") + " ||").MessageEmbed
	_, err = session.ChannelMessageSendEmbed(ctx.Channel().ID, spoilerembed)
	return err
}
