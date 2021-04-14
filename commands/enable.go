package commands

import (
	"strings"

	ctx "git.randomchars.net/Reviath/RemiliaScarlet/CommandHandler/Context"
	"git.randomchars.net/Reviath/RemiliaScarlet/sql"
	"github.com/bwmarrin/discordgo"
)

type Enable struct {
}

func (e Enable) Execute(ctx ctx.Ctx, session *discordgo.Session) error {
	db := sql.Connect()

	type Tag struct {
		isblocked string
	}

	var tag Tag

	if sql.CheckLanguage(ctx.Guild().ID) == "tr" {
		perms, err := session.State.UserChannelPermissions(ctx.Author().ID, ctx.Channel().ID)
		if err == nil && !(int(perms)&discordgo.PermissionAdministrator == discordgo.PermissionAdministrator) {
			_, _ = session.ChannelMessageSend(ctx.Channel().ID, "Bu komutu kullanmak için yönetici yetkisine sahip olmalısın.")

			return nil
		}

		if len(strings.Join(ctx.Args(), " ")) < 1 {
			_, _ = session.ChannelMessageSend(ctx.Channel().ID, "Bir komut belirtmelisin.")

			return nil
		}

		args := ctx.Args()[0]

		if args == "afk" {
			err = db.QueryRow("SELECT isblocked FROM disabledcommands WHERE commandname ='" + args + "' AND guildid ='" + ctx.Guild().ID + "'").Scan(&tag.isblocked)
			if err == nil {
				if tag.isblocked == "True" {
					delete, err := db.Query("DELETE FROM disabledcommands WHERE guildid ='" + ctx.Guild().ID + "' AND commandname ='" + args + "'")

					if err != nil {
						_, _ = session.ChannelMessageSend(ctx.Channel().ID, "Bir hata oluştu.")
						return nil
					}
					_, _ = session.ChannelMessageSend(ctx.Channel().ID, args+" komutu tekrar kullanılabilir!")

					defer delete.Close()
					return nil
				}
				_, _ = session.ChannelMessageSend(ctx.Channel().ID, "Bu komut engellenmemiş.")
				return nil
			} else {
				_, _ = session.ChannelMessageSend(ctx.Channel().ID, "Bu komut engellenmemiş.")

				return nil
			}
		} else if args == "author" {
			err = db.QueryRow("SELECT isblocked FROM disabledcommands WHERE commandname ='" + args + "' AND guildid ='" + ctx.Guild().ID + "'").Scan(&tag.isblocked)
			if err == nil {
				if tag.isblocked == "True" {
					delete, err := db.Query("DELETE FROM disabledcommands WHERE guildid ='" + ctx.Guild().ID + "' AND commandname ='" + args + "'")

					if err != nil {
						_, _ = session.ChannelMessageSend(ctx.Channel().ID, "Bir hata oluştu.")
						return nil
					}
					_, _ = session.ChannelMessageSend(ctx.Channel().ID, args+" komutu tekrar kullanılabilir!")

					defer delete.Close()
					return nil
				}
				_, _ = session.ChannelMessageSend(ctx.Channel().ID, "Bu komut engellenmemiş.")
				return nil
			} else {
				_, _ = session.ChannelMessageSend(ctx.Channel().ID, "Bu komut engellenmemiş.")

				return nil
			}
		} else if args == "avatar" {
			err = db.QueryRow("SELECT isblocked FROM disabledcommands WHERE commandname ='" + args + "' AND guildid ='" + ctx.Guild().ID + "'").Scan(&tag.isblocked)
			if err == nil {
				if tag.isblocked == "True" {
					delete, err := db.Query("DELETE FROM disabledcommands WHERE guildid ='" + ctx.Guild().ID + "' AND commandname ='" + args + "'")

					if err != nil {
						_, _ = session.ChannelMessageSend(ctx.Channel().ID, "Bir hata oluştu.")
						return nil
					}
					_, _ = session.ChannelMessageSend(ctx.Channel().ID, args+" komutu tekrar kullanılabilir!")

					defer delete.Close()
					return nil
				}
				_, _ = session.ChannelMessageSend(ctx.Channel().ID, "Bu komut engellenmemiş.")
				return nil
			} else {
				_, _ = session.ChannelMessageSend(ctx.Channel().ID, "Bu komut engellenmemiş.")

				return nil
			}
		} else if args == "ban" {
			err = db.QueryRow("SELECT isblocked FROM disabledcommands WHERE commandname ='" + args + "' AND guildid ='" + ctx.Guild().ID + "'").Scan(&tag.isblocked)
			if err == nil {
				if tag.isblocked == "True" {
					delete, err := db.Query("DELETE FROM disabledcommands WHERE guildid ='" + ctx.Guild().ID + "' AND commandname ='" + args + "'")

					if err != nil {
						_, _ = session.ChannelMessageSend(ctx.Channel().ID, "Bir hata oluştu.")
						return nil
					}
					_, _ = session.ChannelMessageSend(ctx.Channel().ID, args+" komutu tekrar kullanılabilir!")

					defer delete.Close()
					return nil
				}
				_, _ = session.ChannelMessageSend(ctx.Channel().ID, "Bu komut engellenmemiş.")
				return nil
			} else {
				_, _ = session.ChannelMessageSend(ctx.Channel().ID, "Bu komut engellenmemiş.")

				return nil
			}
		} else if args == "embed" {
			err = db.QueryRow("SELECT isblocked FROM disabledcommands WHERE commandname ='" + args + "' AND guildid ='" + ctx.Guild().ID + "'").Scan(&tag.isblocked)
			if err == nil {
				if tag.isblocked == "True" {
					delete, err := db.Query("DELETE FROM disabledcommands WHERE guildid ='" + ctx.Guild().ID + "' AND commandname ='" + args + "'")

					if err != nil {
						_, _ = session.ChannelMessageSend(ctx.Channel().ID, "Bir hata oluştu.")
						return nil
					}
					_, _ = session.ChannelMessageSend(ctx.Channel().ID, args+" komutu tekrar kullanılabilir!")

					defer delete.Close()
					return nil
				}
				_, _ = session.ChannelMessageSend(ctx.Channel().ID, "Bu komut engellenmemiş.")
				return nil
			} else {
				_, _ = session.ChannelMessageSend(ctx.Channel().ID, "Bu komut engellenmemiş.")

				return nil
			}
		} else if args == "Channel_info" {
			err = db.QueryRow("SELECT isblocked FROM disabledcommands WHERE commandname ='" + args + "' AND guildid ='" + ctx.Guild().ID + "'").Scan(&tag.isblocked)
			if err == nil {
				if tag.isblocked == "True" {
					delete, err := db.Query("DELETE FROM disabledcommands WHERE guildid ='" + ctx.Guild().ID + "' AND commandname ='" + args + "'")

					if err != nil {
						_, _ = session.ChannelMessageSend(ctx.Channel().ID, "Bir hata oluştu.")
						return nil
					}
					_, _ = session.ChannelMessageSend(ctx.Channel().ID, args+" komutu tekrar kullanılabilir!")

					defer delete.Close()
					return nil
				}
				_, _ = session.ChannelMessageSend(ctx.Channel().ID, "Bu komut engellenmemiş.")
				return nil
			} else {
				_, _ = session.ChannelMessageSend(ctx.Channel().ID, "Bu komut engellenmemiş.")

				return nil
			}
		} else if args == "hug" {
			err = db.QueryRow("SELECT isblocked FROM disabledcommands WHERE commandname ='" + args + "' AND guildid ='" + ctx.Guild().ID + "'").Scan(&tag.isblocked)
			if err == nil {
				if tag.isblocked == "True" {
					delete, err := db.Query("DELETE FROM disabledcommands WHERE guildid ='" + ctx.Guild().ID + "' AND commandname ='" + args + "'")

					if err != nil {
						_, _ = session.ChannelMessageSend(ctx.Channel().ID, "Bir hata oluştu.")
						return nil
					}
					_, _ = session.ChannelMessageSend(ctx.Channel().ID, args+" komutu tekrar kullanılabilir!")

					defer delete.Close()
					return nil
				}
				_, _ = session.ChannelMessageSend(ctx.Channel().ID, "Bu komut engellenmemiş.")
				return nil
			} else {
				_, _ = session.ChannelMessageSend(ctx.Channel().ID, "Bu komut engellenmemiş.")

				return nil
			}
		} else if args == "kick" {
			err = db.QueryRow("SELECT isblocked FROM disabledcommands WHERE commandname ='" + args + "' AND guildid ='" + ctx.Guild().ID + "'").Scan(&tag.isblocked)
			if err == nil {
				if tag.isblocked == "True" {
					delete, err := db.Query("DELETE FROM disabledcommands WHERE guildid ='" + ctx.Guild().ID + "' AND commandname ='" + args + "'")

					if err != nil {
						_, _ = session.ChannelMessageSend(ctx.Channel().ID, "Bir hata oluştu.")
						return nil
					}
					_, _ = session.ChannelMessageSend(ctx.Channel().ID, args+" komutu tekrar kullanılabilir!")

					defer delete.Close()
					return nil
				}
				_, _ = session.ChannelMessageSend(ctx.Channel().ID, "Bu komut engellenmemiş.")
				return nil
			} else {
				_, _ = session.ChannelMessageSend(ctx.Channel().ID, "Bu komut engellenmemiş.")

				return nil
			}
		} else if args == "kiss" {
			err = db.QueryRow("SELECT isblocked FROM disabledcommands WHERE commandname ='" + args + "' AND guildid ='" + ctx.Guild().ID + "'").Scan(&tag.isblocked)
			if err == nil {
				if tag.isblocked == "True" {
					delete, err := db.Query("DELETE FROM disabledcommands WHERE guildid ='" + ctx.Guild().ID + "' AND commandname ='" + args + "'")

					if err != nil {
						_, _ = session.ChannelMessageSend(ctx.Channel().ID, "Bir hata oluştu.")
						return nil
					}
					_, _ = session.ChannelMessageSend(ctx.Channel().ID, args+" komutu tekrar kullanılabilir!")

					defer delete.Close()
					return nil
				}
				_, _ = session.ChannelMessageSend(ctx.Channel().ID, "Bu komut engellenmemiş.")
				return nil
			} else {
				_, _ = session.ChannelMessageSend(ctx.Channel().ID, "Bu komut engellenmemiş.")

				return nil
			}
		} else if args == "ping" {
			err = db.QueryRow("SELECT isblocked FROM disabledcommands WHERE commandname ='" + args + "' AND guildid ='" + ctx.Guild().ID + "'").Scan(&tag.isblocked)
			if err == nil {
				if tag.isblocked == "True" {
					delete, err := db.Query("DELETE FROM disabledcommands WHERE guildid ='" + ctx.Guild().ID + "' AND commandname ='" + args + "'")

					if err != nil {
						_, _ = session.ChannelMessageSend(ctx.Channel().ID, "Bir hata oluştu.")
						return nil
					}
					_, _ = session.ChannelMessageSend(ctx.Channel().ID, args+" komutu tekrar kullanılabilir!")

					defer delete.Close()
					return nil
				}
				_, _ = session.ChannelMessageSend(ctx.Channel().ID, "Bu komut engellenmemiş.")
				return nil
			} else {
				_, _ = session.ChannelMessageSend(ctx.Channel().ID, "Bu komut engellenmemiş.")

				return nil
			}
		} else if args == "roles" {
			err = db.QueryRow("SELECT isblocked FROM disabledcommands WHERE commandname ='" + args + "' AND guildid ='" + ctx.Guild().ID + "'").Scan(&tag.isblocked)
			if err == nil {
				if tag.isblocked == "True" {
					delete, err := db.Query("DELETE FROM disabledcommands WHERE guildid ='" + ctx.Guild().ID + "' AND commandname ='" + args + "'")

					if err != nil {
						_, _ = session.ChannelMessageSend(ctx.Channel().ID, "Bir hata oluştu.")
						return nil
					}
					_, _ = session.ChannelMessageSend(ctx.Channel().ID, args+" komutu tekrar kullanılabilir!")

					defer delete.Close()
					return nil
				}
				_, _ = session.ChannelMessageSend(ctx.Channel().ID, "Bu komut engellenmemiş.")
				return nil
			} else {
				_, _ = session.ChannelMessageSend(ctx.Channel().ID, "Bu komut engellenmemiş.")

				return nil
			}
		} else if args == "settings" {
			err = db.QueryRow("SELECT isblocked FROM disabledcommands WHERE commandname ='" + args + "' AND guildid ='" + ctx.Guild().ID + "'").Scan(&tag.isblocked)
			if err == nil {
				if tag.isblocked == "True" {
					delete, err := db.Query("DELETE FROM disabledcommands WHERE guildid ='" + ctx.Guild().ID + "' AND commandname ='" + args + "'")

					if err != nil {
						_, _ = session.ChannelMessageSend(ctx.Channel().ID, "Bir hata oluştu.")
						return nil
					}
					_, _ = session.ChannelMessageSend(ctx.Channel().ID, args+" komutu tekrar kullanılabilir!")

					defer delete.Close()
					return nil
				}
				_, _ = session.ChannelMessageSend(ctx.Channel().ID, "Bu komut engellenmemiş.")
				return nil
			} else {
				_, _ = session.ChannelMessageSend(ctx.Channel().ID, "Bu komut engellenmemiş.")

				return nil
			}
		} else if args == "slap" {
			err = db.QueryRow("SELECT isblocked FROM disabledcommands WHERE commandname ='" + args + "' AND guildid ='" + ctx.Guild().ID + "'").Scan(&tag.isblocked)
			if err == nil {
				if tag.isblocked == "True" {
					delete, err := db.Query("DELETE FROM disabledcommands WHERE guildid ='" + ctx.Guild().ID + "' AND commandname ='" + args + "'")

					if err != nil {
						_, _ = session.ChannelMessageSend(ctx.Channel().ID, "Bir hata oluştu.")
						return nil
					}
					_, _ = session.ChannelMessageSend(ctx.Channel().ID, args+" komutu tekrar kullanılabilir!")

					defer delete.Close()
					return nil
				}
				_, _ = session.ChannelMessageSend(ctx.Channel().ID, "Bu komut engellenmemiş.")
				return nil
			} else {
				_, _ = session.ChannelMessageSend(ctx.Channel().ID, "Bu komut engellenmemiş.")

				return nil
			}
		} else if args == "spoiler" {
			err = db.QueryRow("SELECT isblocked FROM disabledcommands WHERE commandname ='" + args + "' AND guildid ='" + ctx.Guild().ID + "'").Scan(&tag.isblocked)
			if err == nil {
				if tag.isblocked == "True" {
					delete, err := db.Query("DELETE FROM disabledcommands WHERE guildid ='" + ctx.Guild().ID + "' AND commandname ='" + args + "'")

					if err != nil {
						_, _ = session.ChannelMessageSend(ctx.Channel().ID, "Bir hata oluştu.")
						return nil
					}
					_, _ = session.ChannelMessageSend(ctx.Channel().ID, args+" komutu tekrar kullanılabilir!")

					defer delete.Close()
					return nil
				}
				_, _ = session.ChannelMessageSend(ctx.Channel().ID, "Bu komut engellenmemiş.")
				return nil
			} else {
				_, _ = session.ChannelMessageSend(ctx.Channel().ID, "Bu komut engellenmemiş.")

				return nil
			}
		} else if args == "start_vote" {
			err = db.QueryRow("SELECT isblocked FROM disabledcommands WHERE commandname ='" + args + "' AND guildid ='" + ctx.Guild().ID + "'").Scan(&tag.isblocked)
			if err == nil {
				if tag.isblocked == "True" {
					delete, err := db.Query("DELETE FROM disabledcommands WHERE guildid ='" + ctx.Guild().ID + "' AND commandname ='" + args + "'")

					if err != nil {
						_, _ = session.ChannelMessageSend(ctx.Channel().ID, "Bir hata oluştu.")
						return nil
					}
					_, _ = session.ChannelMessageSend(ctx.Channel().ID, args+" komutu tekrar kullanılabilir!")

					defer delete.Close()
					return nil
				}
				_, _ = session.ChannelMessageSend(ctx.Channel().ID, "Bu komut engellenmemiş.")
				return nil
			} else {
				_, _ = session.ChannelMessageSend(ctx.Channel().ID, "Bu komut engellenmemiş.")

				return nil
			}
		} else if args == "stats" {
			err = db.QueryRow("SELECT isblocked FROM disabledcommands WHERE commandname ='" + args + "' AND guildid ='" + ctx.Guild().ID + "'").Scan(&tag.isblocked)
			if err == nil {
				if tag.isblocked == "True" {
					delete, err := db.Query("DELETE FROM disabledcommands WHERE guildid ='" + ctx.Guild().ID + "' AND commandname ='" + args + "'")

					if err != nil {
						_, _ = session.ChannelMessageSend(ctx.Channel().ID, "Bir hata oluştu.")
						return nil
					}
					_, _ = session.ChannelMessageSend(ctx.Channel().ID, args+" komutu tekrar kullanılabilir!")

					defer delete.Close()
					return nil
				}
				_, _ = session.ChannelMessageSend(ctx.Channel().ID, "Bu komut engellenmemiş.")
				return nil
			} else {
				_, _ = session.ChannelMessageSend(ctx.Channel().ID, "Bu komut engellenmemiş.")

				return nil
			}
		} else if args == "unban" {
			err = db.QueryRow("SELECT isblocked FROM disabledcommands WHERE commandname ='" + args + "' AND guildid ='" + ctx.Guild().ID + "'").Scan(&tag.isblocked)
			if err == nil {
				if tag.isblocked == "True" {
					delete, err := db.Query("DELETE FROM disabledcommands WHERE guildid ='" + ctx.Guild().ID + "' AND commandname ='" + args + "'")

					if err != nil {
						_, _ = session.ChannelMessageSend(ctx.Channel().ID, "Bir hata oluştu.")
						return nil
					}
					_, _ = session.ChannelMessageSend(ctx.Channel().ID, args+" komutu tekrar kullanılabilir!")

					defer delete.Close()
					return nil
				}
				_, _ = session.ChannelMessageSend(ctx.Channel().ID, "Bu komut engellenmemiş.")
				return nil
			} else {
				_, _ = session.ChannelMessageSend(ctx.Channel().ID, "Bu komut engellenmemiş.")

				return nil
			}
		} else {
			_, _ = session.ChannelMessageSend(ctx.Channel().ID, "Bir komut belirtmelisin.")

		}
		return nil
	}

	perms, err := session.State.UserChannelPermissions(ctx.Author().ID, ctx.Channel().ID)
	if err == nil && !(int(perms)&discordgo.PermissionAdministrator == discordgo.PermissionAdministrator) {
		_, _ = session.ChannelMessageSend(ctx.Channel().ID, "You need administrator permission to run this command.")

		return nil
	}

	if len(strings.Join(ctx.Args(), " ")) < 1 {
		_, _ = session.ChannelMessageSend(ctx.Channel().ID, "You need to specify a command.")

		return nil
	}

	var args string

	if strings.Join(ctx.Args(), " ") == "" {
		_, _ = session.ChannelMessageSend(ctx.Channel().ID, "You need to specify the command.")

		return nil
	}

	args = ctx.Args()[0]

	if args == "afk" {
		err = db.QueryRow("SELECT isblocked FROM disabledcommands WHERE commandname ='" + args + "' AND guildid ='" + ctx.Guild().ID + "'").Scan(&tag.isblocked)
		if err == nil {
			if tag.isblocked == "True" {
				delete, err := db.Query("DELETE FROM disabledcommands WHERE guildid ='" + ctx.Guild().ID + "' AND commandname ='" + args + "'")

				if err != nil {
					_, _ = session.ChannelMessageSend(ctx.Channel().ID, "An error occurred.")
					return nil
				}
				_, _ = session.ChannelMessageSend(ctx.Channel().ID, "Enabled "+args)

				defer delete.Close()
				return nil
			}
			_, _ = session.ChannelMessageSend(ctx.Channel().ID, "This command is not disabled.")
			return nil
		} else {
			_, _ = session.ChannelMessageSend(ctx.Channel().ID, "This command is not disabled.")

			return nil
		}
	} else if args == "author" {
		err = db.QueryRow("SELECT isblocked FROM disabledcommands WHERE commandname ='" + args + "' AND guildid ='" + ctx.Guild().ID + "'").Scan(&tag.isblocked)
		if err == nil {
			if tag.isblocked == "True" {
				delete, err := db.Query("DELETE FROM disabledcommands WHERE guildid ='" + ctx.Guild().ID + "' AND commandname ='" + args + "'")

				if err != nil {
					_, _ = session.ChannelMessageSend(ctx.Channel().ID, "An error occurred.")
					return nil
				}
				_, _ = session.ChannelMessageSend(ctx.Channel().ID, "Enabled "+args)

				defer delete.Close()

			}
			_, _ = session.ChannelMessageSend(ctx.Channel().ID, "This command is not disabled.")
			return nil
		} else {
			_, _ = session.ChannelMessageSend(ctx.Channel().ID, "This command is not disabled.")

			return nil
		}
	} else if args == "avatar" {
		err = db.QueryRow("SELECT isblocked FROM disabledcommands WHERE commandname ='" + args + "' AND guildid ='" + ctx.Guild().ID + "'").Scan(&tag.isblocked)
		if err == nil {
			if tag.isblocked == "True" {
				delete, err := db.Query("DELETE FROM disabledcommands WHERE guildid ='" + ctx.Guild().ID + "' AND commandname ='" + args + "'")

				if err != nil {
					_, _ = session.ChannelMessageSend(ctx.Channel().ID, "An error occurred.")
					return nil
				}
				_, _ = session.ChannelMessageSend(ctx.Channel().ID, "Enabled "+args)

				defer delete.Close()

			}
			_, _ = session.ChannelMessageSend(ctx.Channel().ID, "This command is not disabled.")
			return nil
		} else {
			_, _ = session.ChannelMessageSend(ctx.Channel().ID, "This command is not disabled.")

			return nil
		}
	} else if args == "ban" {
		err = db.QueryRow("SELECT isblocked FROM disabledcommands WHERE commandname ='" + args + "' AND guildid ='" + ctx.Guild().ID + "'").Scan(&tag.isblocked)
		if err == nil {
			if tag.isblocked == "True" {
				delete, err := db.Query("DELETE FROM disabledcommands WHERE guildid ='" + ctx.Guild().ID + "' AND commandname ='" + args + "'")

				if err != nil {
					_, _ = session.ChannelMessageSend(ctx.Channel().ID, "An error occurred.")
					return nil
				}
				_, _ = session.ChannelMessageSend(ctx.Channel().ID, "Enabled "+args)

				defer delete.Close()

			}
			_, _ = session.ChannelMessageSend(ctx.Channel().ID, "This command is not disabled.")
			return nil
		} else {
			_, _ = session.ChannelMessageSend(ctx.Channel().ID, "This command is not disabled.")

			return nil
		}
	} else if args == "embed" {
		err = db.QueryRow("SELECT isblocked FROM disabledcommands WHERE commandname ='" + args + "' AND guildid ='" + ctx.Guild().ID + "'").Scan(&tag.isblocked)
		if err == nil {
			if tag.isblocked == "True" {
				delete, err := db.Query("DELETE FROM disabledcommands WHERE guildid ='" + ctx.Guild().ID + "' AND commandname ='" + args + "'")

				if err != nil {
					_, _ = session.ChannelMessageSend(ctx.Channel().ID, "An error occurred.")
					return nil
				}
				_, _ = session.ChannelMessageSend(ctx.Channel().ID, "Enabled "+args)

				defer delete.Close()

			}
			_, _ = session.ChannelMessageSend(ctx.Channel().ID, "This command is not disabled.")
			return nil
		} else {
			_, _ = session.ChannelMessageSend(ctx.Channel().ID, "This command is not disabled.")

			return nil
		}
	} else if args == "Channel_info" {
		err = db.QueryRow("SELECT isblocked FROM disabledcommands WHERE commandname ='" + args + "' AND guildid ='" + ctx.Guild().ID + "'").Scan(&tag.isblocked)
		if err == nil {
			if tag.isblocked == "True" {
				delete, err := db.Query("DELETE FROM disabledcommands WHERE guildid ='" + ctx.Guild().ID + "' AND commandname ='" + args + "'")

				if err != nil {
					_, _ = session.ChannelMessageSend(ctx.Channel().ID, "An error occurred.")
					return nil
				}
				_, _ = session.ChannelMessageSend(ctx.Channel().ID, "Enabled "+args)

				defer delete.Close()

			}
			_, _ = session.ChannelMessageSend(ctx.Channel().ID, "This command is not disabled.")
			return nil
		} else {
			_, _ = session.ChannelMessageSend(ctx.Channel().ID, "This command is not disabled.")

			return nil
		}
	} else if args == "hug" {
		err = db.QueryRow("SELECT isblocked FROM disabledcommands WHERE commandname ='" + args + "' AND guildid ='" + ctx.Guild().ID + "'").Scan(&tag.isblocked)
		if err == nil {
			if tag.isblocked == "True" {
				delete, err := db.Query("DELETE FROM disabledcommands WHERE guildid ='" + ctx.Guild().ID + "' AND commandname ='" + args + "'")

				if err != nil {
					_, _ = session.ChannelMessageSend(ctx.Channel().ID, "An error occurred.")
					return nil
				}
				_, _ = session.ChannelMessageSend(ctx.Channel().ID, "Enabled "+args)

				defer delete.Close()

			}
			_, _ = session.ChannelMessageSend(ctx.Channel().ID, "This command is not disabled.")
			return nil
		} else {
			_, _ = session.ChannelMessageSend(ctx.Channel().ID, "This command is not disabled.")

			return nil
		}
	} else if args == "kick" {
		err = db.QueryRow("SELECT isblocked FROM disabledcommands WHERE commandname ='" + args + "' AND guildid ='" + ctx.Guild().ID + "'").Scan(&tag.isblocked)
		if err == nil {
			if tag.isblocked == "True" {
				delete, err := db.Query("DELETE FROM disabledcommands WHERE guildid ='" + ctx.Guild().ID + "' AND commandname ='" + args + "'")

				if err != nil {
					_, _ = session.ChannelMessageSend(ctx.Channel().ID, "An error occurred.")
					return nil
				}
				_, _ = session.ChannelMessageSend(ctx.Channel().ID, "Enabled "+args)

				defer delete.Close()

			}
			_, _ = session.ChannelMessageSend(ctx.Channel().ID, "This command is not disabled.")
			return nil
		} else {
			_, _ = session.ChannelMessageSend(ctx.Channel().ID, "This command is not disabled.")

			return nil
		}
	} else if args == "kiss" {
		err = db.QueryRow("SELECT isblocked FROM disabledcommands WHERE commandname ='" + args + "' AND guildid ='" + ctx.Guild().ID + "'").Scan(&tag.isblocked)
		if err == nil {
			if tag.isblocked == "True" {
				delete, err := db.Query("DELETE FROM disabledcommands WHERE guildid ='" + ctx.Guild().ID + "' AND commandname ='" + args + "'")

				if err != nil {
					_, _ = session.ChannelMessageSend(ctx.Channel().ID, "An error occurred.")
					return nil
				}
				_, _ = session.ChannelMessageSend(ctx.Channel().ID, "Enabled "+args)

				defer delete.Close()

			}
			_, _ = session.ChannelMessageSend(ctx.Channel().ID, "This command is not disabled.")
			return nil
		} else {
			_, _ = session.ChannelMessageSend(ctx.Channel().ID, "This command is not disabled.")

			return nil
		}
	} else if args == "ping" {
		err = db.QueryRow("SELECT isblocked FROM disabledcommands WHERE commandname ='" + args + "' AND guildid ='" + ctx.Guild().ID + "'").Scan(&tag.isblocked)
		if err == nil {
			if tag.isblocked == "True" {
				delete, err := db.Query("DELETE FROM disabledcommands WHERE guildid ='" + ctx.Guild().ID + "' AND commandname ='" + args + "'")

				if err != nil {
					_, _ = session.ChannelMessageSend(ctx.Channel().ID, "An error occurred.")
					return nil
				}
				_, _ = session.ChannelMessageSend(ctx.Channel().ID, "Enabled "+args)

				defer delete.Close()

			}
			_, _ = session.ChannelMessageSend(ctx.Channel().ID, "This command is not disabled.")
			return nil
		} else {
			_, _ = session.ChannelMessageSend(ctx.Channel().ID, "This command is not disabled.")

			return nil
		}
	} else if args == "roles" {
		err = db.QueryRow("SELECT isblocked FROM disabledcommands WHERE commandname ='" + args + "' AND guildid ='" + ctx.Guild().ID + "'").Scan(&tag.isblocked)
		if err == nil {
			if tag.isblocked == "True" {
				delete, err := db.Query("DELETE FROM disabledcommands WHERE guildid ='" + ctx.Guild().ID + "' AND commandname ='" + args + "'")

				if err != nil {
					_, _ = session.ChannelMessageSend(ctx.Channel().ID, "An error occurred.")
					return nil
				}
				_, _ = session.ChannelMessageSend(ctx.Channel().ID, "Enabled "+args)

				defer delete.Close()

			}
			_, _ = session.ChannelMessageSend(ctx.Channel().ID, "This command is not disabled.")
			return nil
		} else {
			_, _ = session.ChannelMessageSend(ctx.Channel().ID, "This command is not disabled.")

			return nil
		}
	} else if args == "settings" {
		err = db.QueryRow("SELECT isblocked FROM disabledcommands WHERE commandname ='" + args + "' AND guildid ='" + ctx.Guild().ID + "'").Scan(&tag.isblocked)
		if err == nil {
			if tag.isblocked == "True" {
				delete, err := db.Query("DELETE FROM disabledcommands WHERE guildid ='" + ctx.Guild().ID + "' AND commandname ='" + args + "'")

				if err != nil {
					_, _ = session.ChannelMessageSend(ctx.Channel().ID, "An error occurred.")
					return nil
				}
				_, _ = session.ChannelMessageSend(ctx.Channel().ID, "Enabled "+args)

				defer delete.Close()

			}
			_, _ = session.ChannelMessageSend(ctx.Channel().ID, "This command is not disabled.")
			return nil
		} else {
			_, _ = session.ChannelMessageSend(ctx.Channel().ID, "This command is not disabled.")

			return nil
		}
	} else if args == "slap" {
		err = db.QueryRow("SELECT isblocked FROM disabledcommands WHERE commandname ='" + args + "' AND guildid ='" + ctx.Guild().ID + "'").Scan(&tag.isblocked)
		if err == nil {
			if tag.isblocked == "True" {
				delete, err := db.Query("DELETE FROM disabledcommands WHERE guildid ='" + ctx.Guild().ID + "' AND commandname ='" + args + "'")

				if err != nil {
					_, _ = session.ChannelMessageSend(ctx.Channel().ID, "An error occurred.")
					return nil
				}
				_, _ = session.ChannelMessageSend(ctx.Channel().ID, "Enabled "+args)

				defer delete.Close()

			}
			_, _ = session.ChannelMessageSend(ctx.Channel().ID, "This command is not disabled.")
			return nil
		} else {
			_, _ = session.ChannelMessageSend(ctx.Channel().ID, "This command is not disabled.")

			return nil
		}
	} else if args == "spoiler" {
		err = db.QueryRow("SELECT isblocked FROM disabledcommands WHERE commandname ='" + args + "' AND guildid ='" + ctx.Guild().ID + "'").Scan(&tag.isblocked)
		if err == nil {
			if tag.isblocked == "True" {
				delete, err := db.Query("DELETE FROM disabledcommands WHERE guildid ='" + ctx.Guild().ID + "' AND commandname ='" + args + "'")

				if err != nil {
					_, _ = session.ChannelMessageSend(ctx.Channel().ID, "An error occurred.")
					return nil
				}
				_, _ = session.ChannelMessageSend(ctx.Channel().ID, "Enabled "+args)

				defer delete.Close()

			}
			_, _ = session.ChannelMessageSend(ctx.Channel().ID, "This command is not disabled.")
			return nil
		} else {
			_, _ = session.ChannelMessageSend(ctx.Channel().ID, "This command is not disabled.")

			return nil
		}
	} else if args == "start_vote" {
		err = db.QueryRow("SELECT isblocked FROM disabledcommands WHERE commandname ='" + args + "' AND guildid ='" + ctx.Guild().ID + "'").Scan(&tag.isblocked)
		if err == nil {
			if tag.isblocked == "True" {
				delete, err := db.Query("DELETE FROM disabledcommands WHERE guildid ='" + ctx.Guild().ID + "' AND commandname ='" + args + "'")

				if err != nil {
					_, _ = session.ChannelMessageSend(ctx.Channel().ID, "An error occurred.")
					return nil
				}
				_, _ = session.ChannelMessageSend(ctx.Channel().ID, "Enabled "+args)

				defer delete.Close()

			}
			_, _ = session.ChannelMessageSend(ctx.Channel().ID, "This command is not disabled.")
			return nil
		} else {
			_, _ = session.ChannelMessageSend(ctx.Channel().ID, "This command is not disabled.")

			return nil
		}
	} else if args == "stats" {
		err = db.QueryRow("SELECT isblocked FROM disabledcommands WHERE commandname ='" + args + "' AND guildid ='" + ctx.Guild().ID + "'").Scan(&tag.isblocked)
		if err == nil {
			if tag.isblocked == "True" {
				delete, err := db.Query("DELETE FROM disabledcommands WHERE guildid ='" + ctx.Guild().ID + "' AND commandname ='" + args + "'")

				if err != nil {
					_, _ = session.ChannelMessageSend(ctx.Channel().ID, "An error occurred.")
					return nil
				}
				_, _ = session.ChannelMessageSend(ctx.Channel().ID, "Enabled "+args)

				defer delete.Close()

			}
			_, _ = session.ChannelMessageSend(ctx.Channel().ID, "This command is not disabled.")
			return nil
		} else {
			_, _ = session.ChannelMessageSend(ctx.Channel().ID, "This command is not disabled.")

			return nil
		}
	} else if args == "unban" {
		err = db.QueryRow("SELECT isblocked FROM disabledcommands WHERE commandname ='" + args + "' AND guildid ='" + ctx.Guild().ID + "'").Scan(&tag.isblocked)
		if err == nil {
			if tag.isblocked == "True" {
				delete, err := db.Query("DELETE FROM disabledcommands WHERE guildid ='" + ctx.Guild().ID + "' AND commandname ='" + args + "'")

				if err != nil {
					_, _ = session.ChannelMessageSend(ctx.Channel().ID, "An error occurred.")
					return nil
				}
				_, _ = session.ChannelMessageSend(ctx.Channel().ID, "Enabled "+args)

				defer delete.Close()

			}
			_, _ = session.ChannelMessageSend(ctx.Channel().ID, "This command is not disabled.")
			return nil
		} else {
			_, _ = session.ChannelMessageSend(ctx.Channel().ID, "This command is not disabled.")

			return nil
		}
	} else {
		_, _ = session.ChannelMessageSend(ctx.Channel().ID, "You need to specify the command.")
		return nil
	}
}
