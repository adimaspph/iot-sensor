package config

import (
	"context"
	"iot-sensor/internal/delivery/messaging"
	"os"
	"os/signal"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type BootstrapConfig struct {
	App      *echo.Echo
	Log      *logrus.Logger
	Validate *validator.Validate
	Config   *viper.Viper
	Mqtt     *mqtt.Client
}

func Bootstrap(config *BootstrapConfig) {
	// setup usecase

}
