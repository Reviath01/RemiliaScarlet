package webfuncs

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	jsoniter "github.com/json-iterator/go"
	"golang.org/x/oauth2"
)

func CallbackHandler(c *gin.Context, state string, conf *oauth2.Config) {
	if c.Request.FormValue("state") != state {
		c.Redirect(http.StatusTemporaryRedirect, "/")
	}
	token, err := conf.Exchange(context.TODO(), c.Query("code"))
	if err != nil {
		fmt.Println("An error occurred while getting token: " + err.Error())
		return
	}
	tokenJSON, err := jsoniter.Marshal(token)
	if err != nil {
		fmt.Println(err)
		return
	}
	c.SetCookie("key", string(tokenJSON), 60*60*24*7, "/", "localhost", false, false)
	c.Redirect(http.StatusTemporaryRedirect, "/")

}
