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
	perms, err := session.State.UserChannelPermissions(ctx.Author().ID, ctx.Channel().ID)
	if err == nil && (int(perms)&discordgo.PermissionAdministrator == discordgo.PermissionAdministrator) == false {
        _, err := session.ChannelMessageSend(ctx.Channel().ID, "You need administrator permission to run this command.")
	return err
    }

	if len(strings.Join(ctx.Args()," ")) < 1 {
		_, err := session.ChannelMessageSend(ctx.Channel().ID, "You need to specify a command.")
		return err
	}

	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/remilia")

    if err != nil {
        panic(err.Error())
    }

	type Tag struct {
		isblocked string `json:"isblocked"`
	}

	var tag Tag

	if strings.Join(ctx.Args(), " ") == "afk" {	
		err = db.QueryRow("SELECT isblocked FROM disabledcommands WHERE commandname ='" + strings.Join(ctx.Args(), " ") + "' AND guildid ='" + ctx.Guild().ID + "'").Scan(&tag.isblocked)
		if err == nil {
			if tag.isblocked == "True" {
				delete, err := db.Query("DELETE FROM disabledcommands WHERE guildid ='" + ctx.Guild().ID + "' AND commandname ='" + strings.Join(ctx.Args(), " ") + "'")

				if err != nil {
					_, err = session.ChannelMessageSend(ctx.Channel().ID, "An error occured.")
					return err
				}

				defer delete.Close()

			} else {
				_, err = session.ChannelMessageSend(ctx.Channel().ID, "This command is not disabled.")
				return err
			}
		} else {
			_, err = session.ChannelMessageSend(ctx.Channel().ID, "This command is not disabled.")
			return err
		}
	} else if strings.Join(ctx.Args(), " ") == "author" {	
		err = db.QueryRow("SELECT isblocked FROM disabledcommands WHERE commandname ='" + strings.Join(ctx.Args(), " ") + "' AND guildid ='" + ctx.Guild().ID + "'").Scan(&tag.isblocked)
		if err == nil {
			if tag.isblocked == "True" {
				delete, err := db.Query("DELETE FROM disabledcommands WHERE guildid ='" + ctx.Guild().ID + "' AND commandname ='" + strings.Join(ctx.Args(), " ") + "'")

				if err != nil {
					_, err = session.ChannelMessageSend(ctx.Channel().ID, "An error occured.")
					return err
				}

				defer delete.Close()

			} else {
				_, err = session.ChannelMessageSend(ctx.Channel().ID, "This command is not disabled.")
				return err
			}
		} else {
			_, err = session.ChannelMessageSend(ctx.Channel().ID, "This command is not disabled.")
			return err
		}
	} else if strings.Join(ctx.Args(), " ") == "avatar" {	
		err = db.QueryRow("SELECT isblocked FROM disabledcommands WHERE commandname ='" + strings.Join(ctx.Args(), " ") + "' AND guildid ='" + ctx.Guild().ID + "'").Scan(&tag.isblocked)
		if err == nil {
			if tag.isblocked == "True" {
				delete, err := db.Query("DELETE FROM disabledcommands WHERE guildid ='" + ctx.Guild().ID + "' AND commandname ='" + strings.Join(ctx.Args(), " ") + "'")

				if err != nil {
					_, err = session.ChannelMessageSend(ctx.Channel().ID, "An error occured.")
					return err
				}

				defer delete.Close()

			} else {
				_, err = session.ChannelMessageSend(ctx.Channel().ID, "This command is not disabled.")
				return err
			}
		} else {
			_, err = session.ChannelMessageSend(ctx.Channel().ID, "This command is not disabled.")
			return err
		}
	} else if strings.Join(ctx.Args(), " ") == "ban" {	
		err = db.QueryRow("SELECT isblocked FROM disabledcommands WHERE commandname ='" + strings.Join(ctx.Args(), " ") + "' AND guildid ='" + ctx.Guild().ID + "'").Scan(&tag.isblocked)
		if err == nil {
			if tag.isblocked == "True" {
				delete, err := db.Query("DELETE FROM disabledcommands WHERE guildid ='" + ctx.Guild().ID + "' AND commandname ='" + strings.Join(ctx.Args(), " ") + "'")

				if err != nil {
					_, err = session.ChannelMessageSend(ctx.Channel().ID, "An error occured.")
					return err
				}

				defer delete.Close()

			} else {
				_, err = session.ChannelMessageSend(ctx.Channel().ID, "This command is not disabled.")
				return err
			}
		} else {
			_, err = session.ChannelMessageSend(ctx.Channel().ID, "This command is not disabled.")
			return err
		}
	} else if strings.Join(ctx.Args(), " ") == "embed" {	
		err = db.QueryRow("SELECT isblocked FROM disabledcommands WHERE commandname ='" + strings.Join(ctx.Args(), " ") + "' AND guildid ='" + ctx.Guild().ID + "'").Scan(&tag.isblocked)
		if err == nil {
			if tag.isblocked == "True" {
				delete, err := db.Query("DELETE FROM disabledcommands WHERE guildid ='" + ctx.Guild().ID + "' AND commandname ='" + strings.Join(ctx.Args(), " ") + "'")

				if err != nil {
					_, err = session.ChannelMessageSend(ctx.Channel().ID, "An error occured.")
					return err
				}

				defer delete.Close()

			} else {
				_, err = session.ChannelMessageSend(ctx.Channel().ID, "This command is not disabled.")
				return err
			}
		} else {
			_, err = session.ChannelMessageSend(ctx.Channel().ID, "This command is not disabled.")
			return err
		}
	} else if strings.Join(ctx.Args(), " ") == "Channel_info" {	
		err = db.QueryRow("SELECT isblocked FROM disabledcommands WHERE commandname ='" + strings.Join(ctx.Args(), " ") + "' AND guildid ='" + ctx.Guild().ID + "'").Scan(&tag.isblocked)
		if err == nil {
			if tag.isblocked == "True" {
				delete, err := db.Query("DELETE FROM disabledcommands WHERE guildid ='" + ctx.Guild().ID + "' AND commandname ='" + strings.Join(ctx.Args(), " ") + "'")

				if err != nil {
					_, err = session.ChannelMessageSend(ctx.Channel().ID, "An error occured.")
					return err
				}

				defer delete.Close()

			} else {
				_, err = session.ChannelMessageSend(ctx.Channel().ID, "This command is not disabled.")
				return err
			}
		} else {
			_, err = session.ChannelMessageSend(ctx.Channel().ID, "This command is not disabled.")
			return err
		}
	} else if strings.Join(ctx.Args(), " ") == "hug" {	
		err = db.QueryRow("SELECT isblocked FROM disabledcommands WHERE commandname ='" + strings.Join(ctx.Args(), " ") + "' AND guildid ='" + ctx.Guild().ID + "'").Scan(&tag.isblocked)
		if err == nil {
			if tag.isblocked == "True" {
				delete, err := db.Query("DELETE FROM disabledcommands WHERE guildid ='" + ctx.Guild().ID + "' AND commandname ='" + strings.Join(ctx.Args(), " ") + "'")

				if err != nil {
					_, err = session.ChannelMessageSend(ctx.Channel().ID, "An error occured.")
					return err
				}

				defer delete.Close()

			} else {
				_, err = session.ChannelMessageSend(ctx.Channel().ID, "This command is not disabled.")
				return err
			}
		} else {
			_, err = session.ChannelMessageSend(ctx.Channel().ID, "This command is not disabled.")
			return err
		}
	} else if strings.Join(ctx.Args(), " ") == "kick" {	
		err = db.QueryRow("SELECT isblocked FROM disabledcommands WHERE commandname ='" + strings.Join(ctx.Args(), " ") + "' AND guildid ='" + ctx.Guild().ID + "'").Scan(&tag.isblocked)
		if err == nil {
			if tag.isblocked == "True" {
				delete, err := db.Query("DELETE FROM disabledcommands WHERE guildid ='" + ctx.Guild().ID + "' AND commandname ='" + strings.Join(ctx.Args(), " ") + "'")

				if err != nil {
					_, err = session.ChannelMessageSend(ctx.Channel().ID, "An error occured.")
					return err
				}

				defer delete.Close()

			} else {
				_, err = session.ChannelMessageSend(ctx.Channel().ID, "This command is not disabled.")
				return err
			}
		} else {
			_, err = session.ChannelMessageSend(ctx.Channel().ID, "This command is not disabled.")
			return err
		}
	} else if strings.Join(ctx.Args(), " ") == "afk" {	
		err = db.QueryRow("SELECT isblocked FROM disabledcommands WHERE commandname ='" + strings.Join(ctx.Args(), " ") + "' AND guildid ='" + ctx.Guild().ID + "'").Scan(&tag.isblocked)
		if err == nil {
			if tag.isblocked == "True" {
				delete, err := db.Query("DELETE FROM disabledcommands WHERE guildid ='" + ctx.Guild().ID + "' AND commandname ='" + strings.Join(ctx.Args(), " ") + "'")

				if err != nil {
					_, err = session.ChannelMessageSend(ctx.Channel().ID, "An error occured.")
					return err
				}

				defer delete.Close()

			} else {
				_, err = session.ChannelMessageSend(ctx.Channel().ID, "This command is not disabled.")
				return err
			}
		} else {
			_, err = session.ChannelMessageSend(ctx.Channel().ID, "This command is not disabled.")
			return err
		}
	} else if strings.Join(ctx.Args(), " ") == "kiss" {	
		err = db.QueryRow("SELECT isblocked FROM disabledcommands WHERE commandname ='" + strings.Join(ctx.Args(), " ") + "' AND guildid ='" + ctx.Guild().ID + "'").Scan(&tag.isblocked)
		if err == nil {
			if tag.isblocked == "True" {
				delete, err := db.Query("DELETE FROM disabledcommands WHERE guildid ='" + ctx.Guild().ID + "' AND commandname ='" + strings.Join(ctx.Args(), " ") + "'")

				if err != nil {
					_, err = session.ChannelMessageSend(ctx.Channel().ID, "An error occured.")
					return err
				}

				defer delete.Close()

			} else {
				_, err = session.ChannelMessageSend(ctx.Channel().ID, "This command is not disabled.")
				return err
			}
		} else {
			_, err = session.ChannelMessageSend(ctx.Channel().ID, "This command is not disabled.")
			return err
		}
	} else if strings.Join(ctx.Args(), " ") == "ping" {	
		err = db.QueryRow("SELECT isblocked FROM disabledcommands WHERE commandname ='" + strings.Join(ctx.Args(), " ") + "' AND guildid ='" + ctx.Guild().ID + "'").Scan(&tag.isblocked)
		if err == nil {
			if tag.isblocked == "True" {
				delete, err := db.Query("DELETE FROM disabledcommands WHERE guildid ='" + ctx.Guild().ID + "' AND commandname ='" + strings.Join(ctx.Args(), " ") + "'")

				if err != nil {
					_, err = session.ChannelMessageSend(ctx.Channel().ID, "An error occured.")
					return err
				}

				defer delete.Close()

			} else {
				_, err = session.ChannelMessageSend(ctx.Channel().ID, "This command is not disabled.")
				return err
			}
		} else {
			_, err = session.ChannelMessageSend(ctx.Channel().ID, "This command is not disabled.")
			return err
		}
	} else if strings.Join(ctx.Args(), " ") == "roles" {	
		err = db.QueryRow("SELECT isblocked FROM disabledcommands WHERE commandname ='" + strings.Join(ctx.Args(), " ") + "' AND guildid ='" + ctx.Guild().ID + "'").Scan(&tag.isblocked)
		if err == nil {
			if tag.isblocked == "True" {
				delete, err := db.Query("DELETE FROM disabledcommands WHERE guildid ='" + ctx.Guild().ID + "' AND commandname ='" + strings.Join(ctx.Args(), " ") + "'")

				if err != nil {
					_, err = session.ChannelMessageSend(ctx.Channel().ID, "An error occured.")
					return err
				}

				defer delete.Close()

			} else {
				_, err = session.ChannelMessageSend(ctx.Channel().ID, "This command is not disabled.")
				return err
			}
		} else {
			_, err = session.ChannelMessageSend(ctx.Channel().ID, "This command is not disabled.")
			return err
		}
	} else if strings.Join(ctx.Args(), " ") == "settings" {	
		err = db.QueryRow("SELECT isblocked FROM disabledcommands WHERE commandname ='" + strings.Join(ctx.Args(), " ") + "' AND guildid ='" + ctx.Guild().ID + "'").Scan(&tag.isblocked)
		if err == nil {
			if tag.isblocked == "True" {
				delete, err := db.Query("DELETE FROM disabledcommands WHERE guildid ='" + ctx.Guild().ID + "' AND commandname ='" + strings.Join(ctx.Args(), " ") + "'")

				if err != nil {
					_, err = session.ChannelMessageSend(ctx.Channel().ID, "An error occured.")
					return err
				}

				defer delete.Close()

			} else {
				_, err = session.ChannelMessageSend(ctx.Channel().ID, "This command is not disabled.")
				return err
			}
		} else {
			_, err = session.ChannelMessageSend(ctx.Channel().ID, "This command is not disabled.")
			return err
		}
	} else if strings.Join(ctx.Args(), " ") == "slap" {	
		err = db.QueryRow("SELECT isblocked FROM disabledcommands WHERE commandname ='" + strings.Join(ctx.Args(), " ") + "' AND guildid ='" + ctx.Guild().ID + "'").Scan(&tag.isblocked)
		if err == nil {
			if tag.isblocked == "True" {
				delete, err := db.Query("DELETE FROM disabledcommands WHERE guildid ='" + ctx.Guild().ID + "' AND commandname ='" + strings.Join(ctx.Args(), " ") + "'")

				if err != nil {
					_, err = session.ChannelMessageSend(ctx.Channel().ID, "An error occured.")
					return err
				}

				defer delete.Close()

			} else {
				_, err = session.ChannelMessageSend(ctx.Channel().ID, "This command is not disabled.")
				return err
			}
		} else {
			_, err = session.ChannelMessageSend(ctx.Channel().ID, "This command is not disabled.")
			return err
		}
	} else if strings.Join(ctx.Args(), " ") == "spoiler" {	
		err = db.QueryRow("SELECT isblocked FROM disabledcommands WHERE commandname ='" + strings.Join(ctx.Args(), " ") + "' AND guildid ='" + ctx.Guild().ID + "'").Scan(&tag.isblocked)
		if err == nil {
			if tag.isblocked == "True" {
				delete, err := db.Query("DELETE FROM disabledcommands WHERE guildid ='" + ctx.Guild().ID + "' AND commandname ='" + strings.Join(ctx.Args(), " ") + "'")

				if err != nil {
					_, err = session.ChannelMessageSend(ctx.Channel().ID, "An error occured.")
					return err
				}

				defer delete.Close()

			} else {
				_, err = session.ChannelMessageSend(ctx.Channel().ID, "This command is not disabled.")
				return err
			}
		} else {
			_, err = session.ChannelMessageSend(ctx.Channel().ID, "This command is not disabled.")
			return err
		}
	} else if strings.Join(ctx.Args(), " ") == "start_vote" {	
		err = db.QueryRow("SELECT isblocked FROM disabledcommands WHERE commandname ='" + strings.Join(ctx.Args(), " ") + "' AND guildid ='" + ctx.Guild().ID + "'").Scan(&tag.isblocked)
		if err == nil {
			if tag.isblocked == "True" {
				delete, err := db.Query("DELETE FROM disabledcommands WHERE guildid ='" + ctx.Guild().ID + "' AND commandname ='" + strings.Join(ctx.Args(), " ") + "'")

				if err != nil {
					_, err = session.ChannelMessageSend(ctx.Channel().ID, "An error occured.")
					return err
				}

				defer delete.Close()

			} else {
				_, err = session.ChannelMessageSend(ctx.Channel().ID, "This command is not disabled.")
				return err
			}
		} else {
			_, err = session.ChannelMessageSend(ctx.Channel().ID, "This command is not disabled.")
			return err
		}
	} else if strings.Join(ctx.Args(), " ") == "stats" {	
		err = db.QueryRow("SELECT isblocked FROM disabledcommands WHERE commandname ='" + strings.Join(ctx.Args(), " ") + "' AND guildid ='" + ctx.Guild().ID + "'").Scan(&tag.isblocked)
		if err == nil {
			if tag.isblocked == "True" {
				delete, err := db.Query("DELETE FROM disabledcommands WHERE guildid ='" + ctx.Guild().ID + "' AND commandname ='" + strings.Join(ctx.Args(), " ") + "'")

				if err != nil {
					_, err = session.ChannelMessageSend(ctx.Channel().ID, "An error occured.")
					return err
				}

				defer delete.Close()

			} else {
				_, err = session.ChannelMessageSend(ctx.Channel().ID, "This command is not disabled.")
				return err
			}
		} else {
			_, err = session.ChannelMessageSend(ctx.Channel().ID, "This command is not disabled.")
			return err
		}
	} else if strings.Join(ctx.Args(), " ") == "unban" {	
		err = db.QueryRow("SELECT isblocked FROM disabledcommands WHERE commandname ='" + strings.Join(ctx.Args(), " ") + "' AND guildid ='" + ctx.Guild().ID + "'").Scan(&tag.isblocked)
		if err == nil {
			if tag.isblocked == "True" {
				delete, err := db.Query("DELETE FROM disabledcommands WHERE guildid ='" + ctx.Guild().ID + "' AND commandname ='" + strings.Join(ctx.Args(), " ") + "'")

				if err != nil {
					_, err = session.ChannelMessageSend(ctx.Channel().ID, "An error occured.")
					return err
				}

				defer delete.Close()

			} else {
				_, err = session.ChannelMessageSend(ctx.Channel().ID, "This command is not disabled.")
				return err
			}
		} else {
			_, err = session.ChannelMessageSend(ctx.Channel().ID, "This command is not disabled.")
			return err
		}
	} else {
		_, err = session.ChannelMessageSend(ctx.Channel().ID, "You need to specify the command.")
	}

	return db.Close()
}