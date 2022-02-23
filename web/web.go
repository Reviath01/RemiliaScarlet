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
	server.Static("/guild/assets", "./web/public/assets")

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

	server.GET("/contact", func(c *gin.Context) {
		webfuncs.ContactHandler(c, cli)
	})

	server.GET("/panel", func(c *gin.Context) {
		webfuncs.PanelHandler(c, cli, conf, session)
	})

	server.GET("/guild/:guildid", func(c *gin.Context) {
		webfuncs.GuildHandler(c, session, conf, cli)
	})

	server.POST("/send", webfuncs.SendHandler)
	server.POST("/resetautorole/:guildid", func(c *gin.Context) {
		webfuncs.ResetAutoroleHandler(c, session, conf)
	})

	server.POST("/autorole/:guildid", func(c *gin.Context) {
		webfuncs.AutoroleHandler(c, session, conf)
	})

	server.POST("/leavechannel/:guildid", func(c *gin.Context) {
		webfuncs.LeaveChannelHandler(c, session, conf)
	})

	server.POST("/resetleavechannel/:guildid", func(c *gin.Context) {
		webfuncs.ResetLeaveChannelHandler(c, session, conf)
	})

	server.POST("/leavemessage/:guildid", func(c *gin.Context) {
		webfuncs.LeaveMessageHandler(c, session, conf)
	})

	server.POST("/resetleavemessage/:guildid", func(c *gin.Context) {
		webfuncs.ResetLeaveMessageHandler(c, session, conf)
	})

	server.POST("/log/:guildid", func(c *gin.Context) {
		webfuncs.LogHandler(c, session, conf)
	})

	server.POST("/resetlog/:guildid", func(c *gin.Context) {
		webfuncs.ResetLogHandler(c, session, conf)
	})

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
