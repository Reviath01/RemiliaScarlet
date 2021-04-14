package handler

import "errors"

var (
	ErrBotBlocked = errors.New("handler: The given author was a bot and the IgnoreBots setting is true")

	ErrCommandAlreadyRegistered = errors.New("handler: Another command was already registered by this name")

	ErrCommandNotFound = errors.New("handler: Command not found")

	ErrDataUnavailable = errors.New("handler: Necessary data couldn't be fetched")

	ErrDMOnly = errors.New("handler: DM-Only command on guild")

	ErrGuildOnly = errors.New("handler: Guild-Only command in DMs")

	ErrOwnerOnly = errors.New("handler: Owner-Only command")

	ErrSelfInsufficientPermissions = errors.New("handler: Insufficient permissions for the bot")

	ErrUserInsufficientPermissions = errors.New("handler: Insufficient permissions for the user")
)
