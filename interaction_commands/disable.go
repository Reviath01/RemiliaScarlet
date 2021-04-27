package interaction_commands

import (
	"fmt"

	multiplexer "git.randomchars.net/Reviath/RemiliaScarlet/Multiplexer"
	"git.randomchars.net/Reviath/RemiliaScarlet/interactions"
	"git.randomchars.net/Reviath/RemiliaScarlet/sql"
	"github.com/bwmarrin/discordgo"
)

func DisableCommand(session *discordgo.Session, interaction interactions.Interaction) interactions.InteractionResponse {
	db := sql.Connect()
	defer db.Close()

	type Tag struct {
		isblocked string
	}

	var tag Tag

	if sql.CheckLanguage(interaction.GuildID) == "tr" {
		if !multiplexer.CheckAdministratorPermission(session, interaction.Member.User.ID, interaction.ChannelID) {
			return multiplexer.CreateResponse("Yeterli yetkiye sahip değilsin.")
		}

		args := interaction.Data.Options[0].Value.(string)

		if args == "afk" {
			err := db.QueryRow(fmt.Sprintf("SELECT isblocked FROM disabledcommands WHERE commandname ='afk' AND guildid ='%s'", interaction.GuildID)).Scan(&tag.isblocked)
			if err == nil {
				return multiplexer.CreateResponse("Bu komut zaten engellenmiş.")
			} else {
				insert, _ := db.Query(fmt.Sprintf("INSERT INTO disabledcommands (isblocked, commandname, guildid) VALUES ('True', 'afk', '%s')", interaction.GuildID))

				defer insert.Close()

				return multiplexer.CreateResponse("Belirtilen komut başarıyla engellendi")
			}
		} else if args == "author" {
			err := db.QueryRow(fmt.Sprintf("SELECT isblocked FROM disabledcommands WHERE commandname ='author' AND guildid ='%s'", interaction.GuildID)).Scan(&tag.isblocked)
			if err == nil {
				return multiplexer.CreateResponse("Bu komut zaten engellenmiş.")
			} else {
				insert, _ := db.Query(fmt.Sprintf("INSERT INTO disabledcommands (isblocked, commandname, guildid) VALUES ('True', 'author', '%s')", interaction.GuildID))

				defer insert.Close()

				return multiplexer.CreateResponse("Belirtilen komut başarıyla engellendi")
			}
		} else if args == "avatar" {
			err := db.QueryRow(fmt.Sprintf("SELECT isblocked FROM disabledcommands WHERE commandname ='avatar' AND guildid ='%s'", interaction.GuildID)).Scan(&tag.isblocked)
			if err == nil {
				return multiplexer.CreateResponse("Bu komut zaten engellenmiş.")
			} else {
				insert, _ := db.Query(fmt.Sprintf("INSERT INTO disabledcommands (isblocked, commandname, guildid) VALUES ('True', 'avatar', '%s')", interaction.GuildID))

				defer insert.Close()

				return multiplexer.CreateResponse("Belirtilen komut başarıyla engellendi")
			}
		} else if args == "ban" {
			err := db.QueryRow(fmt.Sprintf("SELECT isblocked FROM disabledcommands WHERE commandname ='ban' AND guildid ='%s'", interaction.GuildID)).Scan(&tag.isblocked)
			if err == nil {
				return multiplexer.CreateResponse("Bu komut zaten engellenmiş.")
			} else {
				insert, _ := db.Query(fmt.Sprintf("INSERT INTO disabledcommands (isblocked, commandname, guildid) VALUES ('True', 'ban', '%s')", interaction.GuildID))

				defer insert.Close()

				return multiplexer.CreateResponse("Belirtilen komut başarıyla engellendi")
			}
		} else if args == "embed" {
			err := db.QueryRow(fmt.Sprintf("SELECT isblocked FROM disabledcommands WHERE commandname ='embed' AND guildid ='%s'", interaction.GuildID)).Scan(&tag.isblocked)
			if err == nil {
				return multiplexer.CreateResponse("Bu komut zaten engellenmiş.")
			} else {
				insert, _ := db.Query(fmt.Sprintf("INSERT INTO disabledcommands (isblocked, commandname, guildid) VALUES ('True', 'embed', '%s')", interaction.GuildID))

				defer insert.Close()

				return multiplexer.CreateResponse("Belirtilen komut başarıyla engellendi")
			}
		} else if args == "guild_info" {
			err := db.QueryRow(fmt.Sprintf("SELECT isblocked FROM disabledcommands WHERE commandname ='guild_info' AND guildid ='%s'", interaction.GuildID)).Scan(&tag.isblocked)
			if err == nil {
				return multiplexer.CreateResponse("Bu komut zaten engellenmiş.")
			} else {
				insert, _ := db.Query(fmt.Sprintf("INSERT INTO disabledcommands (isblocked, commandname, guildid) VALUES ('True', 'guild_info', '%s')", interaction.GuildID))

				defer insert.Close()

				return multiplexer.CreateResponse("Belirtilen komut başarıyla engellendi")
			}
		} else if args == "hug" {
			err := db.QueryRow(fmt.Sprintf("SELECT isblocked FROM disabledcommands WHERE commandname ='hug' AND guildid ='%s'", interaction.GuildID)).Scan(&tag.isblocked)
			if err == nil {
				return multiplexer.CreateResponse("Bu komut zaten engellenmiş.")
			} else {
				insert, _ := db.Query(fmt.Sprintf("INSERT INTO disabledcommands (isblocked, commandname, guildid) VALUES ('True', 'hug', '%s')", interaction.GuildID))

				defer insert.Close()

				return multiplexer.CreateResponse("Belirtilen komut başarıyla engellendi")
			}
		} else if args == "icon" {
			err := db.QueryRow(fmt.Sprintf("SELECT isblocked FROM disabledcommands WHERE commandname ='icon' AND guildid ='%s'", interaction.GuildID)).Scan(&tag.isblocked)
			if err == nil {
				return multiplexer.CreateResponse("Bu komut zaten engellenmiş.")
			} else {
				insert, _ := db.Query(fmt.Sprintf("INSERT INTO disabledcommands (isblocked, commandname, guildid) VALUES ('True', 'icon', '%s')", interaction.GuildID))

				defer insert.Close()

				return multiplexer.CreateResponse("Belirtilen komut başarıyla engellendi")
			}
		} else if args == "kick" {
			err := db.QueryRow(fmt.Sprintf("SELECT isblocked FROM disabledcommands WHERE commandname ='kick' AND guildid ='%s'", interaction.GuildID)).Scan(&tag.isblocked)
			if err == nil {
				return multiplexer.CreateResponse("Bu komut zaten engellenmiş.")
			} else {
				insert, _ := db.Query(fmt.Sprintf("INSERT INTO disabledcommands (isblocked, commandname, guildid) VALUES ('True', 'kick', '%s')", interaction.GuildID))

				defer insert.Close()

				return multiplexer.CreateResponse("Belirtilen komut başarıyla engellendi")
			}
		} else if args == "kiss" {
			err := db.QueryRow(fmt.Sprintf("SELECT isblocked FROM disabledcommands WHERE commandname ='kiss' AND guildid ='%s'", interaction.GuildID)).Scan(&tag.isblocked)
			if err == nil {
				return multiplexer.CreateResponse("Bu komut zaten engellenmiş.")
			} else {
				insert, _ := db.Query(fmt.Sprintf("INSERT INTO disabledcommands (isblocked, commandname, guildid) VALUES ('True', 'kiss', '%s')", interaction.GuildID))

				defer insert.Close()

				return multiplexer.CreateResponse("Belirtilen komut başarıyla engellendi")
			}
		} else if args == "ping" {
			err := db.QueryRow(fmt.Sprintf("SELECT isblocked FROM disabledcommands WHERE commandname ='ping' AND guildid ='%s'", interaction.GuildID)).Scan(&tag.isblocked)
			if err == nil {
				return multiplexer.CreateResponse("Bu komut zaten engellenmiş.")
			} else {
				insert, _ := db.Query(fmt.Sprintf("INSERT INTO disabledcommands (isblocked, commandname, guildid) VALUES ('True', 'ping', '%s')", interaction.GuildID))

				defer insert.Close()

				return multiplexer.CreateResponse("Belirtilen komut başarıyla engellendi")
			}
		} else if args == "roles" {
			err := db.QueryRow(fmt.Sprintf("SELECT isblocked FROM disabledcommands WHERE commandname ='roles' AND guildid ='%s'", interaction.GuildID)).Scan(&tag.isblocked)
			if err == nil {
				return multiplexer.CreateResponse("Bu komut zaten engellenmiş.")
			} else {
				insert, _ := db.Query(fmt.Sprintf("INSERT INTO disabledcommands (isblocked, commandname, guildid) VALUES ('True', 'roles', '%s')", interaction.GuildID))

				defer insert.Close()

				return multiplexer.CreateResponse("Belirtilen komut başarıyla engellendi")
			}
		} else if args == "settings" {
			err := db.QueryRow(fmt.Sprintf("SELECT isblocked FROM disabledcommands WHERE commandname ='settings' AND guildid ='%s'", interaction.GuildID)).Scan(&tag.isblocked)
			if err == nil {
				return multiplexer.CreateResponse("Bu komut zaten engellenmiş.")
			} else {
				insert, _ := db.Query(fmt.Sprintf("INSERT INTO disabledcommands (isblocked, commandname, guildid) VALUES ('True', 'settings', '%s')", interaction.GuildID))

				defer insert.Close()

				return multiplexer.CreateResponse("Belirtilen komut başarıyla engellendi")
			}
		} else if args == "slap" {
			err := db.QueryRow(fmt.Sprintf("SELECT isblocked FROM disabledcommands WHERE commandname ='slap' AND guildid ='%s'", interaction.GuildID)).Scan(&tag.isblocked)
			if err == nil {
				return multiplexer.CreateResponse("Bu komut zaten engellenmiş.")
			} else {
				insert, _ := db.Query(fmt.Sprintf("INSERT INTO disabledcommands (isblocked, commandname, guildid) VALUES ('True', 'slap', '%s')", interaction.GuildID))

				defer insert.Close()

				return multiplexer.CreateResponse("Belirtilen komut başarıyla engellendi")
			}
		} else if args == "spoiler" {
			err := db.QueryRow(fmt.Sprintf("SELECT isblocked FROM disabledcommands WHERE commandname ='spoiler' AND guildid ='%s'", interaction.GuildID)).Scan(&tag.isblocked)
			if err == nil {
				return multiplexer.CreateResponse("Bu komut zaten engellenmiş.")
			} else {
				insert, _ := db.Query(fmt.Sprintf("INSERT INTO disabledcommands (isblocked, commandname, guildid) VALUES ('True', 'spoiler', '%s')", interaction.GuildID))

				defer insert.Close()

				return multiplexer.CreateResponse("Belirtilen komut başarıyla engellendi")
			}
		} else if args == "start_vote" {
			err := db.QueryRow(fmt.Sprintf("SELECT isblocked FROM disabledcommands WHERE commandname ='start_vote' AND guildid ='%s'", interaction.GuildID)).Scan(&tag.isblocked)
			if err == nil {
				return multiplexer.CreateResponse("Bu komut zaten engellenmiş.")
			} else {
				insert, _ := db.Query(fmt.Sprintf("INSERT INTO disabledcommands (isblocked, commandname, guildid) VALUES ('True', 'start_vote', '%s')", interaction.GuildID))

				defer insert.Close()

				return multiplexer.CreateResponse("Belirtilen komut başarıyla engellendi")
			}
		} else if args == "stats" {
			err := db.QueryRow(fmt.Sprintf("SELECT isblocked FROM disabledcommands WHERE commandname ='stats' AND guildid ='%s'", interaction.GuildID)).Scan(&tag.isblocked)
			if err == nil {
				return multiplexer.CreateResponse("Bu komut zaten engellenmiş.")
			} else {
				insert, _ := db.Query(fmt.Sprintf("INSERT INTO disabledcommands (isblocked, commandname, guildid) VALUES ('True', 'stats', '%s')", interaction.GuildID))

				defer insert.Close()

				return multiplexer.CreateResponse("Belirtilen komut başarıyla engellendi")
			}
		} else if args == "unban" {
			err := db.QueryRow(fmt.Sprintf("SELECT isblocked FROM disabledcommands WHERE commandname ='unban' AND guildid ='%s'", interaction.GuildID)).Scan(&tag.isblocked)
			if err == nil {
				return multiplexer.CreateResponse("Bu komut zaten engellenmiş.")
			} else {
				insert, _ := db.Query(fmt.Sprintf("INSERT INTO disabledcommands (isblocked, commandname, guildid) VALUES ('True', 'unban', '%s')", interaction.GuildID))

				defer insert.Close()

				return multiplexer.CreateResponse("Belirtilen komut başarıyla engellendi")
			}
		} else {
			return multiplexer.CreateResponse("You need to specify a command.")
		}
	}

	if !multiplexer.CheckAdministratorPermission(session, interaction.Member.User.ID, interaction.ChannelID) {
		return multiplexer.CreateResponse("You don't have enough permission")
	}

	args := interaction.Data.Options[0].Value.(string)

	if args == "afk" {
		err := db.QueryRow(fmt.Sprintf("SELECT isblocked FROM disabledcommands WHERE commandname ='afk' AND guildid ='%s'", interaction.GuildID)).Scan(&tag.isblocked)
		if err == nil {
			if tag.isblocked == "True" {
				return multiplexer.CreateResponse("This command is already blocked.")
			}
		} else {
			insert, _ := db.Query(fmt.Sprintf("INSERT INTO disabledcommands (isblocked, commandname, guildid) VALUES ('True', 'afk', '%s')", interaction.GuildID))

			defer insert.Close()

			return multiplexer.CreateResponse("Successfully disabled specified command.")
		}
	} else if args == "author" {
		err := db.QueryRow(fmt.Sprintf("SELECT isblocked FROM disabledcommands WHERE commandname ='author' AND guildid ='%s'", interaction.GuildID)).Scan(&tag.isblocked)
		if err == nil {
			if tag.isblocked == "True" {
				return multiplexer.CreateResponse("This command is already blocked.")
			}
		} else {
			insert, _ := db.Query(fmt.Sprintf("INSERT INTO disabledcommands (isblocked, commandname, guildid) VALUES ('True', 'author', '%s')", interaction.GuildID))

			defer insert.Close()

			return multiplexer.CreateResponse("Successfully disabled specified command.")
		}
	} else if args == "avatar" {
		err := db.QueryRow(fmt.Sprintf("SELECT isblocked FROM disabledcommands WHERE commandname ='avatar' AND guildid ='%s'", interaction.GuildID)).Scan(&tag.isblocked)
		if err == nil {
			if tag.isblocked == "True" {
				return multiplexer.CreateResponse("This command is already blocked.")
			}
		} else {
			insert, _ := db.Query(fmt.Sprintf("INSERT INTO disabledcommands (isblocked, commandname, guildid) VALUES ('True', 'avatar', '%s')", interaction.GuildID))

			defer insert.Close()

			return multiplexer.CreateResponse("Successfully disabled specified command.")
		}
	} else if args == "ban" {
		err := db.QueryRow(fmt.Sprintf("SELECT isblocked FROM disabledcommands WHERE commandname ='ban' AND guildid ='%s'", interaction.GuildID)).Scan(&tag.isblocked)
		if err == nil {
			if tag.isblocked == "True" {
				return multiplexer.CreateResponse("This command is already blocked.")
			}
		} else {
			insert, _ := db.Query(fmt.Sprintf("INSERT INTO disabledcommands (isblocked, commandname, guildid) VALUES ('True', 'ban', '%s')", interaction.GuildID))

			defer insert.Close()

			return multiplexer.CreateResponse("Successfully disabled specified command.")
		}
	} else if args == "embed" {
		err := db.QueryRow(fmt.Sprintf("SELECT isblocked FROM disabledcommands WHERE commandname ='embed' AND guildid ='%s'", interaction.GuildID)).Scan(&tag.isblocked)
		if err == nil {
			if tag.isblocked == "True" {
				return multiplexer.CreateResponse("This command is already blocked.")
			}
		} else {
			insert, _ := db.Query(fmt.Sprintf("INSERT INTO disabledcommands (isblocked, commandname, guildid) VALUES ('True', 'embed', '%s')", interaction.GuildID))

			defer insert.Close()

			return multiplexer.CreateResponse("Successfully disabled specified command.")
		}
	} else if args == "guild_info" {
		err := db.QueryRow(fmt.Sprintf("SELECT isblocked FROM disabledcommands WHERE commandname ='guild_info' AND guildid ='%s'", interaction.GuildID)).Scan(&tag.isblocked)
		if err == nil {
			if tag.isblocked == "True" {
				return multiplexer.CreateResponse("This command is already blocked.")
			}
		} else {
			insert, _ := db.Query(fmt.Sprintf("INSERT INTO disabledcommands (isblocked, commandname, guildid) VALUES ('True', 'guild_info', '%s')", interaction.GuildID))

			defer insert.Close()

			return multiplexer.CreateResponse("Successfully disabled specified command.")
		}
	} else if args == "hug" {
		err := db.QueryRow(fmt.Sprintf("SELECT isblocked FROM disabledcommands WHERE commandname ='hug' AND guildid ='%s'", interaction.GuildID)).Scan(&tag.isblocked)
		if err == nil {
			if tag.isblocked == "True" {
				return multiplexer.CreateResponse("This command is already blocked.")
			}
		} else {
			insert, _ := db.Query(fmt.Sprintf("INSERT INTO disabledcommands (isblocked, commandname, guildid) VALUES ('True', 'hug', '%s')", interaction.GuildID))

			defer insert.Close()

			return multiplexer.CreateResponse("Successfully disabled specified command.")
		}
	} else if args == "icon" {
		err := db.QueryRow(fmt.Sprintf("SELECT isblocked FROM disabledcommands WHERE commandname ='icon' AND guildid ='%s'", interaction.GuildID)).Scan(&tag.isblocked)
		if err == nil {
			if tag.isblocked == "True" {
				return multiplexer.CreateResponse("This command is already blocked.")
			}
		} else {
			insert, _ := db.Query(fmt.Sprintf("INSERT INTO disabledcommands (isblocked, commandname, guildid) VALUES ('True', 'icon', '%s')", interaction.GuildID))

			defer insert.Close()

			return multiplexer.CreateResponse("Successfully disabled specified command.")
		}
	} else if args == "kick" {
		err := db.QueryRow(fmt.Sprintf("SELECT isblocked FROM disabledcommands WHERE commandname ='kick' AND guildid ='%s'", interaction.GuildID)).Scan(&tag.isblocked)
		if err == nil {
			if tag.isblocked == "True" {
				return multiplexer.CreateResponse("This command is already blocked.")
			}
		} else {
			insert, _ := db.Query(fmt.Sprintf("INSERT INTO disabledcommands (isblocked, commandname, guildid) VALUES ('True', 'kick', '%s')", interaction.GuildID))

			defer insert.Close()
			return multiplexer.CreateResponse("Successfully disabled specified command.")
		}
	} else if args == "kiss" {
		err := db.QueryRow(fmt.Sprintf("SELECT isblocked FROM disabledcommands WHERE commandname ='kiss' AND guildid ='%s'", interaction.GuildID)).Scan(&tag.isblocked)
		if err == nil {
			if tag.isblocked == "True" {
				return multiplexer.CreateResponse("This command is already blocked.")
			}
		} else {
			insert, _ := db.Query(fmt.Sprintf("INSERT INTO disabledcommands (isblocked, commandname, guildid) VALUES ('True', 'kiss', '%s')", interaction.GuildID))

			defer insert.Close()

			return multiplexer.CreateResponse("Successfully disabled specified command.")
		}
	} else if args == "ping" {
		err := db.QueryRow(fmt.Sprintf("SELECT isblocked FROM disabledcommands WHERE commandname ='ping' AND guildid ='%s'", interaction.GuildID)).Scan(&tag.isblocked)
		if err == nil {
			if tag.isblocked == "True" {
				return multiplexer.CreateResponse("This command is already blocked.")
			}
		} else {
			insert, _ := db.Query(fmt.Sprintf("INSERT INTO disabledcommands (isblocked, commandname, guildid) VALUES ('True', 'ping', '%s')", interaction.GuildID))

			defer insert.Close()

			return multiplexer.CreateResponse("Successfully disabled specified command.")
		}
	} else if args == "roles" {
		err := db.QueryRow(fmt.Sprintf("SELECT isblocked FROM disabledcommands WHERE commandname ='roles' AND guildid ='%s'", interaction.GuildID)).Scan(&tag.isblocked)
		if err == nil {
			if tag.isblocked == "True" {
				return multiplexer.CreateResponse("This command is already blocked.")
			}
		} else {
			insert, _ := db.Query(fmt.Sprintf("INSERT INTO disabledcommands (isblocked, commandname, guildid) VALUES ('True', 'roles', '%s')", interaction.GuildID))

			defer insert.Close()

			return multiplexer.CreateResponse("Successfully disabled specified command.")
		}
	} else if args == "settings" {
		err := db.QueryRow(fmt.Sprintf("SELECT isblocked FROM disabledcommands WHERE commandname ='settings' AND guildid ='%s'", interaction.GuildID)).Scan(&tag.isblocked)
		if err == nil {
			if tag.isblocked == "True" {
				return multiplexer.CreateResponse("This command is already blocked.")
			}
		} else {
			insert, _ := db.Query(fmt.Sprintf("INSERT INTO disabledcommands (isblocked, commandname, guildid) VALUES ('True', 'settings', '%s')", interaction.GuildID))

			defer insert.Close()

			return multiplexer.CreateResponse("Successfully disabled specified command.")
		}
	} else if args == "slap" {
		err := db.QueryRow(fmt.Sprintf("SELECT isblocked FROM disabledcommands WHERE commandname ='slap' AND guildid ='%s'", interaction.GuildID)).Scan(&tag.isblocked)
		if err == nil {
			if tag.isblocked == "True" {
				return multiplexer.CreateResponse("This command is already blocked.")
			}
		} else {
			insert, _ := db.Query(fmt.Sprintf("INSERT INTO disabledcommands (isblocked, commandname, guildid) VALUES ('True', 'slap', '%s')", interaction.GuildID))

			defer insert.Close()

			return multiplexer.CreateResponse("Successfully disabled specified command.")
		}
	} else if args == "spoiler" {
		err := db.QueryRow(fmt.Sprintf("SELECT isblocked FROM disabledcommands WHERE commandname ='spoiler' AND guildid ='%s'", interaction.GuildID)).Scan(&tag.isblocked)
		if err == nil {
			if tag.isblocked == "True" {
				return multiplexer.CreateResponse("This command is already blocked.")
			}
		} else {
			insert, _ := db.Query(fmt.Sprintf("INSERT INTO disabledcommands (isblocked, commandname, guildid) VALUES ('True', 'spoiler', '%s')", interaction.GuildID))

			defer insert.Close()

			return multiplexer.CreateResponse("Successfully disabled specified command.")
		}
	} else if args == "start_vote" {
		err := db.QueryRow(fmt.Sprintf("SELECT isblocked FROM disabledcommands WHERE commandname ='start_vote' AND guildid ='%s'", interaction.GuildID)).Scan(&tag.isblocked)
		if err == nil {
			if tag.isblocked == "True" {
				return multiplexer.CreateResponse("This command is already blocked.")
			}
		} else {
			insert, _ := db.Query(fmt.Sprintf("INSERT INTO disabledcommands (isblocked, commandname, guildid) VALUES ('True', 'start_vote', '%s')", interaction.GuildID))

			defer insert.Close()

			return multiplexer.CreateResponse("Successfully disabled specified command.")
		}
	} else if args == "stats" {
		err := db.QueryRow(fmt.Sprintf("SELECT isblocked FROM disabledcommands WHERE commandname ='stats' AND guildid ='%s'", interaction.GuildID)).Scan(&tag.isblocked)
		if err == nil {
			if tag.isblocked == "True" {
				return multiplexer.CreateResponse("This command is already blocked.")
			}
		} else {
			insert, _ := db.Query(fmt.Sprintf("INSERT INTO disabledcommands (isblocked, commandname, guildid) VALUES ('True', 'stats', '%s')", interaction.GuildID))

			defer insert.Close()

			return multiplexer.CreateResponse("Successfully disabled specified command.")
		}
	} else if args == "unban" {
		err := db.QueryRow(fmt.Sprintf("SELECT isblocked FROM disabledcommands WHERE commandname ='unban' AND guildid ='%s'", interaction.GuildID)).Scan(&tag.isblocked)
		if err == nil {
			if tag.isblocked == "True" {
				return multiplexer.CreateResponse("This command is already blocked.")
			}
		} else {
			insert, _ := db.Query(fmt.Sprintf("INSERT INTO disabledcommands (isblocked, commandname, guildid) VALUES ('True', 'unban', '%s')", interaction.GuildID))

			defer insert.Close()

			return multiplexer.CreateResponse("Successfully disabled specified command.")
		}
	}
	return multiplexer.CreateResponse("You need to specify a command.")
}
