package newerrors

import "errors"

var (
	ErrBotBlocked = errors.New("I'm not responding to bots")
	ErrCommandAlreadyRegistered = errors.New("There is a command already named with this name.")
	ErrCommandNotFound = errors.New("I can't find that command.")
	ErrDataUnavailable = errors.New("Necessary data couldn't be fetched.")
	ErrDMOnly = errors.New("This command is only for DMs")
	ErrGuildOnly = errors.New("This command is only available for guilds.")
	ErrOwnerOnly = errors.New("You don't own this bot.")
	ErrSelfInsufficientPermissions = errors.New("I don't have enough permission.")
	ErrUserInsufficientPermissions = errors.New("You don't have enough permission to run this command.")
)
