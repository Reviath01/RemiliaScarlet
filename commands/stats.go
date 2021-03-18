package commands

import (
    "strconv"
    "github.com/bwmarrin/discordgo"
    "time"
    "fmt"
    "runtime"
   	ctx "git.randomchars.net/Reviath/handlers/Context"
    embedutil "git.randomchars.net/Reviath/embed-util"
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
    statembed := embedutil.NewEmbed().
            SetTitle("Stats").
            SetColor(0x00ff03).
            AddField("Uptime", getDuration(time.Since(startTime))).
            AddField("Go version", runtime.Version()).
            AddField("Discordgo version", discordgo.VERSION).
            AddField("Server size", strconv.Itoa(len(session.State.Guilds))).
            AddField("Goroutines", strconv.Itoa(runtime.NumGoroutine())).MessageEmbed
        _, err := session.ChannelMessageSendEmbed(ctx.Channel().ID, statembed)
        return err
}
