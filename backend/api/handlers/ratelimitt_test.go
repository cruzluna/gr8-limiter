package handlers

import (
	"io"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"

	"github.com/KnlnKS/gr8-limiter/services/gr8-limiter/internal/cache"
)

func TestRateLimitHandler(t *testing.T) {
	app := fiber.New()
	api := app.Group("/api").Group("/v1")

	// api/v1/ratelimit
	api.Post("/ratelimit", HandleRateLimit)

	req := httptest.NewRequest("POST", "/ratelimit", nil)
	resp, _ := app.Test(req)
	assert.Equal(t, 404, resp.StatusCode, "Route not found.")

	req = httptest.NewRequest("POST", "/api/v1/ratelimit", nil)
	resp, _ = app.Test(req)
	assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode, "Missing api key header")

	req.Header.Set("X-Api-Key", "dummyApiKey")

	cache.InitApiKeyCache(1)

	resp, _ = app.Test(req)
	body, _ := io.ReadAll(resp.Body)
	assert.Equal(t, "API key is incorrect type.", string(body), "Incorrect API key header format")

	// TODO: mock database
	// TODO: Concurrent http requests
}
