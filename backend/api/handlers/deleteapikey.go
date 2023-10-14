package handlers

import (
	"context"
	"log"

	"github.com/gofiber/fiber/v2"

	"github.com/KnlnKS/gr8-limiter/services/gr8-limiter/internal/cache"
	"github.com/KnlnKS/gr8-limiter/services/gr8-limiter/internal/database"
)

func HandleDeleteApiKey(c *fiber.Ctx) error {
	// TODO: validate user has auth to delete api key
	id, ok := c.GetReqHeaders()["Api-Key"]
	if !ok {
		return c.Status(fiber.StatusBadRequest).SendString("API key missing from header.")
	}

	// 1. Remove API key from API go cache
	cache.ApiKeyCache.Remove(id)

	// 2. Remove from DB
	if database.Conn.IsApiKeyInTable(context.Background(), id) {
		err := database.Conn.DeleteByApiKey(context.Background(), id)
		if err != nil {
			log.Println(err)
			return c.Status(fiber.StatusBadRequest).
				SendString("Something went wrong while deleting Api Key")
		}

		return c.SendStatus(fiber.StatusOK)

	}

	return c.Status(fiber.StatusBadRequest).
		SendString("API key does not exist in database")
}
