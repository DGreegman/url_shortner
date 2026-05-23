package handlers

import (
	"context"
	"time"
	"url_shortner/internal/database"

	"github.com/gofiber/fiber/v2"
)


func Redirect(c *fiber.Ctx) error {
	code := c.Params("code")

	if code == "" {
		return fiber.NewError(fiber.StatusBadRequest, "Missing Code")
	}

	// find URL in DB 

	urlData, err := database.FindUrlByCode(code)

	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, "URL Not Found")
	}

	// Check if URL is expired

	if !urlData.ExpireAt.IsZero() && time.Now().After(urlData.ExpireAt) {
		return fiber.NewError(fiber.StatusGone, "URL Expired")
	}

	// Increment click count (optional, can be done asynchronously)

	_, _ = database.DB.Exec(
		context.Background(),
		`UPDATE urls SET clicks = clicks + 1 WHERE code = $1`,
		code,
	)

	// Redirect to target URL

	return c.Redirect(urlData.TargetUrl, fiber.StatusMovedPermanently)
}