package handlers

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"

	"github.com/KnlnKS/gr8-limiter/services/gr8-limiter/services"
)

func HandleRateLimit(c *fiber.Ctx) error {
	id, ok := c.GetReqHeaders()["Api-Key"]
	if !ok {
		return c.Status(fiber.StatusBadRequest).SendString("API key missing from header.")
	}

	limitString := c.Get("limit")
	windowSizeString := c.Get("windowSize")

	limit, err := strconv.ParseInt(limitString, 10, 64)
	if err != nil {
		limit = services.RedisDefaultLimit
	}

	windowSize, err := strconv.ParseInt(windowSizeString, 10, 64)
	if err != nil {
		windowSize = services.RedisDefaultWindowSize
	}

	s, err := services.NewRateLimterConfig(
		id,
		services.WithLimit(limit),
		services.WithWindowSize(windowSize),
	)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Unable to config rate limiter")
	}

	fmt.Println(s)

	// validate headers

	// 1. check if api key in cache
	// 2. not in cache, check in db

	return c.SendString("Hello, World 👋!")
}
