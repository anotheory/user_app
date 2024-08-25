package main

import (
	"user_app/middleware"
	"user_app/router"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New(fiber.Config{
		ErrorHandler: middleware.ErrorHandler,
	})

	app = router.InitRouter(app)

	app.Use(cors.New())
	app.Listen(":5000")
}
