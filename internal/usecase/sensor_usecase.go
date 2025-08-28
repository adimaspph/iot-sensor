package usecase

import (
	"context"
	"fmt"
	"iot-sensor/internal/gateway/messaging"
	"iot-sensor/internal/model"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

type SensorUsecase struct {
	Log             *logrus.Logger
	Validate        *validator.Validate
	sensorPublisher *messaging.SensorPublisher
}

func NewSensorUsecase(
	logger *logrus.Logger,
	validate *validator.Validate,
	sensorPublisher *messaging.SensorPublisher,
) *SensorUsecase {
	return &SensorUsecase{
		Log:             logger,
		Validate:        validate,
		sensorPublisher: sensorPublisher,
	}
}

func (u *SensorUsecase) ChangeInterval(ctx context.Context, request *model.ChangeIntervalRequest) (msg string, err error) {
	// validate
	if err := u.Validate.Struct(request); err != nil {
		u.Log.WithError(err).Error("failed to validate request body")
		return "", echo.ErrBadRequest
	}

	interval := time.Duration(request.Interval) * time.Second
	u.sensorPublisher.SetInterval(interval)

	msg = fmt.Sprintf("sensor publish interval updated successfully to %v second", request.Interval)
	return msg, nil
}
