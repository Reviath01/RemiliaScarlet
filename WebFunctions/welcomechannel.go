package webfuncs

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	multiplexer "git.randomchars.net/Reviath/RemiliaScarlet/Multiplexer"
	"git.randomchars.net/Reviath/RemiliaScarlet/sql"
	"github.com/bwmarrin/discordgo"
	"github.com/gin-gonic/gin"
	jsoniter "github.com/json-iterator/go"
	"golang.org/x/oauth2"
)

func WelcomeChannelHandler(c *gin.Context, session *discordgo.Session, conf *oauth2.Config) {
	val, _ := c.Cookie("key")
	switch val {
	case "":
		c.Redirect(http.StatusMovedPermanently, "/login")
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
				c.Redirect(http.StatusMovedPermanently, "/")
			}

			if !multiplexer.CheckAdministratorPermission(session, user.ID, guild.ID) {
				c.HTML(200, "error.html", gin.H{
					"is404":       "false",
					"description": "You don't have enough permission to access here!",
					"error":       "Unauthorized",
				})
				return
			}
			ch := c.Request.FormValue("welcomechannel")
			channel, err := session.State.Channel(ch)
			if err != nil {
				c.HTML(200, "error.html", gin.H{
					"is404":       "false",
					"description": "Cannot find channel \"" + ch + "\"",
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
			c.Redirect(http.StatusMovedPermanently, "/guild/"+c.Param("guildid"))
		}
	}
}

func ResetWelcomeChannelHandler(c *gin.Context, session *discordgo.Session, conf *oauth2.Config) {
	val, _ := c.Cookie("key")
	switch val {
	case "":
		c.Redirect(http.StatusMovedPermanently, "/login")
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
				c.Redirect(http.StatusMovedPermanently, "/")
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
			c.Redirect(http.StatusMovedPermanently, "/guild/"+guild.ID)
		}
	}
}
