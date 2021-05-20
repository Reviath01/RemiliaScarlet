package webfuncs

import "github.com/gin-gonic/gin"

func NoRouteHandler(c *gin.Context) {
	c.HTML(404, "error.html", gin.H{
		"is404":       "true",
		"description": "The page you are looking for may have been removed or temporarily unavailable",
		"error":       "404",
	})
}
