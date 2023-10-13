package handlers

import (
	"context"
	"strconv"

	"github.com/gofiber/fiber/v2"

	"github.com/KnlnKS/gr8-limiter/services/gr8-limiter/services/ratelimit"
)

func HandleRateLimit(c *fiber.Ctx) error {
	id, ok := c.GetReqHeaders()["Api-Key"]
	if !ok {
		return c.Status(fiber.StatusBadRequest).SendString("API key missing from header.")
	}

	// TODO:
	// 1. check if api key in cache
	// 2. not in cache, check in db

	limitString := c.Get("limit")
	windowSizeString := c.Get("windowSize")

	// validate headers
	limit, err := strconv.ParseInt(limitString, 10, 64)
	if err != nil {
		limit = ratelimit.RedisDefaultLimit
	}

	windowSize, err := strconv.ParseInt(windowSizeString, 10, 64)
	if err != nil {
		windowSize = ratelimit.RedisDefaultWindowSize
	}

	s, err := ratelimit.NewRateLimterConfig(
		id,
		ratelimit.WithLimit(limit),
		ratelimit.WithWindowSize(windowSize),
	)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Unable to config rate limiter")
	}

	// TODO: how should we handle context
	if s.RateLimit(context.TODO()) {
		return c.SendStatus(fiber.StatusOK)
	}
	return c.SendStatus(fiber.StatusTooManyRequests)
}
