package main

import (
	"context"
	"items/internal/database"
	"items/internal/http"

	"github.com/miguel-rosa/go-city-server/inernal/database"

	"github.com/gin-gonic/gin"
	"github.com/miguel-rosa/go-city-server/inernal/http"
)

func main() {
	ctx := context.Background()

	connectionString := "postgressql://post:p0stgr3s@db:5432/posts"
	conn, err := database.NewConnection(ctx, connectionString)

	if err != nil {
		panic(err)
	}

	defer conn.Close()

	g := gin.New()
	g.Use(gin.Recovery())
	http.Configure()
	http.SetRoutes(g)
	g.Run(":3000")
}
