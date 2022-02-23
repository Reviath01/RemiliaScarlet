package webfuncs

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"git.randomchars.net/Reviath/RemiliaScarlet/sql"
	"github.com/bwmarrin/discordgo"
	"github.com/gin-gonic/gin"
	jsoniter "github.com/json-iterator/go"
	"golang.org/x/oauth2"
)

func ContactHandler(c *gin.Context, cli *discordgo.User, conf *oauth2.Config) {
	val, _ := c.Cookie("key")
	switch val {
	case "":
		c.HTML(200, "contact.html", gin.H{
			"login":       "nil",
			"botavatar":   cli.AvatarURL("1024"),
			"botusername": cli.Username,
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
			panic(err.Error())
		}
		c.HTML(200, "contact.html", gin.H{
			"login":       "yes",
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

func SendHandler(c *gin.Context) {
	r := c.Request
	name := r.FormValue("name")
	message := r.FormValue("message")
	db := sql.Connect()

	defer db.Close()
	insert, err := db.Query("INSERT INTO messages (name, message) VALUES (?, ?)", name, message)
	if err != nil {
		c.Redirect(http.StatusMovedPermanently, "/contact")
	}
	defer insert.Close()
	c.Redirect(http.StatusMovedPermanently, "/contact")
}
