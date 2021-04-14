package handler

import "errors"

var (
	ErrBotBlocked = errors.New("I'm ignoring bots.")

	ErrCommandAlreadyRegistered = errors.New("Another command was already registered by this name")

	ErrCommandNotFound = errors.New("Command not found")

	ErrDataUnavailable = errors.New("Necessary data couldn't be fetched")

	ErrDMOnly = errors.New("This command is only available on DMs")

	ErrGuildOnly = errors.New("This command is only available on guilds")

	ErrOwnerOnly = errors.New("Owner-Only command")

	ErrSelfInsufficientPermissions = errors.New("I don't have enough permission.")

	ErrUserInsufficientPermissions = errors.New("You don't have enough permission.")
)
