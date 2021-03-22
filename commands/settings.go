package commands

import (
	ctx "git.randomchars.net/Reviath/handlers/Context"
	"github.com/bwmarrin/discordgo"
    embedutil "git.randomchars.net/Reviath/embed-util"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type Settings struct {

}

func (s Settings) Execute(ctx ctx.Ctx, session *discordgo.Session) error {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/remilia")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	type Tag struct {
		welcomechannelid string `json:"channelid"`
		leavechannelid string `json:"channelid"`
		roleid string `json:"roleid"`
	}

	var tag Tag
	var welcomechannel string
	var leavechannel string
	var autorole string

	err = db.QueryRow("SELECT channelid FROM welcomechannel WHERE guildid ='" + ctx.Guild().ID + "'").Scan(&tag.welcomechannelid)
	if err == nil {
		welcomechannel = "<#" + tag.welcomechannelid + "> (" + tag.welcomechannelid + ")"
	} else {
		welcomechannel = "Not existing."
	}

	err = db.QueryRow("SELECT channelid FROM leavechannel WHERE guildid ='" + ctx.Guild().ID + "'").Scan(&tag.leavechannelid)
	if err == nil {
		leavechannel = "<#" + tag.leavechannelid + "> (" + tag.leavechannelid + ")"
	} else {
		leavechannel = "Not existing."
	}

	err = db.QueryRow("SELECT roleid FROM autorole WHERE guildid ='" + ctx.Guild().ID + "'").Scan(&tag.roleid)
	if err == nil {
		autorole = "<@&" + tag.roleid + "> (" + tag.roleid + ")"
	} else {
		autorole = "Not existing."
	}

	embed := embedutil.NewEmbed().
        SetTitle(ctx.Guild().Name + " Settings").
		AddField("Welcome Channel:", welcomechannel).
		AddField("Leave Channel:", leavechannel).
		AddField("Autorole:", autorole).MessageEmbed
        
	_, err := session.ChannelMessageSendEmbed(ctx.Channel().ID, embed)
	return err
}
