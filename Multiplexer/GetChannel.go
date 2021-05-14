package multiplexer

//Get channel function for commands.
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
