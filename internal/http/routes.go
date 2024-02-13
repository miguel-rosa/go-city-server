package http

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type person struct {
	name     string
	quantity int
}

func respTest(ctx *gin.Context) {

	// response.quantity = 2

	ctx.JSON(http.StatusOK, gin.H{
		"name": "miguel",
	})
}

func SetRoutes(g *gin.Engine) {
	g.GET("/posts", GetAll)
	g.GET("/post/:id", GetPost)
	g.POST("/post", PostPosts)
	g.POST("/test", respTest)

	g.GET("/health", func(c *gin.Context) {
		fmt.Printf("OK!")
		c.Status(http.StatusOK)
	})
}
