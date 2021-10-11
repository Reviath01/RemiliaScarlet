package webfuncs

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SupportHandler(c *gin.Context) {
	c.Redirect(http.StatusTemporaryRedirect, "https://github.com/Reviath01/RemiliaScarlet/issues/new")
}
