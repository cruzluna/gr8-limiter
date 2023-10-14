package router

import (
	"github.com/gofiber/fiber/v2"

	"github.com/KnlnKS/gr8-limiter/services/gr8-limiter/api/handlers"
)

func PublicRoutes(router *fiber.App) {
	api := router.Group("/api")
	v1 := api.Group("/v1")

	// api/v1/ratelimit
	v1.Post("/ratelimit", handlers.HandleRateLimit)

	// api/v1/apikey
	v1.Delete("/apikey", handlers.HandleDeleteApiKey)

	router.Use(func(c *fiber.Ctx) error {
		return c.Status(404).SendString("Not found.")
	})
}
