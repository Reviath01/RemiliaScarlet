package webfuncs

import (
	"github.com/bwmarrin/discordgo"
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
)

func MainHandler(c *gin.Context, cli *discordgo.User, conf *oauth2.Config, session *discordgo.Session) {
	val, _ := c.Cookie("key")
	switch val {
	case "":
		c.HTML(200, "index.html", gin.H{
			"login":       "nil",
			"botavatar":   cli.AvatarURL("1024"),
			"botusername": cli.Username,
		})
	default:
		c.HTML(200, "index.html", gin.H{
			"login":       "yes",
			"botavatar":   cli.AvatarURL("1024"),
			"botusername": cli.Username,
		})
	}
}
