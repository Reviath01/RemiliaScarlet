package webfuncs

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
)

func LoginHandler(c *gin.Context, conf *oauth2.Config, state string) {
	c.Redirect(http.StatusTemporaryRedirect, conf.AuthCodeURL(state))
}
