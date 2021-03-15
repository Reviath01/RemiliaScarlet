package types


import "github.com/bwmarrin/discordgo"

type CommandFunc func(Context, []string) error

type DebugFunc func(string)

type HelpCommandFunc func(Context, []string, []*Command, []string) error
type PrerunFunc func(Context, *Command, []string) bool
type OnErrorFunc func(Context, *Command, []string, error)

type CommandType int

type HelpCommand struct {
	Aliases         []string
	Name            string
	SelfPermissions int
	UserPermissions int
	Run             HelpCommandFunc
}

type CommandHandler struct {
	enabled          bool
	checkPermissions bool
	ignoreBots       bool
	useState         bool

	owners   []string
	prefixes []string

	commands    []*Command
	helpCommand *HelpCommand

	debugFunc   DebugFunc
	onErrorFunc OnErrorFunc
	prerunFunc  PrerunFunc
}

type Command struct {
	Aliases     []string
	Description string
	Name        string

	Hidden    bool
	OwnerOnly bool

	SelfPermissions int
	UserPermissions int

	Run CommandFunc

	Type CommandType
}

type Context struct {
	Handler *CommandHandler

	Channel *discordgo.Channel

	Guild *discordgo.Guild

	Member *discordgo.Member

	Message *discordgo.Message

	Session *discordgo.Session

	User *discordgo.User
}

const (
	CommandTypeEverywhere CommandType = iota

	CommandTypeGuild

	CommandTypePrivate
)
