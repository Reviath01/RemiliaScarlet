package webfuncs

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"git.randomchars.net/Reviath/RemiliaScarlet/config"
	"git.randomchars.net/Reviath/RemiliaScarlet/sql"
	"github.com/bwmarrin/discordgo"
	"github.com/gin-gonic/gin"
	jsoniter "github.com/json-iterator/go"
	"golang.org/x/oauth2"
)

func AdminHandler(c *gin.Context, conf *oauth2.Config, cli *discordgo.User) {
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
			c.Redirect(http.StatusTemporaryRedirect, "/")
		}

		if user.ID != config.Owner {
			c.HTML(200, "error.html", gin.H{
				"is404":       "false",
				"description": "You don't have enough permission to access here!",
				"error":       "Unauthorized",
			})
			return
		}

		db := sql.Connect()
		defer db.Close()
		row, err := db.Query("SELECT message, name FROM messages")
		if err != nil {
			c.HTML(200, "error.html", gin.H{
				"is404":       "false",
				"description": "Unexpected error on database!",
				"error":       "Database",
			})
		}
		defer row.Close()

		type Row struct {
			Content string
			Sender  string
		}

		type RowCollection struct {
			Rows []Row
		}

		var1 := RowCollection{}
		var2 := Row{}

		for row.Next() {
			err = row.Scan(&var2.Content, &var2.Sender)
			if err != nil {
				panic(err.Error())
			}
			var1.Rows = append(var1.Rows, var2)
		}
		c.HTML(200, "admin.html", gin.H{
			"messages":    var1.Rows,
			"botavatar":   cli.AvatarURL("1024"),
			"botusername": cli.Username,
			"user": UserInfo{
				Name:          user.Username,
				ID:            user.ID,
				AvatarURL:     user.AvatarURL("4096"),
				Discriminator: user.Discriminator,
				Bot:           user.Bot,
			},
		})
	}
}
