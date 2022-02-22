package webfuncs

import (
	"net/http"

	"git.randomchars.net/Reviath/RemiliaScarlet/sql"
	"github.com/bwmarrin/discordgo"
	"github.com/gin-gonic/gin"
)

func ContactHandler(c *gin.Context, cli *discordgo.User) {
	val, _ := c.Cookie("key")
	switch val {
	case "":
		c.HTML(200, "contact.html", gin.H{
			"login":       "nil",
			"botavatar":   cli.AvatarURL("1024"),
			"botusername": cli.Username,
		})
	default:
		c.HTML(200, "contact.html", gin.H{
			"login":       "yes",
			"botavatar":   cli.AvatarURL("1024"),
			"botusername": cli.Username,
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
