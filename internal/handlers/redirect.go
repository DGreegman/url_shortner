package handlers

import (
	"context"
	"time"
	"url_shortner/internal/database"
	"url_shortner/internal/utils"

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


	go func ()  {
		ip := c.IP()
		userAgent := c.Get("User-Agent")
		referer := c.Get("Referer")
		deviceType := utils.DetectDeviceType(userAgent)
		country := utils.GetCountry(ip)

		_ = database.LogClickEvent(
			context.Background(),
			urlData.ID,
			ip,
			userAgent,
			referer,
			deviceType,
			country,
		)
		
	}()

	// Increment click count (optional, can be done asynchronously)
	_, _ = database.DB.Exec(
		context.Background(),
		`UPDATE urls SET clicks = clicks + 1 WHERE code = $1`,
		code,
	)



	statusCode := fiber.StatusFound // default 302

	switch urlData.RedirectType {
	case "301":
		statusCode = fiber.StatusMovedPermanently
	case "302":
		statusCode = fiber.StatusFound
	case "307":
		statusCode = fiber.StatusTemporaryRedirect
	default:
		statusCode = fiber.StatusFound
	}

	// Redirect to target URL

	return c.Redirect(urlData.TargetUrl, statusCode)
}