package config

import (
	"iot-sensor/internal/delivery/http"
	"iot-sensor/internal/delivery/http/route"
	"iot-sensor/internal/gateway/messaging"
	"iot-sensor/internal/usecase"
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
	// setup MQTT broker
	v := config.Config
	topic := v.GetString("MQTT_TOPIC")
	id1 := v.GetString("SENSOR_ID_1")
	id2 := v.GetInt("SENSOR_ID_2")
	sensorType := v.GetString("SENSOR_TYPE")
	interval := v.GetInt("PUBLISH_INTERVAL")
	pubInterval := time.Duration(interval) * time.Second
	min := v.GetFloat64("SENSOR_VALUE_MIN")
	max := v.GetFloat64("SENSOR_VALUE_MAX")
	sensorPubliser := messaging.NewSensorPublisher(
		config.Log,
		*config.Mqtt,
		topic,
		id1,
		id2,
		sensorType,
		pubInterval,
		min,
		max,
	)

	// setup use cases
	sensorUseCase := usecase.NewSensorUsecase(config.Log, config.Validate, sensorPubliser)

	// setup controller
	sensorController := http.NewSensorController(sensorUseCase, config.Log)

	routeConfig := route.RouteConfig{
		App:              config.App,
		SensorController: sensorController,
	}
	routeConfig.Setup()

	sensorPubliser.Start()
}
