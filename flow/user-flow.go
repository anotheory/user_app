package flow

import (
	"user_app/db"
	"user_app/models/entity"
	"user_app/models/exception"

	"github.com/gofiber/fiber/v2"
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
	c.JSON(userEntity)
	return nil
}
