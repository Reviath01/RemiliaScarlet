package multiplexer

import (
	"github.com/bwmarrin/discordgo"
)

func CheckBanPermission(s *discordgo.Session, userID string, channelID string) bool {
	p, _ := s.UserChannelPermissions(userID, channelID)
	return p&discordgo.PermissionBanMembers == discordgo.PermissionBanMembers
}

func CheckAdministratorPermission(s *discordgo.Session, userID string, channelID string) bool {
	p, _ := s.UserChannelPermissions(userID, channelID)
	return p&discordgo.PermissionAdministrator == discordgo.PermissionAdministrator
}

func CheckManageMessagesPermission(s *discordgo.Session, userID string, channelID string) bool {
	p, _ := s.UserChannelPermissions(userID, channelID)
	return p&discordgo.PermissionManageMessages == discordgo.PermissionManageMessages
}

func CheckKickPermission(s *discordgo.Session, userID string, channelID string) bool {
	p, _ := s.UserChannelPermissions(userID, channelID)
	return p&discordgo.PermissionKickMembers == discordgo.PermissionKickMembers
}
