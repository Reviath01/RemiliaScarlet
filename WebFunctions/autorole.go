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

func AutoRoleHandler(c *gin.Context, session *discordgo.Session, conf *oauth2.Config) {
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
			role, err := session.State.Role(guild.ID, c.Param("roleid"))

			if err != nil {
				c.HTML(200, "error.html", gin.H{
					"is404":       "false",
					"description": "Cannot find role \"" + c.Param("roleid") + "\" on specified guild.",
					"error":       "Invalid role",
				})
				return
			}

			db := sql.Connect()
			defer db.Close()
			insert, _ := db.Query(fmt.Sprintf("INSERT INTO autorole (roleid, guildid) VALUES ('%s', '%s')", role.ID, guild.ID))
			defer insert.Close()

			c.Redirect(http.StatusTemporaryRedirect, "/guild/"+c.Param("guildid"))
		}
	}
}
