package router

import (
	"github.com/gofiber/fiber/v2"
)

type UserRouter struct {
	Router fiber.Router
}

func (c *UserRouter) Init(app *fiber.App) {
	c.Router = app.Group("/api/users")
}
