package main

import (
	"user_app/middleware"
	"user_app/models/exception"
	"user_app/models/healthcheck"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New(fiber.Config{
		ErrorHandler: middleware.ErrorHandler,
	})
	app.Get("/healthcheck", func(c *fiber.Ctx) error {
		c.JSON(healthcheck.HealthcheckResponse{
			Version: "latest",
			Message: "I'm fine, thank you :)",
		})
		return nil
	})

	app.Get("/api/users", func(c *fiber.Ctx) error {
		err := exception.NewValidationException("test message")
		return err
	})

	app.Use(cors.New())
	app.Listen(":5000")
}
