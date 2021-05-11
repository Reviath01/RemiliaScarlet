package web

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"git.randomchars.net/Reviath/RemiliaScarlet/config"
	"github.com/bwmarrin/discordgo"
	"github.com/gin-gonic/gin"
	jsoniter "github.com/json-iterator/go"
	"github.com/pollen5/discord-oauth2"
	"golang.org/x/oauth2"
)

func Listen(session *discordgo.Session) {
	config.ReadConfig()
	gin.SetMode(gin.ReleaseMode)
	var state = "random"
	server := gin.Default()
	server.LoadHTMLGlob("web/public/*.html")
	server.Static("/css", "./web/public/css")

	cli, err := session.User(config.ClientID)
	if err != nil {
		fmt.Println("An error occurred while getting client user on web panel (please be sure that you wrote ClientID correct)")
		time.Sleep(1 * time.Second)
		fmt.Println(err.Error())
		os.Exit(1)
	}
	avatarURL := cli.AvatarURL("1024")
	userName := cli.Username
	botID := cli.ID

	conf := &oauth2.Config{
		RedirectURL:  fmt.Sprintf("%s/callback", config.WebURL),
		ClientID:     botID,
		ClientSecret: config.ClientSecret,
		Scopes:       []string{discord.ScopeIdentify, discord.ScopeGuilds},
		Endpoint:     discord.Endpoint,
	}

	server.NoRoute(func(c *gin.Context) {
		c.HTML(200, "404.html", gin.H{})
	})

	server.GET("/login", func(c *gin.Context) {
		c.Redirect(http.StatusTemporaryRedirect, conf.AuthCodeURL(state))
	})

	server.GET("/callback", func(c *gin.Context) {
		if c.Request.FormValue("state") != state {
			c.HTML(200, "index.html", gin.H{
				"user":        "nil",
				"botavatar":   avatarURL,
				"botusername": userName,
				"botlink":     fmt.Sprintf("https://discord.com/users/%s", botID),
				"botid":       botID,
			})
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
			c.HTML(200, "index.html", gin.H{
				"login":       "nil",
				"botavatar":   avatarURL,
				"botusername": userName,
				"botlink":     fmt.Sprintf("https://discord.com/users/%s", botID),
				"botid":       botID,
			})
		}
	})

	server.GET("/invite", func(c *gin.Context) {
		c.Redirect(http.StatusTemporaryRedirect, fmt.Sprintf("https://discord.com/oauth2/authorize?client_id=%s&scope=bot+applications.commands&permissions=8", config.ClientID))
	})

	server.GET("/support", func(c *gin.Context) {
		c.Redirect(http.StatusTemporaryRedirect, "https://discord.gg/zVVWWDtSr2")
	})

	server.GET("/", func(c *gin.Context) {
		val, _ := c.Cookie("key")
		switch val {
		case "":
			c.HTML(200, "index.html", gin.H{
				"login":       "nil",
				"botavatar":   avatarURL,
				"botusername": userName,
				"botlink":     fmt.Sprintf("https://discord.com/users/%s", botID),
				"botid":       botID,
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

			var guilds []discordgo.Guild

			data2, _ := ioutil.ReadAll(res2.Body)
			err = json.Unmarshal(data2, &guilds)
			if err != nil {
				fmt.Println("An error occurred on api: " + err.Error())
				return
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
				"botavatar":   avatarURL,
				"botusername": userName,
				"botlink":     fmt.Sprintf("https://discord.com/users/%s", botID),
				"botid":       botID,
			})
		}
	})

	fmt.Printf("Attempting to run website at \"%s\" \n", config.WebURL)
	time.Sleep(1 * time.Second)
	if err := server.Run(config.WebURL[7:]); err != nil {
		fmt.Printf("Cannot run website at \"%s\", running on http://localhost:8000.\n", config.WebURL)
		server.Run(":8000")
	}
}
