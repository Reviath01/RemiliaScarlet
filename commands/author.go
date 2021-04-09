package commands

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"

	ctx "git.randomchars.net/Reviath/RemiliaScarlet/Context"
	embedutil "git.randomchars.net/Reviath/embed-util"
	"github.com/bwmarrin/discordgo"
	_ "github.com/go-sql-driver/mysql"
)

type Author struct {
}

func (a Author) Execute(ctx ctx.Ctx, session *discordgo.Session) error {
	file, err := ioutil.ReadFile("./config.json")

	if err != nil {
		return nil
	}

	type configStruct struct {
		Token     string `json:"Token"`
		BotPrefix string `json:"BotPrefix"`
		Presence  string `json:"Presence"`
		Owner     string `json:"Owner"`
	}

	var (
		config *configStruct
	)

	err = json.Unmarshal(file, &config)

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/remilia")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	type Tag struct {
		isblocked string `json:"isblocked"`
		lang      string `json:"language"`
	}

	var tag Tag

	err = db.QueryRow("SELECT language FROM languages WHERE guildid ='" + ctx.Guild().ID + "'").Scan(&tag.lang)
	if err == nil && tag.lang == "tr" {
		err = db.QueryRow("SELECT isblocked FROM disabledcommands WHERE commandname ='author' AND guildid ='" + ctx.Guild().ID + "'").Scan(&tag.isblocked)

		if err == nil {
			if tag.isblocked == "True" {
				_, err = session.ChannelMessageSend(ctx.Channel().ID, "Bu komut bu sunucuda engellenmiş.")
				if err != nil {
					return nil
				}
				return err
			}
		}

		u, err := session.User(config.Owner)
		if err != nil {
			return nil
		}
		authorembed := embedutil.NewEmbed().
			SetColor(0x007bff).
			AddField("Sahibim:", "<@"+u.ID+"> (["+u.Username+"#"+u.Discriminator+"](https://discord.com/users/"+u.ID+"))").MessageEmbed
		_, err = session.ChannelMessageSendEmbed(ctx.Channel().ID, authorembed)

		if err != nil {
			return nil
		}
		return err
	}

	err = db.QueryRow("SELECT isblocked FROM disabledcommands WHERE commandname ='author' AND guildid ='" + ctx.Guild().ID + "'").Scan(&tag.isblocked)

	if err == nil {
		if tag.isblocked == "True" {
			_, err = session.ChannelMessageSend(ctx.Channel().ID, "This command is blocked on this guild.")
			if err != nil {
				return nil
			}
			return err
		}
	}
	u, err := session.User(config.Owner)
	if err != nil {
		return nil
	}

	authorembed := embedutil.NewEmbed().
		SetColor(0x007bff).
		AddField("My Author:", "<@"+u.ID+"> (["+u.Username+"#"+u.Discriminator+"](https://discord.com/users/"+u.ID+"))").MessageEmbed
	_, err = session.ChannelMessageSendEmbed(ctx.Channel().ID, authorembed)

	if err != nil {
		return nil
	}

	return err
}
