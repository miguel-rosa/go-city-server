package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetRoutes(g *gin.Engine) {
	g.POST("/posts", PostPosts)
	g.GET("/posts/:id", GetPost)
	g.GET("/posts/", GetAll)

	g.GET("/health", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})
}
