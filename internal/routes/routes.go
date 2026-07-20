package routes

import (
	"url_shortner/internal/handlers"

	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(app *fiber.App) {
	api := app.Group("/api/v1")
	api.Post("/shorten", handlers.ShortenURL)
	api.Get("/analytics/:code", handlers.GetAnalytics)

	// Redirect route (catch-all for short codes, MUST be registered last)
	app.Get("/:code", handlers.Redirect)
}