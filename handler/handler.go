package handler

import (
	"errors"
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func (c *CommandHandler) GetCheckPermissions(enable bool) bool {
	return c.checkPermissions
}

func (c *CommandHandler) SetCheckPermissions(enable bool) {
	c.checkPermissions = enable
}

func (c *CommandHandler) GetEnable() bool {
	return c.enabled
}

func (c *CommandHandler) SetEnable(enable bool) {
	c.enabled = enable
}

func (c *CommandHandler) GetIgnoreBots() bool {
	return c.ignoreBots
}

func (c *CommandHandler) SetIgnoreBots(enable bool) {
	c.ignoreBots = enable
}

func (c *CommandHandler) GetUseState() bool {
	return c.useState
}

func (c *CommandHandler) SetUseState(enable bool) {
	c.useState = enable
}

func (c *CommandHandler) AddPrefix(prefix string) {
	c.prefixes = append(c.prefixes, prefix)
}

func (c *CommandHandler) RemovePrefix(prefix string) {
	for i, v := range c.prefixes {
		if v == prefix {
			copy(c.prefixes[i:], c.prefixes[i+1:])
			c.prefixes[len(c.prefixes)-1] = ""
			c.prefixes = c.prefixes[:len(c.prefixes)-1]

			break
		}
	}
}

func (c *CommandHandler) GetAllPrefixes() []string {
	return c.prefixes
}

func (c *CommandHandler) SetAllPrefixes(prefixes []string) {
	c.prefixes = prefixes
}

func (c *CommandHandler) ClearDebugFunc() {
	c.debugFunc = nil
}

func (c *CommandHandler) GetDebugFunc() DebugFunc {
	return c.debugFunc
}

func (c *CommandHandler) SetDebugFunc(df DebugFunc) {
	c.debugFunc = df
}

func (c *CommandHandler) ClearOnErrorFunc() {
	c.onErrorFunc = nil
}

func (c *CommandHandler) GetOnErrorFunc() OnErrorFunc {
	return c.onErrorFunc
}

func (c *CommandHandler) SetOnErrorFunc(oef OnErrorFunc) {
	c.onErrorFunc = oef
}

func (c *CommandHandler) ClearPrerunFunc() {
	c.prerunFunc = func(_ Context, _ *Command, _ []string) bool {
		return true
	}
}

func (c *CommandHandler) GetPrerunFunc() PrerunFunc {
	return c.prerunFunc
}

func (c *CommandHandler) SetPrerunFunc(prf PrerunFunc) {
	c.prerunFunc = prf
}

func (c *CommandHandler) AddCommand(name, desc string, aliases []string, owneronly, hidden bool, selfperms, userperms int, cmdtype CommandType, run CommandFunc) error {
	for _, v := range c.commands {
		if v.Name == name {
			return ErrCommandAlreadyRegistered
		}
	}

	if c.helpCommand != nil && c.helpCommand.Name == name {
		return ErrCommandAlreadyRegistered
	}

	c.commands = append(c.commands, &Command{
		Aliases:         aliases,
		Description:     desc,
		Hidden:          hidden,
		Name:            name,
		OwnerOnly:       owneronly,
		SelfPermissions: selfperms,
		UserPermissions: userperms,
		Run:             run,
		Type:            cmdtype,
	})

	return nil
}

func (c *CommandHandler) RemoveCommand(name string) error {
	for i, v := range c.commands {
		if v.Name == name {
			copy(c.commands[i:], c.commands[i+1:])
			c.commands[len(c.commands)-1] = nil
			c.commands = c.commands[:len(c.commands)-1]

			return nil
		}
	}

	return ErrCommandNotFound
}

func (c *CommandHandler) GetHelpCommand() *HelpCommand {
	return c.helpCommand
}

func (c *CommandHandler) SetHelpCommand(name string, aliases []string, selfperms, userperms int, function HelpCommandFunc) error {
	for _, v := range c.commands {
		if v.Name == name {
			return ErrCommandAlreadyRegistered
		}
	}

	c.helpCommand = &HelpCommand{
		Aliases:         aliases,
		Name:            name,
		SelfPermissions: selfperms,
		UserPermissions: userperms,
		Run:             function,
	}

	return nil
}

func (c *CommandHandler) ClearHelpCommand() {
	c.helpCommand = nil
}

func (c *CommandHandler) AddOwner(id string) {
	c.owners = append(c.owners, id)
}

func (c *CommandHandler) RemoveOwner(id string) {
	for i, v := range c.owners {
		if v == id {
			c.owners = append(c.owners[:i], c.owners[i+1:]...)
			break
		}
	}
}

func (c *CommandHandler) SetOwners(ids []string) {
	c.owners = ids
}

func (c *CommandHandler) GetOwners() []string {
	return c.owners
}

func (c *CommandHandler) IsOwner(id string) bool {
	for _, o := range c.owners {
		if id == o {
			return true
		}
	}

	return false
}

func (c *CommandHandler) debugLog(format string, a ...interface{}) {
	if c.debugFunc != nil {
		c.debugFunc(fmt.Sprintf(format, a...))
	}
}

func (c *CommandHandler) throwError(context Context, command *Command, args []string, err error) {
	if c.onErrorFunc != nil {
		c.onErrorFunc(context, command, args, err)
	}
}

func permissionCheck(session *discordgo.Session, member *discordgo.Member, guild *discordgo.Guild, channel *discordgo.Channel, necessaryPermissions int, useState bool) error {
	if necessaryPermissions == 0 {
		return nil
	}

	var permissions int

	if member.User.ID == guild.OwnerID {
		return nil
	}

	permissions |= int(guild.Roles[0].Permissions)

	for _, roleID := range member.Roles {
		var (
			role *discordgo.Role
			err  error
		)

		if session.StateEnabled && useState {
			role, err = session.State.Role(guild.ID, roleID)
			if err != nil {
				return err
			}
		} else {
			roles, err := session.GuildRoles(guild.ID)
			if err != nil {
				return err
			}

			for _, v := range roles {
				if v.ID == roleID {
					role = v
					break
				}
			}

			if role == nil {
				return ErrDataUnavailable
			}
		}

		if role.Permissions&discordgo.PermissionAdministrator == discordgo.PermissionAdministrator {
			return nil
		}

		permissions |= int(role.Permissions)
	}

	for _, overwrite := range channel.PermissionOverwrites {
		if overwrite.ID == member.User.ID {
			permissions |= int(overwrite.Allow)
			permissions = permissions &^ int(overwrite.Deny)
		}

		for _, roleID := range member.Roles {
			if overwrite.ID == roleID {
				permissions |= int(overwrite.Allow)
				permissions = permissions &^ int(overwrite.Deny)
			}
		}
	}

	if permissions&necessaryPermissions != necessaryPermissions {
		return errors.New("insufficient permissions")
	}

	return nil
}

func (c *CommandHandler) MessageHandler(s *discordgo.Session, event *discordgo.MessageCreate) {
	if !c.enabled || event.Author.ID == s.State.User.ID {
		return
	}

	c.debugLog("Received message (%s) by user \"%s\" (%s): \"%s\"", event.Message.ID, event.Author.String(), event.Author.ID, event.Message.Content)

	var (
		has    bool
		prefix string

		command *Command
		help    *HelpCommand

		context Context

		channel    *discordgo.Channel
		guild      *discordgo.Guild
		member     *discordgo.Member
		selfMember *discordgo.Member

		err error
	)

	context.Handler = c
	context.Message = event.Message
	context.Session = s
	context.User = event.Author

	for i := 0; i < len(c.prefixes); i++ {
		prefix = c.prefixes[i]

		if strings.HasPrefix(event.Message.Content, prefix) {
			has = true
			break
		}
	}

	if !has {
		c.debugLog("Message %s doesn't contain any of the prefixes", event.Message.ID)
		return
	}

	content := strings.Split(strings.TrimPrefix(event.Message.Content, prefix), " ")
	c.debugLog("Parsed message \"%s\": \"%s\"", event.Message.ID, content)

	if content[0] == c.helpCommand.Name {
		help = c.helpCommand
	}

	if help == nil {
		for _, v := range c.helpCommand.Aliases {
			if content[0] == v {
				help = c.helpCommand
			}
		}
	}

	if help != nil {
		if channel, err = s.Channel(event.ChannelID); err != nil {
			c.debugLog("Failed to fetch the current channel: \"%s\"", err.Error())
			c.throwError(context, &Command{
				Name:            c.helpCommand.Name,
				SelfPermissions: c.helpCommand.SelfPermissions,
				UserPermissions: c.helpCommand.UserPermissions,
				Type:            CommandTypeEverywhere,
			}, content[1:], ErrDataUnavailable)
			return
		}

		context.Channel = channel

		if channel.Type == discordgo.ChannelTypeDM {
			if c.prerunFunc != nil && !c.prerunFunc(context, &Command{
				Name:            c.helpCommand.Name,
				SelfPermissions: c.helpCommand.SelfPermissions,
				UserPermissions: c.helpCommand.UserPermissions,
				Type:            CommandTypeEverywhere,
			}, content[1:]) {
				return
			}

			if err = help.Run(context, content[1:], c.commands, c.prefixes); err != nil {
				c.throwError(context, &Command{
					Name:            c.helpCommand.Name,
					SelfPermissions: c.helpCommand.SelfPermissions,
					UserPermissions: c.helpCommand.UserPermissions,
					Type:            CommandTypeEverywhere,
				}, content[1:], err)
			}

			return
		}

		if guild, err = s.Guild(event.GuildID); err != nil {
			c.debugLog("Failed to fetch the current guild: \"%s\"", err.Error())
			c.throwError(context, &Command{
				Name:            c.helpCommand.Name,
				SelfPermissions: c.helpCommand.SelfPermissions,
				UserPermissions: c.helpCommand.UserPermissions,
				Type:            CommandTypeEverywhere,
			}, content[1:], ErrDataUnavailable)
			return
		}

		if member, err = s.GuildMember(event.GuildID, event.Author.ID); err != nil {
			c.debugLog("Failed to fetch the user as a guild member: \"%s\"", err.Error())
			c.throwError(context, &Command{
				Name:            c.helpCommand.Name,
				SelfPermissions: c.helpCommand.SelfPermissions,
				UserPermissions: c.helpCommand.UserPermissions,
				Type:            CommandTypeEverywhere,
			}, content[1:], ErrDataUnavailable)
			return
		}

		context.Guild = guild
		context.Member = member

		if selfMember, err = s.GuildMember(event.GuildID, s.State.User.ID); err != nil {
			c.debugLog("Failed to fetch the bot as a guild member: \"%s\"", err.Error())
			c.throwError(context, command, content[1:], ErrDataUnavailable)
			return
		}

		if c.checkPermissions {
			if err = permissionCheck(s, member, guild, channel, help.UserPermissions, c.useState); err != nil {
				c.debugLog("Insufficient permissions (user): \"%s\"", err.Error())
				c.throwError(context, command, content[1:], ErrUserInsufficientPermissions)
				return
			}

			if err = permissionCheck(s, selfMember, guild, channel, help.SelfPermissions, c.useState); err != nil {
				c.debugLog("Insufficient permissions (bot): \"%s\"", err.Error())
				c.throwError(context, command, content[1:], ErrUserInsufficientPermissions)
				return
			}
		}

		if !c.prerunFunc(context, &Command{
			Name:            c.helpCommand.Name,
			SelfPermissions: c.helpCommand.SelfPermissions,
			UserPermissions: c.helpCommand.UserPermissions,
			Type:            CommandTypeEverywhere,
		}, content[1:]) {
			return
		}

		if err = help.Run(context, content[1:], c.commands, c.prefixes); err != nil {
			c.throwError(context, command, content[1:], err)
		}

		return
	}

	for _, v := range c.commands {
		if content[0] == v.Name {
			command = v
			break
		}

		for _, alias := range v.Aliases {
			if content[0] == alias {
				command = v
				break
			}
		}
	}

	if command == nil {
		c.throwError(context, nil, content[1:], ErrCommandNotFound)
		return
	}

	if command.OwnerOnly && !c.IsOwner(event.Author.ID) {
		c.debugLog("The user tried to run an owner-only command")
		c.throwError(context, command, content[1:], ErrOwnerOnly)
		return
	}

	if channel, err = s.Channel(event.ChannelID); err != nil {
		c.debugLog("Failed to fetch the current channel: \"%s\"", err.Error())
		c.throwError(context, command, content[1:], ErrDataUnavailable)
		return
	}

	context.Channel = channel

	if channel.Type == discordgo.ChannelTypeDM {
		if command.Type == CommandTypeGuild {
			c.debugLog("The user tried to execute a guild-only command in the DMs")
			c.throwError(context, command, content[1:], ErrGuildOnly)
			return
		}

		if c.prerunFunc != nil && !c.prerunFunc(context, command, content[1:]) {
			return
		}

		if err = command.Run(context, content[1:]); err != nil {
			c.throwError(context, command, content[1:], err)
		}

		return
	}

	if guild, err = s.Guild(event.GuildID); err != nil {
		c.debugLog("Failed to fetch the current guild: \"%s\"", err.Error())
		c.throwError(context, command, content[1:], ErrDataUnavailable)
		return
	}

	if member, err = s.GuildMember(event.GuildID, event.Author.ID); err != nil {
		c.debugLog("Failed to fetch the user as a guild member: \"%s\"", err.Error())
		c.throwError(context, command, content[1:], ErrDataUnavailable)
		return
	}

	if command.Type == CommandTypePrivate && guild != nil {
		c.debugLog("The user tried to execute a DM-only command outside the DMs")
		c.throwError(context, command, content[1:], ErrDMOnly)
		return
	}

	if command.Type == CommandTypeGuild && guild == nil {
		c.debugLog("The user tried to execute a guild-only command outside a guild")
		c.throwError(context, command, content[1:], ErrGuildOnly)
		return
	}

	context.Guild = guild
	context.Member = member

	if selfMember, err = s.GuildMember(event.GuildID, s.State.User.ID); err != nil {
		c.debugLog("Failed to fetch the bot as a guild member: \"%s\"", err.Error())
		c.throwError(context, command, content[1:], ErrDataUnavailable)
		return
	}

	if c.checkPermissions {
		if err = permissionCheck(s, member, guild, channel, command.UserPermissions, c.useState); err != nil {
			c.debugLog("Insufficient permissions (user): \"%s\"", err.Error())
			c.throwError(context, command, content[1:], ErrUserInsufficientPermissions)
			return
		}

		if err = permissionCheck(s, selfMember, guild, channel, command.SelfPermissions, c.useState); err != nil {
			c.debugLog("Insufficient permissions (bot): \"%s\"", err.Error())
			c.throwError(context, command, content[1:], ErrUserInsufficientPermissions)
			return
		}
	}

	if c.prerunFunc != nil && !c.prerunFunc(context, command, content[1:]) {
		return
	}

	if err = command.Run(context, content[1:]); err != nil {
		c.throwError(context, command, content[1:], err)
	}
}
