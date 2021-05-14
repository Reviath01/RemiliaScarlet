package multiplexer

// GetChannel function provides to get a channel with id or mention
func GetChannel(channel string) string {
	switch len(channel) {
	case 18:
		return channel
	case 21:
		return channel[2:][:18]
	default:
		return channel
	}
}
