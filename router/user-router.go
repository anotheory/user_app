package router

import (
	"user_app/flow"

	"github.com/gofiber/fiber/v2"
)

func InitUserRouter(app *fiber.App) {
	router := app.Group("/api/users")

	router.Get("/:username", flow.UserFlow{}.GetUser)
	router.Post("/", flow.UserFlow{}.CreateUser)
}
