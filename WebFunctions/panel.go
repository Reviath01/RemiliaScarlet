package webfuncs

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	multiplexer "git.randomchars.net/Reviath/RemiliaScarlet/Multiplexer"
	"git.randomchars.net/Reviath/RemiliaScarlet/config"
	"github.com/bwmarrin/discordgo"
	"github.com/gin-gonic/gin"
	jsoniter "github.com/json-iterator/go"
	"golang.org/x/oauth2"
)

func PanelHandler(c *gin.Context, cli *discordgo.User, conf *oauth2.Config, session *discordgo.Session) {
	config.ReadConfig()
	val, _ := c.Cookie("key")
	switch val {
	case "":
		c.Redirect(http.StatusTemporaryRedirect, "/login")
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
		var guilds2 []discordgo.Guild

		data2, _ := ioutil.ReadAll(res2.Body)
		err = json.Unmarshal(data2, &allguilds)
		if err != nil {
			fmt.Println("An error occurred on api: " + err.Error())
			return
		}

		for _, g := range allguilds {
			_, err := session.State.Guild(g.ID)
			if err == nil {
				if multiplexer.CheckAdministratorPermission(session, user.ID, g.ID) {
					guilds = append(guilds, g)
				}
			}
		}

		for _, g := range allguilds {
			_, err := session.State.Guild(g.ID)
			p := g.Permissions
			if p&discordgo.PermissionAdministrator == discordgo.PermissionAdministrator {
				if err != nil {
					guilds2 = append(guilds2, g)
				}
			}
		}

		c.HTML(200, "panel.html", gin.H{
			"login": "yes",
			"user": UserInfo{
				Name:          user.Username,
				ID:            user.ID,
				AvatarURL:     user.AvatarURL("4096"),
				Discriminator: user.Discriminator,
				Bot:           user.Bot,
			},
			"guilds":      guilds,
			"guilds2":     guilds2,
			"botavatar":   cli.AvatarURL("1024"),
			"botusername": cli.Username,
			"botlink":     fmt.Sprintf("https://discord.com/users/%s", cli.ID),
			"botid":       cli.ID,
			"owner":       config.Owner,
		})
	}
}
