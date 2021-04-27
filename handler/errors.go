package handler

import "errors"

var (
	ErrBotBlocked = errors.New("bots are ignored")

	ErrCommandAlreadyRegistered = errors.New("another command was already registered by this name")

	ErrCommandNotFound = errors.New("command not found")

	ErrDataUnavailable = errors.New("necessary data couldn't be fetched")

	ErrDMOnly = errors.New("dm-only command on guild")

	ErrGuildOnly = errors.New("guild-only command on dm")

	ErrOwnerOnly = errors.New("owner-only command")

	ErrSelfInsufficientPermissions = errors.New("missing permission(s) for bot")

	ErrUserInsufficientPermissions = errors.New("missing permission(s) for user")
)
