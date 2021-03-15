package state

import (
	multiplexer "git.randomchars.net/remilia-scarlet-command-handler"
	"github.com/bwmarrin/discordgo"
	"time"
)

var (
	version  = "v1.0.3"
	revision = "unknown"
	start    time.Time
)

var InviteURL string

func Version() string { return version }

func Revision() string { return revision }

func Uptime() time.Duration { return time.Since(start) }

var (
	ExitCode     = make(chan int)
	DiscordReady = make(chan bool)
)

var Multiplexer = multiplexer.New()

var RawSession, _ = discordgo.New()

var ShardSessions []*discordgo.Session

var Application *discordgo.Application

func init() {
	start = time.Now()
}
