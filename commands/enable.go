package commands

import (
	"fmt"
	"strings"

	multiplexer "git.randomchars.net/Reviath/RemiliaScarlet/Multiplexer"
	CommandHandler "git.randomchars.net/Reviath/RemiliaScarlet/handler"
	"git.randomchars.net/Reviath/RemiliaScarlet/sql"
)

func EnableCommand(ctx CommandHandler.Context, _ []string) error {
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
		} else if args == "author" {
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
		} else if args == "avatar" {
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
		} else if args == "ban" {
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
		} else if args == "embed" {
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
		} else if args == "Channel_info" {
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
		} else if args == "hug" {
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
		} else if args == "kick" {
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
		} else if args == "kiss" {
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
		} else if args == "ping" {
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
		} else if args == "roles" {
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
		} else if args == "settings" {
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
		} else if args == "slap" {
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
		} else if args == "spoiler" {
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
		} else if args == "start_vote" {
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
		} else if args == "stats" {
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
		} else if args == "unban" {
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
		} else {
			ctx.Reply("Bir komut belirtmelisin.")

		}
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

	if args == "afk" {
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
	} else if args == "author" {
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
	} else if args == "avatar" {
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
	} else if args == "ban" {
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
	} else if args == "embed" {
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
	} else if args == "Channel_info" {
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
	} else if args == "hug" {
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
	} else if args == "kick" {
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
	} else if args == "kiss" {
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
	} else if args == "ping" {
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
	} else if args == "roles" {
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
	} else if args == "settings" {
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
	} else if args == "slap" {
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
	} else if args == "spoiler" {
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
	} else if args == "start_vote" {
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
	} else if args == "stats" {
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
	} else if args == "unban" {
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
	} else {
		ctx.Reply("You need to specify the command.")
		return nil
	}

	return db.Close()
}
