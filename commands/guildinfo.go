package commands

import (
	"database/sql"
	"strconv"

	ctx "git.randomchars.net/Reviath/RemiliaScarlet/CommandHandler/Context"
	embedutil "git.randomchars.net/Reviath/RemiliaScarlet/EmbedUtil"
	"github.com/bwmarrin/discordgo"
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
		isblocked string
		lang      string
	}

	var tag Tag

	err = db.QueryRow("SELECT language FROM languages WHERE guildid ='" + ctx.Guild().ID + "'").Scan(&tag.lang)
	if err == nil && tag.lang == "tr" {
		err = db.QueryRow("SELECT isblocked FROM disabledcommands WHERE commandname ='guild_info' AND guildid ='" + ctx.Guild().ID + "'").Scan(&tag.isblocked)

		if err == nil {
			if tag.isblocked == "True" {
				_, _ = session.ChannelMessageSend(ctx.Channel().ID, "Bu komut bu sunucuda engellenmiş.")

				return nil
			}
		}

		embed := embedutil.NewEmbed().
			SetColor(0xefff00).
			AddField("Sunucu İsmi", ctx.Guild().Name).
			AddField("Kişi Sayısı", strconv.Itoa(ctx.Guild().MemberCount)).
			AddField("Bölge", ctx.Guild().Region).
			AddField("Sunucu Sahibi", "<@"+ctx.Guild().OwnerID+">").
			AddField("Sunucu Sahibinin ID'si", ctx.Guild().OwnerID).
			AddField("Afk Süresi", strconv.Itoa(ctx.Guild().AfkTimeout)).
			AddField("ID:", ctx.Guild().ID).
			AddField("Yer", ctx.Guild().PreferredLocale).MessageEmbed
		_, _ = session.ChannelMessageSendEmbed(ctx.Channel().ID, embed)

		return nil
	}

	err = db.QueryRow("SELECT isblocked FROM disabledcommands WHERE commandname ='guild_info' AND guildid ='" + ctx.Guild().ID + "'").Scan(&tag.isblocked)

	if err == nil {
		if tag.isblocked == "True" {
			_, _ = session.ChannelMessageSend(ctx.Channel().ID, "This command is blocked on this guild.")

			return nil
		}
	}

	embed := embedutil.NewEmbed().
		SetColor(0xefff00).
		AddField("Guild Name", ctx.Guild().Name).
		AddField("Member Count", strconv.Itoa(ctx.Guild().MemberCount)).
		AddField("Region", ctx.Guild().Region).
		AddField("Guild Owner", "<@"+ctx.Guild().OwnerID+">").
		AddField("Guild Owner ID", ctx.Guild().OwnerID).
		AddField("Afk Timeout", strconv.Itoa(ctx.Guild().AfkTimeout)).
		AddField("ID:", ctx.Guild().ID).
		AddField("Locale", ctx.Guild().PreferredLocale).MessageEmbed
	_, _ = session.ChannelMessageSendEmbed(ctx.Channel().ID, embed)

	return nil
}
