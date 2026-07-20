package middleware

import (
	"context"
	"url_shortner/internal/database"
	"url_shortner/internal/utils"

	"github.com/gofiber/fiber/v2"
)

// CaptureClickAnalytics middleware captures analytics data and stores it

func CaptureClickAnalytics(linkID int) fiber.Handler{
	return  func(c *fiber.Ctx) error {
		// Extract headers

		ip := utils.GetClientIP(
			c.Get("X-Forwarded-For"),
			c.Get("X-Real-IP"),
			c.IP(),
		)

		userAgent := c.Get("User-Agent")
		referer := c.Get("Referer")
		deviceType := utils.DetectDeviceType(userAgent)
		country := utils.GetCountry(ip)

		// Log click event asynchronously (don't block the redirect)
		go func() {
			ctx := context.Background()
			_ = database.LogClickEvent(ctx, linkID, ip, userAgent, referer, deviceType, country)

		}()
		return c.Next()

	}
}