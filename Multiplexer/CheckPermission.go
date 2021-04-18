package multiplexer

import "github.com/bwmarrin/discordgo"

func CheckPermission(s *discordgo.Session, userID string, channelID string) bool {
	p, err := s.State.UserChannelPermissions(userID, channelID)
	return (err == nil) && (p&discordgo.PermissionBanMembers == discordgo.PermissionBanMembers)
}
