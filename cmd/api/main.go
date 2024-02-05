package main

import (
	"context"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/miguel-rosa/go-city-server/internal/database"
	"github.com/miguel-rosa/go-city-server/internal/http"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	ctx := context.Background()

	connectionString, exists := os.LookupEnv("DB_CONNECTION_STRING")

	if !exists {
		log.Print("Missing DB_CONNECTION_STRING on env")

	}
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
