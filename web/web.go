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

func init() {
	gin.SetMode(gin.ReleaseMode)
}

var (
	Reset  = "\033[0m"
	Green  = "\033[32m"
	Yellow = "\033[43m"
	Red    = "\033[41m"
	server *gin.Engine
)

// Listen returns to web panel that running on specified port
func Listen(session *discordgo.Session) {
	config.ReadConfig()
	var state = uuid.New().String()
	server = gin.New()
	server.LoadHTMLGlob("web/public/*.html")
	server.Static("/assets", "./web/public/assets")
	server.Static("/guild/assets", "./web/public/assets")
	server.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf(Red+"Web:    "+Reset+"| %v |"+Yellow+" %3d "+Reset+"| %13v | %15s |"+Green+" %-7s  "+Reset+"\"%s\" \n",
			param.TimeStamp.Format("2006/01/02 - 15:04:05"),
			param.StatusCode,
			param.Latency,
			param.ClientIP,
			param.Method,
			param.Path,
		)
	}))
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
		webfuncs.ContactHandler(c, cli, conf)
	})

	server.GET("/panel", func(c *gin.Context) {
		webfuncs.PanelHandler(c, cli, conf, session)
	})

	server.GET("/guild/:guildid", func(c *gin.Context) {
		webfuncs.GuildHandler(c, session, conf, cli)
	})

	server.GET("/admin", func(c *gin.Context) {
		webfuncs.AdminHandler(c, conf, cli)
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

	server.POST("/welcomechannel/:guildid", func(c *gin.Context) {
		webfuncs.WelcomeChannelHandler(c, session, conf)
	})

	server.POST("/resetwelcomechannel/:guildid", func(c *gin.Context) {
		webfuncs.ResetWelcomeChannelHandler(c, session, conf)
	})

	server.POST("/welcomemessage/:guildid", func(c *gin.Context) {
		webfuncs.WelcomeMessageHandler(c, session, conf)
	})

	server.POST("/resetwelcomemessage/:guildid", func(c *gin.Context) {
		webfuncs.ResetWelcomeMessageHandler(c, session, conf)
	})

	server.GET("/", func(c *gin.Context) {
		webfuncs.MainHandler(c, cli, conf)
	})

	fmt.Printf("Attempting to run website at \"%s:%s\" \n", config.BaseURL, config.WebPort)
	time.Sleep(1 * time.Second)
	if err := server.Run(":" + config.WebPort); err != nil {
		fmt.Printf("Cannot run website at \"%s\" port.\n", config.WebPort)
		os.Exit(0)
	}
}
