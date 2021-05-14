package multiplexer

import (
	"github.com/bwmarrin/discordgo"
)

// CheckBanPermission checks user's permissions on channel, returns to "true" or "false"
func CheckBanPermission(s *discordgo.Session, userID string, channelID string) bool {
	p, _ := s.UserChannelPermissions(userID, channelID)
	return p&discordgo.PermissionBanMembers == discordgo.PermissionBanMembers
}

// CheckAdministratorPermission checks user's permissions on channel, returns to "true" or "false"
func CheckAdministratorPermission(s *discordgo.Session, userID string, channelID string) bool {
	p, _ := s.UserChannelPermissions(userID, channelID)
	return p&discordgo.PermissionAdministrator == discordgo.PermissionAdministrator
}

// CheckManageMessagesPermission checks user's permissions on channel, returns to "true" or "false"
func CheckManageMessagesPermission(s *discordgo.Session, userID string, channelID string) bool {
	p, _ := s.UserChannelPermissions(userID, channelID)
	return p&discordgo.PermissionManageMessages == discordgo.PermissionManageMessages
}

// CheckKickPermission checks user's permissions on channel, returns to "true" or "false"
func CheckKickPermission(s *discordgo.Session, userID string, channelID string) bool {
	p, _ := s.UserChannelPermissions(userID, channelID)
	return p&discordgo.PermissionKickMembers == discordgo.PermissionKickMembers
}
