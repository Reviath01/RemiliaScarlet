package commands

import (
	"fmt"
	"strings"

	multiplexer "git.randomchars.net/Reviath/RemiliaScarlet/Multiplexer"
	CommandHandler "git.randomchars.net/Reviath/RemiliaScarlet/handler"
	"git.randomchars.net/Reviath/RemiliaScarlet/sql"
)

// EnableCommand is a handler for enable command
func EnableCommand(ctx CommandHandler.Context, _ []string) error {
	db := sql.Connect()
	defer db.Close()
	type Tag struct {
		isblocked string
	}

	var tag Tag

	switch sql.CheckLanguage(ctx.Guild.ID) {
	case "tr":
		if !multiplexer.CheckAdministratorPermission(ctx.Session, ctx.Message.Author.ID, ctx.Guild.ID) {
			ctx.Reply("Yeterli yetkiye sahip değilsin.")
			return nil
		}
		if len(strings.Join(multiplexer.GetArgs(ctx.Message.Content, multiplexer.GetPrefix()), " ")) < 1 {
			ctx.Reply("Bir komut belirtmelisin.")

			return nil
		}

		args := multiplexer.GetArgs(ctx.Message.Content, multiplexer.GetPrefix())[0]

		switch args {
		case "afk":
			err := db.QueryRow(fmt.Sprintf("SELECT isblocked FROM disabledcommands WHERE commandname ='%s' AND guildid ='%s'", args, ctx.Guild.ID)).Scan(&tag.isblocked)
			if err == nil {
				if tag.isblocked == "True" {
					delete, err := db.Query(fmt.Sprintf("DELETE FROM disabledcommands WHERE guildid ='%s' AND commandname ='%s'", ctx.Guild.ID, args))

					if err != nil {
						ctx.Reply("Bir hata oluştu.")

						return nil
					}
					ctx.Reply(args + " komutu tekrar kullanılabilir!")

					defer delete.Close()

				} else {
					ctx.Reply("Bu komut engellenmemiş.")

					return nil
				}
			} else {
				ctx.Reply("Bu komut engellenmemiş.")

				return nil
			}
		case "author":
			err := db.QueryRow(fmt.Sprintf("SELECT isblocked FROM disabledcommands WHERE commandname ='%s' AND guildid ='%s'", args, ctx.Guild.ID)).Scan(&tag.isblocked)
			if err == nil {
				if tag.isblocked == "True" {
					delete, err := db.Query(fmt.Sprintf("DELETE FROM disabledcommands WHERE guildid ='%s' AND commandname ='%s'", ctx.Guild.ID, args))

					if err != nil {
						ctx.Reply("Bir hata oluştu.")

						return nil
					}
					ctx.Reply(args + " komutu tekrar kullanılabilir!")

					defer delete.Close()

				} else {
					ctx.Reply("Bu komut engellenmemiş.")

					return nil
				}
			} else {
				ctx.Reply("Bu komut engellenmemiş.")

				return nil
			}
		case "avatar":
			err := db.QueryRow(fmt.Sprintf("SELECT isblocked FROM disabledcommands WHERE commandname ='%s' AND guildid ='%s'", args, ctx.Guild.ID)).Scan(&tag.isblocked)
			if err == nil {
				if tag.isblocked == "True" {
					delete, err := db.Query(fmt.Sprintf("DELETE FROM disabledcommands WHERE guildid ='%s' AND commandname ='%s'", ctx.Guild.ID, args))

					if err != nil {
						ctx.Reply("Bir hata oluştu.")

						return nil
					}
					ctx.Reply(args + " komutu tekrar kullanılabilir!")

					defer delete.Close()

				} else {
					ctx.Reply("Bu komut engellenmemiş.")

					return nil
				}
			} else {
				ctx.Reply("Bu komut engellenmemiş.")

				return nil
			}
		case "ban":
			err := db.QueryRow(fmt.Sprintf("SELECT isblocked FROM disabledcommands WHERE commandname ='%s' AND guildid ='%s'", args, ctx.Guild.ID)).Scan(&tag.isblocked)
			if err == nil {
				if tag.isblocked == "True" {
					delete, err := db.Query(fmt.Sprintf("DELETE FROM disabledcommands WHERE guildid ='%s' AND commandname ='%s'", ctx.Guild.ID, args))

					if err != nil {
						ctx.Reply("Bir hata oluştu.")

						return nil
					}
					ctx.Reply(args + " komutu tekrar kullanılabilir!")

					defer delete.Close()

				} else {
					ctx.Reply("Bu komut engellenmemiş.")

					return nil
				}
			} else {
				ctx.Reply("Bu komut engellenmemiş.")

				return nil
			}
		case "embed":
			err := db.QueryRow(fmt.Sprintf("SELECT isblocked FROM disabledcommands WHERE commandname ='%s' AND guildid ='%s'", args, ctx.Guild.ID)).Scan(&tag.isblocked)
			if err == nil {
				if tag.isblocked == "True" {
					delete, err := db.Query(fmt.Sprintf("DELETE FROM disabledcommands WHERE guildid ='%s' AND commandname ='%s'", ctx.Guild.ID, args))

					if err != nil {
						ctx.Reply("Bir hata oluştu.")

						return nil
					}
					ctx.Reply(args + " komutu tekrar kullanılabilir!")

					defer delete.Close()

				} else {
					ctx.Reply("Bu komut engellenmemiş.")

					return nil
				}
			} else {
				ctx.Reply("Bu komut engellenmemiş.")

				return nil
			}
		case "Channel_info":
			err := db.QueryRow(fmt.Sprintf("SELECT isblocked FROM disabledcommands WHERE commandname ='%s' AND guildid ='%s'", args, ctx.Guild.ID)).Scan(&tag.isblocked)
			if err == nil {
				if tag.isblocked == "True" {
					delete, err := db.Query(fmt.Sprintf("DELETE FROM disabledcommands WHERE guildid ='%s' AND commandname ='%s'", ctx.Guild.ID, args))

					if err != nil {
						ctx.Reply("Bir hata oluştu.")

						return nil
					}
					ctx.Reply(args + " komutu tekrar kullanılabilir!")

					defer delete.Close()

				} else {
					ctx.Reply("Bu komut engellenmemiş.")

					return nil
				}
			} else {
				ctx.Reply("Bu komut engellenmemiş.")

				return nil
			}
		case "hug":
			err := db.QueryRow(fmt.Sprintf("SELECT isblocked FROM disabledcommands WHERE commandname ='%s' AND guildid ='%s'", args, ctx.Guild.ID)).Scan(&tag.isblocked)
			if err == nil {
				if tag.isblocked == "True" {
					delete, err := db.Query(fmt.Sprintf("DELETE FROM disabledcommands WHERE guildid ='%s' AND commandname ='%s'", ctx.Guild.ID, args))

					if err != nil {
						ctx.Reply("Bir hata oluştu.")

						return nil
					}
					ctx.Reply(args + " komutu tekrar kullanılabilir!")

					defer delete.Close()

				} else {
					ctx.Reply("Bu komut engellenmemiş.")

					return nil
				}
			} else {
				ctx.Reply("Bu komut engellenmemiş.")

				return nil
			}
		case "kick":
			err := db.QueryRow(fmt.Sprintf("SELECT isblocked FROM disabledcommands WHERE commandname ='%s' AND guildid ='%s'", args, ctx.Guild.ID)).Scan(&tag.isblocked)
			if err == nil {
				if tag.isblocked == "True" {
					delete, err := db.Query(fmt.Sprintf("DELETE FROM disabledcommands WHERE guildid ='%s' AND commandname ='%s'", ctx.Guild.ID, args))

					if err != nil {
						ctx.Reply("Bir hata oluştu.")

						return nil
					}
					ctx.Reply(args + " komutu tekrar kullanılabilir!")

					defer delete.Close()

				} else {
					ctx.Reply("Bu komut engellenmemiş.")

					return nil
				}
			} else {
				ctx.Reply("Bu komut engellenmemiş.")

				return nil
			}
		case "kiss":
			err := db.QueryRow(fmt.Sprintf("SELECT isblocked FROM disabledcommands WHERE commandname ='%s' AND guildid ='%s'", args, ctx.Guild.ID)).Scan(&tag.isblocked)
			if err == nil {
				if tag.isblocked == "True" {
					delete, err := db.Query(fmt.Sprintf("DELETE FROM disabledcommands WHERE guildid ='%s' AND commandname ='%s'", ctx.Guild.ID, args))

					if err != nil {
						ctx.Reply("Bir hata oluştu.")

						return nil
					}
					ctx.Reply(args + " komutu tekrar kullanılabilir!")

					defer delete.Close()

				} else {
					ctx.Reply("Bu komut engellenmemiş.")

					return nil
				}
			} else {
				ctx.Reply("Bu komut engellenmemiş.")

				return nil
			}
		case "ping":
			err := db.QueryRow(fmt.Sprintf("SELECT isblocked FROM disabledcommands WHERE commandname ='%s' AND guildid ='%s'", args, ctx.Guild.ID)).Scan(&tag.isblocked)
			if err == nil {
				if tag.isblocked == "True" {
					delete, err := db.Query(fmt.Sprintf("DELETE FROM disabledcommands WHERE guildid ='%s' AND commandname ='%s'", ctx.Guild.ID, args))

					if err != nil {
						ctx.Reply("Bir hata oluştu.")

						return nil
					}
					ctx.Reply(args + " komutu tekrar kullanılabilir!")

					defer delete.Close()

				} else {
					ctx.Reply("Bu komut engellenmemiş.")

					return nil
				}
			} else {
				ctx.Reply("Bu komut engellenmemiş.")

				return nil
			}
		case "roles":
			err := db.QueryRow(fmt.Sprintf("SELECT isblocked FROM disabledcommands WHERE commandname ='%s' AND guildid ='%s'", args, ctx.Guild.ID)).Scan(&tag.isblocked)
			if err == nil {
				if tag.isblocked == "True" {
					delete, err := db.Query(fmt.Sprintf("DELETE FROM disabledcommands WHERE guildid ='%s' AND commandname ='%s'", ctx.Guild.ID, args))

					if err != nil {
						ctx.Reply("Bir hata oluştu.")

						return nil
					}
					ctx.Reply(args + " komutu tekrar kullanılabilir!")

					defer delete.Close()

				} else {
					ctx.Reply("Bu komut engellenmemiş.")

					return nil
				}
			} else {
				ctx.Reply("Bu komut engellenmemiş.")

				return nil
			}
		case "settings":
			err := db.QueryRow(fmt.Sprintf("SELECT isblocked FROM disabledcommands WHERE commandname ='%s' AND guildid ='%s'", args, ctx.Guild.ID)).Scan(&tag.isblocked)
			if err == nil {
				if tag.isblocked == "True" {
					delete, err := db.Query(fmt.Sprintf("DELETE FROM disabledcommands WHERE guildid ='%s' AND commandname ='%s'", ctx.Guild.ID, args))

					if err != nil {
						ctx.Reply("Bir hata oluştu.")

						return nil
					}
					ctx.Reply(args + " komutu tekrar kullanılabilir!")

					defer delete.Close()

				} else {
					ctx.Reply("Bu komut engellenmemiş.")

					return nil
				}
			} else {
				ctx.Reply("Bu komut engellenmemiş.")

				return nil
			}
		case "slap":
			err := db.QueryRow(fmt.Sprintf("SELECT isblocked FROM disabledcommands WHERE commandname ='%s' AND guildid ='%s'", args, ctx.Guild.ID)).Scan(&tag.isblocked)
			if err == nil {
				if tag.isblocked == "True" {
					delete, err := db.Query(fmt.Sprintf("DELETE FROM disabledcommands WHERE guildid ='%s' AND commandname ='%s'", ctx.Guild.ID, args))

					if err != nil {
						ctx.Reply("Bir hata oluştu.")

						return nil
					}
					ctx.Reply(args + " komutu tekrar kullanılabilir!")

					defer delete.Close()

				} else {
					ctx.Reply("Bu komut engellenmemiş.")

					return nil
				}
			} else {
				ctx.Reply("Bu komut engellenmemiş.")

				return nil
			}
		case "spoiler":
			err := db.QueryRow(fmt.Sprintf("SELECT isblocked FROM disabledcommands WHERE commandname ='%s' AND guildid ='%s'", args, ctx.Guild.ID)).Scan(&tag.isblocked)
			if err == nil {
				if tag.isblocked == "True" {
					delete, err := db.Query(fmt.Sprintf("DELETE FROM disabledcommands WHERE guildid ='%s' AND commandname ='%s'", ctx.Guild.ID, args))

					if err != nil {
						ctx.Reply("Bir hata oluştu.")

						return nil
					}
					ctx.Reply(args + " komutu tekrar kullanılabilir!")

					defer delete.Close()

				} else {
					ctx.Reply("Bu komut engellenmemiş.")

					return nil
				}
			} else {
				ctx.Reply("Bu komut engellenmemiş.")

				return nil
			}
		case "start_vote":
			err := db.QueryRow(fmt.Sprintf("SELECT isblocked FROM disabledcommands WHERE commandname ='%s' AND guildid ='%s'", args, ctx.Guild.ID)).Scan(&tag.isblocked)
			if err == nil {
				if tag.isblocked == "True" {
					delete, err := db.Query(fmt.Sprintf("DELETE FROM disabledcommands WHERE guildid ='%s' AND commandname ='%s'", ctx.Guild.ID, args))

					if err != nil {
						ctx.Reply("Bir hata oluştu.")

						return nil
					}
					ctx.Reply(args + " komutu tekrar kullanılabilir!")

					defer delete.Close()

				} else {
					ctx.Reply("Bu komut engellenmemiş.")

					return nil
				}
			} else {
				ctx.Reply("Bu komut engellenmemiş.")

				return nil
			}
		case "stats":
			err := db.QueryRow(fmt.Sprintf("SELECT isblocked FROM disabledcommands WHERE commandname ='%s' AND guildid ='%s'", args, ctx.Guild.ID)).Scan(&tag.isblocked)
			if err == nil {
				if tag.isblocked == "True" {
					delete, err := db.Query(fmt.Sprintf("DELETE FROM disabledcommands WHERE guildid ='%s' AND commandname ='%s'", ctx.Guild.ID, args))

					if err != nil {
						ctx.Reply("Bir hata oluştu.")

						return nil
					}
					ctx.Reply(args + " komutu tekrar kullanılabilir!")

					defer delete.Close()

				} else {
					ctx.Reply("Bu komut engellenmemiş.")

					return nil
				}
			} else {
				ctx.Reply("Bu komut engellenmemiş.")

				return nil
			}
		case "unban":
			err := db.QueryRow(fmt.Sprintf("SELECT isblocked FROM disabledcommands WHERE commandname ='%s' AND guildid ='%s'", args, ctx.Guild.ID)).Scan(&tag.isblocked)
			if err == nil {
				if tag.isblocked == "True" {
					delete, err := db.Query(fmt.Sprintf("DELETE FROM disabledcommands WHERE guildid ='%s' AND commandname ='%s'", ctx.Guild.ID, args))

					if err != nil {
						ctx.Reply("Bir hata oluştu.")

						return nil
					}
					ctx.Reply(args + " komutu tekrar kullanılabilir!")

					defer delete.Close()

				} else {
					ctx.Reply("Bu komut engellenmemiş.")

					return nil
				}
			} else {
				ctx.Reply("Bu komut engellenmemiş.")

				return nil
			}
		default:
			ctx.Reply("Bir komut belirtmelisin.")
		}
		return nil
	default:
		if !multiplexer.CheckAdministratorPermission(ctx.Session, ctx.Message.Author.ID, ctx.Guild.ID) {
			ctx.Reply("You don't have enough permission.")
			return nil
		}
		if len(strings.Join(multiplexer.GetArgs(ctx.Message.Content, multiplexer.GetPrefix()), " ")) < 1 {
			ctx.Reply("You need to specify a command.")

			return nil
		}

		var args string

		if strings.Join(multiplexer.GetArgs(ctx.Message.Content, multiplexer.GetPrefix()), " ") == "" {
			ctx.Reply("You need to specify the command.")

			return nil
		}

		args = multiplexer.GetArgs(ctx.Message.Content, multiplexer.GetPrefix())[0]

		switch args {
		case "afk":
			err := db.QueryRow(fmt.Sprintf("SELECT isblocked FROM disabledcommands WHERE commandname ='%s' AND guildid ='%s'", args, ctx.Guild.ID)).Scan(&tag.isblocked)
			if err == nil {
				if tag.isblocked == "True" {
					delete, err := db.Query(fmt.Sprintf("DELETE FROM disabledcommands WHERE guildid ='%s' AND commandname ='%s'", ctx.Guild.ID, args))

					if err != nil {
						ctx.Reply("An error occurred.")

						return nil
					}
					ctx.Reply("Enabled " + args)

					defer delete.Close()

				} else {
					ctx.Reply("This command is not disabled.")

					return nil
				}
			} else {
				ctx.Reply("This command is not disabled.")

				return nil
			}
		case "author":
			err := db.QueryRow(fmt.Sprintf("SELECT isblocked FROM disabledcommands WHERE commandname ='%s' AND guildid ='%s'", args, ctx.Guild.ID)).Scan(&tag.isblocked)
			if err == nil {
				if tag.isblocked == "True" {
					delete, err := db.Query(fmt.Sprintf("DELETE FROM disabledcommands WHERE guildid ='%s' AND commandname ='%s'", ctx.Guild.ID, args))

					if err != nil {
						ctx.Reply("An error occurred.")

						return nil
					}
					ctx.Reply("Enabled " + args)

					defer delete.Close()

				} else {
					ctx.Reply("This command is not disabled.")

					return nil
				}
			} else {
				ctx.Reply("This command is not disabled.")

				return nil
			}
		case "avatar":
			err := db.QueryRow(fmt.Sprintf("SELECT isblocked FROM disabledcommands WHERE commandname ='%s' AND guildid ='%s'", args, ctx.Guild.ID)).Scan(&tag.isblocked)
			if err == nil {
				if tag.isblocked == "True" {
					delete, err := db.Query(fmt.Sprintf("DELETE FROM disabledcommands WHERE guildid ='%s' AND commandname ='%s'", ctx.Guild.ID, args))

					if err != nil {
						ctx.Reply("An error occurred.")

						return nil
					}
					ctx.Reply("Enabled " + args)

					defer delete.Close()

				} else {
					ctx.Reply("This command is not disabled.")

					return nil
				}
			} else {
				ctx.Reply("This command is not disabled.")

				return nil
			}
		case "ban":
			err := db.QueryRow(fmt.Sprintf("SELECT isblocked FROM disabledcommands WHERE commandname ='%s' AND guildid ='%s'", args, ctx.Guild.ID)).Scan(&tag.isblocked)
			if err == nil {
				if tag.isblocked == "True" {
					delete, err := db.Query(fmt.Sprintf("DELETE FROM disabledcommands WHERE guildid ='%s' AND commandname ='%s'", ctx.Guild.ID, args))

					if err != nil {
						ctx.Reply("An error occurred.")

						return nil
					}
					ctx.Reply("Enabled " + args)

					defer delete.Close()

				} else {
					ctx.Reply("This command is not disabled.")

					return nil
				}
			} else {
				ctx.Reply("This command is not disabled.")

				return nil
			}
		case "embed":
			err := db.QueryRow(fmt.Sprintf("SELECT isblocked FROM disabledcommands WHERE commandname ='%s' AND guildid ='%s'", args, ctx.Guild.ID)).Scan(&tag.isblocked)
			if err == nil {
				if tag.isblocked == "True" {
					delete, err := db.Query(fmt.Sprintf("DELETE FROM disabledcommands WHERE guildid ='%s' AND commandname ='%s'", ctx.Guild.ID, args))

					if err != nil {
						ctx.Reply("An error occurred.")

						return nil
					}
					ctx.Reply("Enabled " + args)

					defer delete.Close()

				} else {
					ctx.Reply("This command is not disabled.")

					return nil
				}
			} else {
				ctx.Reply("This command is not disabled.")

				return nil
			}
		case "Channel_info":
			err := db.QueryRow(fmt.Sprintf("SELECT isblocked FROM disabledcommands WHERE commandname ='%s' AND guildid ='%s'", args, ctx.Guild.ID)).Scan(&tag.isblocked)
			if err == nil {
				if tag.isblocked == "True" {
					delete, err := db.Query(fmt.Sprintf("DELETE FROM disabledcommands WHERE guildid ='%s' AND commandname ='%s'", ctx.Guild.ID, args))

					if err != nil {
						ctx.Reply("An error occurred.")

						return nil
					}
					ctx.Reply("Enabled " + args)

					defer delete.Close()

				} else {
					ctx.Reply("This command is not disabled.")

					return nil
				}
			} else {
				ctx.Reply("This command is not disabled.")

				return nil
			}
		case "hug":
			err := db.QueryRow(fmt.Sprintf("SELECT isblocked FROM disabledcommands WHERE commandname ='%s' AND guildid ='%s'", args, ctx.Guild.ID)).Scan(&tag.isblocked)
			if err == nil {
				if tag.isblocked == "True" {
					delete, err := db.Query(fmt.Sprintf("DELETE FROM disabledcommands WHERE guildid ='%s' AND commandname ='%s'", ctx.Guild.ID, args))

					if err != nil {
						ctx.Reply("An error occurred.")

						return nil
					}
					ctx.Reply("Enabled " + args)

					defer delete.Close()

				} else {
					ctx.Reply("This command is not disabled.")

					return nil
				}
			} else {
				ctx.Reply("This command is not disabled.")

				return nil
			}
		case "kick":
			err := db.QueryRow(fmt.Sprintf("SELECT isblocked FROM disabledcommands WHERE commandname ='%s' AND guildid ='%s'", args, ctx.Guild.ID)).Scan(&tag.isblocked)
			if err == nil {
				if tag.isblocked == "True" {
					delete, err := db.Query(fmt.Sprintf("DELETE FROM disabledcommands WHERE guildid ='%s' AND commandname ='%s'", ctx.Guild.ID, args))

					if err != nil {
						ctx.Reply("An error occurred.")

						return nil
					}
					ctx.Reply("Enabled " + args)

					defer delete.Close()

				} else {
					ctx.Reply("This command is not disabled.")

					return nil
				}
			} else {
				ctx.Reply("This command is not disabled.")

				return nil
			}
		case "kiss":
			err := db.QueryRow(fmt.Sprintf("SELECT isblocked FROM disabledcommands WHERE commandname ='%s' AND guildid ='%s'", args, ctx.Guild.ID)).Scan(&tag.isblocked)
			if err == nil {
				if tag.isblocked == "True" {
					delete, err := db.Query(fmt.Sprintf("DELETE FROM disabledcommands WHERE guildid ='%s' AND commandname ='%s'", ctx.Guild.ID, args))

					if err != nil {
						ctx.Reply("An error occurred.")

						return nil
					}
					ctx.Reply("Enabled " + args)

					defer delete.Close()

				} else {
					ctx.Reply("This command is not disabled.")

					return nil
				}
			} else {
				ctx.Reply("This command is not disabled.")

				return nil
			}
		case "ping":
			err := db.QueryRow(fmt.Sprintf("SELECT isblocked FROM disabledcommands WHERE commandname ='%s' AND guildid ='%s'", args, ctx.Guild.ID)).Scan(&tag.isblocked)
			if err == nil {
				if tag.isblocked == "True" {
					delete, err := db.Query(fmt.Sprintf("DELETE FROM disabledcommands WHERE guildid ='%s' AND commandname ='%s'", ctx.Guild.ID, args))

					if err != nil {
						ctx.Reply("An error occurred.")

						return nil
					}
					ctx.Reply("Enabled " + args)

					defer delete.Close()

				} else {
					ctx.Reply("This command is not disabled.")

					return nil
				}
			} else {
				ctx.Reply("This command is not disabled.")

				return nil
			}
		case "roles":
			err := db.QueryRow(fmt.Sprintf("SELECT isblocked FROM disabledcommands WHERE commandname ='%s' AND guildid ='%s'", args, ctx.Guild.ID)).Scan(&tag.isblocked)
			if err == nil {
				if tag.isblocked == "True" {
					delete, err := db.Query(fmt.Sprintf("DELETE FROM disabledcommands WHERE guildid ='%s' AND commandname ='%s'", ctx.Guild.ID, args))

					if err != nil {
						ctx.Reply("An error occurred.")

						return nil
					}
					ctx.Reply("Enabled " + args)

					defer delete.Close()

				} else {
					ctx.Reply("This command is not disabled.")

					return nil
				}
			} else {
				ctx.Reply("This command is not disabled.")

				return nil
			}
		case "settings":
			err := db.QueryRow(fmt.Sprintf("SELECT isblocked FROM disabledcommands WHERE commandname ='%s' AND guildid ='%s'", args, ctx.Guild.ID)).Scan(&tag.isblocked)
			if err == nil {
				if tag.isblocked == "True" {
					delete, err := db.Query(fmt.Sprintf("DELETE FROM disabledcommands WHERE guildid ='%s' AND commandname ='%s'", ctx.Guild.ID, args))

					if err != nil {
						ctx.Reply("An error occurred.")

						return nil
					}
					ctx.Reply("Enabled " + args)

					defer delete.Close()

				} else {
					ctx.Reply("This command is not disabled.")

					return nil
				}
			} else {
				ctx.Reply("This command is not disabled.")

				return nil
			}
		case "slap":
			err := db.QueryRow(fmt.Sprintf("SELECT isblocked FROM disabledcommands WHERE commandname ='%s' AND guildid ='%s'", args, ctx.Guild.ID)).Scan(&tag.isblocked)
			if err == nil {
				if tag.isblocked == "True" {
					delete, err := db.Query(fmt.Sprintf("DELETE FROM disabledcommands WHERE guildid ='%s' AND commandname ='%s'", ctx.Guild.ID, args))

					if err != nil {
						ctx.Reply("An error occurred.")

						return nil
					}
					ctx.Reply("Enabled " + args)

					defer delete.Close()

				} else {
					ctx.Reply("This command is not disabled.")

					return nil
				}
			} else {
				ctx.Reply("This command is not disabled.")

				return nil
			}
		case "spoiler":
			err := db.QueryRow(fmt.Sprintf("SELECT isblocked FROM disabledcommands WHERE commandname ='%s' AND guildid ='%s'", args, ctx.Guild.ID)).Scan(&tag.isblocked)
			if err == nil {
				if tag.isblocked == "True" {
					delete, err := db.Query(fmt.Sprintf("DELETE FROM disabledcommands WHERE guildid ='%s' AND commandname ='%s'", ctx.Guild.ID, args))

					if err != nil {
						ctx.Reply("An error occurred.")

						return nil
					}
					ctx.Reply("Enabled " + args)

					defer delete.Close()

				} else {
					ctx.Reply("This command is not disabled.")

					return nil
				}
			} else {
				ctx.Reply("This command is not disabled.")

				return nil
			}
		case "start_vote":
			err := db.QueryRow(fmt.Sprintf("SELECT isblocked FROM disabledcommands WHERE commandname ='%s' AND guildid ='%s'", args, ctx.Guild.ID)).Scan(&tag.isblocked)
			if err == nil {
				if tag.isblocked == "True" {
					delete, err := db.Query(fmt.Sprintf("DELETE FROM disabledcommands WHERE guildid ='%s' AND commandname ='%s'", ctx.Guild.ID, args))

					if err != nil {
						ctx.Reply("An error occurred.")

						return nil
					}
					ctx.Reply("Enabled " + args)

					defer delete.Close()

				} else {
					ctx.Reply("This command is not disabled.")

					return nil
				}
			} else {
				ctx.Reply("This command is not disabled.")

				return nil
			}
		case "stats":
			err := db.QueryRow(fmt.Sprintf("SELECT isblocked FROM disabledcommands WHERE commandname ='%s' AND guildid ='%s'", args, ctx.Guild.ID)).Scan(&tag.isblocked)
			if err == nil {
				if tag.isblocked == "True" {
					delete, err := db.Query(fmt.Sprintf("DELETE FROM disabledcommands WHERE guildid ='%s' AND commandname ='%s'", ctx.Guild.ID, args))

					if err != nil {
						ctx.Reply("An error occurred.")

						return nil
					}
					ctx.Reply("Enabled " + args)

					defer delete.Close()

				} else {
					ctx.Reply("This command is not disabled.")

					return nil
				}
			} else {
				ctx.Reply("This command is not disabled.")

				return nil
			}
		case "unban":
			err := db.QueryRow(fmt.Sprintf("SELECT isblocked FROM disabledcommands WHERE commandname ='%s' AND guildid ='%s'", args, ctx.Guild.ID)).Scan(&tag.isblocked)
			if err == nil {
				if tag.isblocked == "True" {
					delete, err := db.Query(fmt.Sprintf("DELETE FROM disabledcommands WHERE guildid ='%s' AND commandname ='%s'", ctx.Guild.ID, args))

					if err != nil {
						ctx.Reply("An error occurred.")

						return nil
					}
					ctx.Reply("Enabled " + args)

					defer delete.Close()

				} else {
					ctx.Reply("This command is not disabled.")

					return nil
				}
			} else {
				ctx.Reply("This command is not disabled.")

				return nil
			}
		default:
			ctx.Reply("You need to specify the command.")
			return nil

		}
	}
	return db.Close()
}
