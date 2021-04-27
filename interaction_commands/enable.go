package interaction_commands

import (
	"fmt"

	multiplexer "git.randomchars.net/Reviath/RemiliaScarlet/Multiplexer"
	"git.randomchars.net/Reviath/RemiliaScarlet/interactions"
	"git.randomchars.net/Reviath/RemiliaScarlet/sql"
	"github.com/bwmarrin/discordgo"
)

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

		if args == "afk" {
			err := db.QueryRow(fmt.Sprintf("SELECT isblocked FROM disabledcommands WHERE commandname ='%s' AND guildid ='%s'", args, interaction.GuildID)).Scan(&tag.isblocked)
			if err == nil {
				if tag.isblocked == "True" {
					delete, err := db.Query(fmt.Sprintf("DELETE FROM disabledcommands WHERE guildid ='%s' AND commandname ='%s'", interaction.GuildID, args))

					if err != nil {
						return multiplexer.CreateResponse("Bir hata oluştu")
					}
					defer delete.Close()
					return multiplexer.CreateResponse(args + " komutu tekrar kullanılabilir.")

				} else {
					return multiplexer.CreateResponse("Bu komut engellenmemiş")
				}
			} else {
				return multiplexer.CreateResponse("Bu komut engellenmemiş.")
			}
		} else if args == "author" {
			err := db.QueryRow(fmt.Sprintf("SELECT isblocked FROM disabledcommands WHERE commandname ='%s' AND guildid ='%s'", args, interaction.GuildID)).Scan(&tag.isblocked)
			if err == nil {
				if tag.isblocked == "True" {
					delete, err := db.Query(fmt.Sprintf("DELETE FROM disabledcommands WHERE guildid ='%s' AND commandname ='%s'", interaction.GuildID, args))

					if err != nil {
						return multiplexer.CreateResponse("Bir hata oluştu")
					}
					defer delete.Close()
					return multiplexer.CreateResponse(args + " komutu tekrar kullanılabilir.")

				} else {
					return multiplexer.CreateResponse("Bu komut engellenmemiş")
				}
			} else {
				return multiplexer.CreateResponse("Bu komut engellenmemiş.")
			}
		} else if args == "avatar" {
			err := db.QueryRow(fmt.Sprintf("SELECT isblocked FROM disabledcommands WHERE commandname ='%s' AND guildid ='%s'", args, interaction.GuildID)).Scan(&tag.isblocked)
			if err == nil {
				if tag.isblocked == "True" {
					delete, err := db.Query(fmt.Sprintf("DELETE FROM disabledcommands WHERE guildid ='%s' AND commandname ='%s'", interaction.GuildID, args))

					if err != nil {
						return multiplexer.CreateResponse("Bir hata oluştu")
					}
					defer delete.Close()
					return multiplexer.CreateResponse(args + " komutu tekrar kullanılabilir.")

				} else {
					return multiplexer.CreateResponse("Bu komut engellenmemiş")
				}
			} else {
				return multiplexer.CreateResponse("Bu komut engellenmemiş.")
			}
		} else if args == "ban" {
			err := db.QueryRow(fmt.Sprintf("SELECT isblocked FROM disabledcommands WHERE commandname ='%s' AND guildid ='%s'", args, interaction.GuildID)).Scan(&tag.isblocked)
			if err == nil {
				if tag.isblocked == "True" {
					delete, err := db.Query(fmt.Sprintf("DELETE FROM disabledcommands WHERE guildid ='%s' AND commandname ='%s'", interaction.GuildID, args))

					if err != nil {
						return multiplexer.CreateResponse("Bir hata oluştu")
					}
					defer delete.Close()
					return multiplexer.CreateResponse(args + " komutu tekrar kullanılabilir.")

				} else {
					return multiplexer.CreateResponse("Bu komut engellenmemiş")
				}
			} else {
				return multiplexer.CreateResponse("Bu komut engellenmemiş.")
			}
		} else if args == "embed" {
			err := db.QueryRow(fmt.Sprintf("SELECT isblocked FROM disabledcommands WHERE commandname ='%s' AND guildid ='%s'", args, interaction.GuildID)).Scan(&tag.isblocked)
			if err == nil {
				if tag.isblocked == "True" {
					delete, err := db.Query(fmt.Sprintf("DELETE FROM disabledcommands WHERE guildid ='%s' AND commandname ='%s'", interaction.GuildID, args))

					if err != nil {
						return multiplexer.CreateResponse("Bir hata oluştu")
					}
					defer delete.Close()
					return multiplexer.CreateResponse(args + " komutu tekrar kullanılabilir.")

				} else {
					return multiplexer.CreateResponse("Bu komut engellenmemiş")
				}
			} else {
				return multiplexer.CreateResponse("Bu komut engellenmemiş.")
			}
		} else if args == "Channel_info" {
			err := db.QueryRow(fmt.Sprintf("SELECT isblocked FROM disabledcommands WHERE commandname ='%s' AND guildid ='%s'", args, interaction.GuildID)).Scan(&tag.isblocked)
			if err == nil {
				if tag.isblocked == "True" {
					delete, err := db.Query(fmt.Sprintf("DELETE FROM disabledcommands WHERE guildid ='%s' AND commandname ='%s'", interaction.GuildID, args))

					if err != nil {
						return multiplexer.CreateResponse("Bir hata oluştu")
					}
					defer delete.Close()
					return multiplexer.CreateResponse(args + " komutu tekrar kullanılabilir.")

				} else {
					return multiplexer.CreateResponse("Bu komut engellenmemiş")
				}
			} else {
				return multiplexer.CreateResponse("Bu komut engellenmemiş.")
			}
		} else if args == "hug" {
			err := db.QueryRow(fmt.Sprintf("SELECT isblocked FROM disabledcommands WHERE commandname ='%s' AND guildid ='%s'", args, interaction.GuildID)).Scan(&tag.isblocked)
			if err == nil {
				if tag.isblocked == "True" {
					delete, err := db.Query(fmt.Sprintf("DELETE FROM disabledcommands WHERE guildid ='%s' AND commandname ='%s'", interaction.GuildID, args))

					if err != nil {
						return multiplexer.CreateResponse("Bir hata oluştu")
					}
					defer delete.Close()
					return multiplexer.CreateResponse(args + " komutu tekrar kullanılabilir.")

				} else {
					return multiplexer.CreateResponse("Bu komut engellenmemiş")
				}
			} else {
				return multiplexer.CreateResponse("Bu komut engellenmemiş.")
			}
		} else if args == "kick" {
			err := db.QueryRow(fmt.Sprintf("SELECT isblocked FROM disabledcommands WHERE commandname ='%s' AND guildid ='%s'", args, interaction.GuildID)).Scan(&tag.isblocked)
			if err == nil {
				if tag.isblocked == "True" {
					delete, err := db.Query(fmt.Sprintf("DELETE FROM disabledcommands WHERE guildid ='%s' AND commandname ='%s'", interaction.GuildID, args))

					if err != nil {
						return multiplexer.CreateResponse("Bir hata oluştu")
					}
					defer delete.Close()
					return multiplexer.CreateResponse(args + " komutu tekrar kullanılabilir.")

				} else {
					return multiplexer.CreateResponse("Bu komut engellenmemiş")
				}
			} else {
				return multiplexer.CreateResponse("Bu komut engellenmemiş.")
			}
		} else if args == "kiss" {
			err := db.QueryRow(fmt.Sprintf("SELECT isblocked FROM disabledcommands WHERE commandname ='%s' AND guildid ='%s'", args, interaction.GuildID)).Scan(&tag.isblocked)
			if err == nil {
				if tag.isblocked == "True" {
					delete, err := db.Query(fmt.Sprintf("DELETE FROM disabledcommands WHERE guildid ='%s' AND commandname ='%s'", interaction.GuildID, args))

					if err != nil {
						return multiplexer.CreateResponse("Bir hata oluştu")
					}
					defer delete.Close()
					return multiplexer.CreateResponse(args + " komutu tekrar kullanılabilir.")

				} else {
					return multiplexer.CreateResponse("Bu komut engellenmemiş")
				}
			} else {
				return multiplexer.CreateResponse("Bu komut engellenmemiş.")
			}
		} else if args == "ping" {
			err := db.QueryRow(fmt.Sprintf("SELECT isblocked FROM disabledcommands WHERE commandname ='%s' AND guildid ='%s'", args, interaction.GuildID)).Scan(&tag.isblocked)
			if err == nil {
				if tag.isblocked == "True" {
					delete, err := db.Query(fmt.Sprintf("DELETE FROM disabledcommands WHERE guildid ='%s' AND commandname ='%s'", interaction.GuildID, args))

					if err != nil {
						return multiplexer.CreateResponse("Bir hata oluştu")
					}
					defer delete.Close()
					return multiplexer.CreateResponse(args + " komutu tekrar kullanılabilir.")

				} else {
					return multiplexer.CreateResponse("Bu komut engellenmemiş")
				}
			} else {
				return multiplexer.CreateResponse("Bu komut engellenmemiş.")
			}
		} else if args == "roles" {
			err := db.QueryRow(fmt.Sprintf("SELECT isblocked FROM disabledcommands WHERE commandname ='%s' AND guildid ='%s'", args, interaction.GuildID)).Scan(&tag.isblocked)
			if err == nil {
				if tag.isblocked == "True" {
					delete, err := db.Query(fmt.Sprintf("DELETE FROM disabledcommands WHERE guildid ='%s' AND commandname ='%s'", interaction.GuildID, args))

					if err != nil {
						return multiplexer.CreateResponse("Bir hata oluştu")
					}
					defer delete.Close()
					return multiplexer.CreateResponse(args + " komutu tekrar kullanılabilir.")

				} else {
					return multiplexer.CreateResponse("Bu komut engellenmemiş")
				}
			} else {
				return multiplexer.CreateResponse("Bu komut engellenmemiş.")
			}
		} else if args == "settings" {
			err := db.QueryRow(fmt.Sprintf("SELECT isblocked FROM disabledcommands WHERE commandname ='%s' AND guildid ='%s'", args, interaction.GuildID)).Scan(&tag.isblocked)
			if err == nil {
				if tag.isblocked == "True" {
					delete, err := db.Query(fmt.Sprintf("DELETE FROM disabledcommands WHERE guildid ='%s' AND commandname ='%s'", interaction.GuildID, args))

					if err != nil {
						return multiplexer.CreateResponse("Bir hata oluştu")
					}
					defer delete.Close()
					return multiplexer.CreateResponse(args + " komutu tekrar kullanılabilir.")

				} else {
					return multiplexer.CreateResponse("Bu komut engellenmemiş")
				}
			} else {
				return multiplexer.CreateResponse("Bu komut engellenmemiş.")
			}
		} else if args == "slap" {
			err := db.QueryRow(fmt.Sprintf("SELECT isblocked FROM disabledcommands WHERE commandname ='%s' AND guildid ='%s'", args, interaction.GuildID)).Scan(&tag.isblocked)
			if err == nil {
				if tag.isblocked == "True" {
					delete, err := db.Query(fmt.Sprintf("DELETE FROM disabledcommands WHERE guildid ='%s' AND commandname ='%s'", interaction.GuildID, args))

					if err != nil {
						return multiplexer.CreateResponse("Bir hata oluştu")
					}
					defer delete.Close()
					return multiplexer.CreateResponse(args + " komutu tekrar kullanılabilir.")

				} else {
					return multiplexer.CreateResponse("Bu komut engellenmemiş")
				}
			} else {
				return multiplexer.CreateResponse("Bu komut engellenmemiş.")
			}
		} else if args == "spoiler" {
			err := db.QueryRow(fmt.Sprintf("SELECT isblocked FROM disabledcommands WHERE commandname ='%s' AND guildid ='%s'", args, interaction.GuildID)).Scan(&tag.isblocked)
			if err == nil {
				if tag.isblocked == "True" {
					delete, err := db.Query(fmt.Sprintf("DELETE FROM disabledcommands WHERE guildid ='%s' AND commandname ='%s'", interaction.GuildID, args))

					if err != nil {
						return multiplexer.CreateResponse("Bir hata oluştu")
					}
					defer delete.Close()
					return multiplexer.CreateResponse(args + " komutu tekrar kullanılabilir.")

				} else {
					return multiplexer.CreateResponse("Bu komut engellenmemiş")
				}
			} else {
				return multiplexer.CreateResponse("Bu komut engellenmemiş.")
			}
		} else if args == "start_vote" {
			err := db.QueryRow(fmt.Sprintf("SELECT isblocked FROM disabledcommands WHERE commandname ='%s' AND guildid ='%s'", args, interaction.GuildID)).Scan(&tag.isblocked)
			if err == nil {
				if tag.isblocked == "True" {
					delete, err := db.Query(fmt.Sprintf("DELETE FROM disabledcommands WHERE guildid ='%s' AND commandname ='%s'", interaction.GuildID, args))

					if err != nil {
						return multiplexer.CreateResponse("Bir hata oluştu")
					}
					defer delete.Close()
					return multiplexer.CreateResponse(args + " komutu tekrar kullanılabilir.")

				} else {
					return multiplexer.CreateResponse("Bu komut engellenmemiş")
				}
			} else {
				return multiplexer.CreateResponse("Bu komut engellenmemiş.")
			}
		} else if args == "stats" {
			err := db.QueryRow(fmt.Sprintf("SELECT isblocked FROM disabledcommands WHERE commandname ='%s' AND guildid ='%s'", args, interaction.GuildID)).Scan(&tag.isblocked)
			if err == nil {
				if tag.isblocked == "True" {
					delete, err := db.Query(fmt.Sprintf("DELETE FROM disabledcommands WHERE guildid ='%s' AND commandname ='%s'", interaction.GuildID, args))

					if err != nil {
						return multiplexer.CreateResponse("Bir hata oluştu")
					}
					defer delete.Close()
					return multiplexer.CreateResponse(args + " komutu tekrar kullanılabilir.")

				} else {
					return multiplexer.CreateResponse("Bu komut engellenmemiş")
				}
			} else {
				return multiplexer.CreateResponse("Bu komut engellenmemiş.")
			}
		} else if args == "unban" {
			err := db.QueryRow(fmt.Sprintf("SELECT isblocked FROM disabledcommands WHERE commandname ='%s' AND guildid ='%s'", args, interaction.GuildID)).Scan(&tag.isblocked)
			if err == nil {
				if tag.isblocked == "True" {
					delete, err := db.Query(fmt.Sprintf("DELETE FROM disabledcommands WHERE guildid ='%s' AND commandname ='%s'", interaction.GuildID, args))

					if err != nil {
						return multiplexer.CreateResponse("Bir hata oluştu")
					}
					defer delete.Close()
					return multiplexer.CreateResponse(args + " komutu tekrar kullanılabilir.")

				} else {
					return multiplexer.CreateResponse("Bu komut engellenmemiş")
				}
			} else {
				return multiplexer.CreateResponse("Bu komut engellenmemiş.")
			}
		} else {
			return multiplexer.CreateResponse("Bir komut belirtmelisin")
		}
	}
	if !multiplexer.CheckAdministratorPermission(session, interaction.Member.User.ID, interaction.ChannelID) {
		return multiplexer.CreateResponse("You don't have enough permission.")
	}

	args := interaction.Data.Options[0].Value.(string)

	if args == "afk" {
		err := db.QueryRow(fmt.Sprintf("SELECT isblocked FROM disabledcommands WHERE commandname ='%s' AND guildid ='%s'", args, interaction.GuildID)).Scan(&tag.isblocked)
		if err == nil {
			if tag.isblocked == "True" {
				delete, err := db.Query(fmt.Sprintf("DELETE FROM disabledcommands WHERE guildid ='%s' AND commandname ='%s'", interaction.GuildID, args))

				if err != nil {
					return multiplexer.CreateResponse("An error occurred")
				}

				defer delete.Close()

				return multiplexer.CreateResponse("Enabled " + args)
			} else {
				return multiplexer.CreateResponse("This command is not disabled.")
			}
		} else {
			return multiplexer.CreateResponse("This command is not disabled.")
		}
	} else if args == "author" {
		err := db.QueryRow(fmt.Sprintf("SELECT isblocked FROM disabledcommands WHERE commandname ='%s' AND guildid ='%s'", args, interaction.GuildID)).Scan(&tag.isblocked)
		if err == nil {
			if tag.isblocked == "True" {
				delete, err := db.Query(fmt.Sprintf("DELETE FROM disabledcommands WHERE guildid ='%s' AND commandname ='%s'", interaction.GuildID, args))

				if err != nil {
					return multiplexer.CreateResponse("An error occurred")
				}

				defer delete.Close()

				return multiplexer.CreateResponse("Enabled " + args)
			} else {
				return multiplexer.CreateResponse("This command is not disabled.")
			}
		} else {
			return multiplexer.CreateResponse("This command is not disabled.")
		}
	} else if args == "avatar" {
		err := db.QueryRow(fmt.Sprintf("SELECT isblocked FROM disabledcommands WHERE commandname ='%s' AND guildid ='%s'", args, interaction.GuildID)).Scan(&tag.isblocked)
		if err == nil {
			if tag.isblocked == "True" {
				delete, err := db.Query(fmt.Sprintf("DELETE FROM disabledcommands WHERE guildid ='%s' AND commandname ='%s'", interaction.GuildID, args))

				if err != nil {
					return multiplexer.CreateResponse("An error occurred")
				}

				defer delete.Close()

				return multiplexer.CreateResponse("Enabled " + args)
			} else {
				return multiplexer.CreateResponse("This command is not disabled.")
			}
		} else {
			return multiplexer.CreateResponse("This command is not disabled.")
		}
	} else if args == "ban" {
		err := db.QueryRow(fmt.Sprintf("SELECT isblocked FROM disabledcommands WHERE commandname ='%s' AND guildid ='%s'", args, interaction.GuildID)).Scan(&tag.isblocked)
		if err == nil {
			if tag.isblocked == "True" {
				delete, err := db.Query(fmt.Sprintf("DELETE FROM disabledcommands WHERE guildid ='%s' AND commandname ='%s'", interaction.GuildID, args))

				if err != nil {
					return multiplexer.CreateResponse("An error occurred")
				}

				defer delete.Close()

				return multiplexer.CreateResponse("Enabled " + args)
			} else {
				return multiplexer.CreateResponse("This command is not disabled.")
			}
		} else {
			return multiplexer.CreateResponse("This command is not disabled.")
		}
	} else if args == "embed" {
		err := db.QueryRow(fmt.Sprintf("SELECT isblocked FROM disabledcommands WHERE commandname ='%s' AND guildid ='%s'", args, interaction.GuildID)).Scan(&tag.isblocked)
		if err == nil {
			if tag.isblocked == "True" {
				delete, err := db.Query(fmt.Sprintf("DELETE FROM disabledcommands WHERE guildid ='%s' AND commandname ='%s'", interaction.GuildID, args))

				if err != nil {
					return multiplexer.CreateResponse("An error occurred")
				}

				defer delete.Close()

				return multiplexer.CreateResponse("Enabled " + args)
			} else {
				return multiplexer.CreateResponse("This command is not disabled.")
			}
		} else {
			return multiplexer.CreateResponse("This command is not disabled.")
		}
	} else if args == "Channel_info" {
		err := db.QueryRow(fmt.Sprintf("SELECT isblocked FROM disabledcommands WHERE commandname ='%s' AND guildid ='%s'", args, interaction.GuildID)).Scan(&tag.isblocked)
		if err == nil {
			if tag.isblocked == "True" {
				delete, err := db.Query(fmt.Sprintf("DELETE FROM disabledcommands WHERE guildid ='%s' AND commandname ='%s'", interaction.GuildID, args))

				if err != nil {
					return multiplexer.CreateResponse("An error occurred")
				}

				defer delete.Close()

				return multiplexer.CreateResponse("Enabled " + args)
			} else {
				return multiplexer.CreateResponse("This command is not disabled.")
			}
		} else {
			return multiplexer.CreateResponse("This command is not disabled.")
		}
	} else if args == "hug" {
		err := db.QueryRow(fmt.Sprintf("SELECT isblocked FROM disabledcommands WHERE commandname ='%s' AND guildid ='%s'", args, interaction.GuildID)).Scan(&tag.isblocked)
		if err == nil {
			if tag.isblocked == "True" {
				delete, err := db.Query(fmt.Sprintf("DELETE FROM disabledcommands WHERE guildid ='%s' AND commandname ='%s'", interaction.GuildID, args))

				if err != nil {
					return multiplexer.CreateResponse("An error occurred")
				}

				defer delete.Close()

				return multiplexer.CreateResponse("Enabled " + args)
			} else {
				return multiplexer.CreateResponse("This command is not disabled.")
			}
		} else {
			return multiplexer.CreateResponse("This command is not disabled.")
		}
	} else if args == "kick" {
		err := db.QueryRow(fmt.Sprintf("SELECT isblocked FROM disabledcommands WHERE commandname ='%s' AND guildid ='%s'", args, interaction.GuildID)).Scan(&tag.isblocked)
		if err == nil {
			if tag.isblocked == "True" {
				delete, err := db.Query(fmt.Sprintf("DELETE FROM disabledcommands WHERE guildid ='%s' AND commandname ='%s'", interaction.GuildID, args))

				if err != nil {
					return multiplexer.CreateResponse("An error occurred")
				}

				defer delete.Close()

				return multiplexer.CreateResponse("Enabled " + args)
			} else {
				return multiplexer.CreateResponse("This command is not disabled.")
			}
		} else {
			return multiplexer.CreateResponse("This command is not disabled.")
		}
	} else if args == "kiss" {
		err := db.QueryRow(fmt.Sprintf("SELECT isblocked FROM disabledcommands WHERE commandname ='%s' AND guildid ='%s'", args, interaction.GuildID)).Scan(&tag.isblocked)
		if err == nil {
			if tag.isblocked == "True" {
				delete, err := db.Query(fmt.Sprintf("DELETE FROM disabledcommands WHERE guildid ='%s' AND commandname ='%s'", interaction.GuildID, args))

				if err != nil {
					return multiplexer.CreateResponse("An error occurred")
				}

				defer delete.Close()

				return multiplexer.CreateResponse("Enabled " + args)
			} else {
				return multiplexer.CreateResponse("This command is not disabled.")
			}
		} else {
			return multiplexer.CreateResponse("This command is not disabled.")
		}
	} else if args == "ping" {
		err := db.QueryRow(fmt.Sprintf("SELECT isblocked FROM disabledcommands WHERE commandname ='%s' AND guildid ='%s'", args, interaction.GuildID)).Scan(&tag.isblocked)
		if err == nil {
			if tag.isblocked == "True" {
				delete, err := db.Query(fmt.Sprintf("DELETE FROM disabledcommands WHERE guildid ='%s' AND commandname ='%s'", interaction.GuildID, args))

				if err != nil {
					return multiplexer.CreateResponse("An error occurred")
				}

				defer delete.Close()

				return multiplexer.CreateResponse("Enabled " + args)
			} else {
				return multiplexer.CreateResponse("This command is not disabled.")
			}
		} else {
			return multiplexer.CreateResponse("This command is not disabled.")
		}
	} else if args == "roles" {
		err := db.QueryRow(fmt.Sprintf("SELECT isblocked FROM disabledcommands WHERE commandname ='%s' AND guildid ='%s'", args, interaction.GuildID)).Scan(&tag.isblocked)
		if err == nil {
			if tag.isblocked == "True" {
				delete, err := db.Query(fmt.Sprintf("DELETE FROM disabledcommands WHERE guildid ='%s' AND commandname ='%s'", interaction.GuildID, args))

				if err != nil {
					return multiplexer.CreateResponse("An error occurred")
				}

				defer delete.Close()

				return multiplexer.CreateResponse("Enabled " + args)
			} else {
				return multiplexer.CreateResponse("This command is not disabled.")
			}
		} else {
			return multiplexer.CreateResponse("This command is not disabled.")
		}
	} else if args == "settings" {
		err := db.QueryRow(fmt.Sprintf("SELECT isblocked FROM disabledcommands WHERE commandname ='%s' AND guildid ='%s'", args, interaction.GuildID)).Scan(&tag.isblocked)
		if err == nil {
			if tag.isblocked == "True" {
				delete, err := db.Query(fmt.Sprintf("DELETE FROM disabledcommands WHERE guildid ='%s' AND commandname ='%s'", interaction.GuildID, args))

				if err != nil {
					return multiplexer.CreateResponse("An error occurred")
				}

				defer delete.Close()

				return multiplexer.CreateResponse("Enabled " + args)
			} else {
				return multiplexer.CreateResponse("This command is not disabled.")
			}
		} else {
			return multiplexer.CreateResponse("This command is not disabled.")
		}
	} else if args == "slap" {
		err := db.QueryRow(fmt.Sprintf("SELECT isblocked FROM disabledcommands WHERE commandname ='%s' AND guildid ='%s'", args, interaction.GuildID)).Scan(&tag.isblocked)
		if err == nil {
			if tag.isblocked == "True" {
				delete, err := db.Query(fmt.Sprintf("DELETE FROM disabledcommands WHERE guildid ='%s' AND commandname ='%s'", interaction.GuildID, args))

				if err != nil {
					return multiplexer.CreateResponse("An error occurred")
				}

				defer delete.Close()

				return multiplexer.CreateResponse("Enabled " + args)
			} else {
				return multiplexer.CreateResponse("This command is not disabled.")
			}
		} else {
			return multiplexer.CreateResponse("This command is not disabled.")
		}
	} else if args == "spoiler" {
		err := db.QueryRow(fmt.Sprintf("SELECT isblocked FROM disabledcommands WHERE commandname ='%s' AND guildid ='%s'", args, interaction.GuildID)).Scan(&tag.isblocked)
		if err == nil {
			if tag.isblocked == "True" {
				delete, err := db.Query(fmt.Sprintf("DELETE FROM disabledcommands WHERE guildid ='%s' AND commandname ='%s'", interaction.GuildID, args))

				if err != nil {
					return multiplexer.CreateResponse("An error occurred")
				}

				defer delete.Close()

				return multiplexer.CreateResponse("Enabled " + args)
			} else {
				return multiplexer.CreateResponse("This command is not disabled.")
			}
		} else {
			return multiplexer.CreateResponse("This command is not disabled.")
		}
	} else if args == "start_vote" {
		err := db.QueryRow(fmt.Sprintf("SELECT isblocked FROM disabledcommands WHERE commandname ='%s' AND guildid ='%s'", args, interaction.GuildID)).Scan(&tag.isblocked)
		if err == nil {
			if tag.isblocked == "True" {
				delete, err := db.Query(fmt.Sprintf("DELETE FROM disabledcommands WHERE guildid ='%s' AND commandname ='%s'", interaction.GuildID, args))

				if err != nil {
					return multiplexer.CreateResponse("An error occurred")
				}

				defer delete.Close()

				return multiplexer.CreateResponse("Enabled " + args)
			} else {
				return multiplexer.CreateResponse("This command is not disabled.")
			}
		} else {
			return multiplexer.CreateResponse("This command is not disabled.")
		}
	} else if args == "stats" {
		err := db.QueryRow(fmt.Sprintf("SELECT isblocked FROM disabledcommands WHERE commandname ='%s' AND guildid ='%s'", args, interaction.GuildID)).Scan(&tag.isblocked)
		if err == nil {
			if tag.isblocked == "True" {
				delete, err := db.Query(fmt.Sprintf("DELETE FROM disabledcommands WHERE guildid ='%s' AND commandname ='%s'", interaction.GuildID, args))

				if err != nil {
					return multiplexer.CreateResponse("An error occurred")
				}

				defer delete.Close()

				return multiplexer.CreateResponse("Enabled " + args)
			} else {
				return multiplexer.CreateResponse("This command is not disabled.")
			}
		} else {
			return multiplexer.CreateResponse("This command is not disabled.")
		}
	} else if args == "unban" {
		err := db.QueryRow(fmt.Sprintf("SELECT isblocked FROM disabledcommands WHERE commandname ='%s' AND guildid ='%s'", args, interaction.GuildID)).Scan(&tag.isblocked)
		if err == nil {
			if tag.isblocked == "True" {
				delete, err := db.Query(fmt.Sprintf("DELETE FROM disabledcommands WHERE guildid ='%s' AND commandname ='%s'", interaction.GuildID, args))

				if err != nil {
					return multiplexer.CreateResponse("An error occurred")
				}

				defer delete.Close()

				return multiplexer.CreateResponse("Enabled " + args)
			} else {
				return multiplexer.CreateResponse("This command is not disabled.")
			}
		} else {
			return multiplexer.CreateResponse("This command is not disabled.")
		}
	}
	return multiplexer.CreateResponse("You need to specify the command.")
}
