package handler

// New returns a new CommandHandler
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
