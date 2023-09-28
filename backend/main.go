package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
)

func main() {
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

	ctx := context.Background()

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
