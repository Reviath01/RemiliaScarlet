package webfuncs

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func LogoutHandler(c *gin.Context) {
	val, _ := c.Cookie("key")
	if val != "" {
		c.SetCookie("key", "nil", -1, "/", "localhost", false, false)
		c.Redirect(http.StatusTemporaryRedirect, "/")
	} else {
		c.Redirect(http.StatusTemporaryRedirect, "/")
	}

}
