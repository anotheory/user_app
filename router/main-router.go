package router

import "github.com/gofiber/fiber/v2"

func InitRouter(app *fiber.App) *fiber.App {
	InitHealthcheckRouter(app)
	InitUserRouter(app)

	return app
}
