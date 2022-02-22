package webfuncs

import (
	"net/http"

	"git.randomchars.net/Reviath/RemiliaScarlet/sql"
	"github.com/gin-gonic/gin"
)

func ContactHandler(c *gin.Context) {
	val, _ := c.Cookie("key")
	switch val {
	case "":
		c.HTML(200, "contact.html", gin.H{
			"login": "nil",
		})
	default:
		c.HTML(200, "contact.html", gin.H{
			"login": "yes",
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
