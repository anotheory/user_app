package flow

import (
	"user_app/models/entity"
	"user_app/models/exception"
	"user_app/models/healthcheck"
	"user_app/repository"

	"github.com/gofiber/fiber/v2"
)

type HealthcheckFlow struct{}

func (flow HealthcheckFlow) GetHealthcheck(c *fiber.Ctx) error {
	var err error

	baseRepo := repository.BaseRepository{}
	healthcheckEntity := entity.HealthcheckEntity{}
	dbResults := repository.IBaseRepository.GetById(baseRepo, 1, &healthcheckEntity)
	if dbResults.Error != nil {
		err = exception.NewDatabaseException(dbResults.Error.Error())
		return err
	}
	c.JSON(healthcheck.HealthcheckResponse{
		Version: "latest",
		Message: healthcheckEntity.Message,
	})
	return nil
}
