package commands

import (
	"strings"

	multiplexer "git.randomchars.net/Reviath/RemiliaScarlet/Multiplexer"
	CommandHandler "git.randomchars.net/Reviath/RemiliaScarlet/handler"
	"git.randomchars.net/Reviath/RemiliaScarlet/sql"
)

func DisableCommand(ctx CommandHandler.Context, _ []string) error {
	db := sql.Connect()

	type Tag struct {
		isblocked string
	}

	var tag Tag

	if sql.CheckLanguage(ctx.Guild.ID) == "tr" {

		if len(strings.Join(multiplexer.GetArgs(ctx.Message.Content, multiplexer.GetPrefix()), " ")) < 1 {
			ctx.Reply("Bir komut belirtmelisin.")

			return nil
		}

		args := multiplexer.GetArgs(ctx.Message.Content, multiplexer.GetPrefix())[0]

		if args == "afk" {
			err := db.QueryRow("SELECT isblocked FROM disabledcommands WHERE commandname ='afk' AND guildid ='" + ctx.Guild.ID + "'").Scan(&tag.isblocked)
			if err == nil {
				if tag.isblocked == "True" {
					ctx.Reply("Bu komut zaten engellenmiş.")

					return nil
				}
			} else {
				insert, err := db.Query("INSERT INTO disabledcommands (isblocked, commandname, guildid) VALUES ('True', 'afk', '" + ctx.Guild.ID + "')")

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
			err := db.QueryRow("SELECT isblocked FROM disabledcommands WHERE commandname ='author' AND guildid ='" + ctx.Guild.ID + "'").Scan(&tag.isblocked)
			if err == nil {
				if tag.isblocked == "True" {
					ctx.Reply("Bu komut zaten engellenmiş.")

					return nil
				}
			} else {
				insert, err := db.Query("INSERT INTO disabledcommands (isblocked, commandname, guildid) VALUES ('True', 'author', '" + ctx.Guild.ID + "')")

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
			err := db.QueryRow("SELECT isblocked FROM disabledcommands WHERE commandname ='avatar' AND guildid ='" + ctx.Guild.ID + "'").Scan(&tag.isblocked)
			if err == nil {
				if tag.isblocked == "True" {
					ctx.Reply("Bu komut zaten engellenmiş.")

					return nil
				}
			} else {
				insert, err := db.Query("INSERT INTO disabledcommands (isblocked, commandname, guildid) VALUES ('True', 'avatar', '" + ctx.Guild.ID + "')")

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
			err := db.QueryRow("SELECT isblocked FROM disabledcommands WHERE commandname ='ban' AND guildid ='" + ctx.Guild.ID + "'").Scan(&tag.isblocked)
			if err == nil {
				if tag.isblocked == "True" {
					ctx.Reply("Bu komut zaten engellenmiş.")

					return nil
				}
			} else {
				insert, err := db.Query("INSERT INTO disabledcommands (isblocked, commandname, guildid) VALUES ('True', 'ban', '" + ctx.Guild.ID + "')")

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
			err := db.QueryRow("SELECT isblocked FROM disabledcommands WHERE commandname ='embed' AND guildid ='" + ctx.Guild.ID + "'").Scan(&tag.isblocked)
			if err == nil {
				if tag.isblocked == "True" {
					ctx.Reply("Bu komut zaten engellenmiş.")

					return nil
				}
			} else {
				insert, err := db.Query("INSERT INTO disabledcommands (isblocked, commandname, guildid) VALUES ('True', 'embed', '" + ctx.Guild.ID + "')")

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
			err := db.QueryRow("SELECT isblocked FROM disabledcommands WHERE commandname ='guild_info' AND guildid ='" + ctx.Guild.ID + "'").Scan(&tag.isblocked)
			if err == nil {
				if tag.isblocked == "True" {
					ctx.Reply("Bu komut zaten engellenmiş.")

					return nil
				}
			} else {
				insert, err := db.Query("INSERT INTO disabledcommands (isblocked, commandname, guildid) VALUES ('True', 'guild_info', '" + ctx.Guild.ID + "')")

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
			err := db.QueryRow("SELECT isblocked FROM disabledcommands WHERE commandname ='hug' AND guildid ='" + ctx.Guild.ID + "'").Scan(&tag.isblocked)
			if err == nil {
				if tag.isblocked == "True" {
					ctx.Reply("Bu komut zaten engellenmiş.")

					return nil
				}
			} else {
				insert, err := db.Query("INSERT INTO disabledcommands (isblocked, commandname, guildid) VALUES ('True', 'hug', '" + ctx.Guild.ID + "')")

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
			err := db.QueryRow("SELECT isblocked FROM disabledcommands WHERE commandname ='icon' AND guildid ='" + ctx.Guild.ID + "'").Scan(&tag.isblocked)
			if err == nil {
				if tag.isblocked == "True" {
					ctx.Reply("Bu komut zaten engellenmiş.")

					return nil
				}
			} else {
				insert, err := db.Query("INSERT INTO disabledcommands (isblocked, commandname, guildid) VALUES ('True', 'icon', '" + ctx.Guild.ID + "')")

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
			err := db.QueryRow("SELECT isblocked FROM disabledcommands WHERE commandname ='kick' AND guildid ='" + ctx.Guild.ID + "'").Scan(&tag.isblocked)
			if err == nil {
				if tag.isblocked == "True" {
					ctx.Reply("Bu komut zaten engellenmiş.")

					return nil
				}
			} else {
				insert, err := db.Query("INSERT INTO disabledcommands (isblocked, commandname, guildid) VALUES ('True', 'kick', '" + ctx.Guild.ID + "')")

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
			err := db.QueryRow("SELECT isblocked FROM disabledcommands WHERE commandname ='kiss' AND guildid ='" + ctx.Guild.ID + "'").Scan(&tag.isblocked)
			if err == nil {
				if tag.isblocked == "True" {
					ctx.Reply("Bu komut zaten engellenmiş.")

					return nil
				}
			} else {
				insert, err := db.Query("INSERT INTO disabledcommands (isblocked, commandname, guildid) VALUES ('True', 'kiss', '" + ctx.Guild.ID + "')")

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
			err := db.QueryRow("SELECT isblocked FROM disabledcommands WHERE commandname ='ping' AND guildid ='" + ctx.Guild.ID + "'").Scan(&tag.isblocked)
			if err == nil {
				if tag.isblocked == "True" {
					ctx.Reply("Bu komut zaten engellenmiş.")

					return nil
				}
			} else {
				insert, err := db.Query("INSERT INTO disabledcommands (isblocked, commandname, guildid) VALUES ('True', 'ping', '" + ctx.Guild.ID + "')")

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
			err := db.QueryRow("SELECT isblocked FROM disabledcommands WHERE commandname ='roles' AND guildid ='" + ctx.Guild.ID + "'").Scan(&tag.isblocked)
			if err == nil {
				if tag.isblocked == "True" {
					ctx.Reply("Bu komut zaten engellenmiş.")

					return nil
				}
			} else {
				insert, err := db.Query("INSERT INTO disabledcommands (isblocked, commandname, guildid) VALUES ('True', 'roles', '" + ctx.Guild.ID + "')")

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
			err := db.QueryRow("SELECT isblocked FROM disabledcommands WHERE commandname ='settings' AND guildid ='" + ctx.Guild.ID + "'").Scan(&tag.isblocked)
			if err == nil {
				if tag.isblocked == "True" {
					ctx.Reply("Bu komut zaten engellenmiş.")

					return nil
				}
			} else {
				insert, err := db.Query("INSERT INTO disabledcommands (isblocked, commandname, guildid) VALUES ('True', 'settings', '" + ctx.Guild.ID + "')")

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
			err := db.QueryRow("SELECT isblocked FROM disabledcommands WHERE commandname ='slap' AND guildid ='" + ctx.Guild.ID + "'").Scan(&tag.isblocked)
			if err == nil {
				if tag.isblocked == "True" {
					ctx.Reply("Bu komut zaten engellenmiş.")

					return nil
				}
			} else {
				insert, err := db.Query("INSERT INTO disabledcommands (isblocked, commandname, guildid) VALUES ('True', 'slap', '" + ctx.Guild.ID + "')")

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
			err := db.QueryRow("SELECT isblocked FROM disabledcommands WHERE commandname ='spoiler' AND guildid ='" + ctx.Guild.ID + "'").Scan(&tag.isblocked)
			if err == nil {
				if tag.isblocked == "True" {
					ctx.Reply("Bu komut zaten engellenmiş.")

					return nil
				}
			} else {
				insert, err := db.Query("INSERT INTO disabledcommands (isblocked, commandname, guildid) VALUES ('True', 'spoiler', '" + ctx.Guild.ID + "')")

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
			err := db.QueryRow("SELECT isblocked FROM disabledcommands WHERE commandname ='start_vote' AND guildid ='" + ctx.Guild.ID + "'").Scan(&tag.isblocked)
			if err == nil {
				if tag.isblocked == "True" {
					ctx.Reply("Bu komut zaten engellenmiş.")

					return nil
				}
			} else {
				insert, err := db.Query("INSERT INTO disabledcommands (isblocked, commandname, guildid) VALUES ('True', 'start_vote', '" + ctx.Guild.ID + "')")

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
			err := db.QueryRow("SELECT isblocked FROM disabledcommands WHERE commandname ='stats' AND guildid ='" + ctx.Guild.ID + "'").Scan(&tag.isblocked)
			if err == nil {
				if tag.isblocked == "True" {
					ctx.Reply("Bu komut zaten engellenmiş.")

					return nil
				}
			} else {
				insert, err := db.Query("INSERT INTO disabledcommands (isblocked, commandname, guildid) VALUES ('True', 'stats', '" + ctx.Guild.ID + "')")

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
			err := db.QueryRow("SELECT isblocked FROM disabledcommands WHERE commandname ='unban' AND guildid ='" + ctx.Guild.ID + "'").Scan(&tag.isblocked)
			if err == nil {
				if tag.isblocked == "True" {
					ctx.Reply("Bu komut zaten engellenmiş.")

					return nil
				}
			} else {
				insert, err := db.Query("INSERT INTO disabledcommands (isblocked, commandname, guildid) VALUES ('True', 'unban', '" + ctx.Guild.ID + "')")

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
	if len(strings.Join(multiplexer.GetArgs(ctx.Message.Content, multiplexer.GetPrefix()), " ")) < 1 {
		ctx.Reply("You need to specify a command.")

		return nil
	}

	args := multiplexer.GetArgs(ctx.Message.Content, multiplexer.GetPrefix())[0]

	if args == "afk" {
		err := db.QueryRow("SELECT isblocked FROM disabledcommands WHERE commandname ='afk' AND guildid ='" + ctx.Guild.ID + "'").Scan(&tag.isblocked)
		if err == nil {
			if tag.isblocked == "True" {
				ctx.Reply("This command is already blocked.")

				return nil
			}
		} else {
			insert, err := db.Query("INSERT INTO disabledcommands (isblocked, commandname, guildid) VALUES ('True', 'afk', '" + ctx.Guild.ID + "')")

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
		err := db.QueryRow("SELECT isblocked FROM disabledcommands WHERE commandname ='author' AND guildid ='" + ctx.Guild.ID + "'").Scan(&tag.isblocked)
		if err == nil {
			if tag.isblocked == "True" {
				ctx.Reply("This command is already blocked.")

				return nil
			}
		} else {
			insert, err := db.Query("INSERT INTO disabledcommands (isblocked, commandname, guildid) VALUES ('True', 'author', '" + ctx.Guild.ID + "')")

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
		err := db.QueryRow("SELECT isblocked FROM disabledcommands WHERE commandname ='avatar' AND guildid ='" + ctx.Guild.ID + "'").Scan(&tag.isblocked)
		if err == nil {
			if tag.isblocked == "True" {
				ctx.Reply("This command is already blocked.")

				return nil
			}
		} else {
			insert, err := db.Query("INSERT INTO disabledcommands (isblocked, commandname, guildid) VALUES ('True', 'avatar', '" + ctx.Guild.ID + "')")

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
		err := db.QueryRow("SELECT isblocked FROM disabledcommands WHERE commandname ='ban' AND guildid ='" + ctx.Guild.ID + "'").Scan(&tag.isblocked)
		if err == nil {
			if tag.isblocked == "True" {
				ctx.Reply("This command is already blocked.")

				return nil
			}
		} else {
			insert, err := db.Query("INSERT INTO disabledcommands (isblocked, commandname, guildid) VALUES ('True', 'ban', '" + ctx.Guild.ID + "')")

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
		err := db.QueryRow("SELECT isblocked FROM disabledcommands WHERE commandname ='embed' AND guildid ='" + ctx.Guild.ID + "'").Scan(&tag.isblocked)
		if err == nil {
			if tag.isblocked == "True" {
				ctx.Reply("This command is already blocked.")

				return nil
			}
		} else {
			insert, err := db.Query("INSERT INTO disabledcommands (isblocked, commandname, guildid) VALUES ('True', 'embed', '" + ctx.Guild.ID + "')")

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
		err := db.QueryRow("SELECT isblocked FROM disabledcommands WHERE commandname ='guild_info' AND guildid ='" + ctx.Guild.ID + "'").Scan(&tag.isblocked)
		if err == nil {
			if tag.isblocked == "True" {
				ctx.Reply("This command is already blocked.")

				return nil
			}
		} else {
			insert, err := db.Query("INSERT INTO disabledcommands (isblocked, commandname, guildid) VALUES ('True', 'guild_info', '" + ctx.Guild.ID + "')")

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
		err := db.QueryRow("SELECT isblocked FROM disabledcommands WHERE commandname ='hug' AND guildid ='" + ctx.Guild.ID + "'").Scan(&tag.isblocked)
		if err == nil {
			if tag.isblocked == "True" {
				ctx.Reply("This command is already blocked.")

				return nil
			}
		} else {
			insert, err := db.Query("INSERT INTO disabledcommands (isblocked, commandname, guildid) VALUES ('True', 'hug', '" + ctx.Guild.ID + "')")

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
		err := db.QueryRow("SELECT isblocked FROM disabledcommands WHERE commandname ='icon' AND guildid ='" + ctx.Guild.ID + "'").Scan(&tag.isblocked)
		if err == nil {
			if tag.isblocked == "True" {
				ctx.Reply("This command is already blocked.")

				return nil
			}
		} else {
			insert, err := db.Query("INSERT INTO disabledcommands (isblocked, commandname, guildid) VALUES ('True', 'icon', '" + ctx.Guild.ID + "')")

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
		err := db.QueryRow("SELECT isblocked FROM disabledcommands WHERE commandname ='kick' AND guildid ='" + ctx.Guild.ID + "'").Scan(&tag.isblocked)
		if err == nil {
			if tag.isblocked == "True" {
				ctx.Reply("This command is already blocked.")

				return nil
			}
		} else {
			insert, err := db.Query("INSERT INTO disabledcommands (isblocked, commandname, guildid) VALUES ('True', 'kick', '" + ctx.Guild.ID + "')")

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
		err := db.QueryRow("SELECT isblocked FROM disabledcommands WHERE commandname ='kiss' AND guildid ='" + ctx.Guild.ID + "'").Scan(&tag.isblocked)
		if err == nil {
			if tag.isblocked == "True" {
				ctx.Reply("This command is already blocked.")

				return nil
			}
		} else {
			insert, err := db.Query("INSERT INTO disabledcommands (isblocked, commandname, guildid) VALUES ('True', 'kiss', '" + ctx.Guild.ID + "')")

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
		err := db.QueryRow("SELECT isblocked FROM disabledcommands WHERE commandname ='ping' AND guildid ='" + ctx.Guild.ID + "'").Scan(&tag.isblocked)
		if err == nil {
			if tag.isblocked == "True" {
				ctx.Reply("This command is already blocked.")

				return nil
			}
		} else {
			insert, err := db.Query("INSERT INTO disabledcommands (isblocked, commandname, guildid) VALUES ('True', 'ping', '" + ctx.Guild.ID + "')")

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
		err := db.QueryRow("SELECT isblocked FROM disabledcommands WHERE commandname ='roles' AND guildid ='" + ctx.Guild.ID + "'").Scan(&tag.isblocked)
		if err == nil {
			if tag.isblocked == "True" {
				ctx.Reply("This command is already blocked.")

				return nil
			}
		} else {
			insert, err := db.Query("INSERT INTO disabledcommands (isblocked, commandname, guildid) VALUES ('True', 'roles', '" + ctx.Guild.ID + "')")

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
		err := db.QueryRow("SELECT isblocked FROM disabledcommands WHERE commandname ='settings' AND guildid ='" + ctx.Guild.ID + "'").Scan(&tag.isblocked)
		if err == nil {
			if tag.isblocked == "True" {
				ctx.Reply("This command is already blocked.")

				return nil
			}
		} else {
			insert, err := db.Query("INSERT INTO disabledcommands (isblocked, commandname, guildid) VALUES ('True', 'settings', '" + ctx.Guild.ID + "')")

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
		err := db.QueryRow("SELECT isblocked FROM disabledcommands WHERE commandname ='slap' AND guildid ='" + ctx.Guild.ID + "'").Scan(&tag.isblocked)
		if err == nil {
			if tag.isblocked == "True" {
				ctx.Reply("This command is already blocked.")

				return nil
			}
		} else {
			insert, err := db.Query("INSERT INTO disabledcommands (isblocked, commandname, guildid) VALUES ('True', 'slap', '" + ctx.Guild.ID + "')")

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
		err := db.QueryRow("SELECT isblocked FROM disabledcommands WHERE commandname ='spoiler' AND guildid ='" + ctx.Guild.ID + "'").Scan(&tag.isblocked)
		if err == nil {
			if tag.isblocked == "True" {
				ctx.Reply("This command is already blocked.")

				return nil
			}
		} else {
			insert, err := db.Query("INSERT INTO disabledcommands (isblocked, commandname, guildid) VALUES ('True', 'spoiler', '" + ctx.Guild.ID + "')")

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
		err := db.QueryRow("SELECT isblocked FROM disabledcommands WHERE commandname ='start_vote' AND guildid ='" + ctx.Guild.ID + "'").Scan(&tag.isblocked)
		if err == nil {
			if tag.isblocked == "True" {
				ctx.Reply("This command is already blocked.")

				return nil
			}
		} else {
			insert, err := db.Query("INSERT INTO disabledcommands (isblocked, commandname, guildid) VALUES ('True', 'start_vote', '" + ctx.Guild.ID + "')")

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
		err := db.QueryRow("SELECT isblocked FROM disabledcommands WHERE commandname ='stats' AND guildid ='" + ctx.Guild.ID + "'").Scan(&tag.isblocked)
		if err == nil {
			if tag.isblocked == "True" {
				ctx.Reply("This command is already blocked.")

				return nil
			}
		} else {
			insert, err := db.Query("INSERT INTO disabledcommands (isblocked, commandname, guildid) VALUES ('True', 'stats', '" + ctx.Guild.ID + "')")

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
		err := db.QueryRow("SELECT isblocked FROM disabledcommands WHERE commandname ='unban' AND guildid ='" + ctx.Guild.ID + "'").Scan(&tag.isblocked)
		if err == nil {
			if tag.isblocked == "True" {
				ctx.Reply("This command is already blocked.")

				return nil
			}
		} else {
			insert, err := db.Query("INSERT INTO disabledcommands (isblocked, commandname, guildid) VALUES ('True', 'unban', '" + ctx.Guild.ID + "')")

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
