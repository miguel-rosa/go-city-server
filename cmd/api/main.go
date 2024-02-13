package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

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

func envPortOr(port string) string {
	// If `PORT` variable in environment exists, return it
	if envPort := os.Getenv("PORT"); envPort != "" {
		return ":" + envPort
	}
	// Otherwise, return the value of `port` variable from function argument
	return ":" + port
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

	g.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("%s - [%s] \" %s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))

	g.Use(gin.Recovery())
	http.Configure()
	http.SetRoutes(g)
	var port = envPortOr("3000")
	g.Run(port)
}
