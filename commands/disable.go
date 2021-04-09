package commands

import (
	"database/sql"
	"strings"

	ctx "git.randomchars.net/Reviath/RemiliaScarlet/Context"
	"github.com/bwmarrin/discordgo"
	_ "github.com/go-sql-driver/mysql"
)

type Disable struct {
}

func (d Disable) Execute(ctx ctx.Ctx, session *discordgo.Session) error {

	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/remilia")

	if err != nil {
		panic(err.Error())
	}

	type Tag struct {
		isblocked string `json:"isblocked"`
		lang      string `json:"language"`
	}

	var tag Tag

	err = db.QueryRow("SELECT language FROM languages WHERE guildid ='" + ctx.Guild().ID + "'").Scan(&tag.lang)
	if err == nil {
		if tag.lang == "tr" {
			perms, err := session.State.UserChannelPermissions(ctx.Author().ID, ctx.Channel().ID)
			if err == nil && (int(perms)&discordgo.PermissionAdministrator == discordgo.PermissionAdministrator) == false {
				_, err := session.ChannelMessageSend(ctx.Channel().ID, "Bu komutu kullanmak için yönetici yetkisine sahip olmalısın.")
				if err != nil {
					return nil
				}
				return err
			}

			if len(strings.Join(ctx.Args(), " ")) < 1 {
				_, err := session.ChannelMessageSend(ctx.Channel().ID, "Bir komut belirtmelisin.")

				if err != nil {
					return nil
				}

				return err
			}

			var args string
			args = ctx.Args()[0]

			db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/remilia")

			if err != nil {
				panic(err.Error())
			}

			if args == "afk" {
				err = db.QueryRow("SELECT isblocked FROM disabledcommands WHERE commandname ='afk' AND guildid ='" + ctx.Guild().ID + "'").Scan(&tag.isblocked)
				if err == nil {
					if tag.isblocked == "True" {
						_, err := session.ChannelMessageSend(ctx.Channel().ID, "Bu komut zaten engellenmiş.")

						if err != nil {
							return nil
						}

						return err
					}
				} else {
					insert, err := db.Query("INSERT INTO disabledcommands (isblocked, commandname, guildid) VALUES ('True', 'afk', '" + ctx.Guild().ID + "')")

					if err != nil {
						return nil
					}

					defer insert.Close()

					if err != nil {
						_, err = session.ChannelMessageSend(ctx.Channel().ID, "Bir hata oluştu.")

						if err != nil {
							return nil
						}

						return err
					}
					_, err = session.ChannelMessageSend(ctx.Channel().ID, "Belirtilen komut engellendi.")

					if err != nil {
						return nil
					}

					return err
				}
			} else if args == "author" {
				err = db.QueryRow("SELECT isblocked FROM disabledcommands WHERE commandname ='author' AND guildid ='" + ctx.Guild().ID + "'").Scan(&tag.isblocked)
				if err == nil {
					if tag.isblocked == "True" {
						_, err := session.ChannelMessageSend(ctx.Channel().ID, "Bu komut zaten engellenmiş.")

						if err != nil {
							return nil
						}

						return err
					}
				} else {
					insert, err := db.Query("INSERT INTO disabledcommands (isblocked, commandname, guildid) VALUES ('True', 'author', '" + ctx.Guild().ID + "')")

					if err != nil {
						return nil
					}

					defer insert.Close()

					if err != nil {
						_, err = session.ChannelMessageSend(ctx.Channel().ID, "Bir hata oluştu.")

						if err != nil {
							return nil
						}

						return err
					}
					_, err = session.ChannelMessageSend(ctx.Channel().ID, "Belirtilen komut engellendi.")

					if err != nil {
						return nil
					}

					return err
				}
			} else if args == "avatar" {
				err = db.QueryRow("SELECT isblocked FROM disabledcommands WHERE commandname ='avatar' AND guildid ='" + ctx.Guild().ID + "'").Scan(&tag.isblocked)
				if err == nil {
					if tag.isblocked == "True" {
						_, err := session.ChannelMessageSend(ctx.Channel().ID, "Bu komut zaten engellenmiş.")

						if err != nil {
							return nil
						}

						return err
					}
				} else {
					insert, err := db.Query("INSERT INTO disabledcommands (isblocked, commandname, guildid) VALUES ('True', 'avatar', '" + ctx.Guild().ID + "')")

					if err != nil {
						return nil
					}

					defer insert.Close()

					if err != nil {
						_, err = session.ChannelMessageSend(ctx.Channel().ID, "Bir hata oluştu.")

						if err != nil {
							return nil
						}

						return err
					}
					_, err = session.ChannelMessageSend(ctx.Channel().ID, "Belirtilen komut engellendi.")

					if err != nil {
						return nil
					}

					return err
				}
			} else if args == "ban" {
				err = db.QueryRow("SELECT isblocked FROM disabledcommands WHERE commandname ='ban' AND guildid ='" + ctx.Guild().ID + "'").Scan(&tag.isblocked)
				if err == nil {
					if tag.isblocked == "True" {
						_, err := session.ChannelMessageSend(ctx.Channel().ID, "Bu komut zaten engellenmiş.")

						if err != nil {
							return nil
						}

						return err
					}
				} else {
					insert, err := db.Query("INSERT INTO disabledcommands (isblocked, commandname, guildid) VALUES ('True', 'ban', '" + ctx.Guild().ID + "')")

					if err != nil {
						return nil
					}

					defer insert.Close()

					if err != nil {
						_, err = session.ChannelMessageSend(ctx.Channel().ID, "Bir hata oluştu.")

						if err != nil {
							return nil
						}

						return err
					}
					_, err = session.ChannelMessageSend(ctx.Channel().ID, "Belirtilen komut engellendi.")

					if err != nil {
						return nil
					}

					return err
				}
			} else if args == "embed" {
				err = db.QueryRow("SELECT isblocked FROM disabledcommands WHERE commandname ='embed' AND guildid ='" + ctx.Guild().ID + "'").Scan(&tag.isblocked)
				if err == nil {
					if tag.isblocked == "True" {
						_, err := session.ChannelMessageSend(ctx.Channel().ID, "Bu komut zaten engellenmiş.")

						if err != nil {
							return nil
						}

						return err
					}
				} else {
					insert, err := db.Query("INSERT INTO disabledcommands (isblocked, commandname, guildid) VALUES ('True', 'embed', '" + ctx.Guild().ID + "')")

					if err != nil {
						return nil
					}

					defer insert.Close()

					if err != nil {
						_, err = session.ChannelMessageSend(ctx.Channel().ID, "Bir hata oluştu.")

						if err != nil {
							return nil
						}

						return err
					}
					_, err = session.ChannelMessageSend(ctx.Channel().ID, "Belirtilen komut engellendi.")

					if err != nil {
						return nil
					}

					return err
				}
			} else if args == "guild_info" {
				err = db.QueryRow("SELECT isblocked FROM disabledcommands WHERE commandname ='guild_info' AND guildid ='" + ctx.Guild().ID + "'").Scan(&tag.isblocked)
				if err == nil {
					if tag.isblocked == "True" {
						_, err := session.ChannelMessageSend(ctx.Channel().ID, "Bu komut zaten engellenmiş.")

						if err != nil {
							return nil
						}

						return err
					}
				} else {
					insert, err := db.Query("INSERT INTO disabledcommands (isblocked, commandname, guildid) VALUES ('True', 'guild_info', '" + ctx.Guild().ID + "')")

					if err != nil {
						return nil
					}

					defer insert.Close()

					if err != nil {
						_, err = session.ChannelMessageSend(ctx.Channel().ID, "Bir hata oluştu.")

						if err != nil {
							return nil
						}

						return err
					}
					_, err = session.ChannelMessageSend(ctx.Channel().ID, "Belirtilen komut engellendi.")

					if err != nil {
						return nil
					}

					return err
				}
			} else if args == "hug" {
				err = db.QueryRow("SELECT isblocked FROM disabledcommands WHERE commandname ='hug' AND guildid ='" + ctx.Guild().ID + "'").Scan(&tag.isblocked)
				if err == nil {
					if tag.isblocked == "True" {
						_, err := session.ChannelMessageSend(ctx.Channel().ID, "Bu komut zaten engellenmiş.")

						if err != nil {
							return nil
						}

						return err
					}
				} else {
					insert, err := db.Query("INSERT INTO disabledcommands (isblocked, commandname, guildid) VALUES ('True', 'hug', '" + ctx.Guild().ID + "')")

					if err != nil {
						return nil
					}

					defer insert.Close()

					if err != nil {
						_, err = session.ChannelMessageSend(ctx.Channel().ID, "Bir hata oluştu.")

						if err != nil {
							return nil
						}

						return err
					}
					_, err = session.ChannelMessageSend(ctx.Channel().ID, "Belirtilen komut engellendi.")

					if err != nil {
						return nil
					}

					return err
				}
			} else if args == "icon" {
				err = db.QueryRow("SELECT isblocked FROM disabledcommands WHERE commandname ='icon' AND guildid ='" + ctx.Guild().ID + "'").Scan(&tag.isblocked)
				if err == nil {
					if tag.isblocked == "True" {
						_, err := session.ChannelMessageSend(ctx.Channel().ID, "Bu komut zaten engellenmiş.")

						if err != nil {
							return nil
						}

						return err
					}
				} else {
					insert, err := db.Query("INSERT INTO disabledcommands (isblocked, commandname, guildid) VALUES ('True', 'icon', '" + ctx.Guild().ID + "')")

					if err != nil {
						return nil
					}

					defer insert.Close()

					if err != nil {
						_, err = session.ChannelMessageSend(ctx.Channel().ID, "Bir hata oluştu.")

						if err != nil {
							return nil
						}

						return err
					}
					_, err = session.ChannelMessageSend(ctx.Channel().ID, "Belirtilen komut engellendi.")

					if err != nil {
						return nil
					}

					return err
				}
			} else if args == "kick" {
				err = db.QueryRow("SELECT isblocked FROM disabledcommands WHERE commandname ='kick' AND guildid ='" + ctx.Guild().ID + "'").Scan(&tag.isblocked)
				if err == nil {
					if tag.isblocked == "True" {
						_, err := session.ChannelMessageSend(ctx.Channel().ID, "Bu komut zaten engellenmiş.")

						if err != nil {
							return nil
						}

						return err
					}
				} else {
					insert, err := db.Query("INSERT INTO disabledcommands (isblocked, commandname, guildid) VALUES ('True', 'kick', '" + ctx.Guild().ID + "')")

					if err != nil {
						return nil
					}

					defer insert.Close()

					if err != nil {
						_, err = session.ChannelMessageSend(ctx.Channel().ID, "Bir hata oluştu.")
						return err
					}
					_, err = session.ChannelMessageSend(ctx.Channel().ID, "Belirtilen komut engellendi.")

					if err != nil {
						return nil
					}

					return err
				}
			} else if args == "kiss" {
				err = db.QueryRow("SELECT isblocked FROM disabledcommands WHERE commandname ='kiss' AND guildid ='" + ctx.Guild().ID + "'").Scan(&tag.isblocked)
				if err == nil {
					if tag.isblocked == "True" {
						_, err := session.ChannelMessageSend(ctx.Channel().ID, "Bu komut zaten engellenmiş.")

						if err != nil {
							return nil
						}

						return err
					}
				} else {
					insert, err := db.Query("INSERT INTO disabledcommands (isblocked, commandname, guildid) VALUES ('True', 'kiss', '" + ctx.Guild().ID + "')")

					if err != nil {
						return nil
					}

					defer insert.Close()

					if err != nil {
						_, err = session.ChannelMessageSend(ctx.Channel().ID, "Bir hata oluştu.")
						return err
					}
					_, err = session.ChannelMessageSend(ctx.Channel().ID, "Belirtilen komut engellendi.")

					if err != nil {
						return nil
					}

					return err
				}
			} else if args == "ping" {
				err = db.QueryRow("SELECT isblocked FROM disabledcommands WHERE commandname ='ping' AND guildid ='" + ctx.Guild().ID + "'").Scan(&tag.isblocked)
				if err == nil {
					if tag.isblocked == "True" {
						_, err := session.ChannelMessageSend(ctx.Channel().ID, "Bu komut zaten engellenmiş.")

						if err != nil {
							return nil
						}

						return err
					}
				} else {
					insert, err := db.Query("INSERT INTO disabledcommands (isblocked, commandname, guildid) VALUES ('True', 'ping', '" + ctx.Guild().ID + "')")

					if err != nil {
						return nil
					}

					defer insert.Close()

					if err != nil {
						_, err = session.ChannelMessageSend(ctx.Channel().ID, "Bir hata oluştu.")
						return err
					}
					_, err = session.ChannelMessageSend(ctx.Channel().ID, "Belirtilen komut engellendi.")

					if err != nil {
						return nil
					}

					return err
				}
			} else if args == "roles" {
				err = db.QueryRow("SELECT isblocked FROM disabledcommands WHERE commandname ='roles' AND guildid ='" + ctx.Guild().ID + "'").Scan(&tag.isblocked)
				if err == nil {
					if tag.isblocked == "True" {
						_, err := session.ChannelMessageSend(ctx.Channel().ID, "Bu komut zaten engellenmiş.")

						if err != nil {
							return nil
						}

						return err
					}
				} else {
					insert, err := db.Query("INSERT INTO disabledcommands (isblocked, commandname, guildid) VALUES ('True', 'roles', '" + ctx.Guild().ID + "')")

					if err != nil {
						return nil
					}

					defer insert.Close()

					if err != nil {
						_, err = session.ChannelMessageSend(ctx.Channel().ID, "Bir hata oluştu.")
						return err
					}
					_, err = session.ChannelMessageSend(ctx.Channel().ID, "Belirtilen komut engellendi.")

					if err != nil {
						return nil
					}

					return err
				}
			} else if args == "settings" {
				err = db.QueryRow("SELECT isblocked FROM disabledcommands WHERE commandname ='settings' AND guildid ='" + ctx.Guild().ID + "'").Scan(&tag.isblocked)
				if err == nil {
					if tag.isblocked == "True" {
						_, err := session.ChannelMessageSend(ctx.Channel().ID, "Bu komut zaten engellenmiş.")

						if err != nil {
							return nil
						}

						return err
					}
				} else {
					insert, err := db.Query("INSERT INTO disabledcommands (isblocked, commandname, guildid) VALUES ('True', 'settings', '" + ctx.Guild().ID + "')")

					if err != nil {
						return nil
					}

					defer insert.Close()

					if err != nil {
						_, err = session.ChannelMessageSend(ctx.Channel().ID, "Bir hata oluştu.")
						return err
					}
					_, err = session.ChannelMessageSend(ctx.Channel().ID, "Belirtilen komut engellendi.")

					if err != nil {
						return nil
					}

					return err
				}
			} else if args == "slap" {
				err = db.QueryRow("SELECT isblocked FROM disabledcommands WHERE commandname ='slap' AND guildid ='" + ctx.Guild().ID + "'").Scan(&tag.isblocked)
				if err == nil {
					if tag.isblocked == "True" {
						_, err := session.ChannelMessageSend(ctx.Channel().ID, "Bu komut zaten engellenmiş.")

						if err != nil {
							return nil
						}

						return err
					}
				} else {
					insert, err := db.Query("INSERT INTO disabledcommands (isblocked, commandname, guildid) VALUES ('True', 'slap', '" + ctx.Guild().ID + "')")

					if err != nil {
						return nil
					}

					defer insert.Close()

					if err != nil {
						_, err = session.ChannelMessageSend(ctx.Channel().ID, "Bir hata oluştu.")
						return err
					}
					_, err = session.ChannelMessageSend(ctx.Channel().ID, "Belirtilen komut engellendi.")

					if err != nil {
						return nil
					}

					return err
				}
			} else if args == "spoiler" {
				err = db.QueryRow("SELECT isblocked FROM disabledcommands WHERE commandname ='spoiler' AND guildid ='" + ctx.Guild().ID + "'").Scan(&tag.isblocked)
				if err == nil {
					if tag.isblocked == "True" {
						_, err := session.ChannelMessageSend(ctx.Channel().ID, "Bu komut zaten engellenmiş.")

						if err != nil {
							return nil
						}

						return err
					}
				} else {
					insert, err := db.Query("INSERT INTO disabledcommands (isblocked, commandname, guildid) VALUES ('True', 'spoiler', '" + ctx.Guild().ID + "')")

					if err != nil {
						return nil
					}

					defer insert.Close()

					if err != nil {
						_, err = session.ChannelMessageSend(ctx.Channel().ID, "Bir hata oluştu.")
						return err
					}
					_, err = session.ChannelMessageSend(ctx.Channel().ID, "Belirtilen komut engellendi.")

					if err != nil {
						return nil
					}

					return err
				}
			} else if args == "start_vote" {
				err = db.QueryRow("SELECT isblocked FROM disabledcommands WHERE commandname ='start_vote' AND guildid ='" + ctx.Guild().ID + "'").Scan(&tag.isblocked)
				if err == nil {
					if tag.isblocked == "True" {
						_, err := session.ChannelMessageSend(ctx.Channel().ID, "Bu komut zaten engellenmiş.")

						if err != nil {
							return nil
						}

						return err
					}
				} else {
					insert, err := db.Query("INSERT INTO disabledcommands (isblocked, commandname, guildid) VALUES ('True', 'start_vote', '" + ctx.Guild().ID + "')")

					if err != nil {
						return nil
					}

					defer insert.Close()

					if err != nil {
						_, err = session.ChannelMessageSend(ctx.Channel().ID, "Bir hata oluştu.")
						return err
					}
					_, err = session.ChannelMessageSend(ctx.Channel().ID, "Belirtilen komut engellendi.")

					if err != nil {
						return nil
					}

					return err
				}
			} else if args == "stats" {
				err = db.QueryRow("SELECT isblocked FROM disabledcommands WHERE commandname ='stats' AND guildid ='" + ctx.Guild().ID + "'").Scan(&tag.isblocked)
				if err == nil {
					if tag.isblocked == "True" {
						_, err := session.ChannelMessageSend(ctx.Channel().ID, "Bu komut zaten engellenmiş.")

						if err != nil {
							return nil
						}

						return err
					}
				} else {
					insert, err := db.Query("INSERT INTO disabledcommands (isblocked, commandname, guizldid) VALUES ('True', 'stats', '" + ctx.Guild().ID + "')")

					defer insert.Close()

					if err != nil {
						_, err = session.ChannelMessageSend(ctx.Channel().ID, "Bir hata oluştu.")
						return err
					}
					_, err = session.ChannelMessageSend(ctx.Channel().ID, "Belirtilen komut engellendi.")

					if err != nil {
						return nil
					}

					return err
				}
			} else if args == "unban" {
				err = db.QueryRow("SELECT isblocked FROM disabledcommands WHERE commandname ='unban' AND guildid ='" + ctx.Guild().ID + "'").Scan(&tag.isblocked)
				if err == nil {
					if tag.isblocked == "True" {
						_, err := session.ChannelMessageSend(ctx.Channel().ID, "Bu komut zaten engellenmiş.")

						if err != nil {
							return nil
						}

						return err
					}
				} else {
					insert, err := db.Query("INSERT INTO disabledcommands (isblocked, commandname, guildid) VALUES ('True', 'unban', '" + ctx.Guild().ID + "')")

					if err != nil {
						return nil
					}

					defer insert.Close()

					if err != nil {
						_, err = session.ChannelMessageSend(ctx.Channel().ID, "Bir hata oluştu.")
						return err
					}
					_, err = session.ChannelMessageSend(ctx.Channel().ID, "Belirtilen komut engellendi.")

					if err != nil {
						return nil
					}

					return err
				}
			} else {
				_, err = session.ChannelMessageSend(ctx.Channel().ID, "You need to specify the command.")
				return err
			}
		}
	}

	perms, err := session.State.UserChannelPermissions(ctx.Author().ID, ctx.Channel().ID)
	if err == nil && (int(perms)&discordgo.PermissionAdministrator == discordgo.PermissionAdministrator) == false {
		_, err := session.ChannelMessageSend(ctx.Channel().ID, "You need administrator permission to run this command.")
		if err != nil {
			return nil
		}
		return err
	}

	if len(strings.Join(ctx.Args(), " ")) < 1 {
		_, err := session.ChannelMessageSend(ctx.Channel().ID, "You need to specify a command.")

		if err != nil {
			return nil
		}

		return err
	}

	var args string
	args = ctx.Args()[0]

	if args == "afk" {
		err = db.QueryRow("SELECT isblocked FROM disabledcommands WHERE commandname ='afk' AND guildid ='" + ctx.Guild().ID + "'").Scan(&tag.isblocked)
		if err == nil {
			if tag.isblocked == "True" {
				_, err := session.ChannelMessageSend(ctx.Channel().ID, "This command is already blocked.")

				if err != nil {
					return nil
				}

				return err
			}
		} else {
			insert, err := db.Query("INSERT INTO disabledcommands (isblocked, commandname, guildid) VALUES ('True', 'afk', '" + ctx.Guild().ID + "')")

			if err != nil {
				return nil
			}

			defer insert.Close()

			if err != nil {
				_, err = session.ChannelMessageSend(ctx.Channel().ID, "An error occurred.")

				if err != nil {
					return nil
				}

				return err
			}
			_, err = session.ChannelMessageSend(ctx.Channel().ID, "Disabled afk command.")

			if err != nil {
				return nil
			}

			return err
		}
	} else if args == "author" {
		err = db.QueryRow("SELECT isblocked FROM disabledcommands WHERE commandname ='author' AND guildid ='" + ctx.Guild().ID + "'").Scan(&tag.isblocked)
		if err == nil {
			if tag.isblocked == "True" {
				_, err := session.ChannelMessageSend(ctx.Channel().ID, "This command is already blocked.")

				if err != nil {
					return nil
				}

				return err
			}
		} else {
			insert, err := db.Query("INSERT INTO disabledcommands (isblocked, commandname, guildid) VALUES ('True', 'author', '" + ctx.Guild().ID + "')")

			if err != nil {
				return nil
			}

			defer insert.Close()

			if err != nil {
				_, err = session.ChannelMessageSend(ctx.Channel().ID, "An error occurred.")

				if err != nil {
					return nil
				}

				return err
			}
			_, err = session.ChannelMessageSend(ctx.Channel().ID, "Disabled author command.")

			if err != nil {
				return nil
			}

			return err
		}
	} else if args == "avatar" {
		err = db.QueryRow("SELECT isblocked FROM disabledcommands WHERE commandname ='avatar' AND guildid ='" + ctx.Guild().ID + "'").Scan(&tag.isblocked)
		if err == nil {
			if tag.isblocked == "True" {
				_, err := session.ChannelMessageSend(ctx.Channel().ID, "This command is already blocked.")

				if err != nil {
					return nil
				}

				return err
			}
		} else {
			insert, err := db.Query("INSERT INTO disabledcommands (isblocked, commandname, guildid) VALUES ('True', 'avatar', '" + ctx.Guild().ID + "')")

			if err != nil {
				return nil
			}

			defer insert.Close()

			if err != nil {
				_, err = session.ChannelMessageSend(ctx.Channel().ID, "An error occurred.")

				if err != nil {
					return nil
				}

				return err
			}
			_, err = session.ChannelMessageSend(ctx.Channel().ID, "Disabled avatar command.")

			if err != nil {
				return nil
			}

			return err
		}
	} else if args == "ban" {
		err = db.QueryRow("SELECT isblocked FROM disabledcommands WHERE commandname ='ban' AND guildid ='" + ctx.Guild().ID + "'").Scan(&tag.isblocked)
		if err == nil {
			if tag.isblocked == "True" {
				_, err := session.ChannelMessageSend(ctx.Channel().ID, "This command is already blocked.")

				if err != nil {
					return nil
				}

				return err
			}
		} else {
			insert, err := db.Query("INSERT INTO disabledcommands (isblocked, commandname, guildid) VALUES ('True', 'ban', '" + ctx.Guild().ID + "')")

			if err != nil {
				return nil
			}

			defer insert.Close()

			if err != nil {
				_, err = session.ChannelMessageSend(ctx.Channel().ID, "An error occurred.")

				if err != nil {
					return nil
				}

				return err
			}
			_, err = session.ChannelMessageSend(ctx.Channel().ID, "Disabled ban command.")

			if err != nil {
				return nil
			}

			return err
		}
	} else if args == "embed" {
		err = db.QueryRow("SELECT isblocked FROM disabledcommands WHERE commandname ='embed' AND guildid ='" + ctx.Guild().ID + "'").Scan(&tag.isblocked)
		if err == nil {
			if tag.isblocked == "True" {
				_, err := session.ChannelMessageSend(ctx.Channel().ID, "This command is already blocked.")

				if err != nil {
					return nil
				}

				return err
			}
		} else {
			insert, err := db.Query("INSERT INTO disabledcommands (isblocked, commandname, guildid) VALUES ('True', 'embed', '" + ctx.Guild().ID + "')")

			if err != nil {
				return nil
			}

			defer insert.Close()

			if err != nil {
				_, err = session.ChannelMessageSend(ctx.Channel().ID, "An error occurred.")

				if err != nil {
					return nil
				}

				return err
			}
			_, err = session.ChannelMessageSend(ctx.Channel().ID, "Disabled embed command.")

			if err != nil {
				return nil
			}

			return err
		}
	} else if args == "guild_info" {
		err = db.QueryRow("SELECT isblocked FROM disabledcommands WHERE commandname ='guild_info' AND guildid ='" + ctx.Guild().ID + "'").Scan(&tag.isblocked)
		if err == nil {
			if tag.isblocked == "True" {
				_, err := session.ChannelMessageSend(ctx.Channel().ID, "This command is already blocked.")

				if err != nil {
					return nil
				}

				return err
			}
		} else {
			insert, err := db.Query("INSERT INTO disabledcommands (isblocked, commandname, guildid) VALUES ('True', 'guild_info', '" + ctx.Guild().ID + "')")

			if err != nil {
				return nil
			}

			defer insert.Close()

			if err != nil {
				_, err = session.ChannelMessageSend(ctx.Channel().ID, "An error occurred.")

				if err != nil {
					return nil
				}

				return err
			}
			_, err = session.ChannelMessageSend(ctx.Channel().ID, "Disabled guild_info command.")

			if err != nil {
				return nil
			}

			return err
		}
	} else if args == "hug" {
		err = db.QueryRow("SELECT isblocked FROM disabledcommands WHERE commandname ='hug' AND guildid ='" + ctx.Guild().ID + "'").Scan(&tag.isblocked)
		if err == nil {
			if tag.isblocked == "True" {
				_, err := session.ChannelMessageSend(ctx.Channel().ID, "This command is already blocked.")

				if err != nil {
					return nil
				}

				return err
			}
		} else {
			insert, err := db.Query("INSERT INTO disabledcommands (isblocked, commandname, guildid) VALUES ('True', 'hug', '" + ctx.Guild().ID + "')")

			if err != nil {
				return nil
			}

			defer insert.Close()

			if err != nil {
				_, err = session.ChannelMessageSend(ctx.Channel().ID, "An error occurred.")

				if err != nil {
					return nil
				}

				return err
			}
			_, err = session.ChannelMessageSend(ctx.Channel().ID, "Disabled hug command.")

			if err != nil {
				return nil
			}

			return err
		}
	} else if args == "icon" {
		err = db.QueryRow("SELECT isblocked FROM disabledcommands WHERE commandname ='icon' AND guildid ='" + ctx.Guild().ID + "'").Scan(&tag.isblocked)
		if err == nil {
			if tag.isblocked == "True" {
				_, err := session.ChannelMessageSend(ctx.Channel().ID, "This command is already blocked.")

				if err != nil {
					return nil
				}

				return err
			}
		} else {
			insert, err := db.Query("INSERT INTO disabledcommands (isblocked, commandname, guildid) VALUES ('True', 'icon', '" + ctx.Guild().ID + "')")

			if err != nil {
				return nil
			}

			defer insert.Close()

			if err != nil {
				_, err = session.ChannelMessageSend(ctx.Channel().ID, "An error occurred.")

				if err != nil {
					return nil
				}

				return err
			}
			_, err = session.ChannelMessageSend(ctx.Channel().ID, "Disabled icon command.")

			if err != nil {
				return nil
			}

			return err
		}
	} else if args == "kick" {
		err = db.QueryRow("SELECT isblocked FROM disabledcommands WHERE commandname ='kick' AND guildid ='" + ctx.Guild().ID + "'").Scan(&tag.isblocked)
		if err == nil {
			if tag.isblocked == "True" {
				_, err := session.ChannelMessageSend(ctx.Channel().ID, "This command is already blocked.")

				if err != nil {
					return nil
				}

				return err
			}
		} else {
			insert, err := db.Query("INSERT INTO disabledcommands (isblocked, commandname, guildid) VALUES ('True', 'kick', '" + ctx.Guild().ID + "')")

			if err != nil {
				return nil
			}

			defer insert.Close()

			if err != nil {
				_, err = session.ChannelMessageSend(ctx.Channel().ID, "An error occurred.")
				return err
			}
			_, err = session.ChannelMessageSend(ctx.Channel().ID, "Disabled kick command.")

			if err != nil {
				return nil
			}

			return err
		}
	} else if args == "kiss" {
		err = db.QueryRow("SELECT isblocked FROM disabledcommands WHERE commandname ='kiss' AND guildid ='" + ctx.Guild().ID + "'").Scan(&tag.isblocked)
		if err == nil {
			if tag.isblocked == "True" {
				_, err := session.ChannelMessageSend(ctx.Channel().ID, "This command is already blocked.")

				if err != nil {
					return nil
				}

				return err
			}
		} else {
			insert, err := db.Query("INSERT INTO disabledcommands (isblocked, commandname, guildid) VALUES ('True', 'kiss', '" + ctx.Guild().ID + "')")

			if err != nil {
				return nil
			}

			defer insert.Close()

			if err != nil {
				_, err = session.ChannelMessageSend(ctx.Channel().ID, "An error occurred.")
				return err
			}
			_, err = session.ChannelMessageSend(ctx.Channel().ID, "Disabled kiss command.")

			if err != nil {
				return nil
			}

			return err
		}
	} else if args == "ping" {
		err = db.QueryRow("SELECT isblocked FROM disabledcommands WHERE commandname ='ping' AND guildid ='" + ctx.Guild().ID + "'").Scan(&tag.isblocked)
		if err == nil {
			if tag.isblocked == "True" {
				_, err := session.ChannelMessageSend(ctx.Channel().ID, "This command is already blocked.")

				if err != nil {
					return nil
				}

				return err
			}
		} else {
			insert, err := db.Query("INSERT INTO disabledcommands (isblocked, commandname, guildid) VALUES ('True', 'ping', '" + ctx.Guild().ID + "')")

			if err != nil {
				return nil
			}

			defer insert.Close()

			if err != nil {
				_, err = session.ChannelMessageSend(ctx.Channel().ID, "An error occurred.")
				return err
			}
			_, err = session.ChannelMessageSend(ctx.Channel().ID, "Disabled ping command.")

			if err != nil {
				return nil
			}

			return err
		}
	} else if args == "roles" {
		err = db.QueryRow("SELECT isblocked FROM disabledcommands WHERE commandname ='roles' AND guildid ='" + ctx.Guild().ID + "'").Scan(&tag.isblocked)
		if err == nil {
			if tag.isblocked == "True" {
				_, err := session.ChannelMessageSend(ctx.Channel().ID, "This command is already blocked.")

				if err != nil {
					return nil
				}

				return err
			}
		} else {
			insert, err := db.Query("INSERT INTO disabledcommands (isblocked, commandname, guildid) VALUES ('True', 'roles', '" + ctx.Guild().ID + "')")

			if err != nil {
				return nil
			}

			defer insert.Close()

			if err != nil {
				_, err = session.ChannelMessageSend(ctx.Channel().ID, "An error occurred.")
				return err
			}
			_, err = session.ChannelMessageSend(ctx.Channel().ID, "Disabled roles command.")

			if err != nil {
				return nil
			}

			return err
		}
	} else if args == "settings" {
		err = db.QueryRow("SELECT isblocked FROM disabledcommands WHERE commandname ='settings' AND guildid ='" + ctx.Guild().ID + "'").Scan(&tag.isblocked)
		if err == nil {
			if tag.isblocked == "True" {
				_, err := session.ChannelMessageSend(ctx.Channel().ID, "This command is already blocked.")

				if err != nil {
					return nil
				}

				return err
			}
		} else {
			insert, err := db.Query("INSERT INTO disabledcommands (isblocked, commandname, guildid) VALUES ('True', 'settings', '" + ctx.Guild().ID + "')")

			if err != nil {
				return nil
			}

			defer insert.Close()

			if err != nil {
				_, err = session.ChannelMessageSend(ctx.Channel().ID, "An error occurred.")
				return err
			}
			_, err = session.ChannelMessageSend(ctx.Channel().ID, "Disabled settings command.")

			if err != nil {
				return nil
			}

			return err
		}
	} else if args == "slap" {
		err = db.QueryRow("SELECT isblocked FROM disabledcommands WHERE commandname ='slap' AND guildid ='" + ctx.Guild().ID + "'").Scan(&tag.isblocked)
		if err == nil {
			if tag.isblocked == "True" {
				_, err := session.ChannelMessageSend(ctx.Channel().ID, "This command is already blocked.")

				if err != nil {
					return nil
				}

				return err
			}
		} else {
			insert, err := db.Query("INSERT INTO disabledcommands (isblocked, commandname, guildid) VALUES ('True', 'slap', '" + ctx.Guild().ID + "')")

			if err != nil {
				return nil
			}

			defer insert.Close()

			if err != nil {
				_, err = session.ChannelMessageSend(ctx.Channel().ID, "An error occurred.")
				return err
			}
			_, err = session.ChannelMessageSend(ctx.Channel().ID, "Disabled slap command.")

			if err != nil {
				return nil
			}

			return err
		}
	} else if args == "spoiler" {
		err = db.QueryRow("SELECT isblocked FROM disabledcommands WHERE commandname ='spoiler' AND guildid ='" + ctx.Guild().ID + "'").Scan(&tag.isblocked)
		if err == nil {
			if tag.isblocked == "True" {
				_, err := session.ChannelMessageSend(ctx.Channel().ID, "This command is already blocked.")

				if err != nil {
					return nil
				}

				return err
			}
		} else {
			insert, err := db.Query("INSERT INTO disabledcommands (isblocked, commandname, guildid) VALUES ('True', 'spoiler', '" + ctx.Guild().ID + "')")

			if err != nil {
				return nil
			}

			defer insert.Close()

			if err != nil {
				_, err = session.ChannelMessageSend(ctx.Channel().ID, "An error occurred.")
				return err
			}
			_, err = session.ChannelMessageSend(ctx.Channel().ID, "Disabled spoiler command.")

			if err != nil {
				return nil
			}

			return err
		}
	} else if args == "start_vote" {
		err = db.QueryRow("SELECT isblocked FROM disabledcommands WHERE commandname ='start_vote' AND guildid ='" + ctx.Guild().ID + "'").Scan(&tag.isblocked)
		if err == nil {
			if tag.isblocked == "True" {
				_, err := session.ChannelMessageSend(ctx.Channel().ID, "This command is already blocked.")

				if err != nil {
					return nil
				}

				return err
			}
		} else {
			insert, err := db.Query("INSERT INTO disabledcommands (isblocked, commandname, guildid) VALUES ('True', 'start_vote', '" + ctx.Guild().ID + "')")

			if err != nil {
				return nil
			}

			defer insert.Close()

			if err != nil {
				_, err = session.ChannelMessageSend(ctx.Channel().ID, "An error occurred.")
				return err
			}
			_, err = session.ChannelMessageSend(ctx.Channel().ID, "Disabled start_vote command.")

			if err != nil {
				return nil
			}

			return err
		}
	} else if args == "stats" {
		err = db.QueryRow("SELECT isblocked FROM disabledcommands WHERE commandname ='stats' AND guildid ='" + ctx.Guild().ID + "'").Scan(&tag.isblocked)
		if err == nil {
			if tag.isblocked == "True" {
				_, err := session.ChannelMessageSend(ctx.Channel().ID, "This command is already blocked.")

				if err != nil {
					return nil
				}

				return err
			}
		} else {
			insert, err := db.Query("INSERT INTO disabledcommands (isblocked, commandname, guildid) VALUES ('True', 'stats', '" + ctx.Guild().ID + "')")

			if err != nil {
				return nil
			}

			defer insert.Close()

			if err != nil {
				_, err = session.ChannelMessageSend(ctx.Channel().ID, "An error occurred.")
				return err
			}
			_, err = session.ChannelMessageSend(ctx.Channel().ID, "Disabled stats command.")

			if err != nil {
				return nil
			}

			return err
		}
	} else if args == "unban" {
		err = db.QueryRow("SELECT isblocked FROM disabledcommands WHERE commandname ='unban' AND guildid ='" + ctx.Guild().ID + "'").Scan(&tag.isblocked)
		if err == nil {
			if tag.isblocked == "True" {
				_, err := session.ChannelMessageSend(ctx.Channel().ID, "This command is already blocked.")

				if err != nil {
					return nil
				}

				return err
			}
		} else {
			insert, err := db.Query("INSERT INTO disabledcommands (isblocked, commandname, guildid) VALUES ('True', 'unban', '" + ctx.Guild().ID + "')")

			if err != nil {
				return nil
			}

			defer insert.Close()

			if err != nil {
				_, err = session.ChannelMessageSend(ctx.Channel().ID, "An error occurred.")
				return err
			}
			_, err = session.ChannelMessageSend(ctx.Channel().ID, "Disabled unban command.")

			if err != nil {
				return nil
			}

			return err
		}
	} else {
		_, err = session.ChannelMessageSend(ctx.Channel().ID, "You need to specify the command.")
		return err
	}
	return db.Close()
}
