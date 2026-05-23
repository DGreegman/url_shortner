package handlers

import (
	"context"
	"log"
	"time"
	"url_shortner/internal/database"
	"url_shortner/internal/models"
	"url_shortner/internal/utils"

	"github.com/gofiber/fiber/v2"
)

func ShortenURL(c *fiber.Ctx) error {

	body := new(models.ShortenRequest)

	if err := c.BodyParser(body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}
	
	if body.URL == "" {
		return fiber.NewError(fiber.StatusBadRequest, "URL is required...")
	}

	// Generate short code, save to DB, and return response

	code := utils.GenerateShortCode(6)

	var expiresAt *time.Time 

	if body.ExpireIn > 0 {
		t := time.Now().Add(time.Duration(body.ExpireIn) * time.Second)
		expiresAt = &t
	}

	// Insert into DB

	query := `INSERT INTO urls (code, target_url, expire_at) VALUES ($1, $2, $3) RETURNING code`

	err := database.DB.QueryRow(
		context.Background(),
		query,
		code,
		body.URL,
		expiresAt,
	).Scan(&code)

	if err !=nil {
		log.Printf("failed to insert into database %v", err.Error())
		return fiber.NewError(fiber.StatusInternalServerError, "Database Insert Failed")
	}
	response := fiber.Map{
		"short_url" : "http://localhost:3000/" + code,
	}

	return c.JSON(response)
}