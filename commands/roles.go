package commands

import (
	ctx "git.randomchars.net/Reviath/handlers/Context"
	"github.com/bwmarrin/discordgo"
    embedutil "git.randomchars.net/Reviath/embed-util"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)

type Roles struct {

}

var roles string

func (r Roles) Execute(ctx ctx.Ctx, session *discordgo.Session) error {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/remilia")

    if err != nil {
        panic(err.Error())
    }

    defer db.Close()

    type Tag struct {
        isblocked string `json:"isblocked"`
    }

    var tag Tag

    err = db.QueryRow("SELECT isblocked FROM disabledcommands WHERE commandname ='roles' AND guildid ='" + ctx.Guild().ID + "'").Scan(&tag.isblocked)

    if err == nil {
        if tag.isblocked == "True" {
            _, err = session.ChannelMessageSend(ctx.Channel().ID, "This command is blocked on this guild.")
            
            if err != nil {
                return nil
            }
            
            return err
        }
    }

    for _, i := range ctx.Guild().Roles {
        roles += "<@&" + i.ID + ">" + ", \n"
    }

    embed := embedutil.NewEmbed().
        AddField("Roles:", roles).MessageEmbed
        
	_, err = session.ChannelMessageSendEmbed(ctx.Channel().ID, embed)
	
    if err != nil {
        return nil
    }
    
    return err
}
