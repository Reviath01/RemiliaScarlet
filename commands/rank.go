package commands

import (
	"database/sql"

	ctx "git.randomchars.net/Reviath/RemiliaScarlet/Context"
	embedutil "git.randomchars.net/Reviath/embed-util"
	"github.com/bwmarrin/discordgo"
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
		xp    string
		level string
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

	var maxxp string

	if level == "0" {
		maxxp = "150"
	} else if level == "1" {
		maxxp = "200"
	} else if level == "2" {
		maxxp = "250"
	} else if level == "3" {
		maxxp = "300"
	} else if level == "4" {
		maxxp = "250"
	} else if level == "5" {
		maxxp = "300"
	} else if level == "6" {
		maxxp = "350"
	} else if level == "7" {
		maxxp = "400"
	} else if level == "8" {
		maxxp = "500"
	} else if level == "9" {
		maxxp = "600"
	} else if level == "10" {
		maxxp = "700"
	} else if level == "11" {
		maxxp = "800"
	} else if level == "12" {
		maxxp = "900"
	} else if level == "13" {
		maxxp = "1000"
	} else if level == "14" {
		maxxp = "1100"
	} else if level == "15" {
		maxxp = "1200"
	} else if level == "16" {
		maxxp = "1400"
	} else if level == "17" {
		maxxp = "1600"
	} else if level == "18" {
		maxxp = "1800"
	} else if level == "19" {
		maxxp = "2000"
	} else if level == "20" {
		maxxp = "2200"
	} else if level == "21" {
		maxxp = "2400"
	} else if level == "22" {
		maxxp = "2600"
	} else if level == "23" {
		maxxp = "2800"
	} else if level == "24" {
		maxxp = "3000"
	} else if level == "25" {
		maxxp = "3200"
	} else if level == "26" {
		maxxp = "3400"
	} else if level == "27" {
		maxxp = "3600"
	} else if level == "28" {
		maxxp = "3800"
	} else {
		maxxp = "4000"
	}

	embed := embedutil.NewEmbed().
		SetColor(0xff1000).
		SetTitle("Rank").
		SetDescription(ctx.Author().Username+"#"+ctx.Author().Discriminator).
		AddField("Level", level).
		SetThumbnail(ctx.Author().AvatarURL("1024")).
		AddField("XP", xp+"/"+maxxp).MessageEmbed
	_, _ = session.ChannelMessageSendEmbed(ctx.Channel().ID, embed)

	return nil
}
