package CommandHandler

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"

	commandMap "git.randomchars.net/Reviath/RemiliaScarlet/CommandMap"
	ctx "git.randomchars.net/Reviath/RemiliaScarlet/Context"
	"github.com/bwmarrin/discordgo"
	_ "github.com/go-sql-driver/mysql"
)

func Shift(a []string, i int) ([]string, string) {
	b := a[i]
	copy(a[i:], a[i+1:])
	a[len(a)-1] = ""
	a = a[:len(a)-1]
	return a, b
}

type Handler struct {
	dg      *discordgo.Session
	cmds    *commandMap.Map
	Prefixs string
}

func (h Handler) Handle(s *discordgo.Session, msg *discordgo.MessageCreate) {
	file, err := ioutil.ReadFile("./config.json")

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	type configStruct struct {
		BotPrefix string `json:"BotPrefix"`
	}

	var (
		config *configStruct
	)

	err = json.Unmarshal(file, &config)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	if msg.Author.Bot {
		return
	}

	if msg.Content == "<@!"+s.State.User.ID+">" {
		_, err := s.ChannelMessageSend(msg.ChannelID, "My prefix is "+config.BotPrefix)
		if err != nil {
			return
		}
	}

	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/remilia")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	type Tag struct {
		isafk string `json:"isafk"`
		lang  string `json:"language"`
		xp    string `json:"xp"`
		level string `json:"level"`
	}

	var tag Tag

	err = db.QueryRow("SELECT language FROM languages WHERE guildid ='" + msg.GuildID + "'").Scan(&tag.lang)
	if err == nil {
		if tag.lang == "tr" {
			err = db.QueryRow("SELECT isafk FROM afk WHERE userid ='" + msg.Author.ID + "'").Scan(&tag.isafk)
			if err != nil {
				if strings.HasPrefix(msg.Content, h.Prefixs) {
					args := strings.Split(strings.TrimPrefix(msg.Content, h.Prefixs), " ")
					args, cmd := Shift(args, 0)
					err := h.cmds.Execute(cmd, ctx.New(args, msg, s), s)
					if err != nil {
						_, err = s.ChannelMessageSend(msg.ChannelID, "Bir hata oluştu.")
					}
				}
			} else {
				var level string
				var xp string

				err = db.QueryRow("SELECT level FROM levels WHERE userid ='" + msg.Author.ID + "' AND guildid ='" + msg.GuildID + "'").Scan(&tag.level)

				if err != nil {
					level = "0"
				} else {
					level = tag.level
				}

				err = db.QueryRow("SELECT xp FROM xps WHERE userid ='" + msg.Author.ID + "' AND guildid ='" + msg.GuildID + "'").Scan(&tag.xp)

				if err != nil {
					xp = "0"
				} else {
					xp = tag.xp
				}

				var maxxp string

				if level == "0" {
					maxxp = "150"
				} else if level == "1" {
					maxxp = "200"
				} else if level == "2" {
					maxxp = "250"
				} else if level == "3" {
					maxxp = "300"
				} else if level == "4" {
					maxxp = "250"
				} else if level == "5" {
					maxxp = "300"
				} else if level == "6" {
					maxxp = "350"
				} else if level == "7" {
					maxxp = "400"
				} else if level == "8" {
					maxxp = "500"
				} else if level == "9" {
					maxxp = "600"
				} else if level == "10" {
					maxxp = "700"
				} else if level == "11" {
					maxxp = "800"
				} else if level == "12" {
					maxxp = "900"
				} else if level == "13" {
					maxxp = "1000"
				} else if level == "14" {
					maxxp = "1100"
				} else if level == "15" {
					maxxp = "1200"
				} else if level == "16" {
					maxxp = "1400"
				} else if level == "17" {
					maxxp = "1600"
				} else if level == "18" {
					maxxp = "1800"
				} else if level == "19" {
					maxxp = "2000"
				} else if level == "20" {
					maxxp = "2200"
				} else if level == "21" {
					maxxp = "2400"
				} else if level == "22" {
					maxxp = "2600"
				} else if level == "23" {
					maxxp = "2800"
				} else if level == "24" {
					maxxp = "3000"
				} else if level == "25" {
					maxxp = "3200"
				} else if level == "26" {
					maxxp = "3400"
				} else if level == "27" {
					maxxp = "3600"
				} else if level == "28" {
					maxxp = "3800"
				} else {
					maxxp = "4000"
				}

				if len(msg.Content) > 3 {
					if xp == maxxp {
						update, err := db.Query("UPDATE xps SET xp ='0' WHERE guildid ='" + msg.GuildID + "' AND userid ='" + msg.Author.ID + "'")
						if err != nil {
							return
						}
						defer update.Close()

						newlevel, err := strconv.Atoi(level)

						if err != nil {
							return
						}

						newlevel2 := newlevel + 1

						update2, err := db.Query("UPDATE levels SET level ='" + strconv.Itoa(newlevel2) + "' WHERE guildid ='" + msg.GuildID + "' AND userid ='" + msg.Author.ID + "'")

						if err != nil {
							return
						}

						defer update2.Close()

					}
					insert, err := db.Query("INSERT INTO xps (xp, guildid, userid) VALUES ('3', '" + msg.GuildID + "', " + msg.Author.ID + ")")
					if err != nil {
						return
					}
					insert.Close()
				}
				_, _ = s.ChannelMessageSend(msg.ChannelID, "Tekrar hoş geldin <@"+msg.Author.ID+"> !")
				delete, err := db.Query("DELETE FROM afk WHERE userid ='" + msg.Author.ID + "'")
				if err != nil {
					return
				}
				defer delete.Close()
			}
		}
	}

	err = db.QueryRow("SELECT isafk FROM afk WHERE userid ='" + msg.Author.ID + "'").Scan(&tag.isafk)
	if err != nil {
		var level string
		var xp string

		err = db.QueryRow("SELECT level FROM levels WHERE userid ='" + msg.Author.ID + "' AND guildid ='" + msg.GuildID + "'").Scan(&tag.level)

		if err != nil {
			level = "0"
		} else {
			level = tag.level
		}

		err = db.QueryRow("SELECT xp FROM xps WHERE userid ='" + msg.Author.ID + "' AND guildid ='" + msg.GuildID + "'").Scan(&tag.xp)

		if err != nil {
			xp = "0"
		} else {
			xp = tag.xp
		}

		var maxxp string

		if level == "0" {
			maxxp = "150"
		} else if level == "1" {
			maxxp = "200"
		} else if level == "2" {
			maxxp = "250"
		} else if level == "3" {
			maxxp = "300"
		} else if level == "4" {
			maxxp = "250"
		} else if level == "5" {
			maxxp = "300"
		} else if level == "6" {
			maxxp = "350"
		} else if level == "7" {
			maxxp = "400"
		} else if level == "8" {
			maxxp = "500"
		} else if level == "9" {
			maxxp = "600"
		} else if level == "10" {
			maxxp = "700"
		} else if level == "11" {
			maxxp = "800"
		} else if level == "12" {
			maxxp = "900"
		} else if level == "13" {
			maxxp = "1000"
		} else if level == "14" {
			maxxp = "1100"
		} else if level == "15" {
			maxxp = "1200"
		} else if level == "16" {
			maxxp = "1400"
		} else if level == "17" {
			maxxp = "1600"
		} else if level == "18" {
			maxxp = "1800"
		} else if level == "19" {
			maxxp = "2000"
		} else if level == "20" {
			maxxp = "2200"
		} else if level == "21" {
			maxxp = "2400"
		} else if level == "22" {
			maxxp = "2600"
		} else if level == "23" {
			maxxp = "2800"
		} else if level == "24" {
			maxxp = "3000"
		} else if level == "25" {
			maxxp = "3200"
		} else if level == "26" {
			maxxp = "3400"
		} else if level == "27" {
			maxxp = "3600"
		} else if level == "28" {
			maxxp = "3800"
		} else {
			maxxp = "4000"
		}

		if len(msg.Content) > 3 {
			if xp == maxxp {
				update, err := db.Query("UPDATE xps SET xp ='0' WHERE guildid ='" + msg.GuildID + "' AND userid ='" + msg.Author.ID + "'")
				if err != nil {
					return
				}
				defer update.Close()

				newlevel, err := strconv.Atoi(level)

				if err != nil {
					return
				}

				newlevel2 := newlevel + 1

				update2, err := db.Query("UPDATE levels SET level ='" + strconv.Itoa(newlevel2) + "' WHERE guildid ='" + msg.GuildID + "' AND userid ='" + msg.Author.ID + "'")

				if err != nil {
					return
				}

				defer update2.Close()

			}
			insert, err := db.Query("INSERT INTO xps (xp, guildid, userid) VALUES ('3', '" + msg.GuildID + "', " + msg.Author.ID + ")")
			if err != nil {
				return
			}
			insert.Close()
		}
		if strings.HasPrefix(msg.Content, h.Prefixs) {
			args := strings.Split(strings.TrimPrefix(msg.Content, h.Prefixs), " ")
			args, cmd := Shift(args, 0)
			err := h.cmds.Execute(cmd, ctx.New(args, msg, s), s)
			if err != nil {
				_, err = s.ChannelMessageSend(msg.ChannelID, "An error occurred.")
				if err != nil {
					return
				}
			}
		}
	} else {
		_, _ = s.ChannelMessageSend(msg.ChannelID, "Welcome back <@"+msg.Author.ID+"> , I removed your AFK.")
		delete, err := db.Query("DELETE FROM afk WHERE userid ='" + msg.Author.ID + "'")
		if err != nil {
			return
		}
		defer delete.Close()
	}
}

func (h Handler) GetCommandMap() *commandMap.Map {
	return h.cmds
}

func New(session *discordgo.Session, cmds *commandMap.Map, prefix string) Handler {
	h := Handler{dg: session, cmds: cmds, Prefixs: prefix}
	session.AddHandler(h.Handle)
	return h
}
