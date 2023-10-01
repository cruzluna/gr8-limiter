package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"

	"github.com/KnlnKS/gr8-limiter/services/gr8-limiter/internal/database"
)

func main() {
	ctx := context.Background()

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbUrl := os.Getenv("DB_URL")
	fmt.Println("DB URL: ", dbUrl)
	db, err := database.StartDatabase(ctx, dbUrl)
	if err != nil {
		log.Panicln(err)
	}

	record := database.ApiTableRecord{
		ApiKey: RandomString(11),
		UserId: rand.Int31(),
	}

	dbErr := db.Insert(ctx, record)

	if dbErr != nil {
		log.Panicln(dbErr)
	}

	// os.Exit(1)

	app := fiber.New()

	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		DB:   0,
	})
	defer rdb.Close()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	app.Post("/api-key", func(c *fiber.Ctx) error {
		// generate an api key with you desired config
		return c.SendString("Here is your API Key 👋!" + RandomString(10))
	})

	// 4 reqs per second
	rc := NewRateLimterConfig(RandomString(10), 4, 1)

	for i := 0; i < 20; i++ {
		allowed := rc.rateLimit(ctx, rdb)
		if allowed {
			fmt.Printf("Request %d allowed.\n", i)
		} else {
			fmt.Printf("Request %d blocked.\n", i)
		}

		time.Sleep(100 * time.Millisecond)
	}

	app.Listen(":3000")
}

// throwaway....just for testing
func RandomString(n int) string {
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	s := make([]rune, n)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}
	return string(s)
}
