package web

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	multiplexer "git.randomchars.net/Reviath/RemiliaScarlet/Multiplexer"
	"git.randomchars.net/Reviath/RemiliaScarlet/config"
	"git.randomchars.net/Reviath/RemiliaScarlet/sql"
	"github.com/bwmarrin/discordgo"
	"github.com/gin-gonic/gin"
	jsoniter "github.com/json-iterator/go"
	"github.com/pollen5/discord-oauth2"
	"golang.org/x/oauth2"
)

// Listen returns to web panel that running on specified port
func Listen(session *discordgo.Session) {
	config.ReadConfig()
	gin.SetMode(gin.ReleaseMode)
	var state = "random"
	server := gin.Default()
	server.LoadHTMLGlob("web/public/*.html")
	server.Static("/css", "./web/public/css")
	server.Static("/guild/css", "./web/public/css")

	cli := GetClientUser(session)

	conf := &oauth2.Config{
		RedirectURL:  fmt.Sprintf("%s:%s/callback", config.BaseURL, config.WebPort),
		ClientID:     cli.ID,
		ClientSecret: config.ClientSecret,
		Scopes:       []string{discord.ScopeIdentify, discord.ScopeGuilds},
		Endpoint:     discord.Endpoint,
	}

	server.NoRoute(func(c *gin.Context) {
		c.HTML(404, "error.html", gin.H{
			"is404":       "true",
			"description": "The page you are looking for may have been removed or temporarily unavailable",
			"error":       "404",
		})
	})

	server.GET("/login", func(c *gin.Context) {
		c.Redirect(http.StatusTemporaryRedirect, conf.AuthCodeURL(state))
	})

	server.GET("/callback", func(c *gin.Context) {
		if c.Request.FormValue("state") != state {
			c.Redirect(http.StatusTemporaryRedirect, "/")
		}
		token, err := conf.Exchange(context.TODO(), c.Query("code"))
		if err != nil {
			fmt.Println("An error occurred while getting token: " + err.Error())
			return
		}
		tokenJSON, err := jsoniter.Marshal(token)
		if err != nil {
			fmt.Println(err)
			return
		}
		c.SetCookie("key", string(tokenJSON), 60*60*24*7, "/", "localhost", false, false)
		c.Redirect(http.StatusTemporaryRedirect, "/")
	})

	server.GET("/logout", func(c *gin.Context) {
		val, _ := c.Cookie("key")
		if val != "" {
			c.SetCookie("key", "nil", -1, "/", "localhost", false, false)
			c.Redirect(http.StatusTemporaryRedirect, "/")
		} else {
			c.Redirect(http.StatusTemporaryRedirect, "/")
		}
	})

	server.GET("/invite", func(c *gin.Context) {
		c.Redirect(http.StatusTemporaryRedirect, fmt.Sprintf("https://discord.com/oauth2/authorize?client_id=%s&scope=bot+applications.commands&permissions=8", cli.ID))
	})

	server.GET("/support", func(c *gin.Context) {
		c.Redirect(http.StatusTemporaryRedirect, "https://discord.gg/zVVWWDtSr2")
	})

	server.GET("/api/log/:guildid/:newchannel", func(c *gin.Context) {
		val, _ := c.Cookie("key")
		switch val {
		case "":
			c.Redirect(http.StatusTemporaryRedirect, "/login")
		default:
			guild, err := session.State.Guild(c.Param("guildid"))
			if err != nil {
				c.HTML(200, "error.html", gin.H{
					"is404":       "false",
					"description": "Cannot find guild " + "\"" + c.Param("guildid") + "\"",
					"error":       "Invalid snowflake.",
				})
			} else {
				var token = &oauth2.Token{}
				jsoniter.UnmarshalFromString(fmt.Sprint(val), token)
				res, err := conf.Client(context.TODO(), token).Get("https://discordapp.com/api/v8/users/@me")

				if err != nil || res.StatusCode != 200 {
					fmt.Println("An error occurred on api: " + err.Error())
					return
				}

				var user discordgo.User
				data, _ := ioutil.ReadAll(res.Body)
				err = json.Unmarshal(data, &user)
				if err != nil {
					c.Redirect(http.StatusTemporaryRedirect, "/")
				}

				if !multiplexer.CheckAdministratorPermission(session, user.ID, guild.ID) {
					c.HTML(200, "error.html", gin.H{
						"is404":       "false",
						"description": "You don't have enough permission to access here!",
						"error":       "Unauthorized",
					})
					return
				}
				channel, err := session.State.Channel(c.Param("newchannel"))
				if err != nil {
					c.HTML(200, "error.html", gin.H{
						"is404":       "false",
						"description": "Cannot find channel \"" + c.Param("newchannel") + "\"",
						"error":       "Invalid channel",
					})
					return
				}
				if channel.GuildID == guild.ID {
					db := sql.Connect()
					defer db.Close()
					insert, _ := db.Query(fmt.Sprintf("INSERT INTO log (channelid, guildid) VALUES ('%s', '%s')", channel.ID, guild.ID))
					defer insert.Close()
				} else {
					c.HTML(200, "error.html", gin.H{
						"is404":       "false",
						"description": "This channel is not in this guild.",
						"error":       "Invalid channel",
					})
					return
				}
				c.Redirect(http.StatusTemporaryRedirect, "/guild/"+c.Param("guildid"))
			}
		}
	})

	server.GET("/api/leavechannel/:guildid/:newchannel", func(c *gin.Context) {
		val, _ := c.Cookie("key")
		switch val {
		case "":
			c.Redirect(http.StatusTemporaryRedirect, "/login")
		default:
			guild, err := session.State.Guild(c.Param("guildid"))
			if err != nil {
				c.HTML(200, "error.html", gin.H{
					"is404":       "false",
					"description": "Cannot find guild " + "\"" + c.Param("guildid") + "\"",
					"error":       "Invalid snowflake.",
				})
			} else {
				var token = &oauth2.Token{}
				jsoniter.UnmarshalFromString(fmt.Sprint(val), token)
				res, err := conf.Client(context.TODO(), token).Get("https://discordapp.com/api/v8/users/@me")

				if err != nil || res.StatusCode != 200 {
					fmt.Println("An error occurred on api: " + err.Error())
					return
				}

				var user discordgo.User
				data, _ := ioutil.ReadAll(res.Body)
				err = json.Unmarshal(data, &user)
				if err != nil {
					c.Redirect(http.StatusTemporaryRedirect, "/")
				}

				if !multiplexer.CheckAdministratorPermission(session, user.ID, guild.ID) {
					c.HTML(200, "error.html", gin.H{
						"is404":       "false",
						"description": "You don't have enough permission to access here!",
						"error":       "Unauthorized",
					})
					return
				}
				channel, err := session.State.Channel(c.Param("newchannel"))
				if err != nil {
					c.HTML(200, "error.html", gin.H{
						"is404":       "false",
						"description": "Cannot find channel \"" + c.Param("newchannel") + "\"",
						"error":       "Invalid channel",
					})
					return
				}
				if channel.GuildID == guild.ID {
					db := sql.Connect()
					defer db.Close()
					insert, _ := db.Query(fmt.Sprintf("INSERT INTO leavechannel (channelid, guildid) VALUES ('%s', '%s')", channel.ID, guild.ID))
					defer insert.Close()
				} else {
					c.HTML(200, "error.html", gin.H{
						"is404":       "false",
						"description": "This channel is not in this guild.",
						"error":       "Invalid channel",
					})
					return
				}
				c.Redirect(http.StatusTemporaryRedirect, "/guild/"+c.Param("guildid"))
			}
		}
	})

	server.GET("/api/welcomechannel/:guildid/:newchannel", func(c *gin.Context) {
		val, _ := c.Cookie("key")
		switch val {
		case "":
			c.Redirect(http.StatusTemporaryRedirect, "/login")
		default:
			guild, err := session.State.Guild(c.Param("guildid"))
			if err != nil {
				c.HTML(200, "error.html", gin.H{
					"is404":       "false",
					"description": "Cannot find guild " + "\"" + c.Param("guildid") + "\"",
					"error":       "Invalid snowflake.",
				})
			} else {
				var token = &oauth2.Token{}
				jsoniter.UnmarshalFromString(fmt.Sprint(val), token)
				res, err := conf.Client(context.TODO(), token).Get("https://discordapp.com/api/v8/users/@me")

				if err != nil || res.StatusCode != 200 {
					fmt.Println("An error occurred on api: " + err.Error())
					return
				}

				var user discordgo.User
				data, _ := ioutil.ReadAll(res.Body)
				err = json.Unmarshal(data, &user)
				if err != nil {
					c.Redirect(http.StatusTemporaryRedirect, "/")
				}

				if !multiplexer.CheckAdministratorPermission(session, user.ID, guild.ID) {
					c.HTML(200, "error.html", gin.H{
						"is404":       "false",
						"description": "You don't have enough permission to access here!",
						"error":       "Unauthorized",
					})
					return
				}
				channel, err := session.State.Channel(c.Param("newchannel"))
				if err != nil {
					c.HTML(200, "error.html", gin.H{
						"is404":       "false",
						"description": "Cannot find channel \"" + c.Param("newchannel") + "\"",
						"error":       "Invalid channel",
					})
					return
				}
				if channel.GuildID == guild.ID {
					db := sql.Connect()
					defer db.Close()
					insert, _ := db.Query(fmt.Sprintf("INSERT INTO welcomechannel (channelid, guildid) VALUES ('%s', '%s')", channel.ID, guild.ID))
					defer insert.Close()
				} else {
					c.HTML(200, "error.html", gin.H{
						"is404":       "false",
						"description": "This channel is not in this guild.",
						"error":       "Invalid channel",
					})
					return
				}
				c.Redirect(http.StatusTemporaryRedirect, "/guild/"+c.Param("guildid"))
			}
		}
	})

	server.GET("/resetlog/:guildid", func(c *gin.Context) {
		val, _ := c.Cookie("key")
		switch val {
		case "":
			c.Redirect(http.StatusTemporaryRedirect, "/login")
		default:
			guild, err := session.State.Guild(c.Param("guildid"))
			if err != nil {
				c.HTML(200, "error.html", gin.H{
					"is404":       "false",
					"description": "Cannot find guild " + "\"" + c.Param("guildid") + "\"",
					"error":       "Invalid snowflake.",
				})
			} else {
				var token = &oauth2.Token{}
				jsoniter.UnmarshalFromString(fmt.Sprint(val), token)
				res, err := conf.Client(context.TODO(), token).Get("https://discordapp.com/api/v8/users/@me")

				if err != nil || res.StatusCode != 200 {
					fmt.Println("An error occurred on api: " + err.Error())
					return
				}

				var user discordgo.User
				data, _ := ioutil.ReadAll(res.Body)
				err = json.Unmarshal(data, &user)
				if err != nil {
					c.Redirect(http.StatusTemporaryRedirect, "/")
				}

				if !multiplexer.CheckAdministratorPermission(session, user.ID, guild.ID) {
					c.HTML(200, "error.html", gin.H{
						"is404":       "false",
						"description": "You don't have enough permission to access here!",
						"error":       "Unauthorized",
					})
					return
				}
				db := sql.Connect()
				defer db.Close()
				delete, _ := db.Query(fmt.Sprintf("DELETE FROM log WHERE guildid ='%s'", guild.ID))

				defer delete.Close()
				c.Redirect(http.StatusTemporaryRedirect, "/guild/"+guild.ID)
			}
		}
	})

	server.GET("/resetwelcomechannel/:guildid", func(c *gin.Context) {
		val, _ := c.Cookie("key")
		switch val {
		case "":
			c.Redirect(http.StatusTemporaryRedirect, "/login")
		default:
			guild, err := session.State.Guild(c.Param("guildid"))
			if err != nil {
				c.HTML(200, "error.html", gin.H{
					"is404":       "false",
					"description": "Cannot find guild " + "\"" + c.Param("guildid") + "\"",
					"error":       "Invalid snowflake.",
				})
			} else {
				var token = &oauth2.Token{}
				jsoniter.UnmarshalFromString(fmt.Sprint(val), token)
				res, err := conf.Client(context.TODO(), token).Get("https://discordapp.com/api/v8/users/@me")

				if err != nil || res.StatusCode != 200 {
					fmt.Println("An error occurred on api: " + err.Error())
					return
				}

				var user discordgo.User
				data, _ := ioutil.ReadAll(res.Body)
				err = json.Unmarshal(data, &user)
				if err != nil {
					c.Redirect(http.StatusTemporaryRedirect, "/")
				}

				if !multiplexer.CheckAdministratorPermission(session, user.ID, guild.ID) {
					c.HTML(200, "error.html", gin.H{
						"is404":       "false",
						"description": "You don't have enough permission to access here!",
						"error":       "Unauthorized",
					})
					return
				}
				db := sql.Connect()
				defer db.Close()
				delete, _ := db.Query(fmt.Sprintf("DELETE FROM welcomechannel WHERE guildid ='%s'", guild.ID))

				defer delete.Close()
				c.Redirect(http.StatusTemporaryRedirect, "/guild/"+guild.ID)
			}
		}
	})

	server.GET("/resetwelcomemessage/:guildid", func(c *gin.Context) {
		val, _ := c.Cookie("key")
		switch val {
		case "":
			c.Redirect(http.StatusTemporaryRedirect, "/login")
		default:
			guild, err := session.State.Guild(c.Param("guildid"))
			if err != nil {
				c.HTML(200, "error.html", gin.H{
					"is404":       "false",
					"description": "Cannot find guild " + "\"" + c.Param("guildid") + "\"",
					"error":       "Invalid snowflake.",
				})
			} else {
				var token = &oauth2.Token{}
				jsoniter.UnmarshalFromString(fmt.Sprint(val), token)
				res, err := conf.Client(context.TODO(), token).Get("https://discordapp.com/api/v8/users/@me")

				if err != nil || res.StatusCode != 200 {
					fmt.Println("An error occurred on api: " + err.Error())
					return
				}

				var user discordgo.User
				data, _ := ioutil.ReadAll(res.Body)
				err = json.Unmarshal(data, &user)
				if err != nil {
					c.Redirect(http.StatusTemporaryRedirect, "/")
				}

				if !multiplexer.CheckAdministratorPermission(session, user.ID, guild.ID) {
					c.HTML(200, "error.html", gin.H{
						"is404":       "false",
						"description": "You don't have enough permission to access here!",
						"error":       "Unauthorized",
					})
					return
				}
				db := sql.Connect()
				defer db.Close()
				delete, _ := db.Query(fmt.Sprintf("DELETE FROM welcomemessage WHERE guildid ='%s'", guild.ID))

				defer delete.Close()
				c.Redirect(http.StatusTemporaryRedirect, "/guild/"+guild.ID)
			}
		}
	})

	server.GET("/resetleavemessage/:guildid", func(c *gin.Context) {
		val, _ := c.Cookie("key")
		switch val {
		case "":
			c.Redirect(http.StatusTemporaryRedirect, "/login")
		default:
			guild, err := session.State.Guild(c.Param("guildid"))
			if err != nil {
				c.HTML(200, "error.html", gin.H{
					"is404":       "false",
					"description": "Cannot find guild " + "\"" + c.Param("guildid") + "\"",
					"error":       "Invalid snowflake.",
				})
			} else {
				var token = &oauth2.Token{}
				jsoniter.UnmarshalFromString(fmt.Sprint(val), token)
				res, err := conf.Client(context.TODO(), token).Get("https://discordapp.com/api/v8/users/@me")

				if err != nil || res.StatusCode != 200 {
					fmt.Println("An error occurred on api: " + err.Error())
					return
				}

				var user discordgo.User
				data, _ := ioutil.ReadAll(res.Body)
				err = json.Unmarshal(data, &user)
				if err != nil {
					c.Redirect(http.StatusTemporaryRedirect, "/")
				}

				if !multiplexer.CheckAdministratorPermission(session, user.ID, guild.ID) {
					c.HTML(200, "error.html", gin.H{
						"is404":       "false",
						"description": "You don't have enough permission to access here!",
						"error":       "Unauthorized",
					})
					return
				}
				db := sql.Connect()
				defer db.Close()
				delete, _ := db.Query(fmt.Sprintf("DELETE FROM leavemessage WHERE guildid ='%s'", guild.ID))

				defer delete.Close()
				c.Redirect(http.StatusTemporaryRedirect, "/guild/"+guild.ID)
			}
		}
	})

	server.GET("/resetleavechannel/:guildid", func(c *gin.Context) {
		val, _ := c.Cookie("key")
		switch val {
		case "":
			c.Redirect(http.StatusTemporaryRedirect, "/login")
		default:
			guild, err := session.State.Guild(c.Param("guildid"))
			if err != nil {
				c.HTML(200, "error.html", gin.H{
					"is404":       "false",
					"description": "Cannot find guild " + "\"" + c.Param("guildid") + "\"",
					"error":       "Invalid snowflake.",
				})
			} else {
				var token = &oauth2.Token{}
				jsoniter.UnmarshalFromString(fmt.Sprint(val), token)
				res, err := conf.Client(context.TODO(), token).Get("https://discordapp.com/api/v8/users/@me")

				if err != nil || res.StatusCode != 200 {
					fmt.Println("An error occurred on api: " + err.Error())
					return
				}

				var user discordgo.User
				data, _ := ioutil.ReadAll(res.Body)
				err = json.Unmarshal(data, &user)
				if err != nil {
					c.Redirect(http.StatusTemporaryRedirect, "/")
				}

				if !multiplexer.CheckAdministratorPermission(session, user.ID, guild.ID) {
					c.HTML(200, "error.html", gin.H{
						"is404":       "false",
						"description": "You don't have enough permission to access here!",
						"error":       "Unauthorized",
					})
					return
				}
				db := sql.Connect()
				defer db.Close()
				delete, _ := db.Query(fmt.Sprintf("DELETE FROM leavechannel WHERE guildid ='%s'", guild.ID))

				defer delete.Close()
				c.Redirect(http.StatusTemporaryRedirect, "/guild/"+guild.ID)
			}
		}
	})

	server.GET("/resetautorole/:guildid", func(c *gin.Context) {
		val, _ := c.Cookie("key")
		switch val {
		case "":
			c.Redirect(http.StatusTemporaryRedirect, "/login")
		default:
			guild, err := session.State.Guild(c.Param("guildid"))
			if err != nil {
				c.HTML(200, "error.html", gin.H{
					"is404":       "false",
					"description": "Cannot find guild " + "\"" + c.Param("guildid") + "\"",
					"error":       "Invalid snowflake.",
				})
			} else {
				var token = &oauth2.Token{}
				jsoniter.UnmarshalFromString(fmt.Sprint(val), token)
				res, err := conf.Client(context.TODO(), token).Get("https://discordapp.com/api/v8/users/@me")

				if err != nil || res.StatusCode != 200 {
					fmt.Println("An error occurred on api: " + err.Error())
					return
				}

				var user discordgo.User
				data, _ := ioutil.ReadAll(res.Body)
				err = json.Unmarshal(data, &user)
				if err != nil {
					c.Redirect(http.StatusTemporaryRedirect, "/")
				}

				if !multiplexer.CheckAdministratorPermission(session, user.ID, guild.ID) {
					c.HTML(200, "error.html", gin.H{
						"is404":       "false",
						"description": "You don't have enough permission to access here!",
						"error":       "Unauthorized",
					})
					return
				}
				db := sql.Connect()
				defer db.Close()
				delete, _ := db.Query(fmt.Sprintf("DELETE FROM autorole WHERE guildid ='%s'", guild.ID))

				defer delete.Close()
				c.Redirect(http.StatusTemporaryRedirect, "/guild/"+guild.ID)
			}
		}
	})

	server.GET("/guild/:guildid", func(c *gin.Context) {
		val, _ := c.Cookie("key")
		switch val {
		case "":
			c.Redirect(http.StatusTemporaryRedirect, "/login")
		default:
			guild, err := session.State.Guild(c.Param("guildid"))
			if err != nil {
				c.HTML(200, "error.html", gin.H{
					"is404":       "false",
					"description": "Cannot find guild " + "\"" + c.Param("guildid") + "\"",
					"error":       "Invalid snowflake.",
				})
			} else {
				var token = &oauth2.Token{}
				jsoniter.UnmarshalFromString(fmt.Sprint(val), token)
				res, err := conf.Client(context.TODO(), token).Get("https://discordapp.com/api/v8/users/@me")

				if err != nil || res.StatusCode != 200 {
					fmt.Println("An error occurred on api: " + err.Error())
					return
				}

				var user discordgo.User
				data, _ := ioutil.ReadAll(res.Body)
				err = json.Unmarshal(data, &user)
				if err != nil {
					c.Redirect(http.StatusTemporaryRedirect, "/")
				}

				if !multiplexer.CheckAdministratorPermission(session, user.ID, guild.ID) {
					c.HTML(200, "error.html", gin.H{
						"is404":       "false",
						"description": "You don't have enough permission to access here!",
						"error":       "Unauthorized",
					})
					return
				}

				db := sql.Connect()

				defer db.Close()

				var scansqldata ScanSQLData
				var welcomechannel string
				var leavechannel string
				var autorole string
				var leavemsg string
				var welcomemsg string
				var logchannel string

				err = db.QueryRow(fmt.Sprintf("SELECT channelid FROM welcomechannel WHERE guildid ='%s'", guild.ID)).Scan(&scansqldata.WelcomeChannelID)

				if err == nil {
					welcomechannel = scansqldata.WelcomeChannelID
				} else {
					welcomechannel = "nil"
				}

				err = db.QueryRow(fmt.Sprintf("SELECT channelid FROM leavechannel WHERE guildid ='%s'", guild.ID)).Scan(&scansqldata.LeaveChannelID)

				if err == nil {
					leavechannel = scansqldata.LeaveChannelID
				} else {
					leavechannel = "nil"
				}

				err = db.QueryRow(fmt.Sprintf("SELECT roleid FROM autorole WHERE guildid ='%s'", guild.ID)).Scan(&scansqldata.RoleID)

				if err == nil {
					autorole = scansqldata.RoleID
				} else {
					autorole = "nil"
				}

				err = db.QueryRow(fmt.Sprintf("SELECT message FROM welcomemessage WHERE guildid ='%s'", guild.ID)).Scan(&scansqldata.WelcomeMessage)

				if err == nil {
					welcomemsg = scansqldata.WelcomeMessage
				} else {
					welcomemsg = "nil"
				}

				err = db.QueryRow(fmt.Sprintf("SELECT message FROM leavemessage WHERE guildid ='%s'", guild.ID)).Scan(&scansqldata.LeaveMessage)

				if err == nil {
					leavemsg = scansqldata.LeaveMessage
				} else {
					leavemsg = "nil"
				}

				err = db.QueryRow(fmt.Sprintf("SELECT channelid FROM log WHERE guildid ='%s'", guild.ID)).Scan(&scansqldata.LogID)

				if err == nil {
					logchannel = scansqldata.LogID
				} else {
					logchannel = "nil"
				}

				c.HTML(200, "guild.html", gin.H{
					"guild": guild,
					"settings": Settings{
						WelcomeChannelID: welcomechannel,
						WelcomeMessage:   welcomemsg,
						LeaveChannelID:   leavechannel,
						LeaveMessage:     leavemsg,
						LogID:            logchannel,
						RoleID:           autorole,
					},
					"botavatar":   cli.AvatarURL("1024"),
					"botusername": cli.Username,
				})
			}
		}
	})

	server.GET("/", func(c *gin.Context) {
		val, _ := c.Cookie("key")
		switch val {
		case "":
			c.HTML(200, "index.html", gin.H{
				"login":       "nil",
				"botavatar":   cli.AvatarURL("1024"),
				"botusername": cli.Username,
				"botlink":     fmt.Sprintf("https://discord.com/users/%s", cli.ID),
				"botid":       cli.ID,
			})
		default:
			var token = &oauth2.Token{}
			jsoniter.UnmarshalFromString(fmt.Sprint(val), token)
			res, err := conf.Client(context.TODO(), token).Get("https://discordapp.com/api/v8/users/@me")

			if err != nil || res.StatusCode != 200 {
				fmt.Println("An error occurred on api: " + err.Error())
				return
			}

			var user discordgo.User
			data, _ := ioutil.ReadAll(res.Body)
			err = json.Unmarshal(data, &user)
			if err != nil {
				fmt.Println("An error occurred on api: " + err.Error())
				return
			}
			res2, err := conf.Client(context.TODO(), token).Get("https://discordapp.com/api/v8/users/@me/guilds")

			if err != nil || res2.StatusCode != 200 {
				fmt.Println("An error occurred on api: " + err.Error())
				return
			}

			var allguilds []discordgo.Guild
			var guilds []discordgo.Guild

			data2, _ := ioutil.ReadAll(res2.Body)
			err = json.Unmarshal(data2, &allguilds)
			if err != nil {
				fmt.Println("An error occurred on api: " + err.Error())
				return
			}

			for _, g := range allguilds {
				_, err := session.State.Guild(g.ID)
				if err == nil {
					guilds = append(guilds, g)
				}
			}

			c.HTML(200, "index.html", gin.H{
				"login": "yes",
				"user": UserInfo{
					Name:          user.Username,
					ID:            user.ID,
					AvatarURL:     user.AvatarURL("4096"),
					Discriminator: user.Discriminator,
					Bot:           user.Bot,
				},
				"guilds":      guilds,
				"botavatar":   cli.AvatarURL("1024"),
				"botusername": cli.Username,
				"botlink":     fmt.Sprintf("https://discord.com/users/%s", cli.ID),
				"botid":       cli.ID,
			})
		}
	})

	fmt.Printf("Attempting to run website at \"%s:%s\" \n", config.BaseURL, config.WebPort)
	time.Sleep(1 * time.Second)
	if err := server.Run(":" + config.WebPort); err != nil {
		fmt.Printf("Cannot run website at \"%s\" port, running on 8000.\n", config.WebPort)
		server.Run(":8000")
	}
}
