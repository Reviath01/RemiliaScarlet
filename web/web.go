package web

import (
	"fmt"
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
	server.Static("/css", "./web/public/css")
	server.Static("/guild/css", "./web/public/css")

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

	server.GET("/support", func(c *gin.Context) {
		webfuncs.SupportHandler(c)
	})

	server.GET("/api/log/:guildid/:newchannel", func(c *gin.Context) {
		webfuncs.LogHandler(c, session, conf)
	})

	server.GET("/api/leavechannel/:guildid/:newchannel", func(c *gin.Context) {
		webfuncs.LeaveChannelHandler(c, session, conf)
	})

	server.GET("/api/welcomechannel/:guildid/:newchannel", func(c *gin.Context) {
		webfuncs.WelcomeChannelHandler(c, session, conf)
	})

	server.GET("/api/autorole/:guildid/:roleid", func(c *gin.Context) {
		webfuncs.AutoRoleHandler(c, session, conf)
	})

	server.GET("/resetlog/:guildid", func(c *gin.Context) {
		webfuncs.ResetLogHandler(c, session, conf)
	})

	server.GET("/resetwelcomechannel/:guildid", func(c *gin.Context) {
		webfuncs.ResetWelcomeChannelHandler(c, session, conf)
	})

	server.GET("/resetwelcomemessage/:guildid", func(c *gin.Context) {
		webfuncs.ResetWelcomeMessageHandler(c, session, conf)
	})

	server.GET("/resetleavemessage/:guildid", func(c *gin.Context) {
		webfuncs.ResetLeaveMessageHandler(c, session, conf)
	})

	server.GET("/resetleavechannel/:guildid", func(c *gin.Context) {
		webfuncs.ResetLeaveChannelHandler(c, session, conf)
	})

	server.GET("/resetautorole/:guildid", func(c *gin.Context) {
		webfuncs.ResetAutoRoleHandler(c, session, conf)
	})

	server.GET("/guild/:guildid", func(c *gin.Context) {
		webfuncs.GuildHandler(c, session, conf, cli)
	})

	server.GET("/", func(c *gin.Context) {
		webfuncs.MainHandler(c, cli, conf, session)
	})

	fmt.Printf("Attempting to run website at \"%s:%s\" \n", config.BaseURL, config.WebPort)
	time.Sleep(1 * time.Second)
	if err := server.Run(":" + config.WebPort); err != nil {
		fmt.Printf("Cannot run website at \"%s\" port, running on 8000.\n", config.WebPort)
		server.Run(":8000")
	}
}
