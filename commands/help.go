package commands

import (
	"fmt"
	"sort"
	"strings"

	CommandHandler "git.randomchars.net/Reviath/RemiliaScarlet/handler"
	"github.com/bwmarrin/discordgo"
)

func HelpCommand(context CommandHandler.Context, args []string, commands []*CommandHandler.Command, prefixes []string) error {
	typeCheck := func(chn discordgo.ChannelType, cmd CommandHandler.CommandType) bool {
		switch cmd {
		case CommandHandler.CommandTypeEverywhere:
			return true

		case CommandHandler.CommandTypePrivate:
			if chn == discordgo.ChannelTypeDM {
				return true
			}

			break

		case CommandHandler.CommandTypeGuild:
			if chn == discordgo.ChannelTypeGuildText {
				return true
			}

			break
		}

		return false
	}

	if len(args) >= 1 {
		for _, command := range commands {
			if args[0] != command.Name {
				continue
			}

			if command.Hidden || !typeCheck(context.Channel.Type, command.Type) {
				return nil
			}

			var (
				owneronlystring = "No"
				typestring      = "Anywhere"
			)

			if command.OwnerOnly {
				owneronlystring = "Yes"
			}

			switch command.Type {
			case CommandHandler.CommandTypePrivate:
				typestring = "Private"
				break

			case CommandHandler.CommandTypeGuild:
				typestring = "Guild-only"
				break
			}

			prefixesBuilder := strings.Builder{}
			if len(prefixes) == 1 {
				prefixesBuilder.WriteString(fmt.Sprintf("The prefix is %s", prefixes[0]))
			} else {
				prefixesBuilder.WriteString("The prefixes are ")

				for i, prefix := range prefixes {
					if i+1 == len(prefixes) {
						prefixesBuilder.WriteString(fmt.Sprintf("and %s", prefix))
					} else {
						prefixesBuilder.WriteString(fmt.Sprintf("%s, ", prefix))
					}
				}
			}

			aliases := "**None.**"
			if len(command.Aliases) > 0 {
				aliases = strings.Join(command.Aliases, "`, `")
				aliases = "`" + aliases + "`"
			}

			_, err := context.ReplyEmbed(&discordgo.MessageEmbed{
				Title:       "Help",
				Color:       0x08a4ff,
				Description: fmt.Sprintf("**%s**\nAliases: %s\nOwner only: **%s**\nUsable: **%s**", command.Description, aliases, owneronlystring, typestring),
				Footer: &discordgo.MessageEmbedFooter{
					Text: fmt.Sprintf(" %s.", prefixesBuilder.String()),
				},
			})

			return err
		}

		_, err := context.Reply("Command `" + args[0] + "` doesn't exist.")
		return err
	}

	var (
		count          int
		commandsSorted = make([]*CommandHandler.Command, len(commands))
		embed          = &discordgo.MessageEmbed{
			Title: "Commands",
			Color: 0x08a4ff,
		}
		names = make([]string, len(commands))
	)

	for i, cmd := range commands {
		names[i] = cmd.Name
	}

	sort.Strings(names)

	for i, v := range names {
		for _, v2 := range commands {
			if v2.Name == v {
				commandsSorted[i] = v2
				break
			}
		}

		if commandsSorted[i] == nil {
			return fmt.Errorf("Sort failure")
		}
	}

	for _, cmd := range commandsSorted {
		if !cmd.Hidden && typeCheck(context.Channel.Type, cmd.Type) {
			embed.Fields = append(embed.Fields, &discordgo.MessageEmbedField{
				Name:   cmd.Name,
				Value:  cmd.Description,
				Inline: count%2 == 0,
			})

			count++
		}
	}

	var footer strings.Builder

	if count == 1 {
		footer.WriteString("There is 1 command.")
	} else {
		footer.WriteString(fmt.Sprintf("There are %d commands.", count))
	}

	footer.WriteString(" | ")

	if len(prefixes) == 1 {
		footer.WriteString(fmt.Sprintf("The prefix is %s.", prefixes[0]))
	} else {
		prefixesBuilder := strings.Builder{}

		for i, prefix := range prefixes {
			if i+1 == len(prefixes) {
				prefixesBuilder.WriteString(fmt.Sprintf("and %s", prefix))
			} else {
				prefixesBuilder.WriteString(fmt.Sprintf("%s, ", prefix))
			}
		}

		footer.WriteString(fmt.Sprintf("The prefixes are %s.", prefixesBuilder.String()))
	}

	embed.Footer = &discordgo.MessageEmbedFooter{Text: footer.String()}

	_, err := context.ReplyEmbed(embed)
	return err
}
