package commands

import (
	ctx "git.randomchars.net/Reviath/handlers/Context"
	"github.com/bwmarrin/discordgo"
    embedutil "git.randomchars.net/Reviath/embed-util"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type Rank struct {

}

func (r Rank) Execute(ctx ctx.Ctx, session *discordgo.Session) error {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/remilia")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	type Tag struct {
		xp string `json:"xp"`
		level string `json:"level"`
		lang string `json:"language"`
	}

	var tag Tag
	var level string
	var xp string

	err = db.QueryRow("SELECT level FROM levels WHERE userid ='" + ctx.Author().ID + "' AND guildid ='" + ctx.Guild().ID + "'").Scan(&tag.level)

	if err != nil {
		level = "0"
	} else {
		level = tag.level	
	}

	err = db.QueryRow("SELECT xp FROM xps WHERE userid ='" + ctx.Author().ID + "' AND guildid ='" + ctx.Guild().ID + "'").Scan(&tag.xp)

	if err != nil {
		xp = "0"
	} else {
		xp = tag.xp
	}

	embed := embedutil.NewEmbed().
            SetColor(0xff1000).
            SetTitle("Rank").
			SetDescription(ctx.Author().Username + "#" + ctx.Author().Discriminator).
			AddField("Level", level).
			AddField("XP", xp).MessageEmbed
	_, err = session.ChannelMessageSendEmbed(ctx.Channel().ID, embed)
	
	if err != nil {
		return nil
	}
	
	return err
}
