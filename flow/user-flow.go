package flow

import (
	"user_app/db"
	"user_app/logic"
	"user_app/logic/utility"
	"user_app/models/entity"
	"user_app/models/exception"
	"user_app/models/user"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

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
	return c.Status(200).JSON(userEntity)
}

func (flow UserFlow) CreateUser(c *fiber.Ctx) error {
	request := new(user.CreateUserRequest)
	var err error
	if err = c.BodyParser(request); err != nil {
		err = exception.NewValidationException("Malformed request body.")
		return err
	}
	if err = logic.ValidateCreateUserRequest(*request); err != nil {
		return err
	}
	var dbResults *gorm.DB
	duplicateUserEntity := entity.UserEntity{}
	dbResults = db.DbConnection.Model(&entity.UserEntity{}).Where("username = ? OR email = ? ", request.Username, request.Email).First(&duplicateUserEntity)
	if dbResults.RowsAffected > 0 {
		err = exception.NewValidationException("The requested user info is already exists.")
		return err
	}

	userEntity := entity.UserEntity{
		Username: request.Username,
		Password: utility.CalculateHmacString(request.Password),
		Email:    request.Email,
	}
	dbResults = db.DbConnection.Create(&userEntity)
	if dbResults.Error != nil {
		err = exception.NewDatabaseException(dbResults.Error.Error())
		return err
	}
	return c.Status(200).JSON(userEntity)
}
