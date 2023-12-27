package handlers

import (
	"context"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"

	"github.com/KnlnKS/gr8-limiter/services/gr8-limiter/internal/cache"
	"github.com/KnlnKS/gr8-limiter/services/gr8-limiter/internal/database"
	"github.com/KnlnKS/gr8-limiter/services/gr8-limiter/services/ratelimit"
)

func HandleRateLimit(c *fiber.Ctx) error {
	id, ok := c.GetReqHeaders()["X-Api-Key"]
	if !ok {
		return c.Status(fiber.StatusBadRequest).SendString("API key missing from header.")
	}

	// 1. check if api key in cache
	_, ok = cache.ApiKeyCache.Get(id)
	if !ok {
		// 2. not in cache, check in db
		_, err := uuid.Parse(id)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).SendString("API key is incorrect type.")
		}

		inTable, err := database.Conn.IsApiKeyInTable(context.Background(), id)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("API key database error.")
		}
		if !inTable {
			return c.Status(fiber.StatusBadRequest).SendString("API key is invalid.")
		}
		// inTable, add to cache
		go func() {
			cache.ApiKeyCache.Add(id, true)
		}()
	}

	limitString := c.Get("limit")
	windowSizeString := c.Get("windowSize")

	// validate headers
	limit, err := strconv.ParseInt(limitString, 10, 64)
	// TODO: more validation of rate limiting config, ie window too large
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
