package commands

import (
	ctx "git.randomchars.net/Reviath/handlers/Context"
	"github.com/bwmarrin/discordgo"
    "strings"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)

type Enable struct {

}

func (e Enable) Execute(ctx ctx.Ctx, session *discordgo.Session) error {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/remilia")

    if err != nil {
        panic(err.Error())
    }

	type Tag struct {
		isblocked string `json:"isblocked"`
		lang string `json:"language"`
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

	if len(strings.Join(ctx.Args()," ")) < 1 {
		_, err := session.ChannelMessageSend(ctx.Channel().ID, "Bir komut belirtmelisin.")
		
		if err != nil {
			return nil
		}

		return err
	}

	var args string

	if strings.Join(ctx.Args(), " ") == "" {
		_, err = session.ChannelMessageSend(ctx.Channel().ID, "Bir komut belirtmelisin.")
		if err != nil {
			return nil
		}
		return err
	}

	args = ctx.Args()[0]

	if args == "afk" {	
		err = db.QueryRow("SELECT isblocked FROM disabledcommands WHERE commandname ='" + args + "' AND guildid ='" + ctx.Guild().ID + "'").Scan(&tag.isblocked)
		if err == nil {
			if tag.isblocked == "True" {
				delete, err := db.Query("DELETE FROM disabledcommands WHERE guildid ='" + ctx.Guild().ID + "' AND commandname ='" + args + "'")

				if err != nil {
					_, err = session.ChannelMessageSend(ctx.Channel().ID, "Bir hata oluştu.")
					
					if err != nil {
						return nil
					}

					return err
				} else {
					_, err = session.ChannelMessageSend(ctx.Channel().ID, args + " komutu tekrar kullanılabilir!")
				}

				defer delete.Close()

			} else {
				_, err = session.ChannelMessageSend(ctx.Channel().ID, "Bu komut engellenmemiş.")
				
				if err != nil {
					return nil
				}

				return err
			}
		} else {
			_, err = session.ChannelMessageSend(ctx.Channel().ID, "Bu komut engellenmemiş.")
			
			if err != nil {
				return nil
			}

			return err
		}
	} else if args == "author" {	
		err = db.QueryRow("SELECT isblocked FROM disabledcommands WHERE commandname ='" + args + "' AND guildid ='" + ctx.Guild().ID + "'").Scan(&tag.isblocked)
		if err == nil {
			if tag.isblocked == "True" {
				delete, err := db.Query("DELETE FROM disabledcommands WHERE guildid ='" + ctx.Guild().ID + "' AND commandname ='" + args + "'")

				if err != nil {
					_, err = session.ChannelMessageSend(ctx.Channel().ID, "Bir hata oluştu.")
					
					if err != nil {
						return nil
					}

					return err
				} else {
					_, err = session.ChannelMessageSend(ctx.Channel().ID, args + " komutu tekrar kullanılabilir!")
				}

				defer delete.Close()

			} else {
				_, err = session.ChannelMessageSend(ctx.Channel().ID, "Bu komut engellenmemiş.")
				
				if err != nil {
					return nil
				}

				return err
			}
		} else {
			_, err = session.ChannelMessageSend(ctx.Channel().ID, "Bu komut engellenmemiş.")
			
			if err != nil {
				return nil
			}

			return err
		}
	} else if args == "avatar" {	
		err = db.QueryRow("SELECT isblocked FROM disabledcommands WHERE commandname ='" + args + "' AND guildid ='" + ctx.Guild().ID + "'").Scan(&tag.isblocked)
		if err == nil {
			if tag.isblocked == "True" {
				delete, err := db.Query("DELETE FROM disabledcommands WHERE guildid ='" + ctx.Guild().ID + "' AND commandname ='" + args + "'")

				if err != nil {
					_, err = session.ChannelMessageSend(ctx.Channel().ID, "Bir hata oluştu.")
					
					if err != nil {
						return nil
					}

					return err
				} else {
					_, err = session.ChannelMessageSend(ctx.Channel().ID, args + " komutu tekrar kullanılabilir!")
				}

				defer delete.Close()

			} else {
				_, err = session.ChannelMessageSend(ctx.Channel().ID, "Bu komut engellenmemiş.")
				
				if err != nil {
					return nil
				}

				return err
			}
		} else {
			_, err = session.ChannelMessageSend(ctx.Channel().ID, "Bu komut engellenmemiş.")
			
			if err != nil {
				return nil
			}

			return err
		}
	} else if args == "ban" {	
		err = db.QueryRow("SELECT isblocked FROM disabledcommands WHERE commandname ='" + args + "' AND guildid ='" + ctx.Guild().ID + "'").Scan(&tag.isblocked)
		if err == nil {
			if tag.isblocked == "True" {
				delete, err := db.Query("DELETE FROM disabledcommands WHERE guildid ='" + ctx.Guild().ID + "' AND commandname ='" + args + "'")

				if err != nil {
					_, err = session.ChannelMessageSend(ctx.Channel().ID, "Bir hata oluştu.")
					
					if err != nil {
						return nil
					}

					return err
				} else {
					_, err = session.ChannelMessageSend(ctx.Channel().ID, args + " komutu tekrar kullanılabilir!")
				}

				defer delete.Close()

			} else {
				_, err = session.ChannelMessageSend(ctx.Channel().ID, "Bu komut engellenmemiş.")
				
				if err != nil {
					return nil
				}

				return err
			}
		} else {
			_, err = session.ChannelMessageSend(ctx.Channel().ID, "Bu komut engellenmemiş.")
			
			if err != nil {
				return nil
			}

			return err
		}
	} else if args == "embed" {	
		err = db.QueryRow("SELECT isblocked FROM disabledcommands WHERE commandname ='" + args + "' AND guildid ='" + ctx.Guild().ID + "'").Scan(&tag.isblocked)
		if err == nil {
			if tag.isblocked == "True" {
				delete, err := db.Query("DELETE FROM disabledcommands WHERE guildid ='" + ctx.Guild().ID + "' AND commandname ='" + args + "'")

				if err != nil {
					_, err = session.ChannelMessageSend(ctx.Channel().ID, "Bir hata oluştu.")
					
					if err != nil {
						return nil
					}

					return err
				} else {
					_, err = session.ChannelMessageSend(ctx.Channel().ID, args + " komutu tekrar kullanılabilir!")
				}

				defer delete.Close()

			} else {
				_, err = session.ChannelMessageSend(ctx.Channel().ID, "Bu komut engellenmemiş.")
				
				if err != nil {
					return nil
				}

				return err
			}
		} else {
			_, err = session.ChannelMessageSend(ctx.Channel().ID, "Bu komut engellenmemiş.")
			
			if err != nil {
				return nil
			}

			return err
		}
	} else if args == "Channel_info" {	
		err = db.QueryRow("SELECT isblocked FROM disabledcommands WHERE commandname ='" + args + "' AND guildid ='" + ctx.Guild().ID + "'").Scan(&tag.isblocked)
		if err == nil {
			if tag.isblocked == "True" {
				delete, err := db.Query("DELETE FROM disabledcommands WHERE guildid ='" + ctx.Guild().ID + "' AND commandname ='" + args + "'")

				if err != nil {
					_, err = session.ChannelMessageSend(ctx.Channel().ID, "Bir hata oluştu.")
					
					if err != nil {
						return nil
					}

					return err
				} else {
					_, err = session.ChannelMessageSend(ctx.Channel().ID, args + " komutu tekrar kullanılabilir!")
				}

				defer delete.Close()

			} else {
				_, err = session.ChannelMessageSend(ctx.Channel().ID, "Bu komut engellenmemiş.")
				
				if err != nil {
					return nil
				}

				return err
			}
		} else {
			_, err = session.ChannelMessageSend(ctx.Channel().ID, "Bu komut engellenmemiş.")
			
			if err != nil {
				return nil
			}

			return err
		}
	} else if args == "hug" {	
		err = db.QueryRow("SELECT isblocked FROM disabledcommands WHERE commandname ='" + args + "' AND guildid ='" + ctx.Guild().ID + "'").Scan(&tag.isblocked)
		if err == nil {
			if tag.isblocked == "True" {
				delete, err := db.Query("DELETE FROM disabledcommands WHERE guildid ='" + ctx.Guild().ID + "' AND commandname ='" + args + "'")

				if err != nil {
					_, err = session.ChannelMessageSend(ctx.Channel().ID, "Bir hata oluştu.")
					
					if err != nil {
						return nil
					}

					return err
				} else {
					_, err = session.ChannelMessageSend(ctx.Channel().ID, args + " komutu tekrar kullanılabilir!")
				}

				defer delete.Close()

			} else {
				_, err = session.ChannelMessageSend(ctx.Channel().ID, "Bu komut engellenmemiş.")
				
				if err != nil {
					return nil
				}

				return err
			}
		} else {
			_, err = session.ChannelMessageSend(ctx.Channel().ID, "Bu komut engellenmemiş.")
			
			if err != nil {
				return nil
			}

			return err
		}
	} else if args == "kick" {	
		err = db.QueryRow("SELECT isblocked FROM disabledcommands WHERE commandname ='" + args + "' AND guildid ='" + ctx.Guild().ID + "'").Scan(&tag.isblocked)
		if err == nil {
			if tag.isblocked == "True" {
				delete, err := db.Query("DELETE FROM disabledcommands WHERE guildid ='" + ctx.Guild().ID + "' AND commandname ='" + args + "'")

				if err != nil {
					_, err = session.ChannelMessageSend(ctx.Channel().ID, "Bir hata oluştu.")
					
					if err != nil {
						return nil
					}

					return err
				} else {
					_, err = session.ChannelMessageSend(ctx.Channel().ID, args + " komutu tekrar kullanılabilir!")
				}

				defer delete.Close()

			} else {
				_, err = session.ChannelMessageSend(ctx.Channel().ID, "Bu komut engellenmemiş.")
				
				if err != nil {
					return nil
				}

				return err
			}
		} else {
			_, err = session.ChannelMessageSend(ctx.Channel().ID, "Bu komut engellenmemiş.")
			
			if err != nil {
				return nil
			}

			return err
		}
	} else if args == "afk" {	
		err = db.QueryRow("SELECT isblocked FROM disabledcommands WHERE commandname ='" + args + "' AND guildid ='" + ctx.Guild().ID + "'").Scan(&tag.isblocked)
		if err == nil {
			if tag.isblocked == "True" {
				delete, err := db.Query("DELETE FROM disabledcommands WHERE guildid ='" + ctx.Guild().ID + "' AND commandname ='" + args + "'")

				if err != nil {
					_, err = session.ChannelMessageSend(ctx.Channel().ID, "Bir hata oluştu.")
					
					if err != nil {
						return nil
					}

					return err
				} else {
					_, err = session.ChannelMessageSend(ctx.Channel().ID, args + " komutu tekrar kullanılabilir!")
				}

				defer delete.Close()

			} else {
				_, err = session.ChannelMessageSend(ctx.Channel().ID, "Bu komut engellenmemiş.")
				
				if err != nil {
					return nil
				}

				return err
			}
		} else {
			_, err = session.ChannelMessageSend(ctx.Channel().ID, "Bu komut engellenmemiş.")
			
			if err != nil {
				return nil
			}

			return err
		}
	} else if args == "kiss" {	
		err = db.QueryRow("SELECT isblocked FROM disabledcommands WHERE commandname ='" + args + "' AND guildid ='" + ctx.Guild().ID + "'").Scan(&tag.isblocked)
		if err == nil {
			if tag.isblocked == "True" {
				delete, err := db.Query("DELETE FROM disabledcommands WHERE guildid ='" + ctx.Guild().ID + "' AND commandname ='" + args + "'")

				if err != nil {
					_, err = session.ChannelMessageSend(ctx.Channel().ID, "Bir hata oluştu.")
					
					if err != nil {
						return nil
					}

					return err
				} else {
					_, err = session.ChannelMessageSend(ctx.Channel().ID, args + " komutu tekrar kullanılabilir!")
				}

				defer delete.Close()

			} else {
				_, err = session.ChannelMessageSend(ctx.Channel().ID, "Bu komut engellenmemiş.")
				
				if err != nil {
					return nil
				}

				return err
			}
		} else {
			_, err = session.ChannelMessageSend(ctx.Channel().ID, "Bu komut engellenmemiş.")
			
			if err != nil {
				return nil
			}

			return err
		}
	} else if args == "ping" {	
		err = db.QueryRow("SELECT isblocked FROM disabledcommands WHERE commandname ='" + args + "' AND guildid ='" + ctx.Guild().ID + "'").Scan(&tag.isblocked)
		if err == nil {
			if tag.isblocked == "True" {
				delete, err := db.Query("DELETE FROM disabledcommands WHERE guildid ='" + ctx.Guild().ID + "' AND commandname ='" + args + "'")

				if err != nil {
					_, err = session.ChannelMessageSend(ctx.Channel().ID, "Bir hata oluştu.")
					
					if err != nil {
						return nil
					}

					return err
				} else {
					_, err = session.ChannelMessageSend(ctx.Channel().ID, args + " komutu tekrar kullanılabilir!")
				}

				defer delete.Close()

			} else {
				_, err = session.ChannelMessageSend(ctx.Channel().ID, "Bu komut engellenmemiş.")
				
				if err != nil {
					return nil
				}

				return err
			}
		} else {
			_, err = session.ChannelMessageSend(ctx.Channel().ID, "Bu komut engellenmemiş.")
			
			if err != nil {
				return nil
			}

			return err
		}
	} else if args == "roles" {	
		err = db.QueryRow("SELECT isblocked FROM disabledcommands WHERE commandname ='" + args + "' AND guildid ='" + ctx.Guild().ID + "'").Scan(&tag.isblocked)
		if err == nil {
			if tag.isblocked == "True" {
				delete, err := db.Query("DELETE FROM disabledcommands WHERE guildid ='" + ctx.Guild().ID + "' AND commandname ='" + args + "'")

				if err != nil {
					_, err = session.ChannelMessageSend(ctx.Channel().ID, "Bir hata oluştu.")
					
					if err != nil {
						return nil
					}

					return err
				} else {
					_, err = session.ChannelMessageSend(ctx.Channel().ID, args + " komutu tekrar kullanılabilir!")
				}

				defer delete.Close()

			} else {
				_, err = session.ChannelMessageSend(ctx.Channel().ID, "Bu komut engellenmemiş.")
				
				if err != nil {
					return nil
				}

				return err
			}
		} else {
			_, err = session.ChannelMessageSend(ctx.Channel().ID, "Bu komut engellenmemiş.")
			
			if err != nil {
				return nil
			}

			return err
		}
	} else if args == "settings" {	
		err = db.QueryRow("SELECT isblocked FROM disabledcommands WHERE commandname ='" + args + "' AND guildid ='" + ctx.Guild().ID + "'").Scan(&tag.isblocked)
		if err == nil {
			if tag.isblocked == "True" {
				delete, err := db.Query("DELETE FROM disabledcommands WHERE guildid ='" + ctx.Guild().ID + "' AND commandname ='" + args + "'")

				if err != nil {
					_, err = session.ChannelMessageSend(ctx.Channel().ID, "Bir hata oluştu.")
					
					if err != nil {
						return nil
					}

					return err
				} else {
					_, err = session.ChannelMessageSend(ctx.Channel().ID, args + " komutu tekrar kullanılabilir!")
				}

				defer delete.Close()

			} else {
				_, err = session.ChannelMessageSend(ctx.Channel().ID, "Bu komut engellenmemiş.")
				
				if err != nil {
					return nil
				}

				return err
			}
		} else {
			_, err = session.ChannelMessageSend(ctx.Channel().ID, "Bu komut engellenmemiş.")
			
			if err != nil {
				return nil
			}

			return err
		}
	} else if args == "slap" {	
		err = db.QueryRow("SELECT isblocked FROM disabledcommands WHERE commandname ='" + args + "' AND guildid ='" + ctx.Guild().ID + "'").Scan(&tag.isblocked)
		if err == nil {
			if tag.isblocked == "True" {
				delete, err := db.Query("DELETE FROM disabledcommands WHERE guildid ='" + ctx.Guild().ID + "' AND commandname ='" + args + "'")

				if err != nil {
					_, err = session.ChannelMessageSend(ctx.Channel().ID, "Bir hata oluştu.")
					
					if err != nil {
						return nil
					}

					return err
				} else {
					_, err = session.ChannelMessageSend(ctx.Channel().ID, args + " komutu tekrar kullanılabilir!")
				}

				defer delete.Close()

			} else {
				_, err = session.ChannelMessageSend(ctx.Channel().ID, "Bu komut engellenmemiş.")
				
				if err != nil {
					return nil
				}

				return err
			}
		} else {
			_, err = session.ChannelMessageSend(ctx.Channel().ID, "Bu komut engellenmemiş.")
			
			if err != nil {
				return nil
			}

			return err
		}
	} else if args == "spoiler" {	
		err = db.QueryRow("SELECT isblocked FROM disabledcommands WHERE commandname ='" + args + "' AND guildid ='" + ctx.Guild().ID + "'").Scan(&tag.isblocked)
		if err == nil {
			if tag.isblocked == "True" {
				delete, err := db.Query("DELETE FROM disabledcommands WHERE guildid ='" + ctx.Guild().ID + "' AND commandname ='" + args + "'")

				if err != nil {
					_, err = session.ChannelMessageSend(ctx.Channel().ID, "Bir hata oluştu.")
					
					if err != nil {
						return nil
					}

					return err
				} else {
					_, err = session.ChannelMessageSend(ctx.Channel().ID, args + " komutu tekrar kullanılabilir!")
				}

				defer delete.Close()

			} else {
				_, err = session.ChannelMessageSend(ctx.Channel().ID, "Bu komut engellenmemiş.")
				
				if err != nil {
					return nil
				}

				return err
			}
		} else {
			_, err = session.ChannelMessageSend(ctx.Channel().ID, "Bu komut engellenmemiş.")
			
			if err != nil {
				return nil
			}

			return err
		}
	} else if args == "start_vote" {	
		err = db.QueryRow("SELECT isblocked FROM disabledcommands WHERE commandname ='" + args + "' AND guildid ='" + ctx.Guild().ID + "'").Scan(&tag.isblocked)
		if err == nil {
			if tag.isblocked == "True" {
				delete, err := db.Query("DELETE FROM disabledcommands WHERE guildid ='" + ctx.Guild().ID + "' AND commandname ='" + args + "'")

				if err != nil {
					_, err = session.ChannelMessageSend(ctx.Channel().ID, "Bir hata oluştu.")
					
					if err != nil {
						return nil
					}

					return err
				} else {
					_, err = session.ChannelMessageSend(ctx.Channel().ID, args + " komutu tekrar kullanılabilir!")
				}

				defer delete.Close()

			} else {
				_, err = session.ChannelMessageSend(ctx.Channel().ID, "Bu komut engellenmemiş.")
				
				if err != nil {
					return nil
				}

				return err
			}
		} else {
			_, err = session.ChannelMessageSend(ctx.Channel().ID, "Bu komut engellenmemiş.")
			
			if err != nil {
				return nil
			}

			return err
		}
	} else if args == "stats" {	
		err = db.QueryRow("SELECT isblocked FROM disabledcommands WHERE commandname ='" + args + "' AND guildid ='" + ctx.Guild().ID + "'").Scan(&tag.isblocked)
		if err == nil {
			if tag.isblocked == "True" {
				delete, err := db.Query("DELETE FROM disabledcommands WHERE guildid ='" + ctx.Guild().ID + "' AND commandname ='" + args + "'")

				if err != nil {
					_, err = session.ChannelMessageSend(ctx.Channel().ID, "Bir hata oluştu.")
					
					if err != nil {
						return nil
					}

					return err
				} else {
					_, err = session.ChannelMessageSend(ctx.Channel().ID, args + " komutu tekrar kullanılabilir!")
				}

				defer delete.Close()

			} else {
				_, err = session.ChannelMessageSend(ctx.Channel().ID, "Bu komut engellenmemiş.")
				
				if err != nil {
					return nil
				}

				return err
			}
		} else {
			_, err = session.ChannelMessageSend(ctx.Channel().ID, "Bu komut engellenmemiş.")
			
			if err != nil {
				return nil
			}

			return err
		}
	} else if args == "unban" {	
		err = db.QueryRow("SELECT isblocked FROM disabledcommands WHERE commandname ='" + args + "' AND guildid ='" + ctx.Guild().ID + "'").Scan(&tag.isblocked)
		if err == nil {
			if tag.isblocked == "True" {
				delete, err := db.Query("DELETE FROM disabledcommands WHERE guildid ='" + ctx.Guild().ID + "' AND commandname ='" + args + "'")

				if err != nil {
					_, err = session.ChannelMessageSend(ctx.Channel().ID, "Bir hata oluştu.")
					
					if err != nil {
						return nil
					}

					return err
				} else {
					_, err = session.ChannelMessageSend(ctx.Channel().ID, args + " komutu tekrar kullanılabilir!")
				}

				defer delete.Close()

			} else {
				_, err = session.ChannelMessageSend(ctx.Channel().ID, "Bu komut engellenmemiş.")
				
				if err != nil {
					return nil
				}

				return err
			}
		} else {
			_, err = session.ChannelMessageSend(ctx.Channel().ID, "Bu komut engellenmemiş.")
			
			if err != nil {
				return nil
			}

			return err
		}
	} else {
		_, err = session.ChannelMessageSend(ctx.Channel().ID, "Bir komut belirtmelisin.")
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

	if len(strings.Join(ctx.Args()," ")) < 1 {
		_, err := session.ChannelMessageSend(ctx.Channel().ID, "You need Bir komut belirtmelisin
		if err != nil {
			return nil
		}

		return err
	}

	var args string

	if strings.Join(ctx.Args(), " ") == "" {
		_, err = session.ChannelMessageSend(ctx.Channel().ID, "You need to specify the command.")
		if err != nil {
			return nil
		}
		return err
	}

	args = ctx.Args()[0]

	if args == "afk" {	
		err = db.QueryRow("SELECT isblocked FROM disabledcommands WHERE commandname ='" + args + "' AND guildid ='" + ctx.Guild().ID + "'").Scan(&tag.isblocked)
		if err == nil {
			if tag.isblocked == "True" {
				delete, err := db.Query("DELETE FROM disabledcommands WHERE guildid ='" + ctx.Guild().ID + "' AND commandname ='" + args + "'")

				if err != nil {
					_, err = session.ChannelMessageSend(ctx.Channel().ID, "An error occured.")
					
					if err != nil {
						return nil
					}

					return err
				} else {
					_, err = session.ChannelMessageSend(ctx.Channel().ID, "Enabled " + args)
				}

				defer delete.Close()

			} else {
				_, err = session.ChannelMessageSend(ctx.Channel().ID, "This command is not disabled.")
				
				if err != nil {
					return nil
				}

				return err
			}
		} else {
			_, err = session.ChannelMessageSend(ctx.Channel().ID, "This command is not disabled.")
			
			if err != nil {
				return nil
			}

			return err
		}
	} else if args == "author" {	
		err = db.QueryRow("SELECT isblocked FROM disabledcommands WHERE commandname ='" + args + "' AND guildid ='" + ctx.Guild().ID + "'").Scan(&tag.isblocked)
		if err == nil {
			if tag.isblocked == "True" {
				delete, err := db.Query("DELETE FROM disabledcommands WHERE guildid ='" + ctx.Guild().ID + "' AND commandname ='" + args + "'")

				if err != nil {
					_, err = session.ChannelMessageSend(ctx.Channel().ID, "An error occured.")
					
					if err != nil {
						return nil
					}

					return err
				} else {
					_, err = session.ChannelMessageSend(ctx.Channel().ID, "Enabled " + args)
				}

				defer delete.Close()

			} else {
				_, err = session.ChannelMessageSend(ctx.Channel().ID, "This command is not disabled.")
				
				if err != nil {
					return nil
				}

				return err
			}
		} else {
			_, err = session.ChannelMessageSend(ctx.Channel().ID, "This command is not disabled.")
			
			if err != nil {
				return nil
			}

			return err
		}
	} else if args == "avatar" {	
		err = db.QueryRow("SELECT isblocked FROM disabledcommands WHERE commandname ='" + args + "' AND guildid ='" + ctx.Guild().ID + "'").Scan(&tag.isblocked)
		if err == nil {
			if tag.isblocked == "True" {
				delete, err := db.Query("DELETE FROM disabledcommands WHERE guildid ='" + ctx.Guild().ID + "' AND commandname ='" + args + "'")

				if err != nil {
					_, err = session.ChannelMessageSend(ctx.Channel().ID, "An error occured.")
					
					if err != nil {
						return nil
					}

					return err
				} else {
					_, err = session.ChannelMessageSend(ctx.Channel().ID, "Enabled " + args)
				}

				defer delete.Close()

			} else {
				_, err = session.ChannelMessageSend(ctx.Channel().ID, "This command is not disabled.")
				
				if err != nil {
					return nil
				}

				return err
			}
		} else {
			_, err = session.ChannelMessageSend(ctx.Channel().ID, "This command is not disabled.")
			
			if err != nil {
				return nil
			}

			return err
		}
	} else if args == "ban" {	
		err = db.QueryRow("SELECT isblocked FROM disabledcommands WHERE commandname ='" + args + "' AND guildid ='" + ctx.Guild().ID + "'").Scan(&tag.isblocked)
		if err == nil {
			if tag.isblocked == "True" {
				delete, err := db.Query("DELETE FROM disabledcommands WHERE guildid ='" + ctx.Guild().ID + "' AND commandname ='" + args + "'")

				if err != nil {
					_, err = session.ChannelMessageSend(ctx.Channel().ID, "An error occured.")
					
					if err != nil {
						return nil
					}

					return err
				} else {
					_, err = session.ChannelMessageSend(ctx.Channel().ID, "Enabled " + args)
				}

				defer delete.Close()

			} else {
				_, err = session.ChannelMessageSend(ctx.Channel().ID, "This command is not disabled.")
				
				if err != nil {
					return nil
				}

				return err
			}
		} else {
			_, err = session.ChannelMessageSend(ctx.Channel().ID, "This command is not disabled.")
			
			if err != nil {
				return nil
			}

			return err
		}
	} else if args == "embed" {	
		err = db.QueryRow("SELECT isblocked FROM disabledcommands WHERE commandname ='" + args + "' AND guildid ='" + ctx.Guild().ID + "'").Scan(&tag.isblocked)
		if err == nil {
			if tag.isblocked == "True" {
				delete, err := db.Query("DELETE FROM disabledcommands WHERE guildid ='" + ctx.Guild().ID + "' AND commandname ='" + args + "'")

				if err != nil {
					_, err = session.ChannelMessageSend(ctx.Channel().ID, "An error occured.")
					
					if err != nil {
						return nil
					}

					return err
				} else {
					_, err = session.ChannelMessageSend(ctx.Channel().ID, "Enabled " + args)
				}

				defer delete.Close()

			} else {
				_, err = session.ChannelMessageSend(ctx.Channel().ID, "This command is not disabled.")
				
				if err != nil {
					return nil
				}

				return err
			}
		} else {
			_, err = session.ChannelMessageSend(ctx.Channel().ID, "This command is not disabled.")
			
			if err != nil {
				return nil
			}

			return err
		}
	} else if args == "Channel_info" {	
		err = db.QueryRow("SELECT isblocked FROM disabledcommands WHERE commandname ='" + args + "' AND guildid ='" + ctx.Guild().ID + "'").Scan(&tag.isblocked)
		if err == nil {
			if tag.isblocked == "True" {
				delete, err := db.Query("DELETE FROM disabledcommands WHERE guildid ='" + ctx.Guild().ID + "' AND commandname ='" + args + "'")

				if err != nil {
					_, err = session.ChannelMessageSend(ctx.Channel().ID, "An error occured.")
					
					if err != nil {
						return nil
					}

					return err
				} else {
					_, err = session.ChannelMessageSend(ctx.Channel().ID, "Enabled " + args)
				}

				defer delete.Close()

			} else {
				_, err = session.ChannelMessageSend(ctx.Channel().ID, "This command is not disabled.")
				
				if err != nil {
					return nil
				}

				return err
			}
		} else {
			_, err = session.ChannelMessageSend(ctx.Channel().ID, "This command is not disabled.")
			
			if err != nil {
				return nil
			}

			return err
		}
	} else if args == "hug" {	
		err = db.QueryRow("SELECT isblocked FROM disabledcommands WHERE commandname ='" + args + "' AND guildid ='" + ctx.Guild().ID + "'").Scan(&tag.isblocked)
		if err == nil {
			if tag.isblocked == "True" {
				delete, err := db.Query("DELETE FROM disabledcommands WHERE guildid ='" + ctx.Guild().ID + "' AND commandname ='" + args + "'")

				if err != nil {
					_, err = session.ChannelMessageSend(ctx.Channel().ID, "An error occured.")
					
					if err != nil {
						return nil
					}

					return err
				} else {
					_, err = session.ChannelMessageSend(ctx.Channel().ID, "Enabled " + args)
				}

				defer delete.Close()

			} else {
				_, err = session.ChannelMessageSend(ctx.Channel().ID, "This command is not disabled.")
				
				if err != nil {
					return nil
				}

				return err
			}
		} else {
			_, err = session.ChannelMessageSend(ctx.Channel().ID, "This command is not disabled.")
			
			if err != nil {
				return nil
			}

			return err
		}
	} else if args == "kick" {	
		err = db.QueryRow("SELECT isblocked FROM disabledcommands WHERE commandname ='" + args + "' AND guildid ='" + ctx.Guild().ID + "'").Scan(&tag.isblocked)
		if err == nil {
			if tag.isblocked == "True" {
				delete, err := db.Query("DELETE FROM disabledcommands WHERE guildid ='" + ctx.Guild().ID + "' AND commandname ='" + args + "'")

				if err != nil {
					_, err = session.ChannelMessageSend(ctx.Channel().ID, "An error occured.")
					
					if err != nil {
						return nil
					}

					return err
				} else {
					_, err = session.ChannelMessageSend(ctx.Channel().ID, "Enabled " + args)
				}

				defer delete.Close()

			} else {
				_, err = session.ChannelMessageSend(ctx.Channel().ID, "This command is not disabled.")
				
				if err != nil {
					return nil
				}

				return err
			}
		} else {
			_, err = session.ChannelMessageSend(ctx.Channel().ID, "This command is not disabled.")
			
			if err != nil {
				return nil
			}

			return err
		}
	} else if args == "afk" {	
		err = db.QueryRow("SELECT isblocked FROM disabledcommands WHERE commandname ='" + args + "' AND guildid ='" + ctx.Guild().ID + "'").Scan(&tag.isblocked)
		if err == nil {
			if tag.isblocked == "True" {
				delete, err := db.Query("DELETE FROM disabledcommands WHERE guildid ='" + ctx.Guild().ID + "' AND commandname ='" + args + "'")

				if err != nil {
					_, err = session.ChannelMessageSend(ctx.Channel().ID, "An error occured.")
					
					if err != nil {
						return nil
					}

					return err
				} else {
					_, err = session.ChannelMessageSend(ctx.Channel().ID, "Enabled " + args)
				}

				defer delete.Close()

			} else {
				_, err = session.ChannelMessageSend(ctx.Channel().ID, "This command is not disabled.")
				
				if err != nil {
					return nil
				}

				return err
			}
		} else {
			_, err = session.ChannelMessageSend(ctx.Channel().ID, "This command is not disabled.")
			
			if err != nil {
				return nil
			}

			return err
		}
	} else if args == "kiss" {	
		err = db.QueryRow("SELECT isblocked FROM disabledcommands WHERE commandname ='" + args + "' AND guildid ='" + ctx.Guild().ID + "'").Scan(&tag.isblocked)
		if err == nil {
			if tag.isblocked == "True" {
				delete, err := db.Query("DELETE FROM disabledcommands WHERE guildid ='" + ctx.Guild().ID + "' AND commandname ='" + args + "'")

				if err != nil {
					_, err = session.ChannelMessageSend(ctx.Channel().ID, "An error occured.")
					
					if err != nil {
						return nil
					}

					return err
				} else {
					_, err = session.ChannelMessageSend(ctx.Channel().ID, "Enabled " + args)
				}

				defer delete.Close()

			} else {
				_, err = session.ChannelMessageSend(ctx.Channel().ID, "This command is not disabled.")
				
				if err != nil {
					return nil
				}

				return err
			}
		} else {
			_, err = session.ChannelMessageSend(ctx.Channel().ID, "This command is not disabled.")
			
			if err != nil {
				return nil
			}

			return err
		}
	} else if args == "ping" {	
		err = db.QueryRow("SELECT isblocked FROM disabledcommands WHERE commandname ='" + args + "' AND guildid ='" + ctx.Guild().ID + "'").Scan(&tag.isblocked)
		if err == nil {
			if tag.isblocked == "True" {
				delete, err := db.Query("DELETE FROM disabledcommands WHERE guildid ='" + ctx.Guild().ID + "' AND commandname ='" + args + "'")

				if err != nil {
					_, err = session.ChannelMessageSend(ctx.Channel().ID, "An error occured.")
					
					if err != nil {
						return nil
					}

					return err
				} else {
					_, err = session.ChannelMessageSend(ctx.Channel().ID, "Enabled " + args)
				}

				defer delete.Close()

			} else {
				_, err = session.ChannelMessageSend(ctx.Channel().ID, "This command is not disabled.")
				
				if err != nil {
					return nil
				}

				return err
			}
		} else {
			_, err = session.ChannelMessageSend(ctx.Channel().ID, "This command is not disabled.")
			
			if err != nil {
				return nil
			}

			return err
		}
	} else if args == "roles" {	
		err = db.QueryRow("SELECT isblocked FROM disabledcommands WHERE commandname ='" + args + "' AND guildid ='" + ctx.Guild().ID + "'").Scan(&tag.isblocked)
		if err == nil {
			if tag.isblocked == "True" {
				delete, err := db.Query("DELETE FROM disabledcommands WHERE guildid ='" + ctx.Guild().ID + "' AND commandname ='" + args + "'")

				if err != nil {
					_, err = session.ChannelMessageSend(ctx.Channel().ID, "An error occured.")
					
					if err != nil {
						return nil
					}

					return err
				} else {
					_, err = session.ChannelMessageSend(ctx.Channel().ID, "Enabled " + args)
				}

				defer delete.Close()

			} else {
				_, err = session.ChannelMessageSend(ctx.Channel().ID, "This command is not disabled.")
				
				if err != nil {
					return nil
				}

				return err
			}
		} else {
			_, err = session.ChannelMessageSend(ctx.Channel().ID, "This command is not disabled.")
			
			if err != nil {
				return nil
			}

			return err
		}
	} else if args == "settings" {	
		err = db.QueryRow("SELECT isblocked FROM disabledcommands WHERE commandname ='" + args + "' AND guildid ='" + ctx.Guild().ID + "'").Scan(&tag.isblocked)
		if err == nil {
			if tag.isblocked == "True" {
				delete, err := db.Query("DELETE FROM disabledcommands WHERE guildid ='" + ctx.Guild().ID + "' AND commandname ='" + args + "'")

				if err != nil {
					_, err = session.ChannelMessageSend(ctx.Channel().ID, "An error occured.")
					
					if err != nil {
						return nil
					}

					return err
				} else {
					_, err = session.ChannelMessageSend(ctx.Channel().ID, "Enabled " + args)
				}

				defer delete.Close()

			} else {
				_, err = session.ChannelMessageSend(ctx.Channel().ID, "This command is not disabled.")
				
				if err != nil {
					return nil
				}

				return err
			}
		} else {
			_, err = session.ChannelMessageSend(ctx.Channel().ID, "This command is not disabled.")
			
			if err != nil {
				return nil
			}

			return err
		}
	} else if args == "slap" {	
		err = db.QueryRow("SELECT isblocked FROM disabledcommands WHERE commandname ='" + args + "' AND guildid ='" + ctx.Guild().ID + "'").Scan(&tag.isblocked)
		if err == nil {
			if tag.isblocked == "True" {
				delete, err := db.Query("DELETE FROM disabledcommands WHERE guildid ='" + ctx.Guild().ID + "' AND commandname ='" + args + "'")

				if err != nil {
					_, err = session.ChannelMessageSend(ctx.Channel().ID, "An error occured.")
					
					if err != nil {
						return nil
					}

					return err
				} else {
					_, err = session.ChannelMessageSend(ctx.Channel().ID, "Enabled " + args)
				}

				defer delete.Close()

			} else {
				_, err = session.ChannelMessageSend(ctx.Channel().ID, "This command is not disabled.")
				
				if err != nil {
					return nil
				}

				return err
			}
		} else {
			_, err = session.ChannelMessageSend(ctx.Channel().ID, "This command is not disabled.")
			
			if err != nil {
				return nil
			}

			return err
		}
	} else if args == "spoiler" {	
		err = db.QueryRow("SELECT isblocked FROM disabledcommands WHERE commandname ='" + args + "' AND guildid ='" + ctx.Guild().ID + "'").Scan(&tag.isblocked)
		if err == nil {
			if tag.isblocked == "True" {
				delete, err := db.Query("DELETE FROM disabledcommands WHERE guildid ='" + ctx.Guild().ID + "' AND commandname ='" + args + "'")

				if err != nil {
					_, err = session.ChannelMessageSend(ctx.Channel().ID, "An error occured.")
					
					if err != nil {
						return nil
					}

					return err
				} else {
					_, err = session.ChannelMessageSend(ctx.Channel().ID, "Enabled " + args)
				}

				defer delete.Close()

			} else {
				_, err = session.ChannelMessageSend(ctx.Channel().ID, "This command is not disabled.")
				
				if err != nil {
					return nil
				}

				return err
			}
		} else {
			_, err = session.ChannelMessageSend(ctx.Channel().ID, "This command is not disabled.")
			
			if err != nil {
				return nil
			}

			return err
		}
	} else if args == "start_vote" {	
		err = db.QueryRow("SELECT isblocked FROM disabledcommands WHERE commandname ='" + args + "' AND guildid ='" + ctx.Guild().ID + "'").Scan(&tag.isblocked)
		if err == nil {
			if tag.isblocked == "True" {
				delete, err := db.Query("DELETE FROM disabledcommands WHERE guildid ='" + ctx.Guild().ID + "' AND commandname ='" + args + "'")

				if err != nil {
					_, err = session.ChannelMessageSend(ctx.Channel().ID, "An error occured.")
					
					if err != nil {
						return nil
					}

					return err
				} else {
					_, err = session.ChannelMessageSend(ctx.Channel().ID, "Enabled " + args)
				}

				defer delete.Close()

			} else {
				_, err = session.ChannelMessageSend(ctx.Channel().ID, "This command is not disabled.")
				
				if err != nil {
					return nil
				}

				return err
			}
		} else {
			_, err = session.ChannelMessageSend(ctx.Channel().ID, "This command is not disabled.")
			
			if err != nil {
				return nil
			}

			return err
		}
	} else if args == "stats" {	
		err = db.QueryRow("SELECT isblocked FROM disabledcommands WHERE commandname ='" + args + "' AND guildid ='" + ctx.Guild().ID + "'").Scan(&tag.isblocked)
		if err == nil {
			if tag.isblocked == "True" {
				delete, err := db.Query("DELETE FROM disabledcommands WHERE guildid ='" + ctx.Guild().ID + "' AND commandname ='" + args + "'")

				if err != nil {
					_, err = session.ChannelMessageSend(ctx.Channel().ID, "An error occured.")
					
					if err != nil {
						return nil
					}

					return err
				} else {
					_, err = session.ChannelMessageSend(ctx.Channel().ID, "Enabled " + args)
				}

				defer delete.Close()

			} else {
				_, err = session.ChannelMessageSend(ctx.Channel().ID, "This command is not disabled.")
				
				if err != nil {
					return nil
				}

				return err
			}
		} else {
			_, err = session.ChannelMessageSend(ctx.Channel().ID, "This command is not disabled.")
			
			if err != nil {
				return nil
			}

			return err
		}
	} else if args == "unban" {	
		err = db.QueryRow("SELECT isblocked FROM disabledcommands WHERE commandname ='" + args + "' AND guildid ='" + ctx.Guild().ID + "'").Scan(&tag.isblocked)
		if err == nil {
			if tag.isblocked == "True" {
				delete, err := db.Query("DELETE FROM disabledcommands WHERE guildid ='" + ctx.Guild().ID + "' AND commandname ='" + args + "'")

				if err != nil {
					_, err = session.ChannelMessageSend(ctx.Channel().ID, "An error occured.")
					
					if err != nil {
						return nil
					}

					return err
				} else {
					_, err = session.ChannelMessageSend(ctx.Channel().ID, "Enabled " + args)
				}

				defer delete.Close()

			} else {
				_, err = session.ChannelMessageSend(ctx.Channel().ID, "This command is not disabled.")
				
				if err != nil {
					return nil
				}

				return err
			}
		} else {
			_, err = session.ChannelMessageSend(ctx.Channel().ID, "This command is not disabled.")
			
			if err != nil {
				return nil
			}

			return err
		}
	} else {
		_, err = session.ChannelMessageSend(ctx.Channel().ID, "You need to specify the command.")
	}

	return db.Close()
}