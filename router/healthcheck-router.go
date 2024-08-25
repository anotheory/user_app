package router

import (
	"user_app/flow"

	"github.com/gofiber/fiber/v2"
)

func InitHealthcheckRouter(app *fiber.App) {
	router := app.Group("/healthcheck")

	router.Get("/", flow.HealthcheckFlow{}.GetHealthcheck)
}
