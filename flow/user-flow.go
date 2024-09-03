package flow

import (
	"regexp"
	"user_app/db"
	"user_app/models/entity"
	"user_app/models/exception"
	"user_app/models/user"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var usernamePattern string = "^[a-zA-Z]([\\w]{6,48})[a-zA-Z0-9]$"
var passwordPattern string = "\\s+"

type UserFlow struct{}

func (flow UserFlow) GetUser(c *fiber.Ctx) error {
	username := c.Params("username")
	var err error
	if username == "" {
		err = exception.NewNotFoundException("Requested username is empty.")
		return err
	}

	userEntity := entity.UserEntity{}
	dbResults := db.DbConnection.Model(&entity.UserEntity{}).Preload("UserLoginHistories").Where("username = ?", username).First(&userEntity)
	if dbResults.Error != nil {
		err = exception.NewDatabaseException(dbResults.Error.Error())
		return err
	}
	c.JSON(userEntity)
	return nil
}

func (flow UserFlow) CreateUser(c *fiber.Ctx) error {
	request := new(user.CreateUserRequest)
	var err error
	if err = c.BodyParser(request); err != nil {
		err = exception.NewValidationException("Malformed request body.")
		return err
	}
	var matched bool = false
	validateFunc := validator.New(validator.WithRequiredStructEnabled())
	if err = validateFunc.Struct(request); err != nil {
		err = exception.NewValidationException("Malformed request body.")
		return err
	} else if matched, _ = regexp.MatchString(usernamePattern, request.Username); !matched {
		err = exception.NewValidationException("Wrong format on username.")
		return err
	} else if matched, _ = regexp.MatchString(passwordPattern, request.Password); matched {
		err = exception.NewValidationException("Wrong format on password.")
		return err
	}
	return c.Status(200).JSON(request)
}
