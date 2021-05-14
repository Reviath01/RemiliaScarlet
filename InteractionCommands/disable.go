package interactioncommands

import (
	"fmt"

	multiplexer "git.randomchars.net/Reviath/RemiliaScarlet/Multiplexer"
	"git.randomchars.net/Reviath/RemiliaScarlet/interactions"
	"git.randomchars.net/Reviath/RemiliaScarlet/sql"
	"github.com/bwmarrin/discordgo"
)

// DisableCommand is disable command for interactions
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

		switch args {
		case "afk":
			err := db.QueryRow(fmt.Sprintf("SELECT isblocked FROM disabledcommands WHERE commandname ='afk' AND guildid ='%s'", interaction.GuildID)).Scan(&tag.isblocked)
			if err == nil {
				return multiplexer.CreateResponse("Bu komut zaten engellenmiş.")
			}
			insert, _ := db.Query(fmt.Sprintf("INSERT INTO disabledcommands (isblocked, commandname, guildid) VALUES ('True', 'afk', '%s')", interaction.GuildID))

			defer insert.Close()

			return multiplexer.CreateResponse("Belirtilen komut başarıyla engellendi")
		case "author":
			err := db.QueryRow(fmt.Sprintf("SELECT isblocked FROM disabledcommands WHERE commandname ='author' AND guildid ='%s'", interaction.GuildID)).Scan(&tag.isblocked)
			if err == nil {
				return multiplexer.CreateResponse("Bu komut zaten engellenmiş.")
			}
			insert, _ := db.Query(fmt.Sprintf("INSERT INTO disabledcommands (isblocked, commandname, guildid) VALUES ('True', 'author', '%s')", interaction.GuildID))

			defer insert.Close()

			return multiplexer.CreateResponse("Belirtilen komut başarıyla engellendi")
		case "avatar":
			err := db.QueryRow(fmt.Sprintf("SELECT isblocked FROM disabledcommands WHERE commandname ='avatar' AND guildid ='%s'", interaction.GuildID)).Scan(&tag.isblocked)
			if err == nil {
				return multiplexer.CreateResponse("Bu komut zaten engellenmiş.")
			}
			insert, _ := db.Query(fmt.Sprintf("INSERT INTO disabledcommands (isblocked, commandname, guildid) VALUES ('True', 'avatar', '%s')", interaction.GuildID))

			defer insert.Close()

			return multiplexer.CreateResponse("Belirtilen komut başarıyla engellendi")
		case "ban":
			err := db.QueryRow(fmt.Sprintf("SELECT isblocked FROM disabledcommands WHERE commandname ='ban' AND guildid ='%s'", interaction.GuildID)).Scan(&tag.isblocked)
			if err == nil {
				return multiplexer.CreateResponse("Bu komut zaten engellenmiş.")
			}
			insert, _ := db.Query(fmt.Sprintf("INSERT INTO disabledcommands (isblocked, commandname, guildid) VALUES ('True', 'ban', '%s')", interaction.GuildID))

			defer insert.Close()

			return multiplexer.CreateResponse("Belirtilen komut başarıyla engellendi")
		case "embed":
			err := db.QueryRow(fmt.Sprintf("SELECT isblocked FROM disabledcommands WHERE commandname ='embed' AND guildid ='%s'", interaction.GuildID)).Scan(&tag.isblocked)
			if err == nil {
				return multiplexer.CreateResponse("Bu komut zaten engellenmiş.")
			}
			insert, _ := db.Query(fmt.Sprintf("INSERT INTO disabledcommands (isblocked, commandname, guildid) VALUES ('True', 'embed', '%s')", interaction.GuildID))

			defer insert.Close()

			return multiplexer.CreateResponse("Belirtilen komut başarıyla engellendi")
		case "guild_info":
			err := db.QueryRow(fmt.Sprintf("SELECT isblocked FROM disabledcommands WHERE commandname ='guild_info' AND guildid ='%s'", interaction.GuildID)).Scan(&tag.isblocked)
			if err == nil {
				return multiplexer.CreateResponse("Bu komut zaten engellenmiş.")
			}
			insert, _ := db.Query(fmt.Sprintf("INSERT INTO disabledcommands (isblocked, commandname, guildid) VALUES ('True', 'guild_info', '%s')", interaction.GuildID))

			defer insert.Close()

			return multiplexer.CreateResponse("Belirtilen komut başarıyla engellendi")
		case "hug":
			err := db.QueryRow(fmt.Sprintf("SELECT isblocked FROM disabledcommands WHERE commandname ='hug' AND guildid ='%s'", interaction.GuildID)).Scan(&tag.isblocked)
			if err == nil {
				return multiplexer.CreateResponse("Bu komut zaten engellenmiş.")
			}
			insert, _ := db.Query(fmt.Sprintf("INSERT INTO disabledcommands (isblocked, commandname, guildid) VALUES ('True', 'hug', '%s')", interaction.GuildID))

			defer insert.Close()

			return multiplexer.CreateResponse("Belirtilen komut başarıyla engellendi")
		case "icon":
			err := db.QueryRow(fmt.Sprintf("SELECT isblocked FROM disabledcommands WHERE commandname ='icon' AND guildid ='%s'", interaction.GuildID)).Scan(&tag.isblocked)
			if err == nil {
				return multiplexer.CreateResponse("Bu komut zaten engellenmiş.")
			}
			insert, _ := db.Query(fmt.Sprintf("INSERT INTO disabledcommands (isblocked, commandname, guildid) VALUES ('True', 'icon', '%s')", interaction.GuildID))

			defer insert.Close()

			return multiplexer.CreateResponse("Belirtilen komut başarıyla engellendi")
		case "kick":
			err := db.QueryRow(fmt.Sprintf("SELECT isblocked FROM disabledcommands WHERE commandname ='kick' AND guildid ='%s'", interaction.GuildID)).Scan(&tag.isblocked)
			if err == nil {
				return multiplexer.CreateResponse("Bu komut zaten engellenmiş.")
			}
			insert, _ := db.Query(fmt.Sprintf("INSERT INTO disabledcommands (isblocked, commandname, guildid) VALUES ('True', 'kick', '%s')", interaction.GuildID))

			defer insert.Close()

			return multiplexer.CreateResponse("Belirtilen komut başarıyla engellendi")
		case "kiss":
			err := db.QueryRow(fmt.Sprintf("SELECT isblocked FROM disabledcommands WHERE commandname ='kiss' AND guildid ='%s'", interaction.GuildID)).Scan(&tag.isblocked)
			if err == nil {
				return multiplexer.CreateResponse("Bu komut zaten engellenmiş.")
			}
			insert, _ := db.Query(fmt.Sprintf("INSERT INTO disabledcommands (isblocked, commandname, guildid) VALUES ('True', 'kiss', '%s')", interaction.GuildID))

			defer insert.Close()

			return multiplexer.CreateResponse("Belirtilen komut başarıyla engellendi")
		case "ping":
			err := db.QueryRow(fmt.Sprintf("SELECT isblocked FROM disabledcommands WHERE commandname ='ping' AND guildid ='%s'", interaction.GuildID)).Scan(&tag.isblocked)
			if err == nil {
				return multiplexer.CreateResponse("Bu komut zaten engellenmiş.")
			}
			insert, _ := db.Query(fmt.Sprintf("INSERT INTO disabledcommands (isblocked, commandname, guildid) VALUES ('True', 'ping', '%s')", interaction.GuildID))

			defer insert.Close()

			return multiplexer.CreateResponse("Belirtilen komut başarıyla engellendi")
		case "roles":
			err := db.QueryRow(fmt.Sprintf("SELECT isblocked FROM disabledcommands WHERE commandname ='roles' AND guildid ='%s'", interaction.GuildID)).Scan(&tag.isblocked)
			if err == nil {
				return multiplexer.CreateResponse("Bu komut zaten engellenmiş.")
			}
			insert, _ := db.Query(fmt.Sprintf("INSERT INTO disabledcommands (isblocked, commandname, guildid) VALUES ('True', 'roles', '%s')", interaction.GuildID))

			defer insert.Close()

			return multiplexer.CreateResponse("Belirtilen komut başarıyla engellendi")
		case "settings":
			err := db.QueryRow(fmt.Sprintf("SELECT isblocked FROM disabledcommands WHERE commandname ='settings' AND guildid ='%s'", interaction.GuildID)).Scan(&tag.isblocked)
			if err == nil {
				return multiplexer.CreateResponse("Bu komut zaten engellenmiş.")
			}
			insert, _ := db.Query(fmt.Sprintf("INSERT INTO disabledcommands (isblocked, commandname, guildid) VALUES ('True', 'settings', '%s')", interaction.GuildID))

			defer insert.Close()

			return multiplexer.CreateResponse("Belirtilen komut başarıyla engellendi")
		case "slap":
			err := db.QueryRow(fmt.Sprintf("SELECT isblocked FROM disabledcommands WHERE commandname ='slap' AND guildid ='%s'", interaction.GuildID)).Scan(&tag.isblocked)
			if err == nil {
				return multiplexer.CreateResponse("Bu komut zaten engellenmiş.")
			}
			insert, _ := db.Query(fmt.Sprintf("INSERT INTO disabledcommands (isblocked, commandname, guildid) VALUES ('True', 'slap', '%s')", interaction.GuildID))

			defer insert.Close()

			return multiplexer.CreateResponse("Belirtilen komut başarıyla engellendi")
		case "spoiler":
			err := db.QueryRow(fmt.Sprintf("SELECT isblocked FROM disabledcommands WHERE commandname ='spoiler' AND guildid ='%s'", interaction.GuildID)).Scan(&tag.isblocked)
			if err == nil {
				return multiplexer.CreateResponse("Bu komut zaten engellenmiş.")
			}
			insert, _ := db.Query(fmt.Sprintf("INSERT INTO disabledcommands (isblocked, commandname, guildid) VALUES ('True', 'spoiler', '%s')", interaction.GuildID))

			defer insert.Close()

			return multiplexer.CreateResponse("Belirtilen komut başarıyla engellendi")
		case "start_vote":
			err := db.QueryRow(fmt.Sprintf("SELECT isblocked FROM disabledcommands WHERE commandname ='start_vote' AND guildid ='%s'", interaction.GuildID)).Scan(&tag.isblocked)
			if err == nil {
				return multiplexer.CreateResponse("Bu komut zaten engellenmiş.")
			}
			insert, _ := db.Query(fmt.Sprintf("INSERT INTO disabledcommands (isblocked, commandname, guildid) VALUES ('True', 'start_vote', '%s')", interaction.GuildID))

			defer insert.Close()

			return multiplexer.CreateResponse("Belirtilen komut başarıyla engellendi")
		case "stats":
			err := db.QueryRow(fmt.Sprintf("SELECT isblocked FROM disabledcommands WHERE commandname ='stats' AND guildid ='%s'", interaction.GuildID)).Scan(&tag.isblocked)
			if err == nil {
				return multiplexer.CreateResponse("Bu komut zaten engellenmiş.")
			}
			insert, _ := db.Query(fmt.Sprintf("INSERT INTO disabledcommands (isblocked, commandname, guildid) VALUES ('True', 'stats', '%s')", interaction.GuildID))

			defer insert.Close()

			return multiplexer.CreateResponse("Belirtilen komut başarıyla engellendi")
		case "unban":
			err := db.QueryRow(fmt.Sprintf("SELECT isblocked FROM disabledcommands WHERE commandname ='unban' AND guildid ='%s'", interaction.GuildID)).Scan(&tag.isblocked)
			if err == nil {
				return multiplexer.CreateResponse("Bu komut zaten engellenmiş.")
			}
			insert, _ := db.Query(fmt.Sprintf("INSERT INTO disabledcommands (isblocked, commandname, guildid) VALUES ('True', 'unban', '%s')", interaction.GuildID))

			defer insert.Close()

			return multiplexer.CreateResponse("Belirtilen komut başarıyla engellendi")
		default:
			return multiplexer.CreateResponse("You need to specify a command.")
		}
	}

	if !multiplexer.CheckAdministratorPermission(session, interaction.Member.User.ID, interaction.ChannelID) {
		return multiplexer.CreateResponse("You don't have enough permission")
	}

	args := interaction.Data.Options[0].Value.(string)

	switch args {
	case "afk":
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
	case "author":
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
	case "avatar":
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
	case "ban":
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
	case "embed":
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
	case "guild_info":
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
	case "hug":
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
	case "icon":
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
	case "kick":
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
	case "kiss":
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
	case "ping":
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
	case "roles":
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
	case "settings":
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
	case "slap":
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
	case "spoiler":
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
	case "start_vote":
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
	case "stats":
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
	case "unban":
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
