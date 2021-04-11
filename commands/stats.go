package commands

import (
	"fmt"
	"runtime"
	"strconv"
	"time"

	"database/sql"

	ctx "git.randomchars.net/Reviath/RemiliaScarlet/CommandHandler/Context"
	embedutil "git.randomchars.net/Reviath/RemiliaScarlet/EmbedUtil"
	"github.com/bwmarrin/discordgo"
	_ "github.com/go-sql-driver/mysql"
)

type Stats struct {
}

var startTime = time.Now()

func getDuration(duration time.Duration) string {
	return fmt.Sprintf(
		"%0.2d:%0.2d:%0.2d",
		int(duration.Hours()),
		int(duration.Minutes())%60,
		int(duration.Seconds())%60,
	)
}
func (s Stats) Execute(ctx ctx.Ctx, session *discordgo.Session) error {
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
		err = db.QueryRow("SELECT isblocked FROM disabledcommands WHERE commandname ='stats' AND guildid ='" + ctx.Guild().ID + "'").Scan(&tag.isblocked)

		if err == nil && tag.isblocked == "True" {
			_, _ = session.ChannelMessageSend(ctx.Channel().ID, "Bu komut bu sunucuda engellenmiş.")
			return nil
		}

		statembed := embedutil.NewEmbed().
			SetTitle("İstatistikler").
			SetColor(0x00ff03).
			AddField("Çalışma Süresi", getDuration(time.Since(startTime))).
			AddField("Go Versiyonu", runtime.Version()).
			AddField("Discordgo Versiyonu", discordgo.VERSION).
			AddField("Sunucular", strconv.Itoa(len(session.State.Guilds))).
			AddField("Goroutine'ler", strconv.Itoa(runtime.NumGoroutine())).MessageEmbed
		_, _ = session.ChannelMessageSendEmbed(ctx.Channel().ID, statembed)

		return nil

	}

	err = db.QueryRow("SELECT isblocked FROM disabledcommands WHERE commandname ='stats' AND guildid ='" + ctx.Guild().ID + "'").Scan(&tag.isblocked)

	if err == nil {
		if tag.isblocked == "True" {
			_, _ = session.ChannelMessageSend(ctx.Channel().ID, "This command is blocked on this guild.")

			return nil
		}
	}

	statembed := embedutil.NewEmbed().
		SetTitle("Stats").
		SetColor(0x00ff03).
		AddField("Uptime", getDuration(time.Since(startTime))).
		AddField("Go version", runtime.Version()).
		AddField("Discordgo version", discordgo.VERSION).
		AddField("Server size", strconv.Itoa(len(session.State.Guilds))).
		AddField("Goroutines", strconv.Itoa(runtime.NumGoroutine())).MessageEmbed
	_, _ = session.ChannelMessageSendEmbed(ctx.Channel().ID, statembed)

	return nil
}
