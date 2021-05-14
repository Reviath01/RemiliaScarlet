package interactioncommands

import (
	"fmt"

	multiplexer "git.randomchars.net/Reviath/RemiliaScarlet/Multiplexer"
	"git.randomchars.net/Reviath/RemiliaScarlet/interactions"
	"git.randomchars.net/Reviath/RemiliaScarlet/sql"
	"github.com/bwmarrin/discordgo"
)

// EnableCommand is enable command for interactions
func EnableCommand(session *discordgo.Session, interaction interactions.Interaction) interactions.InteractionResponse {
	db := sql.Connect()
	defer db.Close()

	type Tag struct {
		isblocked string
	}

	var tag Tag

	if sql.CheckLanguage(interaction.GuildID) == "tr" {
		if !multiplexer.CheckAdministratorPermission(session, interaction.Member.User.ID, interaction.ChannelID) {
			return multiplexer.CreateResponse("Yeterli yetkiye sahip değilsin")
		}

		args := interaction.Data.Options[0].Value.(string)

		switch args {
		case "afk":
			err := db.QueryRow(fmt.Sprintf("SELECT isblocked FROM disabledcommands WHERE commandname ='%s' AND guildid ='%s'", args, interaction.GuildID)).Scan(&tag.isblocked)
			if err == nil {
				if tag.isblocked == "True" {
					delete, err := db.Query(fmt.Sprintf("DELETE FROM disabledcommands WHERE guildid ='%s' AND commandname ='%s'", interaction.GuildID, args))
					if err != nil {
						return multiplexer.CreateResponse("Bir hata oluştu")
					}
					defer delete.Close()
					return multiplexer.CreateResponse(args + " komutu tekrar kullanılabilir.")
				}
				return multiplexer.CreateResponse("Bu komut engellenmemiş")
			}
			return multiplexer.CreateResponse("Bu komut engellenmemiş.")
		case "author":
			err := db.QueryRow(fmt.Sprintf("SELECT isblocked FROM disabledcommands WHERE commandname ='%s' AND guildid ='%s'", args, interaction.GuildID)).Scan(&tag.isblocked)
			if err == nil {
				if tag.isblocked == "True" {
					delete, err := db.Query(fmt.Sprintf("DELETE FROM disabledcommands WHERE guildid ='%s' AND commandname ='%s'", interaction.GuildID, args))
					if err != nil {
						return multiplexer.CreateResponse("Bir hata oluştu")
					}
					defer delete.Close()
					return multiplexer.CreateResponse(args + " komutu tekrar kullanılabilir.")
				}
				return multiplexer.CreateResponse("Bu komut engellenmemiş")
			}
			return multiplexer.CreateResponse("Bu komut engellenmemiş.")
		case "avatar":
			err := db.QueryRow(fmt.Sprintf("SELECT isblocked FROM disabledcommands WHERE commandname ='%s' AND guildid ='%s'", args, interaction.GuildID)).Scan(&tag.isblocked)
			if err == nil {
				if tag.isblocked == "True" {
					delete, err := db.Query(fmt.Sprintf("DELETE FROM disabledcommands WHERE guildid ='%s' AND commandname ='%s'", interaction.GuildID, args))
					if err != nil {
						return multiplexer.CreateResponse("Bir hata oluştu")
					}
					defer delete.Close()
					return multiplexer.CreateResponse(args + " komutu tekrar kullanılabilir.")
				}
				return multiplexer.CreateResponse("Bu komut engellenmemiş")
			}
			return multiplexer.CreateResponse("Bu komut engellenmemiş.")
		case "ban":
			err := db.QueryRow(fmt.Sprintf("SELECT isblocked FROM disabledcommands WHERE commandname ='%s' AND guildid ='%s'", args, interaction.GuildID)).Scan(&tag.isblocked)
			if err == nil {
				if tag.isblocked == "True" {
					delete, err := db.Query(fmt.Sprintf("DELETE FROM disabledcommands WHERE guildid ='%s' AND commandname ='%s'", interaction.GuildID, args))
					if err != nil {
						return multiplexer.CreateResponse("Bir hata oluştu")
					}
					defer delete.Close()
					return multiplexer.CreateResponse(args + " komutu tekrar kullanılabilir.")
				}
				return multiplexer.CreateResponse("Bu komut engellenmemiş")
			}
			return multiplexer.CreateResponse("Bu komut engellenmemiş.")
		case "embed":
			err := db.QueryRow(fmt.Sprintf("SELECT isblocked FROM disabledcommands WHERE commandname ='%s' AND guildid ='%s'", args, interaction.GuildID)).Scan(&tag.isblocked)
			if err == nil {
				if tag.isblocked == "True" {
					delete, err := db.Query(fmt.Sprintf("DELETE FROM disabledcommands WHERE guildid ='%s' AND commandname ='%s'", interaction.GuildID, args))
					if err != nil {
						return multiplexer.CreateResponse("Bir hata oluştu")
					}
					defer delete.Close()
					return multiplexer.CreateResponse(args + " komutu tekrar kullanılabilir.")
				}
				return multiplexer.CreateResponse("Bu komut engellenmemiş")
			}
			return multiplexer.CreateResponse("Bu komut engellenmemiş.")
		case "Channel_info":
			err := db.QueryRow(fmt.Sprintf("SELECT isblocked FROM disabledcommands WHERE commandname ='%s' AND guildid ='%s'", args, interaction.GuildID)).Scan(&tag.isblocked)
			if err == nil {
				if tag.isblocked == "True" {
					delete, err := db.Query(fmt.Sprintf("DELETE FROM disabledcommands WHERE guildid ='%s' AND commandname ='%s'", interaction.GuildID, args))
					if err != nil {
						return multiplexer.CreateResponse("Bir hata oluştu")
					}
					defer delete.Close()
					return multiplexer.CreateResponse(args + " komutu tekrar kullanılabilir.")
				}
				return multiplexer.CreateResponse("Bu komut engellenmemiş")
			}
			return multiplexer.CreateResponse("Bu komut engellenmemiş.")
		case "hug":
			err := db.QueryRow(fmt.Sprintf("SELECT isblocked FROM disabledcommands WHERE commandname ='%s' AND guildid ='%s'", args, interaction.GuildID)).Scan(&tag.isblocked)
			if err == nil {
				if tag.isblocked == "True" {
					delete, err := db.Query(fmt.Sprintf("DELETE FROM disabledcommands WHERE guildid ='%s' AND commandname ='%s'", interaction.GuildID, args))
					if err != nil {
						return multiplexer.CreateResponse("Bir hata oluştu")
					}
					defer delete.Close()
					return multiplexer.CreateResponse(args + " komutu tekrar kullanılabilir.")
				}
				return multiplexer.CreateResponse("Bu komut engellenmemiş")
			}
			return multiplexer.CreateResponse("Bu komut engellenmemiş.")
		case "kick":
			err := db.QueryRow(fmt.Sprintf("SELECT isblocked FROM disabledcommands WHERE commandname ='%s' AND guildid ='%s'", args, interaction.GuildID)).Scan(&tag.isblocked)
			if err == nil {
				if tag.isblocked == "True" {
					delete, err := db.Query(fmt.Sprintf("DELETE FROM disabledcommands WHERE guildid ='%s' AND commandname ='%s'", interaction.GuildID, args))
					if err != nil {
						return multiplexer.CreateResponse("Bir hata oluştu")
					}
					defer delete.Close()
					return multiplexer.CreateResponse(args + " komutu tekrar kullanılabilir.")
				}
				return multiplexer.CreateResponse("Bu komut engellenmemiş")
			}
			return multiplexer.CreateResponse("Bu komut engellenmemiş.")
		case "kiss":
			err := db.QueryRow(fmt.Sprintf("SELECT isblocked FROM disabledcommands WHERE commandname ='%s' AND guildid ='%s'", args, interaction.GuildID)).Scan(&tag.isblocked)
			if err == nil {
				if tag.isblocked == "True" {
					delete, err := db.Query(fmt.Sprintf("DELETE FROM disabledcommands WHERE guildid ='%s' AND commandname ='%s'", interaction.GuildID, args))
					if err != nil {
						return multiplexer.CreateResponse("Bir hata oluştu")
					}
					defer delete.Close()
					return multiplexer.CreateResponse(args + " komutu tekrar kullanılabilir.")
				}
				return multiplexer.CreateResponse("Bu komut engellenmemiş")
			}
			return multiplexer.CreateResponse("Bu komut engellenmemiş.")
		case "ping":
			err := db.QueryRow(fmt.Sprintf("SELECT isblocked FROM disabledcommands WHERE commandname ='%s' AND guildid ='%s'", args, interaction.GuildID)).Scan(&tag.isblocked)
			if err == nil {
				if tag.isblocked == "True" {
					delete, err := db.Query(fmt.Sprintf("DELETE FROM disabledcommands WHERE guildid ='%s' AND commandname ='%s'", interaction.GuildID, args))
					if err != nil {
						return multiplexer.CreateResponse("Bir hata oluştu")
					}
					defer delete.Close()
					return multiplexer.CreateResponse(args + " komutu tekrar kullanılabilir.")
				}
				return multiplexer.CreateResponse("Bu komut engellenmemiş")
			}
			return multiplexer.CreateResponse("Bu komut engellenmemiş.")
		case "roles":
			err := db.QueryRow(fmt.Sprintf("SELECT isblocked FROM disabledcommands WHERE commandname ='%s' AND guildid ='%s'", args, interaction.GuildID)).Scan(&tag.isblocked)
			if err == nil {
				if tag.isblocked == "True" {
					delete, err := db.Query(fmt.Sprintf("DELETE FROM disabledcommands WHERE guildid ='%s' AND commandname ='%s'", interaction.GuildID, args))
					if err != nil {
						return multiplexer.CreateResponse("Bir hata oluştu")
					}
					defer delete.Close()
					return multiplexer.CreateResponse(args + " komutu tekrar kullanılabilir.")
				}
				return multiplexer.CreateResponse("Bu komut engellenmemiş")
			}
			return multiplexer.CreateResponse("Bu komut engellenmemiş.")
		case "settings":
			err := db.QueryRow(fmt.Sprintf("SELECT isblocked FROM disabledcommands WHERE commandname ='%s' AND guildid ='%s'", args, interaction.GuildID)).Scan(&tag.isblocked)
			if err == nil {
				if tag.isblocked == "True" {
					delete, err := db.Query(fmt.Sprintf("DELETE FROM disabledcommands WHERE guildid ='%s' AND commandname ='%s'", interaction.GuildID, args))
					if err != nil {
						return multiplexer.CreateResponse("Bir hata oluştu")
					}
					defer delete.Close()
					return multiplexer.CreateResponse(args + " komutu tekrar kullanılabilir.")
				}
				return multiplexer.CreateResponse("Bu komut engellenmemiş")
			}
			return multiplexer.CreateResponse("Bu komut engellenmemiş.")
		case "slap":
			err := db.QueryRow(fmt.Sprintf("SELECT isblocked FROM disabledcommands WHERE commandname ='%s' AND guildid ='%s'", args, interaction.GuildID)).Scan(&tag.isblocked)
			if err == nil {
				if tag.isblocked == "True" {
					delete, err := db.Query(fmt.Sprintf("DELETE FROM disabledcommands WHERE guildid ='%s' AND commandname ='%s'", interaction.GuildID, args))
					if err != nil {
						return multiplexer.CreateResponse("Bir hata oluştu")
					}
					defer delete.Close()
					return multiplexer.CreateResponse(args + " komutu tekrar kullanılabilir.")
				}
				return multiplexer.CreateResponse("Bu komut engellenmemiş")
			}
			return multiplexer.CreateResponse("Bu komut engellenmemiş.")
		case "spoiler":
			err := db.QueryRow(fmt.Sprintf("SELECT isblocked FROM disabledcommands WHERE commandname ='%s' AND guildid ='%s'", args, interaction.GuildID)).Scan(&tag.isblocked)
			if err == nil {
				if tag.isblocked == "True" {
					delete, err := db.Query(fmt.Sprintf("DELETE FROM disabledcommands WHERE guildid ='%s' AND commandname ='%s'", interaction.GuildID, args))
					if err != nil {
						return multiplexer.CreateResponse("Bir hata oluştu")
					}
					defer delete.Close()
					return multiplexer.CreateResponse(args + " komutu tekrar kullanılabilir.")
				}
				return multiplexer.CreateResponse("Bu komut engellenmemiş")
			}
			return multiplexer.CreateResponse("Bu komut engellenmemiş.")
		case "start_vote":
			err := db.QueryRow(fmt.Sprintf("SELECT isblocked FROM disabledcommands WHERE commandname ='%s' AND guildid ='%s'", args, interaction.GuildID)).Scan(&tag.isblocked)
			if err == nil {
				if tag.isblocked == "True" {
					delete, err := db.Query(fmt.Sprintf("DELETE FROM disabledcommands WHERE guildid ='%s' AND commandname ='%s'", interaction.GuildID, args))
					if err != nil {
						return multiplexer.CreateResponse("Bir hata oluştu")
					}
					defer delete.Close()
					return multiplexer.CreateResponse(args + " komutu tekrar kullanılabilir.")
				}
				return multiplexer.CreateResponse("Bu komut engellenmemiş")
			}
			return multiplexer.CreateResponse("Bu komut engellenmemiş.")
		case "stats":
			err := db.QueryRow(fmt.Sprintf("SELECT isblocked FROM disabledcommands WHERE commandname ='%s' AND guildid ='%s'", args, interaction.GuildID)).Scan(&tag.isblocked)
			if err == nil {
				if tag.isblocked == "True" {
					delete, err := db.Query(fmt.Sprintf("DELETE FROM disabledcommands WHERE guildid ='%s' AND commandname ='%s'", interaction.GuildID, args))
					if err != nil {
						return multiplexer.CreateResponse("Bir hata oluştu")
					}
					defer delete.Close()
					return multiplexer.CreateResponse(args + " komutu tekrar kullanılabilir.")
				}
				return multiplexer.CreateResponse("Bu komut engellenmemiş")
			}
			return multiplexer.CreateResponse("Bu komut engellenmemiş.")
		case "unban":
			err := db.QueryRow(fmt.Sprintf("SELECT isblocked FROM disabledcommands WHERE commandname ='%s' AND guildid ='%s'", args, interaction.GuildID)).Scan(&tag.isblocked)
			if err == nil {
				if tag.isblocked == "True" {
					delete, err := db.Query(fmt.Sprintf("DELETE FROM disabledcommands WHERE guildid ='%s' AND commandname ='%s'", interaction.GuildID, args))
					if err != nil {
						return multiplexer.CreateResponse("Bir hata oluştu")
					}
					defer delete.Close()
					return multiplexer.CreateResponse(args + " komutu tekrar kullanılabilir.")
				}
				return multiplexer.CreateResponse("Bu komut engellenmemiş")
			}
			return multiplexer.CreateResponse("Bu komut engellenmemiş.")
		default:
			return multiplexer.CreateResponse("Bir komut belirtmelisin")
		}
	}
	if !multiplexer.CheckAdministratorPermission(session, interaction.Member.User.ID, interaction.ChannelID) {
		return multiplexer.CreateResponse("You don't have enough permission.")
	}

	args := interaction.Data.Options[0].Value.(string)

	switch args {
	case "afk":
		err := db.QueryRow(fmt.Sprintf("SELECT isblocked FROM disabledcommands WHERE commandname ='%s' AND guildid ='%s'", args, interaction.GuildID)).Scan(&tag.isblocked)
		if err == nil {
			if tag.isblocked == "True" {
				delete, err := db.Query(fmt.Sprintf("DELETE FROM disabledcommands WHERE guildid ='%s' AND commandname ='%s'", interaction.GuildID, args))
				if err != nil {
					return multiplexer.CreateResponse("An error occurred")
				}
				defer delete.Close()
				return multiplexer.CreateResponse("Enabled " + args)
			}
			return multiplexer.CreateResponse("This command is not disabled.")
		}
		return multiplexer.CreateResponse("This command is not disabled.")
	case "author":
		err := db.QueryRow(fmt.Sprintf("SELECT isblocked FROM disabledcommands WHERE commandname ='%s' AND guildid ='%s'", args, interaction.GuildID)).Scan(&tag.isblocked)
		if err == nil {
			if tag.isblocked == "True" {
				delete, err := db.Query(fmt.Sprintf("DELETE FROM disabledcommands WHERE guildid ='%s' AND commandname ='%s'", interaction.GuildID, args))
				if err != nil {
					return multiplexer.CreateResponse("An error occurred")
				}
				defer delete.Close()
				return multiplexer.CreateResponse("Enabled " + args)
			}
			return multiplexer.CreateResponse("This command is not disabled.")
		}
		return multiplexer.CreateResponse("This command is not disabled.")
	case "avatar":
		err := db.QueryRow(fmt.Sprintf("SELECT isblocked FROM disabledcommands WHERE commandname ='%s' AND guildid ='%s'", args, interaction.GuildID)).Scan(&tag.isblocked)
		if err == nil {
			if tag.isblocked == "True" {
				delete, err := db.Query(fmt.Sprintf("DELETE FROM disabledcommands WHERE guildid ='%s' AND commandname ='%s'", interaction.GuildID, args))
				if err != nil {
					return multiplexer.CreateResponse("An error occurred")
				}
				defer delete.Close()
				return multiplexer.CreateResponse("Enabled " + args)
			}
			return multiplexer.CreateResponse("This command is not disabled.")
		}
		return multiplexer.CreateResponse("This command is not disabled.")
	case "ban":
		err := db.QueryRow(fmt.Sprintf("SELECT isblocked FROM disabledcommands WHERE commandname ='%s' AND guildid ='%s'", args, interaction.GuildID)).Scan(&tag.isblocked)
		if err == nil {
			if tag.isblocked == "True" {
				delete, err := db.Query(fmt.Sprintf("DELETE FROM disabledcommands WHERE guildid ='%s' AND commandname ='%s'", interaction.GuildID, args))
				if err != nil {
					return multiplexer.CreateResponse("An error occurred")
				}
				defer delete.Close()
				return multiplexer.CreateResponse("Enabled " + args)
			}
			return multiplexer.CreateResponse("This command is not disabled.")
		}
		return multiplexer.CreateResponse("This command is not disabled.")
	case "embed":
		err := db.QueryRow(fmt.Sprintf("SELECT isblocked FROM disabledcommands WHERE commandname ='%s' AND guildid ='%s'", args, interaction.GuildID)).Scan(&tag.isblocked)
		if err == nil {
			if tag.isblocked == "True" {
				delete, err := db.Query(fmt.Sprintf("DELETE FROM disabledcommands WHERE guildid ='%s' AND commandname ='%s'", interaction.GuildID, args))
				if err != nil {
					return multiplexer.CreateResponse("An error occurred")
				}
				defer delete.Close()
				return multiplexer.CreateResponse("Enabled " + args)
			}
			return multiplexer.CreateResponse("This command is not disabled.")
		}
		return multiplexer.CreateResponse("This command is not disabled.")
	case "Channel_info":
		err := db.QueryRow(fmt.Sprintf("SELECT isblocked FROM disabledcommands WHERE commandname ='%s' AND guildid ='%s'", args, interaction.GuildID)).Scan(&tag.isblocked)
		if err == nil {
			if tag.isblocked == "True" {
				delete, err := db.Query(fmt.Sprintf("DELETE FROM disabledcommands WHERE guildid ='%s' AND commandname ='%s'", interaction.GuildID, args))
				if err != nil {
					return multiplexer.CreateResponse("An error occurred")
				}
				defer delete.Close()
				return multiplexer.CreateResponse("Enabled " + args)
			}
			return multiplexer.CreateResponse("This command is not disabled.")
		}
		return multiplexer.CreateResponse("This command is not disabled.")
	case "hug":
		err := db.QueryRow(fmt.Sprintf("SELECT isblocked FROM disabledcommands WHERE commandname ='%s' AND guildid ='%s'", args, interaction.GuildID)).Scan(&tag.isblocked)
		if err == nil {
			if tag.isblocked == "True" {
				delete, err := db.Query(fmt.Sprintf("DELETE FROM disabledcommands WHERE guildid ='%s' AND commandname ='%s'", interaction.GuildID, args))
				if err != nil {
					return multiplexer.CreateResponse("An error occurred")
				}
				defer delete.Close()
				return multiplexer.CreateResponse("Enabled " + args)
			}
			return multiplexer.CreateResponse("This command is not disabled.")
		}
		return multiplexer.CreateResponse("This command is not disabled.")
	case "kick":
		err := db.QueryRow(fmt.Sprintf("SELECT isblocked FROM disabledcommands WHERE commandname ='%s' AND guildid ='%s'", args, interaction.GuildID)).Scan(&tag.isblocked)
		if err == nil {
			if tag.isblocked == "True" {
				delete, err := db.Query(fmt.Sprintf("DELETE FROM disabledcommands WHERE guildid ='%s' AND commandname ='%s'", interaction.GuildID, args))
				if err != nil {
					return multiplexer.CreateResponse("An error occurred")
				}
				defer delete.Close()
				return multiplexer.CreateResponse("Enabled " + args)
			}
			return multiplexer.CreateResponse("This command is not disabled.")
		}
		return multiplexer.CreateResponse("This command is not disabled.")
	case "kiss":
		err := db.QueryRow(fmt.Sprintf("SELECT isblocked FROM disabledcommands WHERE commandname ='%s' AND guildid ='%s'", args, interaction.GuildID)).Scan(&tag.isblocked)
		if err == nil {
			if tag.isblocked == "True" {
				delete, err := db.Query(fmt.Sprintf("DELETE FROM disabledcommands WHERE guildid ='%s' AND commandname ='%s'", interaction.GuildID, args))
				if err != nil {
					return multiplexer.CreateResponse("An error occurred")
				}
				defer delete.Close()
				return multiplexer.CreateResponse("Enabled " + args)
			}
			return multiplexer.CreateResponse("This command is not disabled.")
		}
		return multiplexer.CreateResponse("This command is not disabled.")
	case "ping":
		err := db.QueryRow(fmt.Sprintf("SELECT isblocked FROM disabledcommands WHERE commandname ='%s' AND guildid ='%s'", args, interaction.GuildID)).Scan(&tag.isblocked)
		if err == nil {
			if tag.isblocked == "True" {
				delete, err := db.Query(fmt.Sprintf("DELETE FROM disabledcommands WHERE guildid ='%s' AND commandname ='%s'", interaction.GuildID, args))
				if err != nil {
					return multiplexer.CreateResponse("An error occurred")
				}
				defer delete.Close()
				return multiplexer.CreateResponse("Enabled " + args)
			}
			return multiplexer.CreateResponse("This command is not disabled.")
		}
		return multiplexer.CreateResponse("This command is not disabled.")
	case "roles":
		err := db.QueryRow(fmt.Sprintf("SELECT isblocked FROM disabledcommands WHERE commandname ='%s' AND guildid ='%s'", args, interaction.GuildID)).Scan(&tag.isblocked)
		if err == nil {
			if tag.isblocked == "True" {
				delete, err := db.Query(fmt.Sprintf("DELETE FROM disabledcommands WHERE guildid ='%s' AND commandname ='%s'", interaction.GuildID, args))
				if err != nil {
					return multiplexer.CreateResponse("An error occurred")
				}
				defer delete.Close()
				return multiplexer.CreateResponse("Enabled " + args)
			}
			return multiplexer.CreateResponse("This command is not disabled.")
		}
		return multiplexer.CreateResponse("This command is not disabled.")
	case "settings":
		err := db.QueryRow(fmt.Sprintf("SELECT isblocked FROM disabledcommands WHERE commandname ='%s' AND guildid ='%s'", args, interaction.GuildID)).Scan(&tag.isblocked)
		if err == nil {
			if tag.isblocked == "True" {
				delete, err := db.Query(fmt.Sprintf("DELETE FROM disabledcommands WHERE guildid ='%s' AND commandname ='%s'", interaction.GuildID, args))
				if err != nil {
					return multiplexer.CreateResponse("An error occurred")
				}
				defer delete.Close()
				return multiplexer.CreateResponse("Enabled " + args)
			}
			return multiplexer.CreateResponse("This command is not disabled.")
		}
		return multiplexer.CreateResponse("This command is not disabled.")
	case "slap":
		err := db.QueryRow(fmt.Sprintf("SELECT isblocked FROM disabledcommands WHERE commandname ='%s' AND guildid ='%s'", args, interaction.GuildID)).Scan(&tag.isblocked)
		if err == nil {
			if tag.isblocked == "True" {
				delete, err := db.Query(fmt.Sprintf("DELETE FROM disabledcommands WHERE guildid ='%s' AND commandname ='%s'", interaction.GuildID, args))
				if err != nil {
					return multiplexer.CreateResponse("An error occurred")
				}
				defer delete.Close()
				return multiplexer.CreateResponse("Enabled " + args)
			}
			return multiplexer.CreateResponse("This command is not disabled.")
		}
		return multiplexer.CreateResponse("This command is not disabled.")
	case "spoiler":
		err := db.QueryRow(fmt.Sprintf("SELECT isblocked FROM disabledcommands WHERE commandname ='%s' AND guildid ='%s'", args, interaction.GuildID)).Scan(&tag.isblocked)
		if err == nil {
			if tag.isblocked == "True" {
				delete, err := db.Query(fmt.Sprintf("DELETE FROM disabledcommands WHERE guildid ='%s' AND commandname ='%s'", interaction.GuildID, args))
				if err != nil {
					return multiplexer.CreateResponse("An error occurred")
				}
				defer delete.Close()
				return multiplexer.CreateResponse("Enabled " + args)
			}
			return multiplexer.CreateResponse("This command is not disabled.")
		}
		return multiplexer.CreateResponse("This command is not disabled.")
	case "start_vote":
		err := db.QueryRow(fmt.Sprintf("SELECT isblocked FROM disabledcommands WHERE commandname ='%s' AND guildid ='%s'", args, interaction.GuildID)).Scan(&tag.isblocked)
		if err == nil {
			if tag.isblocked == "True" {
				delete, err := db.Query(fmt.Sprintf("DELETE FROM disabledcommands WHERE guildid ='%s' AND commandname ='%s'", interaction.GuildID, args))
				if err != nil {
					return multiplexer.CreateResponse("An error occurred")
				}
				defer delete.Close()
				return multiplexer.CreateResponse("Enabled " + args)
			}
			return multiplexer.CreateResponse("This command is not disabled.")
		}
		return multiplexer.CreateResponse("This command is not disabled.")
	case "stats":
		err := db.QueryRow(fmt.Sprintf("SELECT isblocked FROM disabledcommands WHERE commandname ='%s' AND guildid ='%s'", args, interaction.GuildID)).Scan(&tag.isblocked)
		if err == nil {
			if tag.isblocked == "True" {
				delete, err := db.Query(fmt.Sprintf("DELETE FROM disabledcommands WHERE guildid ='%s' AND commandname ='%s'", interaction.GuildID, args))
				if err != nil {
					return multiplexer.CreateResponse("An error occurred")
				}
				defer delete.Close()
				return multiplexer.CreateResponse("Enabled " + args)
			}
			return multiplexer.CreateResponse("This command is not disabled.")
		}
		return multiplexer.CreateResponse("This command is not disabled.")
	case "unban":
		err := db.QueryRow(fmt.Sprintf("SELECT isblocked FROM disabledcommands WHERE commandname ='%s' AND guildid ='%s'", args, interaction.GuildID)).Scan(&tag.isblocked)
		if err == nil {
			if tag.isblocked == "True" {
				delete, err := db.Query(fmt.Sprintf("DELETE FROM disabledcommands WHERE guildid ='%s' AND commandname ='%s'", interaction.GuildID, args))
				if err != nil {
					return multiplexer.CreateResponse("An error occurred")
				}
				defer delete.Close()
				return multiplexer.CreateResponse("Enabled " + args)
			}
			return multiplexer.CreateResponse("This command is not disabled.")
		}
		return multiplexer.CreateResponse("This command is not disabled.")
	}
	return multiplexer.CreateResponse("You need to specify the command.")
}
