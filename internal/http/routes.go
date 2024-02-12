package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetRoutes(g *gin.Engine) {
	g.GET("/posts", GetAll)
	g.GET("/post/:id", GetPost)
	g.POST("/post", PostPosts)

	g.GET("/health", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})
}
