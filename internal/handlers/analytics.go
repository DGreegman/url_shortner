package handlers

import (
	"log"
	"url_shortner/internal/database"

	"github.com/gofiber/fiber/v2"
)

func GetAnalytics(c *fiber.Ctx) error {
	code := c.Params("code")

	if code == ""{
		return fiber.NewError(fiber.StatusBadRequest, "Missing code parameter")

	}

	analytics, err := database.GetAnalytics(c.Context(), code)

	if err != nil {
		log.Fatalf("Failed to retrieve analytics for code %s: %s", code, err.Error())
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to retrieve analytics")
		
	}

	return c.JSON(analytics)
}