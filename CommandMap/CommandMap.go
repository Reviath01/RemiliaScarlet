package commandMap

import (
	"database/sql"
	"strings"

	command "git.randomchars.net/Reviath/RemiliaScarlet/Command"
	commandGroup "git.randomchars.net/Reviath/RemiliaScarlet/CommandGroup"
	ctx "git.randomchars.net/Reviath/RemiliaScarlet/Context"
	embedutil "git.randomchars.net/Reviath/embed-util"
	"github.com/bwmarrin/discordgo"
	_ "github.com/go-sql-driver/mysql"
)

type Map struct {
	commands map[string]command.Command
	groups   map[string]commandGroup.Group
}

func (m *Map) RegisterCommandGroup(name string, group commandGroup.Group) {
	if !m.DoesGroupExist(name) && m.CanRegisterGroup(name) {
		m.groups[name] = group
	}
}

func (m *Map) GetGroups() map[string]commandGroup.Group {
	return m.groups
}

func (m *Map) GetGroup(name string) commandGroup.Group {
	if m.DoesGroupExist(name) {
		return m.groups[name]
	}
	return nil
}

func (m *Map) CanRegisterGroup(name string) bool {
	return m.commands[name] == nil && m.GetGroup(name) == nil
}

func (m *Map) DoesGroupExist(name string) bool {
	_, b := m.groups[name]
	return b
}

func (m *Map) Execute(command string, c ctx.Ctx, s *discordgo.Session) error {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/remilia")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	type Tag struct {
		lang string
	}
	var tag Tag
	err = db.QueryRow("SELECT language FROM languages WHERE guildid ='" + c.Guild().ID + "'").Scan(&tag.lang)
	if err == nil {
		if tag.lang == "tr" {
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
					SetDescription(command + " isimli bir komut bulunamadı!").
					SetColor(0xff8c00).MessageEmbed
				_, _ = s.ChannelMessageSendEmbed(c.Channel().ID, nocmd)
				return nil
			}
		}
	} else {
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
	return nil
}

func (m *Map) GetAllCommands() map[string]command.Command {
	cs := m.GetCommands()
	for k, g := range m.GetGroups() {
		for name, cmd := range g.GetCommands() {
			cs[k+" "+name] = cmd
		}
	}
	return cs
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
