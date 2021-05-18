package multiplexer

import "git.randomchars.net/Reviath/RemiliaScarlet/config"

// GetPrefix returns to Bot's prefix from config package
func GetPrefix() string {
	return config.BotPrefix
}
