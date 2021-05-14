package handler

import "errors"

var (
	ErrBotBlocked = errors.New("bots are ignored") // If bots are ignored return to error

	ErrCommandAlreadyRegistered = errors.New("another command was already registered by this name") // If another command is already registered by this name return to error

	ErrCommandNotFound = errors.New("command not found") // If command is not found return to error

	ErrDataUnavailable = errors.New("necessary data couldn't be fetched") // ErrDataUnavailable : If bot cannot fetch data return to error

	ErrDMOnly = errors.New("dm-only command on guild") // ErrDMOnly : If DM only command is used on guild return to error

	ErrGuildOnly = errors.New("guild-only command on dm") // ErrGuildOnly : If Guild only command is used on DM return to error

	ErrOwnerOnly = errors.New("owner-only command") // ErrOwnerOnly : If someone is trying to run owner only command return to error

	ErrSelfInsufficientPermissions = errors.New("missing permission(s) for bot") // ErrSelfInsufficientPermissions : If bot doesn't have enough permission return to error

	ErrUserInsufficientPermissions = errors.New("missing permission(s) for user") // ErrUserInsufficientPermissions : If user doesn't have enough permission return to error
)
