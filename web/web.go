package web

import (
	"fmt"
	"os"
	"time"

	webfuncs "git.randomchars.net/Reviath/RemiliaScarlet/WebFunctions"
	"git.randomchars.net/Reviath/RemiliaScarlet/config"
	"github.com/bwmarrin/discordgo"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/oauth2"
)

// Listen returns to web panel that running on specified port
func Listen(session *discordgo.Session) {
	config.ReadConfig()
	gin.SetMode(gin.ReleaseMode)
	var state = uuid.New().String()
	server := gin.Default()
	server.LoadHTMLGlob("web/public/*.html")
	server.Static("/assets", "./web/public/assets")

	cli := GetClientUser(session)

	conf := &oauth2.Config{
		RedirectURL:  fmt.Sprintf("%s:%s/callback", config.BaseURL, config.WebPort),
		ClientID:     cli.ID,
		ClientSecret: config.ClientSecret,
		Scopes:       []string{"identify", "guilds"},
		Endpoint:     Endpoint,
	}

	server.NoRoute(func(c *gin.Context) {
		webfuncs.NoRouteHandler(c)
	})

	server.GET("/login", func(c *gin.Context) {
		webfuncs.LoginHandler(c, conf, state)
	})

	server.GET("/callback", func(c *gin.Context) {
		webfuncs.CallbackHandler(c, state, conf)
	})

	server.GET("/logout", func(c *gin.Context) {
		webfuncs.LogoutHandler(c)
	})

	server.GET("/invite", func(c *gin.Context) {
		webfuncs.InviteHandler(c, cli)
	})

	server.GET("/contact", webfuncs.ContactHandler)

	server.POST("/send", webfuncs.SendHandler)

	server.GET("/", func(c *gin.Context) {
		webfuncs.MainHandler(c, cli, conf, session)
	})

	fmt.Printf("Attempting to run website at \"%s:%s\" \n", config.BaseURL, config.WebPort)
	time.Sleep(1 * time.Second)
	if err := server.Run(":" + config.WebPort); err != nil {
		fmt.Printf("Cannot run website at \"%s\" port.\n", config.WebPort)
		os.Exit(0)
	}
}
