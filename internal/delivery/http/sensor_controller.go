package http

import (
	"iot-sensor/internal/model"
	"iot-sensor/internal/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

type SensorController struct {
	UseCase *usecase.SensorUsecase
	Log     *logrus.Logger
}

func NewSensorController(useCase *usecase.SensorUsecase, log *logrus.Logger) *SensorController {
	return &SensorController{
		UseCase: useCase,
		Log:     log,
	}
}

func (c SensorController) ChangeInterval(ctx echo.Context) error {
	var request model.ChangeIntervalRequest

	err := ctx.Bind(&request)
	if err != nil {
		c.Log.WithError(err).Error("failed to bind request")
		return err
	}

	message, err := c.UseCase.ChangeInterval(ctx.Request().Context(), &request)
	if err != nil {
		c.Log.WithError(err).Error("failed to create sensor record")
		return err
	}

	return ctx.JSON(http.StatusOK, model.MessageResponse[string]{
		Message: message,
	})
}
