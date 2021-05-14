package handler

import "github.com/bwmarrin/discordgo"

// CommandFunc is function to run all commands
type CommandFunc func(Context, []string) error

// DebugFunc debug function
type DebugFunc func(string)

// HelpCommandFunc is help command function
type HelpCommandFunc func(Context, []string, []*Command, []string) error

// PrerunFunc is a function when command is used
type PrerunFunc func(Context, *Command, []string) bool

// OnErrorFunc is a function if error occurres on commands
type OnErrorFunc func(Context, *Command, []string, error)

// CommandType int
type CommandType int

// HelpCommand struct
type HelpCommand struct {
	Aliases         []string
	Name            string
	SelfPermissions int
	UserPermissions int
	Run             HelpCommandFunc
}

// CommandHandler structure
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

// Command structure
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

// Context structure
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
	CommandTypeEverywhere CommandType = iota // CommandTypeEverywhere : command can be used on everywhere (DMs + Guilds)

	CommandTypeGuild // CommandTypeGuild : command can be only used on guilds (not DMs)

	CommandTypePrivate // CommandTypePrivate : private command type
)
