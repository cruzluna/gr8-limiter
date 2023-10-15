package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/joho/godotenv"

	"github.com/KnlnKS/gr8-limiter/services/gr8-limiter/api/router"
	"github.com/KnlnKS/gr8-limiter/services/gr8-limiter/internal/database"
	"github.com/KnlnKS/gr8-limiter/services/gr8-limiter/services/ratelimit"
)

func main() {
	ctx := context.Background()
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("Error loading .env file: ", err)
	}

	dbUrl := os.Getenv("DB_URL")
	fmt.Println("DB URL: ", dbUrl)
	err = database.Init(ctx, dbUrl)

	if err != nil {
		log.Fatalln("Error: ", err)
	}

	// redis
	ratelimit.Init()
	// cache.Init(10)

	app := fiber.New()

	app.Use(recover.New())

	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path} ${latency}\n",
	}))

	router.PublicRoutes(app)

	app.Listen(":3000")
}
