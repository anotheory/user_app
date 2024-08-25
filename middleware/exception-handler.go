package middleware

import (
	"errors"
	"user_app/models/exception"

	"github.com/gofiber/fiber/v2"
)

var notFoundException *exception.NotFoundException
var validationException *exception.ValidationException
var fiberError *fiber.Error

func ErrorHandler(c *fiber.Ctx, err error) error {
	if errors.As(err, &notFoundException) {
		c.Status(fiber.ErrBadRequest.Code).JSON(err.(*exception.NotFoundException))
	} else if errors.As(err, &validationException) {
		c.Status(fiber.ErrNotFound.Code).JSON(err.(*exception.ValidationException))
	} else {
		if errors.As(err, &fiberError) {
			c.Status(fiberError.Code).SendString(err.Error())
		}
	}
	return nil
}
