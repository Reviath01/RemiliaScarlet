package main

import (
	"../discord"
	"../config"
	"git.randomchars.net/FreeNitori/FreeNitori/nitori/log"
	"../state"
	"go/types"
	"os"
	"os/signal"
	"syscall"
)

var err error
var _ = func() *types.Nil {
	if config.VersionStartup {
		println(state.Version() + " (" + state.Revision() + ")")
		os.Exit(0)
	}
	return nil
}()

func init() {

	log.Infof("%s (%s) early initialization.", state.Version(), state.Revision())

	err = discord.Initialize()
	if err != nil {
		log.Fatalf("Unable to initialize Discord services, %s", err)
		_ = database.Database.Close()
		os.Exit(1)
	}


func main() {

	defer cleanup()

	go discord.Serve()
 
 	log.Info("Begin late initialization.")

	err = discord.LateInitialize()
	if err != nil {
		log.Fatalf("Unable to initialize Discord services, %s", err)
		cleanup()
		os.Exit(1)
	}

	log.Info("Late initialization completed.")

	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP, os.Interrupt, os.Kill)
	go func() {
		var exit int
		defer func() { state.ExitCode <- exit }()
		for {
			currentSignal := <-signalChannel
			switch currentSignal {
			case os.Interrupt:
				exit = 0
				println()
				log.Info("Gracefully exiting.")
				return
			default:
				exit = 0
				log.Info("Gracefully exiting.")
				return
			}
		}
	}()

	exitCode := <-state.ExitCode
	if exitCode != 0 {
		if exitCode == -1 {
			cleanup()
			restart()
		}
		abnormalExit()
		os.Exit(exitCode)
	}
}
