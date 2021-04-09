package commands

import (
	"database/sql"

	ctx "git.randomchars.net/Reviath/RemiliaScarlet/Context"
	embedutil "git.randomchars.net/Reviath/embed-util"
	"github.com/bwmarrin/discordgo"
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
		isblocked string
		lang      string
	}

	var tag Tag

	err = db.QueryRow("SELECT language FROM languages WHERE guildid ='" + ctx.Guild().ID + "'").Scan(&tag.lang)
	if err == nil && tag.lang == "tr" {
		err = db.QueryRow("SELECT isblocked FROM disabledcommands WHERE commandname ='roles' AND guildid ='" + ctx.Guild().ID + "'").Scan(&tag.isblocked)

		if err == nil {
			if tag.isblocked == "True" {
				_, err = session.ChannelMessageSend(ctx.Channel().ID, "Bu komut bu sunucuda engellenmiş.")

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
			AddField("Roller:", roles).MessageEmbed

		_, err = session.ChannelMessageSendEmbed(ctx.Channel().ID, embed)

		if err != nil {
			return nil
		}

		return err
	}

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
