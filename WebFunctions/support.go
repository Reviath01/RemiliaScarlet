package webfuncs

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SupportHandler(c *gin.Context) {
	c.Redirect(http.StatusTemporaryRedirect, "https://discord.com/users/770218429096656917")
}
