package commandMap

import (
	"strings"

	command "git.randomchars.net/Reviath/RemiliaScarlet/CommandHandler/Command"
	commandGroup "git.randomchars.net/Reviath/RemiliaScarlet/CommandHandler/CommandGroup"
	ctx "git.randomchars.net/Reviath/RemiliaScarlet/CommandHandler/Context"
	embedutil "git.randomchars.net/Reviath/RemiliaScarlet/EmbedUtil"
	"github.com/bwmarrin/discordgo"
	_ "github.com/go-sql-driver/mysql"
)

type Map struct {
	commands map[string]command.Command
	groups   map[string]commandGroup.Group
}

func (m *Map) GetGroup(name string) commandGroup.Group {
	if m.DoesGroupExist(name) {
		return m.groups[name]
	}
	return nil
}

func (m *Map) DoesGroupExist(name string) bool {
	_, b := m.groups[name]
	return b
}

func (m *Map) Execute(command string, c ctx.Ctx, s *discordgo.Session) error {
	switch true {
	case m.CanExecute(command):
		return m.commands[strings.ToLower(command)].Execute(c, s)
	case m.DoesGroupExist(command):
		if len(c.Args()) > 0 {
			args, cmd := shift(c.Args(), 0)
			if m.GetGroup(command).CanExecute(cmd) {
				ct := ctx.New(args, c.Message(), s)
				return m.GetGroup(command).Execute(cmd, ct, s)
			}
		}
		fallthrough
	default:
		if len(command) < 1 {
			return nil
		}
		nocmd := embedutil.NewEmbed().
			SetDescription("No command match with: " + command).
			SetColor(0xff8c00).MessageEmbed
		_, _ = s.ChannelMessageSendEmbed(c.Channel().ID, nocmd)

		return nil
	}
}

func (m *Map) RegisterCommand(name string, command command.Command, override bool) {
	if m.CanRegisterCommand(name) || override {
		m.commands[strings.ToLower(name)] = command
	}
}

func (m *Map) CanRegisterCommand(name string) bool {
	return m.commands[name] == nil && m.GetGroup(name) == nil
}

func (m *Map) GetCommands() map[string]command.Command {
	return m.commands
}

func New() *Map {
	return &Map{commands: map[string]command.Command{}, groups: map[string]commandGroup.Group{}}
}

func (m *Map) CanExecute(name string) bool {
	_, ok := m.commands[name]
	return ok
}

func shift(a []string, i int) ([]string, string) {
	b := a[i]
	copy(a[i:], a[i+1:])
	a[len(a)-1] = ""
	a = a[:len(a)-1]
	return a, b
}
