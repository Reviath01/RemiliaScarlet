package commands

import (
	"fmt"
	"strings"

	multiplexer "git.randomchars.net/Reviath/RemiliaScarlet/Multiplexer"
	CommandHandler "git.randomchars.net/Reviath/RemiliaScarlet/handler"
	"git.randomchars.net/Reviath/RemiliaScarlet/sql"
)

func DisableCommand(ctx CommandHandler.Context, _ []string) error {
	db := sql.Connect()
	defer db.Close()

	type Tag struct {
		isblocked string
	}

	var tag Tag

	if sql.CheckLanguage(ctx.Guild.ID) == "tr" {
		if !multiplexer.CheckAdministratorPermission(ctx.Session, ctx.Message.Author.ID, ctx.Channel.ID) {
			ctx.Reply("Yeterli yetkiye sahip değilsin.")
			return nil
		}

		if len(strings.Join(multiplexer.GetArgs(ctx.Message.Content, multiplexer.GetPrefix()), " ")) < 1 {
			ctx.Reply("Bir komut belirtmelisin.")

			return nil
		}

		args := multiplexer.GetArgs(ctx.Message.Content, multiplexer.GetPrefix())[0]

		if args == "afk" {
			err := db.QueryRow(fmt.Sprintf("SELECT isblocked FROM disabledcommands WHERE commandname ='afk' AND guildid ='%s'", ctx.Guild.ID)).Scan(&tag.isblocked)
			if err == nil {
				if tag.isblocked == "True" {
					ctx.Reply("Bu komut zaten engellenmiş.")

					return nil
				}
			} else {
				insert, err := db.Query(fmt.Sprintf("INSERT INTO disabledcommands (isblocked, commandname, guildid) VALUES ('True', 'afk', '%s')", ctx.Guild.ID))

				if err != nil {
					return nil
				}
				defer insert.Close()

				if err != nil {
					ctx.Reply("Bir hata oluştu.")

					return nil
				}
				ctx.Reply("Belirtilen komut engellendi.")

				return nil
			}
		} else if args == "author" {
			err := db.QueryRow(fmt.Sprintf("SELECT isblocked FROM disabledcommands WHERE commandname ='author' AND guildid ='%s'", ctx.Guild.ID)).Scan(&tag.isblocked)
			if err == nil {
				if tag.isblocked == "True" {
					ctx.Reply("Bu komut zaten engellenmiş.")

					return nil
				}
			} else {
				insert, err := db.Query(fmt.Sprintf("INSERT INTO disabledcommands (isblocked, commandname, guildid) VALUES ('True', 'author', '%s')", ctx.Guild.ID))

				if err != nil {
					return nil
				}
				defer insert.Close()

				if err != nil {
					ctx.Reply("Bir hata oluştu.")

					return nil
				}
				ctx.Reply("Belirtilen komut engellendi.")

				return nil
			}
		} else if args == "avatar" {
			err := db.QueryRow(fmt.Sprintf("SELECT isblocked FROM disabledcommands WHERE commandname ='avatar' AND guildid ='%s'", ctx.Guild.ID)).Scan(&tag.isblocked)
			if err == nil {
				if tag.isblocked == "True" {
					ctx.Reply("Bu komut zaten engellenmiş.")

					return nil
				}
			} else {
				insert, err := db.Query(fmt.Sprintf("INSERT INTO disabledcommands (isblocked, commandname, guildid) VALUES ('True', 'avatar', '%s')", ctx.Guild.ID))

				if err != nil {
					return nil
				}
				defer insert.Close()

				if err != nil {
					ctx.Reply("Bir hata oluştu.")

					return nil
				}
				ctx.Reply("Belirtilen komut engellendi.")

				return nil
			}
		} else if args == "ban" {
			err := db.QueryRow(fmt.Sprintf("SELECT isblocked FROM disabledcommands WHERE commandname ='ban' AND guildid ='%s'", ctx.Guild.ID)).Scan(&tag.isblocked)
			if err == nil {
				if tag.isblocked == "True" {
					ctx.Reply("Bu komut zaten engellenmiş.")

					return nil
				}
			} else {
				insert, err := db.Query(fmt.Sprintf("INSERT INTO disabledcommands (isblocked, commandname, guildid) VALUES ('True', 'ban', '%s')", ctx.Guild.ID))

				if err != nil {
					return nil
				}
				defer insert.Close()

				if err != nil {
					ctx.Reply("Bir hata oluştu.")

					return nil
				}
				ctx.Reply("Belirtilen komut engellendi.")

				return nil
			}
		} else if args == "embed" {
			err := db.QueryRow(fmt.Sprintf("SELECT isblocked FROM disabledcommands WHERE commandname ='embed' AND guildid ='%s'", ctx.Guild.ID)).Scan(&tag.isblocked)
			if err == nil {
				if tag.isblocked == "True" {
					ctx.Reply("Bu komut zaten engellenmiş.")

					return nil
				}
			} else {
				insert, err := db.Query(fmt.Sprintf("INSERT INTO disabledcommands (isblocked, commandname, guildid) VALUES ('True', 'embed', '%s')", ctx.Guild.ID))

				if err != nil {
					return nil
				}
				defer insert.Close()

				if err != nil {
					ctx.Reply("Bir hata oluştu.")

					return nil
				}
				ctx.Reply("Belirtilen komut engellendi.")

				return nil
			}
		} else if args == "guild_info" {
			err := db.QueryRow(fmt.Sprintf("SELECT isblocked FROM disabledcommands WHERE commandname ='guild_info' AND guildid ='%s'", ctx.Guild.ID)).Scan(&tag.isblocked)
			if err == nil {
				if tag.isblocked == "True" {
					ctx.Reply("Bu komut zaten engellenmiş.")

					return nil
				}
			} else {
				insert, err := db.Query(fmt.Sprintf("INSERT INTO disabledcommands (isblocked, commandname, guildid) VALUES ('True', 'guild_info', '%s')", ctx.Guild.ID))

				if err != nil {
					return nil
				}
				defer insert.Close()

				if err != nil {
					ctx.Reply("Bir hata oluştu.")

					return nil
				}
				ctx.Reply("Belirtilen komut engellendi.")

				return nil
			}
		} else if args == "hug" {
			err := db.QueryRow(fmt.Sprintf("SELECT isblocked FROM disabledcommands WHERE commandname ='hug' AND guildid ='%s'", ctx.Guild.ID)).Scan(&tag.isblocked)
			if err == nil {
				if tag.isblocked == "True" {
					ctx.Reply("Bu komut zaten engellenmiş.")

					return nil
				}
			} else {
				insert, err := db.Query(fmt.Sprintf("INSERT INTO disabledcommands (isblocked, commandname, guildid) VALUES ('True', 'hug', '%s')", ctx.Guild.ID))

				if err != nil {
					return nil
				}
				defer insert.Close()

				if err != nil {
					ctx.Reply("Bir hata oluştu.")

					return nil
				}
				ctx.Reply("Belirtilen komut engellendi.")

				return nil
			}
		} else if args == "icon" {
			err := db.QueryRow(fmt.Sprintf("SELECT isblocked FROM disabledcommands WHERE commandname ='icon' AND guildid ='%s'", ctx.Guild.ID)).Scan(&tag.isblocked)
			if err == nil {
				if tag.isblocked == "True" {
					ctx.Reply("Bu komut zaten engellenmiş.")

					return nil
				}
			} else {
				insert, err := db.Query(fmt.Sprintf("INSERT INTO disabledcommands (isblocked, commandname, guildid) VALUES ('True', 'icon', '%s')", ctx.Guild.ID))

				if err != nil {
					return nil
				}

				defer insert.Close()

				if err != nil {
					ctx.Reply("Bir hata oluştu.")

					return nil
				}
				ctx.Reply("Belirtilen komut engellendi.")

				return nil
			}
		} else if args == "kick" {
			err := db.QueryRow(fmt.Sprintf("SELECT isblocked FROM disabledcommands WHERE commandname ='kick' AND guildid ='%s'", ctx.Guild.ID)).Scan(&tag.isblocked)
			if err == nil {
				if tag.isblocked == "True" {
					ctx.Reply("Bu komut zaten engellenmiş.")

					return nil
				}
			} else {
				insert, err := db.Query(fmt.Sprintf("INSERT INTO disabledcommands (isblocked, commandname, guildid) VALUES ('True', 'kick', '%s')", ctx.Guild.ID))

				if err != nil {
					return nil
				}

				defer insert.Close()

				if err != nil {
					ctx.Reply("Bir hata oluştu.")
					return nil
				}
				ctx.Reply("Belirtilen komut engellendi.")

				return nil
			}
		} else if args == "kiss" {
			err := db.QueryRow(fmt.Sprintf("SELECT isblocked FROM disabledcommands WHERE commandname ='kiss' AND guildid ='%s'", ctx.Guild.ID)).Scan(&tag.isblocked)
			if err == nil {
				if tag.isblocked == "True" {
					ctx.Reply("Bu komut zaten engellenmiş.")

					return nil
				}
			} else {
				insert, err := db.Query(fmt.Sprintf("INSERT INTO disabledcommands (isblocked, commandname, guildid) VALUES ('True', 'kiss', '%s')", ctx.Guild.ID))

				if err != nil {
					return nil
				}

				defer insert.Close()

				if err != nil {
					ctx.Reply("Bir hata oluştu.")
					return nil
				}
				ctx.Reply("Belirtilen komut engellendi.")

				return nil
			}
		} else if args == "ping" {
			err := db.QueryRow(fmt.Sprintf("SELECT isblocked FROM disabledcommands WHERE commandname ='ping' AND guildid ='%s'", ctx.Guild.ID)).Scan(&tag.isblocked)
			if err == nil {
				if tag.isblocked == "True" {
					ctx.Reply("Bu komut zaten engellenmiş.")

					return nil
				}
			} else {
				insert, err := db.Query(fmt.Sprintf("INSERT INTO disabledcommands (isblocked, commandname, guildid) VALUES ('True', 'ping', '%s')", ctx.Guild.ID))

				if err != nil {
					return nil
				}

				defer insert.Close()

				if err != nil {
					ctx.Reply("Bir hata oluştu.")
					return nil
				}
				ctx.Reply("Belirtilen komut engellendi.")

				return nil
			}
		} else if args == "roles" {
			err := db.QueryRow(fmt.Sprintf("SELECT isblocked FROM disabledcommands WHERE commandname ='roles' AND guildid ='%s'", ctx.Guild.ID)).Scan(&tag.isblocked)
			if err == nil {
				if tag.isblocked == "True" {
					ctx.Reply("Bu komut zaten engellenmiş.")

					return nil
				}
			} else {
				insert, err := db.Query(fmt.Sprintf("INSERT INTO disabledcommands (isblocked, commandname, guildid) VALUES ('True', 'roles', '%s')", ctx.Guild.ID))

				if err != nil {
					return nil
				}

				defer insert.Close()

				if err != nil {
					ctx.Reply("Bir hata oluştu.")
					return nil
				}
				ctx.Reply("Belirtilen komut engellendi.")

				return nil
			}
		} else if args == "settings" {
			err := db.QueryRow(fmt.Sprintf("SELECT isblocked FROM disabledcommands WHERE commandname ='settings' AND guildid ='%s'", ctx.Guild.ID)).Scan(&tag.isblocked)
			if err == nil {
				if tag.isblocked == "True" {
					ctx.Reply("Bu komut zaten engellenmiş.")

					return nil
				}
			} else {
				insert, err := db.Query(fmt.Sprintf("INSERT INTO disabledcommands (isblocked, commandname, guildid) VALUES ('True', 'settings', '%s')", ctx.Guild.ID))

				if err != nil {
					return nil
				}

				defer insert.Close()

				if err != nil {
					ctx.Reply("Bir hata oluştu.")
					return nil
				}
				ctx.Reply("Belirtilen komut engellendi.")

				return nil
			}
		} else if args == "slap" {
			err := db.QueryRow(fmt.Sprintf("SELECT isblocked FROM disabledcommands WHERE commandname ='slap' AND guildid ='%s'", ctx.Guild.ID)).Scan(&tag.isblocked)
			if err == nil {
				if tag.isblocked == "True" {
					ctx.Reply("Bu komut zaten engellenmiş.")

					return nil
				}
			} else {
				insert, err := db.Query(fmt.Sprintf("INSERT INTO disabledcommands (isblocked, commandname, guildid) VALUES ('True', 'slap', '%s')", ctx.Guild.ID))

				if err != nil {
					return nil
				}
				defer insert.Close()

				if err != nil {
					ctx.Reply("Bir hata oluştu.")
					return nil
				}
				ctx.Reply("Belirtilen komut engellendi.")

				return nil
			}
		} else if args == "spoiler" {
			err := db.QueryRow(fmt.Sprintf("SELECT isblocked FROM disabledcommands WHERE commandname ='spoiler' AND guildid ='%s'", ctx.Guild.ID)).Scan(&tag.isblocked)
			if err == nil {
				if tag.isblocked == "True" {
					ctx.Reply("Bu komut zaten engellenmiş.")

					return nil
				}
			} else {
				insert, err := db.Query(fmt.Sprintf("INSERT INTO disabledcommands (isblocked, commandname, guildid) VALUES ('True', 'spoiler', '%s')", ctx.Guild.ID))

				if err != nil {
					return nil
				}
				defer insert.Close()

				if err != nil {
					ctx.Reply("Bir hata oluştu.")
					return nil
				}
				ctx.Reply("Belirtilen komut engellendi.")

				return nil
			}
		} else if args == "start_vote" {
			err := db.QueryRow(fmt.Sprintf("SELECT isblocked FROM disabledcommands WHERE commandname ='start_vote' AND guildid ='%s'", ctx.Guild.ID)).Scan(&tag.isblocked)
			if err == nil {
				if tag.isblocked == "True" {
					ctx.Reply("Bu komut zaten engellenmiş.")

					return nil
				}
			} else {
				insert, err := db.Query(fmt.Sprintf("INSERT INTO disabledcommands (isblocked, commandname, guildid) VALUES ('True', 'start_vote', '%s')", ctx.Guild.ID))

				if err != nil {
					return nil
				}
				defer insert.Close()

				if err != nil {
					ctx.Reply("Bir hata oluştu.")
					return nil
				}
				ctx.Reply("Belirtilen komut engellendi.")

				return nil
			}
		} else if args == "stats" {
			err := db.QueryRow(fmt.Sprintf("SELECT isblocked FROM disabledcommands WHERE commandname ='stats' AND guildid ='%s'", ctx.Guild.ID)).Scan(&tag.isblocked)
			if err == nil {
				if tag.isblocked == "True" {
					ctx.Reply("Bu komut zaten engellenmiş.")

					return nil
				}
			} else {
				insert, err := db.Query(fmt.Sprintf("INSERT INTO disabledcommands (isblocked, commandname, guildid) VALUES ('True', 'stats', '%s')", ctx.Guild.ID))

				if err != nil {
					return nil
				}

				defer insert.Close()

				if err != nil {
					ctx.Reply("Bir hata oluştu.")
					return nil
				}
				ctx.Reply("Belirtilen komut engellendi.")

				return nil
			}
		} else if args == "unban" {
			err := db.QueryRow(fmt.Sprintf("SELECT isblocked FROM disabledcommands WHERE commandname ='unban' AND guildid ='%s'", ctx.Guild.ID)).Scan(&tag.isblocked)
			if err == nil {
				if tag.isblocked == "True" {
					ctx.Reply("Bu komut zaten engellenmiş.")

					return nil
				}
			} else {
				insert, err := db.Query(fmt.Sprintf("INSERT INTO disabledcommands (isblocked, commandname, guildid) VALUES ('True', 'unban', '%s')", ctx.Guild.ID))

				if err != nil {
					return nil
				}
				defer insert.Close()

				if err != nil {
					ctx.Reply("Bir hata oluştu.")
					return nil
				}
				ctx.Reply("Belirtilen komut engellendi.")

				return nil
			}
		} else {
			ctx.Reply("You need to specify the command.")
			return nil
		}
	}

	if !multiplexer.CheckAdministratorPermission(ctx.Session, ctx.Message.Author.ID, ctx.Channel.ID) {
		ctx.Reply("You don't have enough permission.")
		return nil
	}

	if len(strings.Join(multiplexer.GetArgs(ctx.Message.Content, multiplexer.GetPrefix()), " ")) < 1 {
		ctx.Reply("You need to specify a command.")

		return nil
	}

	args := multiplexer.GetArgs(ctx.Message.Content, multiplexer.GetPrefix())[0]

	if args == "afk" {
		err := db.QueryRow(fmt.Sprintf("SELECT isblocked FROM disabledcommands WHERE commandname ='afk' AND guildid ='%s'", ctx.Guild.ID)).Scan(&tag.isblocked)
		if err == nil {
			if tag.isblocked == "True" {
				ctx.Reply("This command is already blocked.")

				return nil
			}
		} else {
			insert, err := db.Query(fmt.Sprintf("INSERT INTO disabledcommands (isblocked, commandname, guildid) VALUES ('True', 'afk', '%s')", ctx.Guild.ID))

			if err != nil {
				return nil
			}

			if err != nil {
				return nil
			}

			defer insert.Close()

			if err != nil {
				ctx.Reply("An error occurred.")

				return nil
			}
			ctx.Reply("Disabled afk command.")

			return nil
		}
	} else if args == "author" {
		err := db.QueryRow(fmt.Sprintf("SELECT isblocked FROM disabledcommands WHERE commandname ='author' AND guildid ='%s'", ctx.Guild.ID)).Scan(&tag.isblocked)
		if err == nil {
			if tag.isblocked == "True" {
				ctx.Reply("This command is already blocked.")

				return nil
			}
		} else {
			insert, err := db.Query(fmt.Sprintf("INSERT INTO disabledcommands (isblocked, commandname, guildid) VALUES ('True', 'author', '%s')", ctx.Guild.ID))

			if err != nil {
				return nil
			}

			if err != nil {
				return nil
			}

			defer insert.Close()

			if err != nil {
				ctx.Reply("An error occurred.")

				return nil
			}
			ctx.Reply("Disabled author command.")

			return nil
		}
	} else if args == "avatar" {
		err := db.QueryRow(fmt.Sprintf("SELECT isblocked FROM disabledcommands WHERE commandname ='avatar' AND guildid ='%s'", ctx.Guild.ID)).Scan(&tag.isblocked)
		if err == nil {
			if tag.isblocked == "True" {
				ctx.Reply("This command is already blocked.")

				return nil
			}
		} else {
			insert, err := db.Query(fmt.Sprintf("INSERT INTO disabledcommands (isblocked, commandname, guildid) VALUES ('True', 'avatar', '%s')", ctx.Guild.ID))

			if err != nil {
				return nil
			}

			if err != nil {
				return nil
			}

			defer insert.Close()

			if err != nil {
				ctx.Reply("An error occurred.")

				return nil
			}
			ctx.Reply("Disabled avatar command.")

			return nil
		}
	} else if args == "ban" {
		err := db.QueryRow(fmt.Sprintf("SELECT isblocked FROM disabledcommands WHERE commandname ='ban' AND guildid ='%s'", ctx.Guild.ID)).Scan(&tag.isblocked)
		if err == nil {
			if tag.isblocked == "True" {
				ctx.Reply("This command is already blocked.")

				return nil
			}
		} else {
			insert, err := db.Query(fmt.Sprintf("INSERT INTO disabledcommands (isblocked, commandname, guildid) VALUES ('True', 'ban', '%s')", ctx.Guild.ID))

			if err != nil {
				return nil
			}

			if err != nil {
				return nil
			}

			defer insert.Close()

			if err != nil {
				ctx.Reply("An error occurred.")

				return nil
			}
			ctx.Reply("Disabled ban command.")

			return nil
		}
	} else if args == "embed" {
		err := db.QueryRow(fmt.Sprintf("SELECT isblocked FROM disabledcommands WHERE commandname ='embed' AND guildid ='%s'", ctx.Guild.ID)).Scan(&tag.isblocked)
		if err == nil {
			if tag.isblocked == "True" {
				ctx.Reply("This command is already blocked.")

				return nil
			}
		} else {
			insert, err := db.Query(fmt.Sprintf("INSERT INTO disabledcommands (isblocked, commandname, guildid) VALUES ('True', 'embed', '%s')", ctx.Guild.ID))

			if err != nil {
				return nil
			}

			if err != nil {
				return nil
			}

			defer insert.Close()

			if err != nil {
				ctx.Reply("An error occurred.")

				return nil
			}
			ctx.Reply("Disabled embed command.")

			return nil
		}
	} else if args == "guild_info" {
		err := db.QueryRow(fmt.Sprintf("SELECT isblocked FROM disabledcommands WHERE commandname ='guild_info' AND guildid ='%s'", ctx.Guild.ID)).Scan(&tag.isblocked)
		if err == nil {
			if tag.isblocked == "True" {
				ctx.Reply("This command is already blocked.")

				return nil
			}
		} else {
			insert, err := db.Query(fmt.Sprintf("INSERT INTO disabledcommands (isblocked, commandname, guildid) VALUES ('True', 'guild_info', '%s')", ctx.Guild.ID))

			if err != nil {
				return nil
			}

			if err != nil {
				return nil
			}

			defer insert.Close()

			if err != nil {
				ctx.Reply("An error occurred.")

				return nil
			}
			ctx.Reply("Disabled guild_info command.")

			return nil
		}
	} else if args == "hug" {
		err := db.QueryRow(fmt.Sprintf("SELECT isblocked FROM disabledcommands WHERE commandname ='hug' AND guildid ='%s'", ctx.Guild.ID)).Scan(&tag.isblocked)
		if err == nil {
			if tag.isblocked == "True" {
				ctx.Reply("This command is already blocked.")

				return nil
			}
		} else {
			insert, err := db.Query(fmt.Sprintf("INSERT INTO disabledcommands (isblocked, commandname, guildid) VALUES ('True', 'hug', '%s')", ctx.Guild.ID))

			if err != nil {
				return nil
			}

			if err != nil {
				return nil
			}

			defer insert.Close()

			if err != nil {
				ctx.Reply("An error occurred.")

				return nil
			}
			ctx.Reply("Disabled hug command.")

			return nil
		}
	} else if args == "icon" {
		err := db.QueryRow(fmt.Sprintf("SELECT isblocked FROM disabledcommands WHERE commandname ='icon' AND guildid ='%s'", ctx.Guild.ID)).Scan(&tag.isblocked)
		if err == nil {
			if tag.isblocked == "True" {
				ctx.Reply("This command is already blocked.")

				return nil
			}
		} else {
			insert, err := db.Query(fmt.Sprintf("INSERT INTO disabledcommands (isblocked, commandname, guildid) VALUES ('True', 'icon', '%s')", ctx.Guild.ID))

			if err != nil {
				return nil
			}

			if err != nil {
				return nil
			}

			defer insert.Close()

			if err != nil {
				ctx.Reply("An error occurred.")

				return nil
			}
			ctx.Reply("Disabled icon command.")

			return nil
		}
	} else if args == "kick" {
		err := db.QueryRow(fmt.Sprintf("SELECT isblocked FROM disabledcommands WHERE commandname ='kick' AND guildid ='%s'", ctx.Guild.ID)).Scan(&tag.isblocked)
		if err == nil {
			if tag.isblocked == "True" {
				ctx.Reply("This command is already blocked.")

				return nil
			}
		} else {
			insert, err := db.Query(fmt.Sprintf("INSERT INTO disabledcommands (isblocked, commandname, guildid) VALUES ('True', 'kick', '%s')", ctx.Guild.ID))

			if err != nil {
				return nil
			}

			if err != nil {
				return nil
			}

			defer insert.Close()

			if err != nil {
				ctx.Reply("An error occurred.")
				return nil
			}
			ctx.Reply("Disabled kick command.")

			return nil
		}
	} else if args == "kiss" {
		err := db.QueryRow(fmt.Sprintf("SELECT isblocked FROM disabledcommands WHERE commandname ='kiss' AND guildid ='%s'", ctx.Guild.ID)).Scan(&tag.isblocked)
		if err == nil {
			if tag.isblocked == "True" {
				ctx.Reply("This command is already blocked.")

				return nil
			}
		} else {
			insert, err := db.Query(fmt.Sprintf("INSERT INTO disabledcommands (isblocked, commandname, guildid) VALUES ('True', 'kiss', '%s')", ctx.Guild.ID))

			if err != nil {
				return nil
			}

			if err != nil {
				return nil
			}

			defer insert.Close()

			if err != nil {
				ctx.Reply("An error occurred.")
				return nil
			}
			ctx.Reply("Disabled kiss command.")

			return nil
		}
	} else if args == "ping" {
		err := db.QueryRow(fmt.Sprintf("SELECT isblocked FROM disabledcommands WHERE commandname ='ping' AND guildid ='%s'", ctx.Guild.ID)).Scan(&tag.isblocked)
		if err == nil {
			if tag.isblocked == "True" {
				ctx.Reply("This command is already blocked.")

				return nil
			}
		} else {
			insert, err := db.Query(fmt.Sprintf("INSERT INTO disabledcommands (isblocked, commandname, guildid) VALUES ('True', 'ping', '%s')", ctx.Guild.ID))

			if err != nil {
				return nil
			}

			if err != nil {
				return nil
			}

			defer insert.Close()

			if err != nil {
				ctx.Reply("An error occurred.")
				return nil
			}
			ctx.Reply("Disabled ping command.")

			return nil
		}
	} else if args == "roles" {
		err := db.QueryRow(fmt.Sprintf("SELECT isblocked FROM disabledcommands WHERE commandname ='roles' AND guildid ='%s'", ctx.Guild.ID)).Scan(&tag.isblocked)
		if err == nil {
			if tag.isblocked == "True" {
				ctx.Reply("This command is already blocked.")

				return nil
			}
		} else {
			insert, err := db.Query(fmt.Sprintf("INSERT INTO disabledcommands (isblocked, commandname, guildid) VALUES ('True', 'roles', '%s')", ctx.Guild.ID))

			if err != nil {
				return nil
			}

			if err != nil {
				return nil
			}

			defer insert.Close()

			if err != nil {
				ctx.Reply("An error occurred.")
				return nil
			}
			ctx.Reply("Disabled roles command.")

			return nil
		}
	} else if args == "settings" {
		err := db.QueryRow(fmt.Sprintf("SELECT isblocked FROM disabledcommands WHERE commandname ='settings' AND guildid ='%s'", ctx.Guild.ID)).Scan(&tag.isblocked)
		if err == nil {
			if tag.isblocked == "True" {
				ctx.Reply("This command is already blocked.")

				return nil
			}
		} else {
			insert, err := db.Query(fmt.Sprintf("INSERT INTO disabledcommands (isblocked, commandname, guildid) VALUES ('True', 'settings', '%s')", ctx.Guild.ID))

			if err != nil {
				return nil
			}

			if err != nil {
				return nil
			}

			defer insert.Close()

			if err != nil {
				ctx.Reply("An error occurred.")
				return nil
			}
			ctx.Reply("Disabled settings command.")

			return nil
		}
	} else if args == "slap" {
		err := db.QueryRow(fmt.Sprintf("SELECT isblocked FROM disabledcommands WHERE commandname ='slap' AND guildid ='%s'", ctx.Guild.ID)).Scan(&tag.isblocked)
		if err == nil {
			if tag.isblocked == "True" {
				ctx.Reply("This command is already blocked.")

				return nil
			}
		} else {
			insert, err := db.Query(fmt.Sprintf("INSERT INTO disabledcommands (isblocked, commandname, guildid) VALUES ('True', 'slap', '%s')", ctx.Guild.ID))

			if err != nil {
				return nil
			}

			if err != nil {
				return nil
			}

			defer insert.Close()

			if err != nil {
				ctx.Reply("An error occurred.")
				return nil
			}
			ctx.Reply("Disabled slap command.")

			return nil
		}
	} else if args == "spoiler" {
		err := db.QueryRow(fmt.Sprintf("SELECT isblocked FROM disabledcommands WHERE commandname ='spoiler' AND guildid ='%s'", ctx.Guild.ID)).Scan(&tag.isblocked)
		if err == nil {
			if tag.isblocked == "True" {
				ctx.Reply("This command is already blocked.")

				return nil
			}
		} else {
			insert, err := db.Query(fmt.Sprintf("INSERT INTO disabledcommands (isblocked, commandname, guildid) VALUES ('True', 'spoiler', '%s')", ctx.Guild.ID))

			if err != nil {
				return nil
			}

			if err != nil {
				return nil
			}

			defer insert.Close()

			if err != nil {
				ctx.Reply("An error occurred.")
				return nil
			}
			ctx.Reply("Disabled spoiler command.")

			return nil
		}
	} else if args == "start_vote" {
		err := db.QueryRow(fmt.Sprintf("SELECT isblocked FROM disabledcommands WHERE commandname ='start_vote' AND guildid ='%s'", ctx.Guild.ID)).Scan(&tag.isblocked)
		if err == nil {
			if tag.isblocked == "True" {
				ctx.Reply("This command is already blocked.")

				return nil
			}
		} else {
			insert, err := db.Query(fmt.Sprintf("INSERT INTO disabledcommands (isblocked, commandname, guildid) VALUES ('True', 'start_vote', '%s')", ctx.Guild.ID))

			if err != nil {
				return nil
			}

			if err != nil {
				return nil
			}

			defer insert.Close()

			if err != nil {
				ctx.Reply("An error occurred.")
				return nil
			}
			ctx.Reply("Disabled start_vote command.")

			return nil
		}
	} else if args == "stats" {
		err := db.QueryRow(fmt.Sprintf("SELECT isblocked FROM disabledcommands WHERE commandname ='stats' AND guildid ='%s'", ctx.Guild.ID)).Scan(&tag.isblocked)
		if err == nil {
			if tag.isblocked == "True" {
				ctx.Reply("This command is already blocked.")

				return nil
			}
		} else {
			insert, err := db.Query(fmt.Sprintf("INSERT INTO disabledcommands (isblocked, commandname, guildid) VALUES ('True', 'stats', '%s')", ctx.Guild.ID))

			if err != nil {
				return nil
			}

			if err != nil {
				return nil
			}

			defer insert.Close()

			if err != nil {
				ctx.Reply("An error occurred.")
				return nil
			}
			ctx.Reply("Disabled stats command.")

			return nil
		}
	} else if args == "unban" {
		err := db.QueryRow(fmt.Sprintf("SELECT isblocked FROM disabledcommands WHERE commandname ='unban' AND guildid ='%s'", ctx.Guild.ID)).Scan(&tag.isblocked)
		if err == nil {
			if tag.isblocked == "True" {
				ctx.Reply("This command is already blocked.")

				return nil
			}
		} else {
			insert, err := db.Query(fmt.Sprintf("INSERT INTO disabledcommands (isblocked, commandname, guildid) VALUES ('True', 'unban', '%s')", ctx.Guild.ID))

			if err != nil {
				return nil
			}

			if err != nil {
				return nil
			}

			defer insert.Close()

			if err != nil {
				ctx.Reply("An error occurred.")
				return nil
			}
			ctx.Reply("Disabled unban command.")

			return nil
		}
	} else {
		ctx.Reply("You need to specify the command.")
		return nil
	}
	return db.Close()
}
