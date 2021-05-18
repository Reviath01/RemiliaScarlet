package multiplexer

import (
	"github.com/bwmarrin/discordgo"
)

// CheckBanPermission checks user's permissions on channel, returns to "true" or "false"
func CheckBanPermission(s *discordgo.Session, userID string, guildid string) bool {
	guild, _ := s.Guild(guildid)
	p, _ := s.UserChannelPermissions(userID, guild.Channels[0].ID)
	return p&discordgo.PermissionBanMembers == discordgo.PermissionBanMembers
}

// CheckAdministratorPermission checks user's permissions on channel, returns to "true" or "false"
func CheckAdministratorPermission(s *discordgo.Session, userID string, guildid string) bool {
	guild, _ := s.Guild(guildid)
	p, _ := s.UserChannelPermissions(userID, guild.Channels[0].ID)
	return p&discordgo.PermissionAdministrator == discordgo.PermissionAdministrator
}

// CheckManageMessagesPermission checks user's permissions on channel, returns to "true" or "false"
func CheckManageMessagesPermission(s *discordgo.Session, userID string, guildid string) bool {
	guild, _ := s.Guild(guildid)
	p, _ := s.UserChannelPermissions(userID, guild.Channels[0].ID)
	return p&discordgo.PermissionManageMessages == discordgo.PermissionManageMessages
}

// CheckKickPermission checks user's permissions on channel, returns to "true" or "false"
func CheckKickPermission(s *discordgo.Session, userID string, guildid string) bool {
	guild, _ := s.Guild(guildid)
	p, _ := s.UserChannelPermissions(userID, guild.Channels[0].ID)
	return p&discordgo.PermissionKickMembers == discordgo.PermissionKickMembers
}
