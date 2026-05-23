package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"url_shortner/internal/database"
	"url_shortner/internal/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)


func main(){

	// connect to DB
	database.Connect()
	database.Migrate()
	
	app := fiber.New()

	// middleware
	app.Use(logger.New())

	// health check
	app.Get("/health", func (c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status": "ok",
			"message": "API is healthy",
		})
	})

	// Register all app routes 

	routes.RegisterRoutes(app)

	// Graceful shutdown
	go func ()  {
		if err := app.Listen(":3000"); err != nil {
			log.Fatalf("Error starting server: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit
	
	
	log.Println("Shutting down server...")

	if err := app.Shutdown(); err != nil {
		log.Fatalf("Error shutting down server: %v", err)
	}
}