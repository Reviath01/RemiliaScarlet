package multiplexer

var (
	BotPrefix string        // BotPrefix is bot's prefix to cut prefix's length
	config    *configStruct // config structure
)

type configStruct struct {
	BotPrefix string `json:"BotPrefix"`
}

// GetPrefix returns to Bot's prefix from config package
func GetPrefix() string {
	return config.BotPrefix
}
