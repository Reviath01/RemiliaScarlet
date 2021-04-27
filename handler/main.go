package handler

import (
	"os"
	"os/signal"
	"syscall"
)

func New(prefixes []string, owners []string, useState, ignoreBots, checkPermssions bool) CommandHandler {
	return CommandHandler{
		enabled:          true,
		prefixes:         prefixes,
		owners:           owners,
		useState:         useState,
		ignoreBots:       ignoreBots,
		checkPermissions: checkPermssions,
	}
}

func WaitForInterrupt() {
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
}
