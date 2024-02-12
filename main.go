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
	if _, exists := os.LookupEnv("RAILWAY_ENVIRONMENT"); exists == false {
		if err := godotenv.Load(); err != nil {
			log.Print("No .env file found")
		}
	}
}

func main() {
	ctx := context.Background()

	connectionString := os.Getenv("DB_CONNECTION_STRING")

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
