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

func GuildHandler(c *gin.Context, session *discordgo.Session, conf *oauth2.Config, cli *discordgo.User) {
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

			var channels []*discordgo.Channel

			for _, channel := range guild.Channels {
				if channel.Type == 0 || channel.Type == 5 {
					channels = append(channels, channel)
				}
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
				"channels": channels,
				"user": UserInfo{
					Name:          user.Username,
					ID:            user.ID,
					AvatarURL:     user.AvatarURL("4096"),
					Discriminator: user.Discriminator,
					Bot:           user.Bot,
				},
				"botavatar":   cli.AvatarURL("1024"),
				"botusername": cli.Username,
			})
		}
	}
}
