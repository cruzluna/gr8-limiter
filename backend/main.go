package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/joho/godotenv"

	"github.com/KnlnKS/gr8-limiter/services/gr8-limiter/api/router"
	"github.com/KnlnKS/gr8-limiter/services/gr8-limiter/internal/analytics"
	"github.com/KnlnKS/gr8-limiter/services/gr8-limiter/internal/cache"
	"github.com/KnlnKS/gr8-limiter/services/gr8-limiter/internal/database"
	"github.com/KnlnKS/gr8-limiter/services/gr8-limiter/services/ratelimit"
)

func main() {
	env := flag.String("env", "local", "Environment flag")
	flag.Parse()

	// go run main.go -env=local|prod
	if *env == "local" {
		err := godotenv.Load()
		if err != nil {
			log.Fatalln("Error loading .env file: ", err)
		}
	}

	dbUrl := os.Getenv("DB_URL")
	fmt.Println("DB URL: ", dbUrl)
	posthogKey := os.Getenv("POSTHOG_KEY")
	fmt.Println("POSTHOG KEY: ", posthogKey)

	// Postgres
	ctx := context.Background()
	err := database.Init(ctx, dbUrl)
	if err != nil {
		log.Fatalln("Error Connecting to Postgres: ", err)
	}

	// redis
	redisUrl := os.Getenv("REDIS_URL")
	ratelimit.Init(redisUrl)

	// API key cache
	cache.InitApiKeyCache(512)

	// posthog
	err = analytics.Init(posthogKey)
	if err != nil {
		log.Fatalln("Error connecting to Posthog: ", err)
	}
	defer analytics.Client.Close()

	app := fiber.New()

	app.Use(recover.New())

	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path} ${latency}\n",
	}))

	router.PublicRoutes(app)

	port := os.Getenv("PORT")
	if port == "" {
		port = "10000"
	}

	if err = app.Listen(":" + port); err != nil {
		log.Fatalln("Unable to start server. ", err)
	}
}
