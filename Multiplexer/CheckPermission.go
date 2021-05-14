package multiplexer

import (
	"github.com/bwmarrin/discordgo"
)

//Checking ban permission for user on specified channel.
func CheckBanPermission(s *discordgo.Session, userID string, channelID string) bool {
	p, _ := s.UserChannelPermissions(userID, channelID)
	return p&discordgo.PermissionBanMembers == discordgo.PermissionBanMembers
}

//Checking administrator permission for user on specified channel.
func CheckAdministratorPermission(s *discordgo.Session, userID string, channelID string) bool {
	p, _ := s.UserChannelPermissions(userID, channelID)
	return p&discordgo.PermissionAdministrator == discordgo.PermissionAdministrator
}

//Checking manage messages permission for user on specified channel.
func CheckManageMessagesPermission(s *discordgo.Session, userID string, channelID string) bool {
	p, _ := s.UserChannelPermissions(userID, channelID)
	return p&discordgo.PermissionManageMessages == discordgo.PermissionManageMessages
}

//Checking kick permission for user on specified channel.
func CheckKickPermission(s *discordgo.Session, userID string, channelID string) bool {
	p, _ := s.UserChannelPermissions(userID, channelID)
	return p&discordgo.PermissionKickMembers == discordgo.PermissionKickMembers
}
