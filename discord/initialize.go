package discord

import (
	"errors"
	"fmt"
	"git.randomchars.net/FreeNitori/FreeNitori/cmd/server/discord/sessioning"
	"../config"
	"git.randomchars.net/FreeNitori/FreeNitori/nitori/log"
	"../state"
	multiplexer "git.randomchars.net/Reviath/remilia-scarlet-command-handler"
	"github.com/bwmarrin/discordgo"
	"strconv"
)

var err error

func Initialize() error {
	multiplexer.SetLogger(log.Logger)
	multiplexer.NoCommandMatched = func(context *multiplexer.Context) {
		if context.HasMention {
			context.SendMessage("<a:KyokoAngryPing:821054351358820372>")
		} else {
			context.SendMessage(fmt.Sprintf("This command does not exist! Issue `%sman` for a list of command manuals.",
				context.Prefix()))
		}
	}
	state.Multiplexer.Prefix = config.BotPrefix
	discordgo.Logger = log.DiscordGoLogger
	state.RawSession.UserAgent = "DiscordBot (FreeNitori " + state.Version() + ")"
	if config.TokenOverride == "" {
		state.RawSession.Token = "Bot " + config.Token
	} else {
		state.RawSession.Token = "Bot " + config.Token
	}
	state.RawSession.ShouldReconnectOnError = true
	state.RawSession.State.MaxMessageCount = 1000
	state.RawSession.Identify.Intents = discordgo.IntentsAll

	return nil
}

func LateInitialize() error {
	err = state.RawSession.Open()
	if err != nil {
		log.Warnf("Unable to open session with all intents, %s, Nitori will now fallback to unprivileged intents, some functionality will be unavailable.", err)
		state.RawSession.Identify.Intents = discordgo.IntentsAllWithoutPrivileged
		err = state.RawSession.Open()
		if err != nil {
			return errors.New("unable to open session with Discord")
		}
	}
	log.Info("Raw session with Discord opened.")
	state.Multiplexer.Administrator, err = state.RawSession.User(strconv.Itoa(config.Config.System.Administrator))
	if err != nil {
		return errors.New("unable to get system administrator")
	}
	for _, id := range config.Config.System.Operator {
		user, err := state.RawSession.User(strconv.Itoa(id))
		if err == nil {
			state.Multiplexer.Operator = append(state.Multiplexer.Operator, user)
		}
	}
	state.Application, err = state.RawSession.Application("@me")
	if err != nil {
		return errors.New("unable to fetch application information")
	}
	state.InviteURL = fmt.Sprintf("https://discord.com/oauth2/authorize?client_id=%s&scope=bot&permissions=8", state.Application.ID)
	go func() {
		for {
			state.DiscordReady <- true
		}
	}()

	log.Infof("Successfully logged in as %s#%s (%s).",
		state.RawSession.State.User.Username,
		state.RawSession.State.User.Discriminator,
		state.RawSession.State.User.ID)
	return nil
}
