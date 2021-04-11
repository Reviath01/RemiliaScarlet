package multiplexer

func GetChannel(channel string) string {
	if len(channel) == 18 {
		return channel
	} else if len(channel) == 21 {
		return channel[2:][:18]
	}
	return channel
}
